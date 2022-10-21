package staker

import (
	"bytes"
	"math/big"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/common/hexutil"
	"github.com/Ankr-network/coqchain/core/contracts"
	"github.com/Ankr-network/coqchain/core/state"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/crypto"
	"github.com/Ankr-network/coqchain/params"
)

var (
	nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.
)

/*
   uint256 public epoch;
   uint256 public threshold;
   uint256 public fineRatio; // 10 = 10%

   address[] signers;
   mapping(address => uint256) public balances;

   struct VoteInfo {
       address vote;
       bool authorize;
   }
   mapping(address => VoteInfo[]) public proposalVotes;
   address[] public votees;
*/

const (
	epochSlot         uint8 = iota // uint256 public epoch;
	thresholdSlot                  // uint256 public threshold;
	fineRatioSlot                  // uint256 public fineRatio;
	signersListSlot                // signers list
	balancesSlot                   // mapping(address => uint256) public balances;
	proposalsVoteslot              // mapping(address => VoteInfo[]) public proposalVotes;
	voteesSlot                     // address[] public votees;
)

func Constructor(statedb *state.StateDB, validatorList []common.Address, config *params.PosaConfig) {
	threshold := config.SealerBalanceThreshold
	if threshold.Uint64() > 0 {
		var signerCnt int64 = 0

		statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{epochSlot}), common.BigToHash(big.NewInt(0).SetUint64(config.Epoch)))
		statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{thresholdSlot}), common.BigToHash(threshold))
		statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{fineRatioSlot}), common.BigToHash(big.NewInt(10))) // 10%

		slotAddrBig := common.BytesToHash(crypto.Keccak256(
			common.LeftPadBytes([]byte{signersListSlot}, 32),
		)).Big()

		for _, addr := range validatorList {
			statedb.SubBalance(addr, threshold)
			// Keccak256(p) + signerCnt
			statedb.SetState(contracts.SlashAddr, common.BigToHash(
				big.NewInt(0).Add(
					slotAddrBig,
					big.NewInt(signerCnt),
				),
			), common.BytesToHash(addr.Bytes()))

			statedb.SetState(contracts.SlashAddr, crypto.Keccak256Hash(
				common.LeftPadBytes(addr.Bytes(), 32), // address
				common.LeftPadBytes([]byte{balancesSlot}, 32),
			), common.BigToHash(threshold))

			signerCnt++
		}
		statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{signersListSlot}), common.BigToHash(big.NewInt(signerCnt)))
		statedb.AddBalance(contracts.SlashAddr, big.NewInt(0).Mul(big.NewInt(signerCnt), threshold))
	}
}

func Vote(statedb *state.StateDB, block *types.Header) {
	if getThreshold(statedb).Cmp(big.NewInt(0)) <= 0 {
		return
	}
	epoch := getEpoch(statedb)
	if big.NewInt(0).Mod(block.Number, epoch) == big.NewInt(0) {
		cleanAllVotee(statedb)
	} else {
		votee := block.Coinbase
		if votee != (common.Address{}) {
			var authorize bool

			signer, err := ecrecover(block)
			if err != nil {
				return
			}

			switch {
			case bytes.Equal(block.Nonce[:], nonceAuthVote):
				authorize = true
			case bytes.Equal(block.Nonce[:], nonceDropVote):
				authorize = false
			default:
				return
			}
			if signersContain(statedb, signer) {
				existVotee := signersContain(statedb, votee)
				if (existVotee && !authorize) || (!existVotee && authorize) && GetBalance(statedb, votee).Cmp(getThreshold(statedb)) >= 0 {
					vote(statedb, votee, authorize, signer)
					checkProposal(statedb)
				}
			}
			return
		}
	}
}

func vote(statedb *state.StateDB, votee common.Address, authorize bool, signer common.Address) {
	if !voteeContain(statedb, votee) {
		addVotee(statedb, votee)
	}

	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.LeftPadBytes([]byte{proposalsVoteslot}, 32),
	)
	len := statedb.GetState(contracts.SlashAddr, key).Big()

	index := crypto.Keccak256Hash(
		key.Bytes(),
	).Big()

	var exit bool
	for i := big.NewInt(0); i.Cmp(len) < 0; i.Add(i, big.NewInt(1)) {
		info := toVoteInfo(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))))
		if info.vote == signer {
			exit = true
			info.authorize = authorize
			statedb.SetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i)), info.toHash())
		}
	}
	if !exit {
		index.Add(index, len)
		info := voteInfo{
			vote:      signer,
			authorize: authorize,
		}
		statedb.SetState(contracts.SlashAddr, common.BigToHash(index), info.toHash())
		statedb.SetState(contracts.SlashAddr, key, common.BigToHash(len.Add(len, big.NewInt(1))))
	}
}

func checkProposal(statedb *state.StateDB) {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	listNum := voteeListNum(statedb)
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		votee := common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes())
		voteeIsSigner := signersContain(statedb, votee)
		if proposalVotesHandle(statedb, voteeIsSigner, votee) {
			cleanVotee(statedb, votee)
			statedb.SetState(contracts.SlashAddr,
				common.BigToHash(big.NewInt(0).Add(index, i)),
				common.BytesToHash([]byte{}),
			)
			removeVotee(statedb, votee)
			if voteeIsSigner {

			}
		}
	}
}

func removeSignerVote(statedb *state.StateDB, votee common.Address) {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	listNum := voteeListNum(statedb)
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		checkProposalVoteeToRemove(statedb, common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes()))
	}
}

func checkProposalVoteeToRemove(statedb *state.StateDB, votee common.Address) {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
	)
	len := statedb.GetState(contracts.SlashAddr, key).Big()

	index := crypto.Keccak256Hash(key.Bytes()).Big()
	for i := big.NewInt(0); i.Cmp(len) < 0; i.Add(i, big.NewInt(1)) {
		if toVoteInfo(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i)))).vote == votee {

			statedb.SetState(
				contracts.SlashAddr,
				common.BigToHash(big.NewInt(0).Add(index, i)),
				statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, len.Sub(len, big.NewInt(1))))),
			)

			statedb.SetState(
				contracts.SlashAddr,
				common.BigToHash(big.NewInt(0).Add(index, len)),
				common.BytesToHash([]byte{}),
			)

			statedb.SetState(contracts.SlashAddr, key, common.BigToHash(len))
		}
	}

	statedb.SetState(contracts.SlashAddr, key, common.BytesToHash([]byte{}))

}

func proposalVotesHandle(statedb *state.StateDB, voteeIsSigner bool, votee common.Address) bool {
	middleNum := big.NewInt(0).Div(singerListNum(statedb), big.NewInt(2))

	res := !voteeIsSigner
	var count = big.NewInt(0)

	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
	)
	len := statedb.GetState(contracts.SlashAddr, key).Big()

	index := crypto.Keccak256Hash(key.Bytes()).Big()
	for i := big.NewInt(0); i.Cmp(len) < 0; i.Add(i, big.NewInt(1)) {
		if toVoteInfo(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i)))).authorize == res {
			count.Add(count, big.NewInt(1))
		}
	}

	if count.Cmp(middleNum) > 0 {
		if res {
			proposalJoinHandle(statedb, votee)
		} else {
			proposalExitHandle(statedb, votee)
		}
		return true
	}
	return false
}

func proposalJoinHandle(statedb *state.StateDB, votee common.Address) {
	if GetBalance(statedb, votee).Cmp(getThreshold(statedb)) >= 0 {
		addSigner(statedb, votee)
	}
}

func proposalExitHandle(statedb *state.StateDB, votee common.Address) {
	removeSigner(statedb, votee)
	balance := GetBalance(statedb, votee)

	balance.Sub(balance,
		big.NewInt(0).Div(
			big.NewInt(0).Mul(getThreshold(statedb), getFineRatio(statedb)),
			big.NewInt(100),
		),
	)

	setBalance(statedb, votee, balance)
}

func getEpoch(statedb *state.StateDB) *big.Int {
	return statedb.GetState(contracts.SlashAddr, common.BytesToHash([]byte{epochSlot})).Big()
}

func getThreshold(statedb *state.StateDB) *big.Int {
	return statedb.GetState(contracts.SlashAddr, common.BytesToHash([]byte{thresholdSlot})).Big()
}

func getFineRatio(statedb *state.StateDB) *big.Int {
	return statedb.GetState(contracts.SlashAddr, common.BytesToHash([]byte{fineRatioSlot})).Big()
}

func singerListNum(statedb *state.StateDB) *big.Int {
	return statedb.GetState(contracts.SlashAddr, common.BytesToHash([]byte{signersListSlot})).Big()
}

func SingerList(statedb *state.StateDB) []common.Address {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{signersListSlot}).Bytes(),
	).Big()
	listNum := singerListNum(statedb)
	list := make([]common.Address, listNum.Uint64())
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		list[i.Int64()] = common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes())
	}
	return list
}

func signersContain(statedb *state.StateDB, address common.Address) bool {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{signersListSlot}).Bytes(),
	).Big()
	listNum := singerListNum(statedb)
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		if common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes()) == address {
			return true
		}
	}
	return false
}

func removeSigner(statedb *state.StateDB, address common.Address) {
	if signersContain(statedb, address) {
		indexes := statedb.GetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32))).Big()

		listSlot := crypto.Keccak256Hash(common.LeftPadBytes([]byte{signersListSlot}, 32)).Big()

		for i := big.NewInt(0); indexes.Cmp(i) > 0; i.Add(i, big.NewInt(1)) {
			if common.BytesToAddress(
				statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, i))).Bytes(),
			) == address {
				statedb.SetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, i)), statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, indexes.Sub(indexes, big.NewInt(1))))))
				statedb.SetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32)), common.BigToHash(indexes))
				break
			}
		}
	}
}

func addSigner(statedb *state.StateDB, address common.Address) {
	if !signersContain(statedb, address) {
		indexes := statedb.GetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32))).Big()

		listSlot := crypto.Keccak256Hash(common.LeftPadBytes([]byte{signersListSlot}, 32)).Big()
		listSlot.Add(listSlot, indexes)

		statedb.SetState(
			contracts.SlashAddr,
			common.BigToHash(listSlot),
			common.BytesToHash(address.Bytes()),
		)

		indexes.Add(indexes, big.NewInt(1))
		statedb.SetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32)), common.BigToHash(indexes))
	}
}

func GetBalance(statedb *state.StateDB, address common.Address) *big.Int {
	return statedb.GetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(address.Bytes(), 32), // address
			common.LeftPadBytes([]byte{balancesSlot}, 32),
		),
	).Big()
}

func setBalance(statedb *state.StateDB, address common.Address, amount *big.Int) {
	statedb.SetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(address.Bytes(), 32), // address
			common.LeftPadBytes([]byte{balancesSlot}, 32),
		),
		common.BytesToHash(amount.Bytes()),
	)
}

func voteeListNum(statedb *state.StateDB) *big.Int {
	return statedb.GetState(contracts.SlashAddr, common.BytesToHash([]byte{voteesSlot})).Big()
}

func voteeList(statedb *state.StateDB) []common.Address {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	listNum := voteeListNum(statedb)
	list := make([]common.Address, listNum.Uint64())
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		list[i.Int64()] = common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes())
	}
	return list
}

func addVotee(statedb *state.StateDB, votee common.Address) {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	len := voteeListNum(statedb)
	statedb.SetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, len)), common.BytesToHash(votee.Bytes()))

	len.Add(len, big.NewInt(1))
	statedb.SetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{voteesSlot}, 32)), common.BigToHash(len))
}

func removeVotee(statedb *state.StateDB, votee common.Address) {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	listNum := voteeListNum(statedb)

	var exist bool
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		if common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes()) == votee {
			exist = true
			statedb.SetState(contracts.SlashAddr,
				common.BigToHash(big.NewInt(0).Add(index, i)),
				statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, listNum.Sub(listNum, big.NewInt(1))))),
			)
		}
	}

	if exist {
		statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{voteesSlot}), common.BigToHash(listNum))
	}
}

func voteeContain(statedb *state.StateDB, votee common.Address) bool {
	return statedb.GetState(contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(votee.Bytes(), 32),
			common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
		)).Big().Cmp(big.NewInt(0)) != 0
}

func cleanAllVotee(statedb *state.StateDB) {
	index := crypto.Keccak256Hash(
		common.BytesToHash([]byte{voteesSlot}).Bytes(),
	).Big()
	listNum := voteeListNum(statedb)
	for i := big.NewInt(0); i.Cmp(listNum) < 0; i.Add(i, big.NewInt(1)) {
		cleanVotee(statedb, common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i))).Bytes()))
		statedb.SetState(contracts.SlashAddr,
			common.BigToHash(big.NewInt(0).Add(index, i)),
			common.BytesToHash([]byte{}),
		)
	}

	statedb.SetState(contracts.SlashAddr, common.BytesToHash([]byte{voteesSlot}), common.BytesToHash([]byte{}))
}

func cleanVotee(statedb *state.StateDB, votee common.Address) {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
	)
	len := statedb.GetState(contracts.SlashAddr, key).Big()

	index := crypto.Keccak256Hash(key.Bytes()).Big()
	for i := big.NewInt(0); i.Cmp(len) < 0; i.Add(i, big.NewInt(1)) {
		statedb.SetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(index, i)), common.BytesToHash([]byte{}))
	}

	statedb.SetState(contracts.SlashAddr, key, common.BytesToHash([]byte{}))
}

func getVoteNum(statedb *state.StateDB, votee common.Address) *big.Int {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
	)

	return statedb.GetState(contracts.SlashAddr, key).Big()
}

func getProposalVoteInfo(statedb *state.StateDB, votee common.Address, index *big.Int) common.Hash {
	voteeKey := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		common.BytesToHash([]byte{proposalsVoteslot}).Bytes(),
	)

	indexKey := crypto.Keccak256Hash(
		voteeKey.Bytes(),
	).Big()

	indexKey.Add(indexKey, index)

	return statedb.GetState(contracts.SlashAddr, common.BigToHash(indexKey))
}

/*
struct VoteInfo {
	address vote;
	bool authorize;
}
*/

type voteInfo struct {
	vote      common.Address
	authorize bool
}

func (p *voteInfo) toHash() common.Hash {
	var bs []byte = make([]byte, 0, len(p.vote.Bytes())+1)
	if p.authorize {
		bs = append(bs, []byte{1}...)
	}
	bs = append(bs, p.vote.Bytes()...)
	return common.BytesToHash(bs)
}

func toVoteInfo(hash common.Hash) *voteInfo {
	bs := hash.Bytes()[len(hash.Bytes())-1-common.AddressLength:]
	authorize := bs[0] == byte(1)
	return &voteInfo{
		vote:      common.BytesToAddress(bs[1:]),
		authorize: authorize,
	}
}
