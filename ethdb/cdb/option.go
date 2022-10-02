package cdb

import (
	"time"

	"github.com/torquem-ch/mdbx-go/mdbx"
)

type Option struct {
	MaxDB      uint64
	MaxReaders uint64
	SyncPeriod time.Duration
	Flags      uint
}

var defaultOption = &Option{
	MaxDB:      128,
	MaxReaders: 6400,
	SyncPeriod: 3 * time.Second,
	Flags:      mdbx.NoReadahead | mdbx.Coalesce | mdbx.Durable,
}
