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

package extdb

import (
	"fmt"
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/log"
	"go.etcd.io/bbolt"
	"gopkg.in/urfave/cli.v1"
)

type AddrMgr struct {
	lock      sync.Mutex
	zeroAddrs map[common.Address]struct{}
	store     *bbolt.DB
}

func NewAddrMgr(ctx *cli.Context) *AddrMgr {
	name := fmt.Sprintf("%s/%s", ctx.GlobalString("datadir"), extenddb)
	db, err := bbolt.Open(name, 0644, nil)
	if err != nil {
		panic(err)
	}

	return &AddrMgr{
		zeroAddrs: make(map[common.Address]struct{}),
		store:     db,
	}
}

func (z *AddrMgr) AddZeroAddr(addr common.Address) {

	z.lock.Lock()
	defer z.lock.Unlock()
	if addr == (common.Address{}) {
		return
	}
	z.zeroAddrs[addr] = struct{}{}
}

func (z *AddrMgr) RemoveZeroAddr(addr common.Address) {
	z.lock.Lock()
	defer z.lock.Unlock()

	if z.ContainZeroAddr(addr) {
		delete(z.zeroAddrs, addr)
	}
}

func (z *AddrMgr) ContainZeroAddr(addr common.Address) bool {
	_, ok := z.zeroAddrs[addr]
	return ok
}

var defaultZeroAddrs *AddrMgr

const (
	extenddb = "extenddb"
)

func InitAddrMgr(ctx *cli.Context) {
	log.Info("init zero fee address db")
	defaultZeroAddrs = NewAddrMgr(ctx)
	// load zero gas fee address
	if err := defaultZeroAddrs.store.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(extenddb))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			log.Info("adding", "address", common.BytesToAddress(k))
			defaultZeroAddrs.zeroAddrs[common.BytesToAddress(k)] = struct{}{}
			return nil
		})
		return nil
	}); err != nil {
		log.Error("init zero fee address db", "error", err)
		return
	}

}

func Close() {

	defer defaultZeroAddrs.store.Close()

	log.Info("close zero fee address db ...")
	if len(defaultZeroAddrs.zeroAddrs) > 0 {
		if err := defaultZeroAddrs.store.Update(func(tx *bbolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(extenddb))
			if err != nil {
				return err
			}
			for addr := range defaultZeroAddrs.zeroAddrs {
				b.Put(addr.Bytes(), []byte{})
			}
			return nil
		}); err != nil {
			log.Error("update zero fee address db", "error", err)
		}
	}
	log.Info("close zero fee address db success")
}

func AddZeroFeeAddress(addr common.Address) {
	defaultZeroAddrs.AddZeroAddr(addr)
}

func RemoveZeroFeeAddress(addr common.Address) {
	defaultZeroAddrs.RemoveZeroAddr(addr)
}

func ContainsZeroFeeAddress(addr common.Address) bool {
	return defaultZeroAddrs.ContainZeroAddr(addr)
}

func ListZeroFeeAddress() []common.Address {
	addrs := make([]common.Address, 0, len(defaultZeroAddrs.zeroAddrs))
	for addr := range defaultZeroAddrs.zeroAddrs {
		addrs = append(addrs, addr)
	}
	return addrs
}
