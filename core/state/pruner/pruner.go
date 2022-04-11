// Copyright 2020 The go-ethereum Authors
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

package pruner

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/trie"
)

const (
	// stateBloomFilePrefix is the filename prefix of state bloom filter.
	stateBloomFilePrefix = "statebloom"

	// stateBloomFilePrefix is the filename suffix of state bloom filter.
	stateBloomFileSuffix = "bf.gz"

	// stateBloomFileTempSuffix is the filename suffix of state bloom filter
	// while it is being written out to detect write aborts.
	stateBloomFileTempSuffix = ".tmp"

	// rangeCompactionThreshold is the minimal deleted entry number for
	// triggering range compaction. It's a quite arbitrary number but just
	// to avoid triggering range compaction because of small deletion.
	rangeCompactionThreshold = 100000
)

var (
	// emptyRoot is the known root hash of an empty trie.
	emptyRoot = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")

	// emptyCode is the known hash of the empty EVM bytecode.
	emptyCode = crypto.Keccak256(nil)
)

// Pruner is an offline tool to prune the stale state with the
// help of the snapshot. The workflow of pruner is very simple:
//
// - iterate the snapshot, reconstruct the relevant state
// - iterate the database, delete all other state entries which
//   don't belong to the target state and the genesis state
//
// It can take several hours(around 2 hours for mainnet) to finish
// the whole pruning work. It's recommended to run this offline tool
// periodically in order to release the disk usage and improve the
// disk read performance to some extent.
type Pruner struct {
	db            ethdb.Database
	stateBloom    *stateBloom
	datadir       string
	trieCachePath string
	headHeader    *types.Header
	snaptree      *snapshot.Tree
	triesInMemory uint64
}

type BlockPruner struct {
	db                  ethdb.Database
	oldAncientPath      string
	newAncientPath      string
	node                *node.Node
	BlockAmountReserved uint64
}

// NewPruner creates the pruner instance.
func NewPruner(db ethdb.Database, datadir, trieCachePath string, bloomSize, triesInMemory uint64) (*Pruner, error) {
	headBlock := rawdb.ReadHeadBlock(db)
	if headBlock == nil {
		return nil, errors.New("Failed to load head block")
	}
	snaptree := snapshot.New(db, trie.NewDatabase(db), 256, headBlock.Root(), false, false)
	// Sanitize the bloom filter size if it's too small.
	if bloomSize < 256 {
		log.Warn("Sanitizing bloomfilter size", "provided(MB)", bloomSize, "updated(MB)", 256)
		bloomSize = 256
	}
	stateBloom, err := newStateBloomWithSize(bloomSize)

	if err != nil {
		return nil, err
	}
	return &Pruner{
		db:            db,
		stateBloom:    stateBloom,
		datadir:       datadir,
		trieCachePath: trieCachePath,
		triesInMemory: triesInMemory,
		headHeader:    headBlock.Header(),
		snaptree:      snaptree,
	}, nil
}

func prune(snaptree *snapshot.Tree, root common.Hash, maindb ethdb.Database, stateBloom *stateBloom, bloomPath string, middleStateRoots map[common.Hash]struct{}, start time.Time) error {
	// Delete all stale trie nodes in the disk. With the help of state bloom
	// the trie nodes(and codes) belong to the active state will be filtered
	// out. A very small part of stale tries will also be filtered because of
	// the false-positive rate of bloom filter. But the assumption is held here
	// that the false-positive is low enough(~0.05%). The probablity of the
	// dangling node is the state root is super low. So the dangling nodes in
	// theory will never ever be visited again.
	var (
		count  int
		size   common.StorageSize
		pstart = time.Now()
		logged = time.Now()
		batch  = maindb.NewBatch()
		iter   = maindb.NewIterator(nil, nil)
	)
	for iter.Next() {
		key := iter.Key()

		// All state entries don't belong to specific state and genesis are deleted here
		// - trie node
		// - legacy contract code
		// - new-scheme contract code
		isCode, codeKey := rawdb.IsCodeKey(key)
		if len(key) == common.HashLength || isCode {
			checkKey := key
			if isCode {
				checkKey = codeKey
			}
			if _, exist := middleStateRoots[common.BytesToHash(checkKey)]; exist {
				log.Debug("Forcibly delete the middle state roots", "hash", common.BytesToHash(checkKey))
			} else {
				if ok, err := stateBloom.Contain(checkKey); err != nil {
					return err
				} else if ok {
					continue
				}
			}
			count += 1
			size += common.StorageSize(len(key) + len(iter.Value()))
			batch.Delete(key)

			var eta time.Duration // Realistically will never remain uninited
			if done := binary.BigEndian.Uint64(key[:8]); done > 0 {
				var (
					left  = math.MaxUint64 - binary.BigEndian.Uint64(key[:8])
					speed = done/uint64(time.Since(pstart)/time.Millisecond+1) + 1 // +1s to avoid division by zero
				)
				eta = time.Duration(left/speed) * time.Millisecond
			}
			if time.Since(logged) > 8*time.Second {
				log.Info("Pruning state data", "nodes", count, "size", size,
					"elapsed", common.PrettyDuration(time.Since(pstart)), "eta", common.PrettyDuration(eta))
				logged = time.Now()
			}
			// Recreate the iterator after every batch commit in order
			// to allow the underlying compactor to delete the entries.
			if batch.ValueSize() >= ethdb.IdealBatchSize {
				batch.Write()
				batch.Reset()

				iter.Release()
				iter = maindb.NewIterator(nil, key)
			}
		}
	}
	if batch.ValueSize() > 0 {
		batch.Write()
		batch.Reset()
	}
	iter.Release()
	log.Info("Pruned state data", "nodes", count, "size", size, "elapsed", common.PrettyDuration(time.Since(pstart)))

	// Pruning is done, now drop the "useless" layers from the snapshot.
	// Firstly, flushing the target layer into the disk. After that all
	// diff layers below the target will all be merged into the disk.
	if root != snaptree.DiskRoot() {
		if err := snaptree.Cap(root, 0); err != nil {
			return err
		}
	}
	// Secondly, flushing the snapshot journal into the disk. All diff
	// layers upon are dropped silently. Eventually the entire snapshot
	// tree is converted into a single disk layer with the pruning target
	// as the root.
	if _, err := snaptree.Journal(root); err != nil {
		return err
	}
	// Delete the state bloom, it marks the entire pruning procedure is
	// finished. If any crashes or manual exit happens before this,
	// `RecoverPruning` will pick it up in the next restarts to redo all
	// the things.
	os.RemoveAll(bloomPath)

	// Start compactions, will remove the deleted data from the disk immediately.
	// Note for small pruning, the compaction is skipped.
	if count >= rangeCompactionThreshold {
		cstart := time.Now()
		for b := 0x00; b <= 0xf0; b += 0x10 {
			var (
				start = []byte{byte(b)}
				end   = []byte{byte(b + 0x10)}
			)
			if b == 0xf0 {
				end = nil
			}
			log.Info("Compacting database", "range", fmt.Sprintf("%#x-%#x", start, end), "elapsed", common.PrettyDuration(time.Since(cstart)))
			if err := maindb.Compact(start, end); err != nil {
				log.Error("Database compaction failed", "error", err)
				return err
			}
		}
		log.Info("Database compaction finished", "elapsed", common.PrettyDuration(time.Since(cstart)))
	}
	log.Info("State pruning successful", "pruned", size, "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}
