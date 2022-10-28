package staker

import (
	"math/big"
	"testing"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/rawdb"
	"github.com/Ankr-network/coqchain/core/state"
	"github.com/Ankr-network/coqchain/params"
	"github.com/stretchr/testify/assert"
)

var testState *state.StateDB
var testConfig = &params.PosaConfig{
	Epoch:                  100,
	SealerBalanceThreshold: big.NewInt(100),
}
var testAddrs = []common.Address{
	common.HexToAddress("0xb156c5aa93B682E6E5bf294eE32160d277b95161"),
	common.HexToAddress("0x6A92F2E354228e866C44419860233Cc23bec0d8A"),
	common.HexToAddress("0xA319658456A578470A6b532F11A8f9f9608cf902"),
}

func testConstructor() {
	db := rawdb.NewMemoryDatabase()
	testState, _ = state.New(common.Hash{}, state.NewDatabase(db), nil)

	Constructor(testState, testAddrs, testConfig)
}

func TestCheck(t *testing.T) {
	testConstructor()
	assert.Equal(t, getEpoch(testState), big.NewInt(int64(testConfig.Epoch)))

	assert.Equal(t, getThreshold(testState), testConfig.SealerBalanceThreshold)

	assert.Equal(t, getFineRatio(testState), big.NewInt(10))

	for i := range testAddrs {
		assert.Equal(t, GetBalance(testState, testAddrs[i]), testConfig.SealerBalanceThreshold)
		assert.True(t, signersContain(testState, testAddrs[i]))
	}

	assert.False(t, signersContain(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594")))

	assert.Equal(t, singerListNum(testState), big.NewInt(int64(len(testAddrs))))

	addSigner(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594"))
	assert.True(t, signersContain(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594")))
	removeSigner(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594"))
	assert.False(t, signersContain(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594")))
	removeSigner(testState, testAddrs[1])

	t.Log(SingerList(testState))

	hash := common.HexToHash("0x000000000000000000000001ab8483f64d9c6d1ecf9b849ae677dd3315835cb2")
	p := toVoteInfo(hash)
	t.Log(p.vote)
	t.Log(p.authorize)

	assert.True(t, p.toHash() == hash)

}

func TestVote(t *testing.T) {
	testConstructor()
	NeedAddSigner1 := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	NeedAddSigner2 := common.HexToAddress("0x24b265fa8B5241c30020EB239d65FE1aCdF97737")
	NeedAddSigner3 := common.HexToAddress("0x94b265fa8B5241c30020EB239d65FE1aCdF97737")
	fakeAddr := common.HexToAddress("0xc805C32B3D9a29E54F6c01d4d0a322697BE23C64")

	testVote(testState, big.NewInt(1), testAddrs[1], false, fakeAddr)
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[1], big.NewInt(0)), common.Hash{})
	testVote(testState, big.NewInt(1), testAddrs[1], false, testAddrs[2])
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[1], big.NewInt(0)), (&voteInfo{vote: testAddrs[2]}).toHash())

	assert.True(t, getVoteNum(testState, testAddrs[1]).Cmp(big.NewInt(1)) == 0)

	testVote(testState, big.NewInt(1), testAddrs[0], false, testAddrs[2])
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[0], big.NewInt(0)), (&voteInfo{vote: testAddrs[2]}).toHash())
	testVote(testState, big.NewInt(1), testAddrs[0], false, testAddrs[1])
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[0], big.NewInt(1)), common.Hash{})
	testVote(testState, big.NewInt(1), testAddrs[0], true, testAddrs[0])
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[0], big.NewInt(0)), common.Hash{})

	testVote(testState, big.NewInt(1), NeedAddSigner1, true, fakeAddr)
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner1, big.NewInt(0)), common.Hash{})
	testVote(testState, big.NewInt(1), NeedAddSigner1, true, testAddrs[1])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner1, big.NewInt(0)), common.Hash{})
	testVote(testState, big.NewInt(1), NeedAddSigner1, false, testAddrs[0])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner1, big.NewInt(1)), common.Hash{})

	testVote(testState, big.NewInt(1), NeedAddSigner2, true, testAddrs[2])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner2, big.NewInt(0)), common.Hash{})
	testVote(testState, big.NewInt(1), NeedAddSigner2, true, testAddrs[1])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner2, big.NewInt(1)), common.Hash{})
	testVote(testState, big.NewInt(1), NeedAddSigner2, false, testAddrs[0])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner2, big.NewInt(2)), common.Hash{})

	setBalance(testState, NeedAddSigner3, testConfig.SealerBalanceThreshold)
	testVote(testState, big.NewInt(1), NeedAddSigner3, true, testAddrs[1])

	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner3, big.NewInt(0)), (&voteInfo{vote: testAddrs[1], authorize: true}).toHash())
	testVote(testState, big.NewInt(1), NeedAddSigner3, false, fakeAddr)
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner3, big.NewInt(1)), common.Hash{})
	testVote(testState, big.NewInt(1), NeedAddSigner3, true, testAddrs[2])
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner3, big.NewInt(0)), common.Hash{})

	t.Log("voted signers: ", SingerList(testState))

	t.Log("voted signers num: ", singerListNum(testState))

	t.Log("testAddrs[0]", GetBalance(testState, testAddrs[0]))
	t.Log("testAddrs[1]", GetBalance(testState, testAddrs[1]))
	t.Log("testAddrs[2]", GetBalance(testState, testAddrs[2]))

	t.Log("NeedAddSigner1", GetBalance(testState, NeedAddSigner1))
	t.Log("NeedAddSigner2", GetBalance(testState, NeedAddSigner2))
	t.Log("NeedAddSigner3", GetBalance(testState, NeedAddSigner3))

	cleanAllVotee(testState)

	assert.True(t, len(voteeList(testState)) == 0)
	assert.True(t, voteeListNum(testState).Cmp(big.NewInt(0)) == 0)

	assert.True(t, getVoteNum(testState, testAddrs[0]).Cmp(big.NewInt(0)) == 0)
	assert.True(t, getVoteNum(testState, testAddrs[1]).Cmp(big.NewInt(0)) == 0)
	assert.True(t, getVoteNum(testState, testAddrs[2]).Cmp(big.NewInt(0)) == 0)
	assert.True(t, getVoteNum(testState, NeedAddSigner1).Cmp(big.NewInt(0)) == 0)
	assert.True(t, getVoteNum(testState, NeedAddSigner2).Cmp(big.NewInt(0)) == 0)
	assert.True(t, getVoteNum(testState, NeedAddSigner3).Cmp(big.NewInt(0)) == 0)

	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[0], big.NewInt(0)), common.Hash{})
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[1], big.NewInt(0)), common.Hash{})
	assert.Equal(t, getProposalVoteInfo(testState, testAddrs[2], big.NewInt(0)), common.Hash{})
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner1, big.NewInt(0)), common.Hash{})
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner2, big.NewInt(0)), common.Hash{})
	assert.Equal(t, getProposalVoteInfo(testState, NeedAddSigner3, big.NewInt(0)), common.Hash{})

}

func testVote(statedb *state.StateDB, blockNo *big.Int, votee common.Address, authorize bool, signer common.Address) {
	if getThreshold(statedb).Cmp(big.NewInt(0)) <= 0 {
		return
	}
	epoch := getEpoch(statedb)
	if big.NewInt(0).Mod(blockNo, epoch) == big.NewInt(0) {
		cleanAllVotee(statedb)
	} else {
		if votee != (common.Address{}) {

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
