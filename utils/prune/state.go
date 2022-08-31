package prune

import (
	"fmt"

	"github.com/Ankr-network/coqchain/cmd/utils"
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/log"
	"github.com/sunvim/utils/tools"
	"go.etcd.io/bbolt"
	"gopkg.in/urfave/cli.v1"
)

const (
	stateName             = "prune_state"
	stateBucketName       = "state"
	stateRecentPruneBlock = "recent_prune_block"
)

type BlockState struct {
	Height uint64
	Root   common.Hash
}

type PruneState struct {
	size  uint64
	blkCh chan *BlockState
	db    *bbolt.DB
}

func New(ctx *cli.Context) *PruneState {
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
		db:    db,
		blkCh: make(chan *BlockState, 128),
		size:  ctx.GlobalUint64(utils.PruneSizeFlag.Name),
	}
}

func InitPruneWorker(ctx *cli.Context) {
	log.Info("pruner initing ...")
	defaultPruneWorker = New(ctx)

	log.Info("pruner init over.")
}

func SetLatestBlock(num uint64, hash common.Hash) {
	defaultPruneWorker.blkCh <- &BlockState{Height: num, Root: hash}
}

func Run() {
	defer defaultPruneWorker.Close()
}

var (
	defaultPruneWorker *PruneState
)

func (p *PruneState) Close() error {
	return p.db.Close()
}
