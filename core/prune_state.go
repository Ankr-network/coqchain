package core

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/common/container"
	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/log"
	"github.com/Ankr-network/coqchain/utils"
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
	db     ethdb.Database
}

type PruneState struct {
	lock       sync.Mutex
	size       uint64
	blkCh      chan *BlockState
	pruneBlock uint64
	buf        container.IHeaderInTailOut
	worker     *workpool.WorkPool
	db         *bbolt.DB
	state      bool
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
			tx.CreateBucketIfNotExists(utils.S2B(stateBucketName))
			return nil
		})
		return &PruneState{
			db:         db,
			pruneBlock: 1,
			buf:        container.New(int(size)),
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

func SetLatestBlock(num uint64, hash common.Hash, db ethdb.Database) {
	dp.blkCh <- &BlockState{Height: num, Root: hash, db: db}
}

func IsRunning() bool {
	return dp.state
}

func Run() {
	go func() {
		dp.state = true
		for v := range dp.blkCh {
			out := dp.buf.Put(new(big.Int).SetBytes(v.Root.Bytes()))
			if out != nil {
				outh := common.BigToHash(out)
				v.db.Delete(outh[:])
			}
		}
	}()
}

var (
	dp *PruneState
)

func (p *PruneState) Close() error {
	p.db.Batch(func(tx *bbolt.Tx) error {
		return nil
	})
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
