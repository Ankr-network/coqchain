// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// base on clique

package dpos

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	_ "github.com/ethereum/go-ethereum/consensus/misc"
	_ "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru"
)

const (
	signerReward       = 50   // the percent of signer reward
	maxSignerSize      = 17   // max signer in one epoch
	storeSnapInterval  = 1024 // when block number mod storeSnapInterval == 0, then store it into database
	inmemorySnapshots  = 1024 // cache snap size
	inmemorySignatures = 4096 //cache the result of ecrecover

	wiggleTime = 500 * time.Millisecond
)

var (
	BlockReward = big.NewInt(3e+18)

	epochLength = uint64(54000) // default every 15 minutes, rechoose signers

	// fill in block.header.nonce
	nonceYesVote = hexutil.MustDecode("0xffffffffffffffff") // agreed
	nonceNoVote  = hexutil.MustDecode("0x0000000000000000") // cancel

	uncleHash = types.CalcUncleHash(nil) // no uncle block

	diffInTurn = big.NewInt(2)
	diffNoTurn = big.NewInt(1)
)

var (
	errUnknownBlock                   = errors.New("Unknown block")
	errInvalidEpochVoting             = errors.New("Voting does not allow in epoch block")
	errInvalidNonEpochExtra           = errors.New("Non epoch block's extra only allow signature field")
	errInvalidEpochExtraSigner        = errors.New("Invalid signers contain in epoch block's extra")
	errInvalidEpochExtraProposal      = errors.New("Invalid proposals contain in epoch block's extra")
	errInvalidVote                    = errors.New("Vote nonce not 0x00..0 or 0xff..f")
	errInvalidEpochVote               = errors.New("Vote nonce in epoch block non-zero")
	errMissingSignature               = errors.New("Extra-data 65 byte signature suffix missing")
	errInvalidEpochSigners            = errors.New("Invalid signer list on epoch block")
	errMismatchingEpochSigners        = errors.New("Mismatching signer list on epoch block")
	errInvalidUncleHash               = errors.New("Non empty uncle hash")
	errInvalidDifficulty              = errors.New("Invalid difficulty")
	errWrongDifficultyAgainstSnap     = errors.New("Wrong difficulty against snapshot")
	errWrongDifficultyAgainstExtra    = errors.New("Wrong difficulty against header.extra")
	errInvalidTimestamp               = errors.New("Invalid timestamp")
	errInvalidVotingChain             = errors.New("Invalid voting chain")
	errUnauthorizedSignerAgainstSnap  = errors.New("Unauthorized signer against snapshot")
	errUnauthorizedSignerAgainstExtra = errors.New("Unauthorized signer against header.extra")
	errRecentlySigned                 = errors.New("Recently signed")
	errMissingBody                    = errors.New("Missing body")
	errWrongEpochNumber               = errors.New("Wrong epoch number")
	errMissingEpochBlock              = errors.New("Missing epoch block during stateless situation")
)

// SignerFn hashes and signs the data to be signed by a backing account.
type SignerFn func(signer accounts.Account, mimeType string, message []byte) ([]byte, error)

type Dpos struct {
	config *params.DposConfig
	db     ethdb.Database // save snapshot

	recents    *lru.ARCCache // save recently snapshot
	signatures *lru.ARCCache // save recently signatures

	proposals map[common.Hash]bool

	signer common.Address
	signFn SignerFn
	lock   sync.RWMutex

	fakeDiff bool // for test, skip check diffculty

	state *state.StateDB
}

func New(config *params.DposConfig, db ethdb.Database) *Dpos {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.EpochInterval == 0 {
		conf.EpochInterval = epochLength
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Dpos{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		proposals:  make(map[common.Hash]bool),
	}
}

func (this *Dpos) Authorize(signer common.Address, signFn SignerFn) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.signer = signer
	this.signFn = signFn
}

func (this *Dpos) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, this.signatures)
}

func (this *Dpos) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return this.verifyHeader(chain, header, nil)
}

func (this *Dpos) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {

	abort := make(chan struct{})

	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := this.verifyHeader(chain, header, headers[:i])
			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

func (this *Dpos) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {

	if len(block.Uncles()) != 0 {
		return errors.New("uncle is not allow")
	}

	return nil
}

func (this *Dpos) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {

	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	epochBlock := (number % this.config.EpochInterval) == 0
	// header.MixDigest record the proposal by signer
	if epochBlock && header.MixDigest != (common.Hash{}) {
		return errInvalidEpochVoting
	}

	if !bytes.Equal(header.Nonce[:], nonceYesVote) && !bytes.Equal(header.Nonce[:], nonceNoVote) {
		return errInvalidVote
	}

	if epochBlock && !bytes.Equal(header.Nonce[:], nonceNoVote) {
		return errInvalidEpochVote
	}

	// check header.extra

	// code rule:
	// include four kind element, which are sign/signers/proposal/delegator
	// first one byte is data length

	// all block extra must prefix with 0x41
	var extras [][]byte

	if header.Extra[0] != 0x41 {
		return errMissingSignature
	}

	if !epochBlock {

		if len(header.Extra) != crypto.SignatureLength+1 {
			return errInvalidNonEpochExtra
		}

	} else {
		extras = unserialize(header.Extra)

		// one signer at least
		if !(len(extras[1])%common.AddressLength == 0 && len(extras[1])/common.AddressLength > 0) {
			return errInvalidEpochExtraSigner
		}

		if !(len(extras[2])%common.HashLength == 0 && len(extras[2])/common.HashLength == len(Proposals)) {
			return errInvalidEpochExtraProposal
		} else {

			proposalCnt := len(extras[2]) / common.HashLength

			for i := 0; i < proposalCnt; i++ {
				proposal := &Proposal{}
				if err := proposal.fromBytes(common.BytesToHash(extras[2][i*common.HashLength : (i+1)*common.HashLength])); err != nil {
					return err
				}
			}
		}
	}

	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}

	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}

	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	if parent.Time+this.config.SlotInterval > header.Time {
		return errInvalidTimestamp
	}

	epochHeader := this.epochOfHeader(chain, header, parents)

	if epochHeader != nil {

		signer, err := ecrecover(header, this.signatures)
		if err != nil {
			return err
		}

		signers, _, _ := parseEpochExtra(epochHeader)
		totalSigners := len(signers)

		validSigner := false
		offset := 0

		for _, _signer := range signers {
			if _signer == signer {
				validSigner = true
			} else if !validSigner {
				offset++
			}
		}

		if !validSigner {
			return errUnauthorizedSignerAgainstExtra
		}

		inturn := (number % uint64(totalSigners)) == uint64(offset)

		if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
			return errWrongDifficultyAgainstExtra
		}

		if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
			return errWrongDifficultyAgainstExtra
		}
	} else if !epochBlock {
		return errMissingEpochBlock
	}

	return nil
}

func (this *Dpos) VerifySeal(chain consensus.ChainHeaderReader, header *types.Header) error {
	return this.verifySeal(chain, header, nil)
}

func (this *Dpos) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {

	extras := unserialize(header.Extra)

	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	epochBlock := (number % this.config.EpochInterval) == 0

	snap, err := this.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	if epochBlock {

		signers := make([]byte, len(snap.PreElectedSigners)*common.AddressLength)
		for i, signer := range snap.preElectedSigners() {
			copy(signers[i*common.AddressLength:], signer[:])
		}

		if !bytes.Equal(signers, extras[1]) {
			return errMismatchingEpochSigners
		}
	}

	signer, err := ecrecover(header, this.signatures)
	if err != nil {
		return err
	}

	if _, ok := snap.ElectedSigners[signer]; !ok {
		return errUnauthorizedSignerAgainstSnap
	}

	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only fail if the current block doesn't shift it out
			if limit := uint64(len(snap.ElectedSigners)/2 + 1); seen > number-limit {
				return errRecentlySigned
			}
		}
	}

	if !this.fakeDiff {
		inturn := snap.inturn(header.Number.Uint64(), signer)
		if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
			return errWrongDifficultyAgainstSnap
		}

		if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
			return errWrongDifficultyAgainstSnap
		}
	}

	return nil
}

func (this *Dpos) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {

	header.MixDigest = common.Hash{}

	// this feild is nil forever in DPOS
	header.Coinbase = common.Address{}

	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()

	if number%this.config.EpochInterval == 0 {
		this.proposals = make(map[common.Hash]bool)
	}

	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	header.Time = parent.Time + this.config.SlotInterval
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	return nil
}

func (this *Dpos) Finalize(chain consensus.ChainHeaderReader, header *types.Header, _state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header) {

	this.state = _state
	blockReward := BlockReward

	signer, _ := ecrecover(header, this.signatures)

	if signer == (common.Address{}) {
		signer = this.signer
	}

	// assign reward to signer
	toSigner := new(big.Int).Set(blockReward)
	toSigner.Mul(toSigner, big.NewInt(signerReward))
	toSigner.Div(toSigner, big.NewInt(100))

	_state.AddBalance(signer, toSigner)

	// assign reward to delegator
	toDelegators := new(big.Int).Set(blockReward)
	toDelegators.Sub(toDelegators, toSigner)

	epochHeader := this.epochOfHeader(chain, header, nil)

	signers, _, delegatorss := parseEpochExtra(epochHeader)

	electedDelegators := make(map[common.Address][]ElectedDelegator)
	for k, delegators := range delegatorss {
		for _, delegator := range delegators {
			electedDelegators[signers[k]] = append(electedDelegators[signers[k]], delegator)
		}
	}

	totalDelegators := len(electedDelegators[signer])

	if totalDelegators > 0 {

		for _, delegator := range electedDelegators[signer] {

			portionAmt := new(big.Float).Mul(new(big.Float).SetInt(toDelegators), big.NewFloat(float64(delegator.Portion)))

			delegatorReward := new(big.Int)
			portionAmt.Int(delegatorReward)

			_state.AddBalance(delegator.Delegator, delegatorReward)
		}
	}
	// update word state tire root
	header.Root = _state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	header.UncleHash = types.CalcUncleHash(nil)
}

func (this *Dpos) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, _state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {

	this.state = _state

	this.Finalize(chain, header, _state, txs, uncles)

	number := header.Number.Uint64()

	snap, err := this.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return nil, err
	}

	if number%this.config.EpochInterval != 0 {
		this.lock.RLock()

		validProposals := make([]common.Hash, 0, len(this.proposals))
		for proposalBytes, yesNo := range this.proposals {
			if snap.validVote(this.signer, proposalBytes, yesNo) {
				validProposals = append(validProposals, proposalBytes)
			}
		}

		if len(validProposals) > 0 {

			r := rand.Intn(len(validProposals))

			for i, proposalBytes := range validProposals {
				if r == i {
					header.MixDigest = proposalBytes

					yesNoVote := this.proposals[proposalBytes]
					if yesNoVote {
						copy(header.Nonce[:], nonceYesVote)
					} else {
						copy(header.Nonce[:], nonceNoVote)
					}
					break
				}
			}
		}
		this.lock.RUnlock()
	}

	header.Difficulty = calcDifficulty(snap, this.signer)

	header.Extra = make([]byte, 0)

	item := bytes.Repeat([]byte{0x00}, crypto.SignatureLength)

	header.Extra = append(header.Extra, VarIntToBytes(item)...)
	header.Extra = append(header.Extra, item...)

	if number%this.config.EpochInterval == 0 {

		item = make([]byte, 0)

		// append signers info
		for _, signer := range snap.preElectedSigners() {
			item = append(item, signer[:]...)
		}

		header.Extra = append(header.Extra, VarIntToBytes(item)...)
		header.Extra = append(header.Extra, item...)

		item = make([]byte, 0)

		// append proposal info
		for _, proposalBytes := range snap.unconfirmedProposals() {
			item = append(item, proposalBytes.Bytes()...)
		}

		header.Extra = append(header.Extra, VarIntToBytes(item)...)
		header.Extra = append(header.Extra, item...)

		item = make([]byte, 0)

		// append delegator info
		for _, signer := range snap.preElectedSigners() {
			delegators := snap.PreElectedDelegators[signer]

			subitem := []byte{}
			for _, delegator := range delegators {
				subitem = append(subitem, delegator.Delegator.Bytes()...)
				subitem = append(subitem, common.FromHex(fmt.Sprintf("%x", math.Float32bits(delegator.Portion)))...)
			}

			item = append(item, VarIntToBytes(subitem)...)
			item = append(item, subitem...)
		}

		header.Extra = append(header.Extra, VarIntToBytes(item)...)
		header.Extra = append(header.Extra, item...)

	}

	return types.NewBlock(header, txs, nil, receipts, new(trie.Trie)), nil
}

func (this *Dpos) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {

	header := block.Header()

	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	this.lock.RLock()
	signer, signFn := this.signer, this.signFn
	this.lock.RUnlock()

	snap, err := this.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	if _, authorized := snap.ElectedSigners[signer]; !authorized {
		return errUnauthorizedSignerAgainstSnap
	}

	for seen, recent := range snap.Recents {
		if recent == signer {
			if limit := uint64(len(snap.ElectedSigners)/2 + 1); number < limit || seen > number-limit {
				log.Info("Signed recently, must wait for others")
				return nil
			}
		}
	}

	delay := time.Unix(int64(header.Time), 0).Sub(time.Now())
	if header.Difficulty.Cmp(diffNoTurn) == 0 {

		wiggle := time.Duration(len(snap.ElectedSigners)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))

		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}

	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeDpos, RLP(header))

	if err != nil {
		return err
	}

	extras := unserialize(header.Extra)
	header.Extra = make([]byte, 0)

	copy(extras[0][:], sighash)
	for _, extra := range extras {
		header.Extra = append(header.Extra, VarIntToBytes(extra)...)
		header.Extra = append(header.Extra, extra...)
	}

	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))

	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

func (this *Dpos) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

func (this *Dpos) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	snap, err := this.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return calcDifficulty(snap, this.signer)
}

func calcDifficulty(snap *Snapshot, signer common.Address) *big.Int {
	if snap.inturn(snap.Number+1, signer) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

func (this *Dpos) APIs(chain consensus.ChainHeaderReader) []rpc.API {

	return []rpc.API{{
		Namespace: "dpos",
		Version:   "1.0",
		Service:   &API{chain: chain, dpos: this},
		Public:    false,
	}}

}

func (this *Dpos) Close() error {
	return nil
}

func (this *Dpos) epochOfHeader(chain consensus.ChainHeaderReader, header *types.Header, _parents []*types.Header) *types.Header {

	var parents []*types.Header
	copy(parents, _parents)

	number := header.Number.Uint64()

	var epochNumber uint64
	if number%this.config.EpochInterval == 0 {
		epochNumber = number - this.config.EpochInterval
	} else {
		epochNumber = number - (number % this.config.EpochInterval)
	}

	searchNumber := number - 1
	searchHash := header.ParentHash

	for searchNumber != epochNumber {

		var header *types.Header

		if len(parents) > 0 {
			header = parents[len(parents)-1]
			parents = parents[:len(parents)-1]
		}

		if header == nil {
			header = chain.GetHeader(searchHash, searchNumber)
		}

		if header == nil {
			return nil
		}
		searchNumber, searchHash = searchNumber-1, header.ParentHash
	}

	if len(parents) > 0 {
		header = parents[len(parents)-1]
		return header
	} else {
		return chain.GetHeaderByHash(searchHash)
	}
}

func (this *Dpos) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {

	var (
		headers []*types.Header
		snap    *Snapshot
	)

	for snap == nil {
		if s, ok := this.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}

		if number == 0 {
			thisHeader := chain.GetHeaderByNumber(number)
			if thisHeader != nil {
				signers, proposals, delegatorss := parseEpochExtra(thisHeader)
				snap = newSnapshot(this.config, this.signatures, number, hash, signers, proposals, delegatorss)
				if err := snap.store(this.db); err != nil {
					return nil, err
				}
				log.Info("Stored epoch snapshot to disk", "number", number, "hash", hash)
				break
			}
		}

		if number%storeSnapInterval == 0 || (number+1)%this.config.EpochInterval == 0 {
			if s, err := loadSnapshot(this.config, this.signatures, this.db, hash); err == nil {
				log.Trace("Loaded voting snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}

		var header *types.Header

		if len(parents) > 0 {
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}

		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}

	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	snap, err := snap.apply(chain.(consensus.ChainReader), headers, this.db, this.state)
	if err != nil {

		return nil, err
	}

	this.recents.Add(snap.Hash, snap)

	if snap.Number%storeSnapInterval == 0 && len(headers) > 0 {
		if err = snap.store(this.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}

	return snap, err
}
