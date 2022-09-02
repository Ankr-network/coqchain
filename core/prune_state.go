package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/log"
	"github.com/Ankr-network/coqchain/rlp"
	"github.com/Ankr-network/coqchain/trie"
	"github.com/Ankr-network/coqchain/utils"
	"github.com/sunvim/utils/tools"
	"github.com/sunvim/utils/workpool"
	"go.etcd.io/bbolt"
	"gopkg.in/urfave/cli.v1"
)

const (
	stateName             = "prune_state"
	stateBucketName       = "state"
	stateRecentPruneBlock = "recent_prune_block"
	defaultWorkSize       = 2
	pruneSize             = 1024
)

type ReadFunc func(key []byte) ([]byte, error)
type DelFunc func(key []byte) error

type BlockState struct {
	Height uint64
	Root   common.Hash
	db     *trie.Database
}

type PruneState struct {
	lock       sync.Mutex
	size       uint64
	blkCh      chan *BlockState
	pruneBlock uint64
	worker     *workpool.WorkPool
	db         *bbolt.DB
}

func NewPruneState(ctx *cli.Context) *PruneState {

	size := ctx.GlobalUint64("prune.size")
	if size >= pruneSize && size%pruneSize == 0 {
		name := fmt.Sprintf("%s/%s", ctx.GlobalString("datadir"), stateName)
		db, err := bbolt.Open(name, 0644, nil)
		if err != nil {
			panic(err)
		}
		db.Update(func(tx *bbolt.Tx) error {
			tx.CreateBucketIfNotExists(tools.StringToBytes(stateBucketName))
			return nil
		})
		return &PruneState{
			db:         db,
			pruneBlock: 1,
			blkCh:      make(chan *BlockState, 64),
			size:       size,
			worker:     workpool.New(defaultWorkSize),
		}
	} else {
		panic("prune.size should be muliple of 1024")
	}
}

func InitPruneWorker(ctx *cli.Context) {
	log.Info("pruner initing ...")
	dp = NewPruneState(ctx)
	log.Info("pruner init over.")
}

func SetLatestBlock(num uint64, hash common.Hash, db *trie.Database) {
	dp.blkCh <- &BlockState{Height: num, Root: hash, db: db}
}

func IsRunning() bool {
	return pruneState
}

func Run() {
	go func() {
		defer func() {
			dp.Close()
		}()

		pruneState = true
		var (
			acc *types.StateAccount
			hs  = make([]byte, 8)
		)

		for v := range dp.blkCh {

			binary.BigEndian.PutUint64(hs, v.Height)
			dp.Put(hs, v.Root[:])

			if v.Height-dp.pruneBlock > dp.size {
				dp.pruneBlock = v.Height

				dp.worker.Do(func() error {
					log.Info("prune", "height", v.Height)
					tr, err := trie.New(v.Root, v.db)
					if err != nil {
						log.Error("init trie", "err", err)
						return nil
					}

					it := trie.NewIterator(tr.NodeIterator(nil))
					for it.Next() {
						rlp.Decode(bytes.NewReader(it.Value), &acc)
						log.Info("interator", "height", v.Height,
							"key", common.Bytes2Hex(it.Key), "val.none", acc.Nonce,
							"val.balance", acc.Balance,
							"val.root", acc.Root,
							"val.codeHash", common.Bytes2Hex(acc.CodeHash))
					}

					return nil
				})
			}
		}
	}()
}

var (
	dp         *PruneState
	pruneState bool
	trpool     = &sync.Pool{
		New: func() interface{} {
			return &trie.Trie{}
		},
	}
)

func (p *PruneState) Close() error {
	return p.db.Close()
}

func (p *PruneState) Put(key, val []byte) error {
	return p.db.Update(func(tx *bbolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists(utils.S2B(stateBucketName))
		return b.Put(key, val)
	})
}

func (p *PruneState) Remove(stx, end []byte) error {
	nstx := binary.BigEndian.Uint64(stx)
	nend := binary.BigEndian.Uint64(end)
	p.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(stateBucketName))
		b.ForEach(func(k, v []byte) error {
			nk := binary.BigEndian.Uint64(k)
			if nk >= nstx && nk < nend {
				b.Delete(k)
			}
			return nil
		})

		return nil
	})
	return nil
}

func (p *PruneState) Range(stx, end []byte) ([][]byte, error) {
	var rs [][]byte
	nstx := binary.BigEndian.Uint64(stx)
	nend := binary.BigEndian.Uint64(end)
	rs = make([][]byte, 0, nend-nstx)

	p.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(stateBucketName))
		b.ForEach(func(k, v []byte) error {
			nk := binary.BigEndian.Uint64(k)
			if nk >= nstx && nk < nend {
				rs = append(rs, v)
			}
			return nil
		})

		return nil
	})
	return rs, nil
}
