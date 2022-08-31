package core

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Ankr-network/coqchain/common"
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
	Get    ReadFunc
	Del    DelFunc
}

type PruneState struct {
	lock       sync.Mutex
	running    bool
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

func SetLatestBlock(num uint64, hash common.Hash, get ReadFunc, del DelFunc) {
	dp.blkCh <- &BlockState{Height: num, Root: hash, Get: get, Del: del}
}

func IsRunning() bool {
	return dp.running
}

func Run() {
	go func() {
		defer func() {
			dp.Close()
		}()

		dp.running = true

		for v := range dp.blkCh {
			dp.lock.Lock()
			if v.Height-dp.pruneBlock > dp.size {
				dp.pruneBlock++

				dp.worker.Do(func() error {
					log.Info("prune", "height", v.Height)
					bs, err := v.Get(v.Root[:])
					if err != nil {
						log.Error("prune state", "error", err, "root", v.Root.Hex())
						return nil
					}
					tr := trpool.Get().(*trie.Trie)
					rlp.Decode(bytes.NewReader(bs), &tr)
					defer trpool.Put(tr)
					it := trie.NewIterator(tr.NodeIterator(nil))
					for it.Next() {
						log.Info("interator", "height", v.Height,
							"key", tools.BytesToStringFast(it.Key), "val", tools.BytesToStringFast(it.Value))
					}

					return nil
				})
			}
			dp.lock.Unlock()
		}
	}()
}

var (
	dp     *PruneState
	trpool = &sync.Pool{
		New: func() interface{} {
			return &trie.Trie{}
		},
	}
)

func (p *PruneState) Close() error {
	return p.db.Close()
}
