package core

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/log"
	"github.com/Ankr-network/coqchain/rlp"
	"github.com/Ankr-network/coqchain/trie"
	"github.com/sunvim/utils/tools"
	"github.com/sunvim/utils/workpool"
	"go.etcd.io/bbolt"
	"gopkg.in/urfave/cli.v1"
)

const (
	stateName             = "prune_state"
	stateBucketName       = "state"
	stateRecentPruneBlock = "recent_prune_block"
	defaultWorkSize       = 4
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
		db:     db,
		blkCh:  make(chan *BlockState, 128),
		size:   ctx.GlobalUint64("prune.size"),
		worker: workpool.New(defaultWorkSize),
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
		)

		for v := range dp.blkCh {
			dp.lock.Lock()
			if v.Height-dp.pruneBlock > dp.size {
				dp.pruneBlock++

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
			dp.lock.Unlock()
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
