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

func TestMain(m *testing.M) {
	db := rawdb.NewMemoryDatabase()
	testState, _ = state.New(common.Hash{}, state.NewDatabase(db), nil)

	Constructor(testState, testAddrs, testConfig)
	m.Run()
}

func TestCheck(t *testing.T) {
	assert.Equal(t, getEpoch(testState), big.NewInt(int64(testConfig.Epoch)))

	assert.Equal(t, getThreshold(testState), testConfig.SealerBalanceThreshold)

	assert.Equal(t, getFineRatio(testState), big.NewInt(10))

	assert.Equal(t, getCycle(testState, big.NewInt(1)), big.NewInt(0))

	assert.Equal(t, checkEpochVoted(testState, big.NewInt(0)), false)

	// setEpochVoted()

	assert.True(t, checkBalanceGreaterThreshold(testState, testAddrs[0]))

	for i := range testAddrs {
		assert.Equal(t, GetBalance(testState, testAddrs[i]), testConfig.SealerBalanceThreshold)
		assert.True(t, signersContain(testState, testAddrs[i]))
	}

	assert.False(t, signersContain(testState, common.HexToAddress("0xE0804972d5535a5764dfdbD432ebE58F0419C594")))

	assert.Equal(t, singerListNum(testState), big.NewInt(int64(len(testAddrs))))

	t.Log(SingerList(testState))
}

func TestVote(t *testing.T) {
	cycle := big.NewInt(0)
	NeedAddSigner1 := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	NeedAddSigner2 := common.HexToAddress("0x24b265fa8B5241c30020EB239d65FE1aCdF97737")
	NeedAddSigner3 := common.HexToAddress("0x94b265fa8B5241c30020EB239d65FE1aCdF97737")
	fakeAddr := common.HexToAddress("0xc805C32B3D9a29E54F6c01d4d0a322697BE23C64")
	vote(testState, cycle, testAddrs[1], false, fakeAddr)
	vote(testState, cycle, testAddrs[1], false, testAddrs[2])

	vote(testState, cycle, testAddrs[0], false, testAddrs[2])
	vote(testState, cycle, testAddrs[0], false, testAddrs[1])
	vote(testState, cycle, testAddrs[0], true, testAddrs[0])

	vote(testState, cycle, NeedAddSigner1, true, fakeAddr)
	vote(testState, cycle, NeedAddSigner1, true, testAddrs[1])
	vote(testState, cycle, NeedAddSigner1, false, testAddrs[0])

	vote(testState, cycle, NeedAddSigner2, true, testAddrs[2])
	vote(testState, cycle, NeedAddSigner2, true, testAddrs[1])
	vote(testState, cycle, NeedAddSigner2, false, testAddrs[0])

	setBalance(testState, NeedAddSigner3, testConfig.SealerBalanceThreshold)
	vote(testState, cycle, NeedAddSigner3, true, testAddrs[1])
	vote(testState, cycle, NeedAddSigner3, false, fakeAddr)
	vote(testState, cycle, NeedAddSigner3, true, testAddrs[2])

	assert.Equal(t, getEpochProposalVoteType(testState, cycle, testAddrs[1]), voteTypeExit)
	assert.Equal(t, getEpochProposalVoteType(testState, cycle, testAddrs[0]), voteTypeExit)
	assert.Equal(t, getEpochProposalVoteType(testState, cycle, NeedAddSigner1), voteTypeJoin)
	assert.Equal(t, getEpochProposalVoteType(testState, cycle, NeedAddSigner2), voteTypeJoin)

	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, testAddrs[1], fakeAddr), voteResUnknow)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, testAddrs[1], testAddrs[2]), voteResAgree)

	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, testAddrs[0], testAddrs[2]), voteResAgree)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, testAddrs[0], testAddrs[1]), voteResAgree)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, testAddrs[0], testAddrs[0]), voteResAgainst)

	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner1, fakeAddr), voteResUnknow)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner1, testAddrs[1]), voteResAgree)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner1, testAddrs[0]), voteResAgainst)

	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner2, testAddrs[2]), voteResAgree)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner2, testAddrs[1]), voteResAgree)
	assert.Equal(t, getEpochProposalVoteMap(testState, cycle, NeedAddSigner2, testAddrs[0]), voteResAgainst)

	assert.Equal(t, getEpochProposalVoteesNumByCycle(testState, cycle), big.NewInt(5))

	t.Log(testAddrs[1], "voted: ", getEpochProposalVoteList(testState, cycle, testAddrs[1]))
	assert.Equal(t, len(getEpochProposalVoteList(testState, cycle, testAddrs[1])), 1)

	t.Log(testAddrs[0], "voted: ", getEpochProposalVoteList(testState, cycle, testAddrs[0]))
	assert.Equal(t, len(getEpochProposalVoteList(testState, cycle, testAddrs[0])), 3)

	t.Log(NeedAddSigner1, "voted: ", getEpochProposalVoteList(testState, cycle, NeedAddSigner1))
	assert.Equal(t, len(getEpochProposalVoteList(testState, cycle, NeedAddSigner1)), 2)

	t.Log(NeedAddSigner2, "voted: ", getEpochProposalVoteList(testState, cycle, NeedAddSigner2))
	assert.Equal(t, len(getEpochProposalVoteList(testState, cycle, NeedAddSigner2)), 3)

	t.Log(NeedAddSigner3, "voted: ", getEpochProposalVoteList(testState, cycle, NeedAddSigner3))
	assert.Equal(t, len(getEpochProposalVoteList(testState, cycle, NeedAddSigner3)), 2)

	processLastVote(testState, big.NewInt(int64(testConfig.Epoch)))
	processLastVote(testState, big.NewInt(int64(testConfig.Epoch)))

	t.Log("voted signers: ", SingerList(testState))

	t.Log("voted signers num: ", singerListNum(testState))

	assert.True(t, checkEpochVoted(testState, big.NewInt(0).Sub(getCycle(testState, big.NewInt(int64(testConfig.Epoch))), big.NewInt(1))))
	assert.Equal(t, singerListNum(testState), big.NewInt(3))

}
