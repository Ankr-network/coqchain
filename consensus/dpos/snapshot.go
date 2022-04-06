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

package dpos

import (
	"bytes"
	"encoding/json"
	_ "errors"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	lru "github.com/hashicorp/golang-lru"
)

const (
	dbSnapPrefix             = "dpos-"
	maxCandidateAndDelegator = 51
)

type Vote struct {
	Signer   common.Address `json:"signer"` // Authorized signer that cast this vote
	Block    uint64         `json:"block"`  // Block number the vote was cast in (expire old votes)
	YesNo    bool           `json:"yesno"`
	Proposal common.Hash    `json:"proposal"` // Proposal bytes
}

type ElectedDelegator struct {
	Delegator common.Address `json:"delegator"`
	Portion   float32        `json:"portion"`
}

// Snapshot is the state of the authorization voting at a given point in time.
type Snapshot struct {
	config   *params.DposConfig // Consensus engine parameters to fine tune behavior
	sigcache *lru.ARCCache

	Number uint64      `json:"number"`
	Hash   common.Hash `json:"hash"`

	ElectedSigners    map[common.Address]uint16   `json:"elected_signers"`
	PreElectedSigners map[common.Address]struct{} `json:"pre_elected_signers"`

	ElectedDelegators    map[common.Address][]ElectedDelegator `json:"elected_delegators"`
	PreElectedDelegators map[common.Address][]ElectedDelegator `json:"pre_elected_delegators"`

	ConfirmedProposals   map[uint8]common.Hash `json:"proposals"`
	UnconfirmedProposals map[uint8]common.Hash `json:"unconfirmed_proposals"`

	Candidates map[common.Address]struct{}       `json:"candidates"`
	Delegators map[common.Address]common.Address `json:"delegators"`

	Recents map[uint64]common.Address `json:"recents"`
	Votes   []*Vote                   `json:"votes"`
	Tally   map[common.Hash]int       `json:"tally"`
}

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent signers, so only ever use if for
// the genesis block.
func newSnapshot(config *params.DposConfig, sigcache *lru.ARCCache, number uint64, hash common.Hash, signers []common.Address, proposals []*Proposal, delegatorss [][]ElectedDelegator) *Snapshot {

	snap := &Snapshot{
		config:   config,
		sigcache: sigcache,
		Number:   number,
		Hash:     hash,

		ElectedSigners:    make(map[common.Address]uint16),
		PreElectedSigners: make(map[common.Address]struct{}),

		ElectedDelegators:    make(map[common.Address][]ElectedDelegator),
		PreElectedDelegators: make(map[common.Address][]ElectedDelegator),

		ConfirmedProposals:   make(map[uint8]common.Hash),
		UnconfirmedProposals: make(map[uint8]common.Hash),

		Candidates: make(map[common.Address]struct{}),
		Delegators: make(map[common.Address]common.Address),

		Recents: make(map[uint64]common.Address),
		Tally:   make(map[common.Hash]int),
	}

	for _, signer := range signers {
		snap.ElectedSigners[signer] = uint16(0)
		snap.ElectedDelegators[signer] = []ElectedDelegator{}

		if number == uint64(0) {
			snap.Candidates[signer] = struct{}{}
		}
	}

	for _, proposal := range proposals {

		hash, err := proposal.toBytes()
		_ = err
		snap.ConfirmedProposals[proposal.Id] = hash
	}

	for k, delegators := range delegatorss {
		for _, delegator := range delegators {
			snap.ElectedDelegators[signers[k]] = append(snap.ElectedDelegators[signers[k]], delegator)

			if number == uint64(0) {
				snap.Delegators[delegator.Delegator] = signers[k]
			}
		}
	}

	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(config *params.DposConfig, sigcache *lru.ARCCache, db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	blob, err := db.Get(append([]byte(dbSnapPrefix), hash[:]...))
	if err != nil {
		return nil, err
	}

	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}

	snap.config = config
	snap.sigcache = sigcache

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte(dbSnapPrefix), s.Hash[:]...), blob)
}

// copy creates a deep copy of the snapshot, though not the individual votes.
func (s *Snapshot) copy() *Snapshot {

	cpy := &Snapshot{
		config:   s.config,
		sigcache: s.sigcache,
		Number:   s.Number,
		Hash:     s.Hash,

		ElectedSigners:    make(map[common.Address]uint16),
		PreElectedSigners: make(map[common.Address]struct{}),

		ElectedDelegators:    make(map[common.Address][]ElectedDelegator),
		PreElectedDelegators: make(map[common.Address][]ElectedDelegator),

		ConfirmedProposals:   make(map[uint8]common.Hash),
		UnconfirmedProposals: make(map[uint8]common.Hash),

		Candidates: make(map[common.Address]struct{}),
		Delegators: make(map[common.Address]common.Address),

		Recents: make(map[uint64]common.Address),
		Votes:   make([]*Vote, len(s.Votes)),
		Tally:   make(map[common.Hash]int),
	}

	for signer, mintCnt := range s.ElectedSigners {
		cpy.ElectedSigners[signer] = mintCnt
	}

	for signer := range s.PreElectedSigners {
		cpy.PreElectedSigners[signer] = struct{}{}
	}

	for signer, electedDelegators := range s.ElectedDelegators {
		if _, exist := cpy.ElectedDelegators[signer]; !exist {
			cpy.ElectedDelegators[signer] = make([]ElectedDelegator, 0)
		}

		for _, electedDelegator := range electedDelegators {
			cpy.ElectedDelegators[signer] = append(cpy.ElectedDelegators[signer], electedDelegator)
		}
	}

	for signer, electedDelegators := range s.PreElectedDelegators {
		if _, exist := cpy.PreElectedDelegators[signer]; !exist {
			cpy.PreElectedDelegators[signer] = make([]ElectedDelegator, 0)
		}

		for _, electedDelegator := range electedDelegators {
			cpy.PreElectedDelegators[signer] = append(cpy.PreElectedDelegators[signer], electedDelegator)
		}
	}

	for candidate := range s.Candidates {
		cpy.Candidates[candidate] = struct{}{}
	}

	for delegator, signer := range s.Delegators {
		cpy.Delegators[delegator] = signer
	}

	for proposalId, proposalBytes := range s.ConfirmedProposals {
		cpy.ConfirmedProposals[proposalId] = proposalBytes
	}

	for proposalId, proposalBytes := range s.UnconfirmedProposals {
		cpy.UnconfirmedProposals[proposalId] = proposalBytes
	}

	for block, signer := range s.Recents {
		cpy.Recents[block] = signer
	}

	for proposalBytes, votes := range s.Tally {
		cpy.Tally[proposalBytes] = votes
	}
	copy(cpy.Votes, s.Votes)

	return cpy
}

func (s *Snapshot) lastVote(signer common.Address, proposalBytes common.Hash) *Vote {

	var lastVote *Vote

	proposal := &Proposal{}
	if err := proposal.fromBytes(proposalBytes); err == nil {
		for _, vote := range s.Votes {
			if vote.Signer == signer && vote.Proposal == proposalBytes {
				lastVote = vote
			}
		}
	}

	return lastVote
}

func (s *Snapshot) validVote(signer common.Address, proposalBytes common.Hash, yesNo bool) bool {

	if _, exist := s.ElectedSigners[signer]; !exist {
		return false
	}

	var lastVote *Vote = s.lastVote(signer, proposalBytes)

	if lastVote != nil {
		return (!lastVote.YesNo && yesNo) || (lastVote.YesNo && !yesNo)
	} else {
		return yesNo
	}
}

func (s *Snapshot) cast(signer common.Address, proposalBytes common.Hash, yesNo bool) bool {

	proposal := &Proposal{}
	if err := proposal.fromBytes(proposalBytes); err != nil {
		return false
	}

	_, exist := getProposal(proposal.Id)
	if exist != nil {
		return false
	}

	if !s.validVote(signer, proposalBytes, yesNo) {
		return false
	}

	if oldVotes, ok := s.Tally[proposalBytes]; ok {
		if yesNo {
			oldVotes++
		} else {
			oldVotes--
		}

		if oldVotes <= 0 {
			delete(s.Tally, proposalBytes)
		} else {
			s.Tally[proposalBytes] = oldVotes
		}
	} else {
		s.Tally[proposalBytes] = 1
	}
	return true
}

func (s *Snapshot) uncast(signer common.Address, proposalBytes common.Hash) bool {

	proposal := &Proposal{}
	if err := proposal.fromBytes(proposalBytes); err != nil {
		return false
	}

	_, exist := getProposal(proposal.Id)
	if exist != nil {
		return false
	}

	votes, ok := s.Tally[proposalBytes]
	if !ok {
		return false
	}

	var lastVote *Vote = s.lastVote(signer, proposalBytes)

	if !lastVote.YesNo {
		return false
	}

	if votes > 1 {
		votes--
		s.Tally[proposalBytes] = votes
	} else {
		delete(s.Tally, proposalBytes)
	}

	return true
}

// signersAscending implements the sort interface to allow sorting a list of addresses
type signersAscending []common.Address

func (s signersAscending) Len() int           { return len(s) }
func (s signersAscending) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }
func (s signersAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type hashIntDesc struct {
	Key   common.Hash
	Value int
}

type manyHashIntDesc []hashIntDesc

func (p manyHashIntDesc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p manyHashIntDesc) Len() int           { return len(p) }
func (p manyHashIntDesc) Less(i, j int) bool { return p[i].Value > p[j].Value }

func hashIntDescSorter(m map[common.Hash]int) manyHashIntDesc {
	p := make(manyHashIntDesc, len(m))
	i := 0
	for k, v := range m {
		p[i] = hashIntDesc{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

type addressIntAsc struct {
	Key   common.Address
	Value int
}

type manyAddressIntAsc []addressIntAsc

func (p manyAddressIntAsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p manyAddressIntAsc) Len() int           { return len(p) }
func (p manyAddressIntAsc) Less(i, j int) bool { return p[i].Value < p[j].Value }

func addressIntAscSorter(m map[common.Address]uint16) manyAddressIntAsc {
	p := make(manyAddressIntAsc, len(m))
	i := 0
	for k, v := range m {
		p[i] = addressIntAsc{k, int(v)}
		i++
	}
	sort.Sort(p)
	return p
}

type addressBigIntDesc struct {
	Key   common.Address
	Value *big.Int
}

type manyAddressBigIntDesc []addressBigIntDesc

func (p manyAddressBigIntDesc) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p manyAddressBigIntDesc) Len() int      { return len(p) }
func (p manyAddressBigIntDesc) Less(i, j int) bool {
	if p[i].Value.Cmp(p[j].Value) < 0 {
		return false
	} else if p[i].Value.Cmp(p[j].Value) > 0 {
		return true
	} else {
		return p[i].Key.String() < p[j].Key.String()
	}
}

func addressBigIntDescSorter(m map[common.Address]*big.Int) manyAddressBigIntDesc {
	p := make(manyAddressBigIntDesc, len(m))
	i := 0
	for k, v := range m {
		p[i] = addressBigIntDesc{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot) apply(chain consensus.ChainReader, headers []*types.Header, db ethdb.Database, _state *state.StateDB) (*Snapshot, error) {

	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}

	// Sanity check that the headers can be applied
	for i := 0; i < len(headers)-1; i++ {
		if headers[i+1].Number.Uint64() != headers[i].Number.Uint64()+1 {
			return nil, errInvalidVotingChain
		}
	}

	if headers[0].Number.Uint64() != s.Number+1 {
		return nil, errInvalidVotingChain
	}

	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	var (
		start  = time.Now()
		logged = time.Now()
	)

	for i, header := range headers {

		snap.Number += 1
		snap.Hash = header.Hash()

		number := header.Number.Uint64()

		if number%s.config.EpochInterval == 0 {
			for kickoutSigner := range snap.ElectedSigners {
				_, exist := snap.PreElectedSigners[kickoutSigner]

				if !exist {
					delete(snap.Candidates, kickoutSigner)

					delete(snap.Delegators, kickoutSigner)

					for delegator, candidate := range snap.Delegators {
						if candidate == kickoutSigner {
							delete(snap.Delegators, delegator)
						}
					}
				}
			}

			for k, v := range snap.UnconfirmedProposals {
				snap.ConfirmedProposals[k] = v
			}

			snap.ElectedSigners = make(map[common.Address]uint16)
			for k := range snap.PreElectedSigners {
				snap.ElectedSigners[k] = 0
			}

			snap.ElectedDelegators = make(map[common.Address][]ElectedDelegator)
			for k, v := range snap.PreElectedDelegators {
				snap.ElectedDelegators[k] = v
			}

			snap.PreElectedDelegators = make(map[common.Address][]ElectedDelegator)
			snap.PreElectedSigners = make(map[common.Address]struct{})
			snap.UnconfirmedProposals = make(map[uint8]common.Hash)

			snap.Votes = nil
			snap.Tally = make(map[common.Hash]int)

		}

		if limit := uint64(len(snap.ElectedSigners)/2 + 1); number >= limit {
			delete(snap.Recents, number-limit)
		}

		signer, err := ecrecover(header, s.sigcache)

		if err != nil {
			return nil, err
		}

		if _, ok := snap.ElectedSigners[signer]; !ok {
			return nil, errUnauthorizedSignerAgainstSnap
		} else {
			snap.ElectedSigners[signer]++
		}

		for _, recent := range snap.Recents {
			if recent == signer {
				return nil, errRecentlySigned
			}
		}

		snap.Recents[number] = signer

		var yesNo bool
		switch {
		case bytes.Equal(header.Nonce[:], nonceYesVote):
			yesNo = true
		case bytes.Equal(header.Nonce[:], nonceNoVote):
			yesNo = false
		default:
			return nil, errInvalidVote
		}

		if snap.cast(signer, header.MixDigest, yesNo) {
			snap.Votes = append(snap.Votes, &Vote{
				Signer:   signer,
				Block:    number,
				Proposal: header.MixDigest,
				YesNo:    yesNo,
			})
		}

		block := chain.GetBlock(header.Hash(), number)
		if block == nil {
			// light node must happen, so dicide to remove light node support
			return nil, errMissingBody
		} else {

			txs := block.Body().Transactions

			ethSigner := types.MakeSigner(chain.Config(), new(big.Int).SetUint64(number))

			for i := 0; i < len(txs); i++ {
				tx := txs[i]

				if *tx.To() == contractAddress {
					action := &Action{}
					if err := action.fromBytes(tx.Data()); err == nil {

						if from, err := ethSigner.Sender(tx); err == nil {

							switch action.Id {
							case becomeCandidate:

								if len(snap.Candidates) < maxCandidateAndDelegator {
									snap.Candidates[from] = struct{}{}
								}

							case becomeDelegator:

								candidate := action.Values[0].(common.Address)

								_, exist := snap.Candidates[candidate]

								if exist {
									//  I know it's stupid and bug obviously, but for performance and snap size
									if len(snap.Candidates) < maxCandidateAndDelegator {
										snap.Delegators[from] = candidate
									}
								}

							case quitCandidate:
								delete(snap.Candidates, from)

								for delegator, candidate := range snap.Delegators {
									if candidate == from {
										delete(snap.Delegators, delegator)
									}
								}
							case quitDelegator:
								delete(snap.Delegators, from)
							}
						}
					}
				}
			}
		}

		if (number+1)%s.config.EpochInterval == 0 {

			statedb, err := state.New(header.Root, state.NewDatabase(db), nil)

			if statedb == nil {
				statedb = _state
			}

			if statedb == nil {

				log.Trace("chiew check apply", err)

				if len(headers)-1 < i+1 {

					return nil, errMissingEpochBlock

				} else {

					electedSigners, _, delegatorss := parseEpochExtra(headers[i+1])
					electedDelegators := make(map[common.Address][]ElectedDelegator)
					for k, delegators := range delegatorss {
						for _, delegator := range delegators {
							electedDelegators[electedSigners[k]] = append(electedDelegators[electedSigners[k]], delegator)
						}
					}

					for _, signer := range electedSigners {
						snap.PreElectedSigners[signer] = struct{}{}
					}
					snap.PreElectedDelegators = electedDelegators
				}
			} else {

				minMintTarget := (int(snap.config.EpochInterval) - 1) / len(snap.ElectedSigners) / 2
				candidateCnt := len(snap.Candidates) - len(snap.ElectedSigners)
				sorted := addressIntAscSorter(snap.ElectedSigners)

				kickoutSigners := make(map[common.Address]struct{}, 0)
				candidateVotes := make(map[common.Address]*big.Int)

				for _, kv := range sorted {

					address := kv.Key
					mintCnt := kv.Value

					if len(kickoutSigners) < candidateCnt && mintCnt < minMintTarget {
						kickoutSigners[address] = struct{}{}
					} else {
						candidateVotes[address] = big.NewInt(0)
					}
				}

				for delegator, candidate := range snap.Delegators {

					if _, exist := kickoutSigners[candidate]; exist {
						continue
					}

					if _, exist := kickoutSigners[delegator]; exist {
						continue
					}

					if _, exist := candidateVotes[candidate]; !exist {
						candidateVotes[candidate] = big.NewInt(0)
					}

					balance := statedb.GetBalance(delegator)

					if balance.Cmp(common.Big0) > 0 {
						candidateVotes[candidate].Add(candidateVotes[candidate], balance)
					}
				}

				newSigners := addressBigIntDescSorter(candidateVotes)

				if len(newSigners) > maxSignerSize {
					newSigners = newSigners[:maxSignerSize]
				}

				for _, newSigner := range newSigners {

					preElectedSigner := newSigner.Key
					snap.PreElectedSigners[preElectedSigner] = struct{}{}

					snap.PreElectedDelegators[preElectedSigner] = []ElectedDelegator{}

					delegators := make(map[common.Address]*big.Int)
					sum := new(big.Int)
					for delegator, candidate := range snap.Delegators {
						if candidate == preElectedSigner {
							delegators[delegator] = statedb.GetBalance(delegator)
							sum.Add(sum, delegators[delegator])
						}
					}

					sortedDelegators := addressBigIntDescSorter(delegators)
					operand2 := new(big.Float).SetInt(sum)
					for i := 0; i < len(sortedDelegators); i++ {
						address := sortedDelegators[i].Key
						operand1 := new(big.Float).SetInt(sortedDelegators[i].Value)
						result := new(big.Float).Quo(operand1, operand2)

						portion, _ := result.Float32()

						snap.PreElectedDelegators[preElectedSigner] = append(snap.PreElectedDelegators[preElectedSigner], ElectedDelegator{address, portion})
					}
				}
			}

			groupProposals := make(map[uint8]map[common.Hash]int)
			for proposalBytes, votes := range snap.Tally {

				proposal := &Proposal{}
				if err := proposal.fromBytes(proposalBytes); err != nil {
					return nil, err
				}

				_, exist := groupProposals[proposal.Id]

				if !exist {
					groupProposals[proposal.Id] = make(map[common.Hash]int)
				}
				groupProposals[proposal.Id][proposalBytes] = votes

			}

			selectedProposals := make(map[uint8]common.Hash)
			for proposalId, proposalVotes := range groupProposals {

				if len(proposalVotes) > 1 {
					sorted := hashIntDescSorter(proposalVotes)

					if sorted[0].Value == sorted[1].Value {
						continue
					}

					selectedProposals[proposalId] = sorted[0].Key
				} else {
					for proposalBytes := range proposalVotes {
						selectedProposals[proposalId] = proposalBytes
					}
				}
			}

			for k, v := range snap.ConfirmedProposals {
				snap.UnconfirmedProposals[k] = v
			}

			for proposalId, proposalBytes := range selectedProposals {
				snap.UnconfirmedProposals[proposalId] = proposalBytes
			}

			if err := snap.store(db); err != nil {
				return nil, err
			}
		}

		if time.Since(logged) > 15*time.Second {
			log.Info("Reconstructing voting history", "processed", i, "total", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
			logged = time.Now()
		}
	}

	if time.Since(start) > 15*time.Second {
		log.Info("Reconstructed voting history", "processed", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
	}

	return snap, nil
}

func (s *Snapshot) preElectedSigners() []common.Address {
	signers := make([]common.Address, 0, len(s.PreElectedSigners))
	for signer := range s.PreElectedSigners {
		signers = append(signers, signer)
	}
	sort.Sort(signersAscending(signers))
	return signers
}

// signers retrieves the list of authorized signers in ascending order.
func (s *Snapshot) electedSigners() []common.Address {
	signers := make([]common.Address, 0, len(s.ElectedSigners))
	for signer := range s.ElectedSigners {
		signers = append(signers, signer)
	}
	sort.Sort(signersAscending(signers))
	return signers
}

func (s *Snapshot) unconfirmedProposals() []common.Hash {
	keys := make([]int, len(s.UnconfirmedProposals))

	ascSortedProposals := make([]common.Hash, 0, len(s.UnconfirmedProposals))

	i := int(0)
	for k := range s.UnconfirmedProposals {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)

	for _, k := range keys {
		ascSortedProposals = append(ascSortedProposals, s.UnconfirmedProposals[uint8(k)])
	}

	return ascSortedProposals
}

func (s *Snapshot) confirmedProposals() []common.Hash {

	keys := make([]int, len(s.ConfirmedProposals))

	sortedProposals := make([]common.Hash, 0, len(s.ConfirmedProposals))

	i := int(0)
	for k := range s.ConfirmedProposals {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)

	for _, k := range keys {
		sortedProposals = append(sortedProposals, s.ConfirmedProposals[uint8(k)])
	}

	return sortedProposals
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Snapshot) inturn(number uint64, signer common.Address) bool {
	signers, offset := s.electedSigners(), 0
	for offset < len(signers) && signers[offset] != signer {
		offset++
	}
	return (number % uint64(len(signers))) == uint64(offset)
}
