package staker

import (
	"bytes"
	"fmt"
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

type voteType uint8

const (
	voteTypeUnknow voteType = iota
	voteTypeJoin
	voteTypeExit
)

type voteRes uint8

const (
	voteResUnknow voteRes = iota
	voteResAgree
	voteResAgainst
)

/*
   enum VoteType {
       UNKNOW,
       JOIN,
       EXIT
   }

   enum VoteRes {
       UNKNOW,
       AGREE,
       AGAINST
   }

   struct Proposal {
       VoteType voteType;
       mapping(address => VoteRes) voteMaps;
       address[] votes;
   }
*/

const (
	epochSlot               uint8 = iota // uint256 public epoch;
	thresholdSlot                        // uint256 public threshold;
	fineRatioSlot                        // uint256 public fineRatio;
	signersListSlot                      // signers list
	signersMapSlot                       // signers mapping
	balancesSlot                         // mapping(address => uint256) public balances;
	epochProposalsSlot                   // mapping(uint256 => mapping(address => Proposal)) public epochProposals;
	epochProposalVoteesSlot              // mapping(uint256 => address[]) public epochProposalVotees;
	epochVotedSlot                       // mapping(uint256 => bool) epochVoted;
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

			// Keccak256(k.p)
			statedb.SetState(contracts.SlashAddr, crypto.Keccak256Hash(
				common.RightPadBytes(addr.Bytes(), 32), // bytes32
				common.LeftPadBytes([]byte{signersMapSlot}, 32),
			), common.HexToHash("0x1"))

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
		processLastVote(statedb, block.Number)
	} else {
		if block.Coinbase != (common.Address{}) {
			var authorize bool

			signer, err := ecrecover(block)
			fmt.Println("staker vote singer:", signer)
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
				vote(statedb, getCycle(statedb, block.Number), block.Coinbase, authorize, signer)
			}
			return
		}
	}
}

func checkBalanceGreaterThreshold(statedb *state.StateDB, address common.Address) bool {
	return GetBalance(statedb, address).Cmp(getThreshold(statedb)) >= 0
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

func vote(statedb *state.StateDB, cycle *big.Int, votee common.Address, authorize bool, signer common.Address) {
	if signersContain(statedb, signer) {
		var vt voteType
		var res voteRes
		if signersContain(statedb, votee) {
			vt = voteTypeExit
			if authorize {
				res = voteResAgainst
			} else {
				res = voteResAgree
			}
		} else {
			vt = voteTypeJoin
			if authorize {
				res = voteResAgree
			} else {
				res = voteResAgainst
			}
		}

		if getEpochProposalVoteType(statedb, cycle, votee) == voteTypeUnknow {
			setEpochProposalVoteType(statedb, cycle, votee, vt)
			addEpochProposalVotees(statedb, cycle, votee)
		}

		if getEpochProposalVoteMap(statedb, cycle, votee, signer) == voteResUnknow {
			setEpochProposalVoteList(statedb, cycle, votee, signer)
		}
		setEpochProposalVoteMap(statedb, cycle, votee, signer, res)
	}
}

// Processing the results of the last round of voting
func processLastVote(statedb *state.StateDB, blockNum *big.Int) {
	cycle := getCycle(statedb, blockNum)

	if cycle.Cmp(big.NewInt(0)) > 0 {
		lastCycle := big.NewInt(0).Sub(cycle, big.NewInt(1))

		LastEpochProposalVotees := getEpochProposalVoteesNumByCycle(statedb, lastCycle)
		if LastEpochProposalVotees.Cmp(big.NewInt(0)) > 0 &&
			!checkEpochVoted(statedb, lastCycle) {

			signersNum := singerListNum(statedb)
			signersMedian := big.NewInt(0).Div(signersNum, big.NewInt(2))
			ProposalListKey := crypto.Keccak256Hash(
				crypto.Keccak256(
					common.LeftPadBytes(lastCycle.Bytes(), 32),
					common.LeftPadBytes([]byte{epochProposalVoteesSlot}, 32),
				)).Big()

			for i := big.NewInt(0); LastEpochProposalVotees.Cmp(i) > 0; i = i.Add(i, big.NewInt(1)) {
				proposalVotee := common.BytesToAddress(statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(ProposalListKey, i))).Bytes())
				agreeNum := big.NewInt(0)
				proposalVoteType := getEpochProposalVoteType(statedb, lastCycle, proposalVotee)
				epochProposalsHandle(statedb, lastCycle, proposalVotee, func(vote common.Address, res voteRes) {
					if res == voteResAgree {
						agreeNum.Add(agreeNum, big.NewInt(1))
					}
				})

				if agreeNum.Cmp(signersMedian) > 0 {
					if proposalVoteType == voteTypeJoin {
						proposalJoinHandle(statedb, proposalVotee)
					} else if proposalVoteType == voteTypeExit {
						proposalExitHandle(statedb, proposalVotee)
					}
				}
			}
			setEpochVoted(statedb, lastCycle, true)
		}
	}
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

func getCycle(statedb *state.StateDB, number *big.Int) *big.Int {
	return big.NewInt(0).Div(number, getEpoch(statedb))
}

func checkEpochVoted(statedb *state.StateDB, cycle *big.Int) bool {
	return statedb.GetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.BigToHash(cycle).Bytes(),
			common.BytesToHash([]byte{epochVotedSlot}).Bytes(),
		),
	).Big().Text(10) == "1"
}

func setEpochVoted(statedb *state.StateDB, cycle *big.Int, res bool) {
	value := common.Hash{}
	if res {
		value = common.BytesToHash([]byte{1})
	}
	statedb.SetState(contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.BigToHash(cycle).Bytes(),
			common.BytesToHash([]byte{epochVotedSlot}).Bytes(),
		),
		value)
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
	return statedb.GetState(contracts.SlashAddr, crypto.Keccak256Hash(
		common.RightPadBytes(address.Bytes(), 32), // bytes32
		common.LeftPadBytes([]byte{signersMapSlot}, 32),
	)) != common.Hash{}
}

func removeSigner(statedb *state.StateDB, address common.Address) {
	if signersContain(statedb, address) {
		statedb.SetState(contracts.SlashAddr, crypto.Keccak256Hash(
			common.RightPadBytes(address.Bytes(), 32), // bytes32
			common.LeftPadBytes([]byte{signersMapSlot}, 32),
		), common.Hash{})

		indexes := statedb.GetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32))).Big()

		listSlot := crypto.Keccak256Hash(common.LeftPadBytes([]byte{signersListSlot}, 32)).Big()
		var exist bool
		for i := big.NewInt(0); indexes.Cmp(i) > 0; i.Add(i, big.NewInt(1)) {
			if common.BytesToAddress(
				statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, i))).Bytes(),
			) == address {
				exist = true
				statedb.SetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, i)), statedb.GetState(contracts.SlashAddr, common.BigToHash(big.NewInt(0).Add(listSlot, big.NewInt(0).Sub(indexes, big.NewInt(1))))))
				break
			}
		}
		if exist {
			indexes.Sub(indexes, big.NewInt(1))
			statedb.SetState(contracts.SlashAddr, common.BytesToHash(common.LeftPadBytes([]byte{signersListSlot}, 32)), common.BigToHash(indexes))
		}
	}
}

func addSigner(statedb *state.StateDB, address common.Address) {
	if !signersContain(statedb, address) {
		statedb.SetState(contracts.SlashAddr, crypto.Keccak256Hash(
			common.RightPadBytes(address.Bytes(), 32), // bytes32
			common.LeftPadBytes([]byte{signersMapSlot}, 32),
		), common.BytesToHash([]byte{1}))

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

func getEpochProposalVoteesNumByCycle(statedb *state.StateDB, cycle *big.Int) *big.Int {
	return statedb.GetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalVoteesSlot}, 32),
		),
	).Big()
}

func addEpochProposalVotees(statedb *state.StateDB, cycle *big.Int, address common.Address) {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(cycle.Bytes(), 32),
		common.LeftPadBytes([]byte{epochProposalVoteesSlot}, 32),
	)

	length := statedb.GetState(
		contracts.SlashAddr,
		key,
	).Big()

	statedb.SetState(
		contracts.SlashAddr,
		common.BigToHash(big.NewInt(0).Add(length, crypto.Keccak256Hash(key.Bytes()).Big())),
		address.Hash(),
	)

	statedb.SetState(
		contracts.SlashAddr,
		key,
		common.BigToHash(big.NewInt(0).Add(length, big.NewInt(1))),
	)
}

func getEpochProposalVoteType(statedb *state.StateDB, cycle *big.Int, votee common.Address) voteType {
	return voteType(statedb.GetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(votee.Bytes(), 32),
			crypto.Keccak256(
				common.LeftPadBytes(cycle.Bytes(), 32),
				common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
			),
		),
	).Big().Uint64())
}

func setEpochProposalVoteType(statedb *state.StateDB, cycle *big.Int, votee common.Address, vt voteType) {
	statedb.SetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(votee.Bytes(), 32),
			crypto.Keccak256(
				common.LeftPadBytes(cycle.Bytes(), 32),
				common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
			),
		),
		common.BigToHash(big.NewInt(int64(vt))),
	)
}

func getEpochProposalVoteMap(statedb *state.StateDB, cycle *big.Int, votee, vote common.Address) voteRes {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		crypto.Keccak256(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
		),
	).Big()

	return voteRes(statedb.GetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(vote.Bytes(), 32),
			common.BigToHash(key.Add(key, big.NewInt(1))).Bytes(),
		)).Big().Uint64())
}

func setEpochProposalVoteMap(statedb *state.StateDB, cycle *big.Int, votee, vote common.Address, vr voteRes) {
	key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		crypto.Keccak256(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
		),
	).Big()

	statedb.SetState(
		contracts.SlashAddr,
		crypto.Keccak256Hash(
			common.LeftPadBytes(vote.Bytes(), 32),
			common.BigToHash(key.Add(key, big.NewInt(1))).Bytes(),
		),
		common.BigToHash(big.NewInt(int64(vr))),
	)
}

func setEpochProposalVoteList(statedb *state.StateDB, cycle *big.Int, votee, vote common.Address) {
	Key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		crypto.Keccak256(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
		),
	).Big()

	ListKey := big.NewInt(0).Add(Key, big.NewInt(2))

	voteeLen := statedb.GetState(
		contracts.SlashAddr,
		common.BigToHash(ListKey),
	).Big()

	statedb.SetState(
		contracts.SlashAddr,
		common.BigToHash(
			big.NewInt(0).Add(
				crypto.Keccak256Hash(
					common.BigToHash(ListKey).Bytes(),
				).Big(),
				voteeLen,
			)),
		common.BytesToHash(vote.Bytes()),
	)

	statedb.SetState(
		contracts.SlashAddr,
		common.BigToHash(ListKey),
		common.BigToHash(voteeLen.Add(voteeLen, big.NewInt(1))),
	)
}

func getEpochProposalVoteList(statedb *state.StateDB, cycle *big.Int, votee common.Address) []common.Address {
	Key := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		crypto.Keccak256(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
		),
	).Big()

	ListKey := big.NewInt(0).Add(Key, big.NewInt(2))

	voteeLen := statedb.GetState(
		contracts.SlashAddr,
		common.BigToHash(ListKey),
	).Big()

	list := make([]common.Address, voteeLen.Uint64())
	for i := big.NewInt(0); i.Cmp(voteeLen) < 0; i.Add(i, big.NewInt(1)) {
		list[i.Uint64()] = common.BytesToAddress(statedb.GetState(
			contracts.SlashAddr,
			common.BigToHash(
				big.NewInt(0).Add(
					crypto.Keccak256Hash(
						common.BigToHash(ListKey).Bytes(),
					).Big(),
					i,
				)),
		).Bytes())
	}
	return list
}

func epochProposalsHandle(statedb *state.StateDB, cycle *big.Int, votee common.Address, callback func(vote common.Address, res voteRes)) {
	epochProposalsVoteeKey := crypto.Keccak256Hash(
		common.LeftPadBytes(votee.Bytes(), 32),
		crypto.Keccak256(
			common.LeftPadBytes(cycle.Bytes(), 32),
			common.LeftPadBytes([]byte{epochProposalsSlot}, 32),
		),
	).Big()

	epochProposalsVoteeMapKey := big.NewInt(0).Add(epochProposalsVoteeKey, big.NewInt(1))
	epochProposalsVoteeListKey := big.NewInt(0).Add(epochProposalsVoteeKey, big.NewInt(2))

	voteeLen := statedb.GetState(
		contracts.SlashAddr,
		common.BigToHash(epochProposalsVoteeListKey),
	).Big()

	for i := big.NewInt(0); i.Cmp(voteeLen) < 0; i.Add(i, big.NewInt(1)) {

		vote := common.BytesToAddress(
			statedb.GetState(
				contracts.SlashAddr,
				common.BigToHash(big.NewInt(0).Add(
					crypto.Keccak256Hash(common.BigToHash(epochProposalsVoteeListKey).Bytes()).Big(),
					i,
				)),
			).Bytes())
		callback(
			vote,
			voteRes(statedb.GetState(
				contracts.SlashAddr,
				crypto.Keccak256Hash(
					common.LeftPadBytes(vote.Bytes(), 32),
					common.LeftPadBytes(epochProposalsVoteeMapKey.Bytes(), 32),
				),
			).Big().Uint64()),
		)

	}
}
