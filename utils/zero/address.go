// Copyright (c) 2022 mobus sunsc0220@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zero

import (
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/log"
	"go.etcd.io/bbolt"
	"gopkg.in/urfave/cli.v1"
)

type ZeroAddrs struct {
	lock    sync.Mutex
	addrs   map[common.Address]struct{}
	rmAddrs map[common.Address]struct{}
	db      *bbolt.DB
	len     int
}

func NewZeroAddrs(ctx *cli.Context) *ZeroAddrs {
	opts := bbolt.DefaultOptions
	opts.InitialMmapSize = 4 * 1024
	db, err := bbolt.Open(ctx.GlobalString("datadir")+"/zero_addrs.db", 0600, opts)
	if err != nil {
		panic(err)
	}
	return &ZeroAddrs{
		addrs: make(map[common.Address]struct{}),
		db:    db,
	}
}

func (z *ZeroAddrs) Add(addr common.Address) {

	z.lock.Lock()
	defer z.lock.Unlock()

	z.addrs[addr] = struct{}{}
	z.len++
}

func (z *ZeroAddrs) Remove(addr common.Address) {
	z.lock.Lock()
	defer z.lock.Unlock()

	if z.Contains(addr) {
		z.rmAddrs[addr] = struct{}{}
		delete(z.addrs, addr)
		z.len--
	}
}

func (z *ZeroAddrs) Contains(addr common.Address) bool {
	_, ok := z.addrs[addr]
	return ok
}

var defaultZeroAddrs *ZeroAddrs

const (
	zeroAddrsBucket = "zero_addrs"
)

func InitZeroFeeAddress(ctx *cli.Context) {

	log.Info("init zero fee address db")

	defaultZeroAddrs = NewZeroAddrs(ctx)
	defaultZeroAddrs.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(zeroAddrsBucket))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, v []byte) error {
			defaultZeroAddrs.Add(common.BytesToAddress(k))
			return nil
		})
	})
}

func Close() {
	log.Info("close zero fee address db ...")
	defaultZeroAddrs.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(zeroAddrsBucket))
		if b == nil {
			return nil
		}

		for addr := range defaultZeroAddrs.addrs {
			b.Put(addr.Bytes(), []byte{})
		}

		b.ForEach(func(k, v []byte) error {
			if _, ok := defaultZeroAddrs.rmAddrs[common.BytesToAddress(k)]; ok {
				return b.Delete(k)
			}
			return nil
		})
		tx.Commit()
		return nil
	})

	if err := defaultZeroAddrs.db.Close(); err != nil {
		log.Warn("close zero fee address db error", "err", err)
	}
	log.Info("close zero fee address db success")
}

func AddZeroFeeAddress(addr common.Address) {
	defaultZeroAddrs.Add(addr)
}

func RemoveZeroFeeAddress(addr common.Address) {
	defaultZeroAddrs.Remove(addr)
}

func ContainsZeroFeeAddress(addr common.Address) bool {
	return defaultZeroAddrs.Contains(addr)
}

func ListZeroFeeAddress() []common.Address {
	addrs := make([]common.Address, 0, len(defaultZeroAddrs.addrs))
	for addr := range defaultZeroAddrs.addrs {
		addrs = append(addrs, addr)
	}
	return addrs
}

func Len() int {
	return defaultZeroAddrs.len
}
