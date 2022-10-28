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
	"github.com/Ankr-network/coqchain/core/contracts"
	"github.com/Ankr-network/coqchain/log"
	"go.etcd.io/bbolt"
)

type AddrMgr struct {
	lock       sync.Mutex
	zeroAddrs  map[common.Address]struct{}
	slashAddrs map[common.Address]struct{}
	store      *bbolt.DB
}

func NewAddrMgr(datadir string) *AddrMgr {
	var (
		db  *bbolt.DB
		err error
	)

	if datadir != "" {
		name := fmt.Sprintf("%s/%s", datadir, extenddb)
		db, err = bbolt.Open(name, 0644, nil)
		if err != nil {
			panic(err)
		}
	}

	am := &AddrMgr{
		zeroAddrs: make(map[common.Address]struct{}),
		store:     db,
	}

	am.zeroAddrs[contracts.SlashContract.Address] = struct{}{}

	return am
}
func (z *AddrMgr) AddSlashAddr(addr common.Address) {

	z.lock.Lock()
	defer z.lock.Unlock()
	if addr == (common.Address{}) {
		return
	}
	z.slashAddrs[addr] = struct{}{}
}

func (z *AddrMgr) RemoveSlashAddr(addr common.Address) {
	z.lock.Lock()
	defer z.lock.Unlock()

	if z.ContainSlashAddr(addr) {
		delete(z.slashAddrs, addr)
	}
}

func (z *AddrMgr) ContainSlashAddr(addr common.Address) bool {
	_, ok := z.slashAddrs[addr]
	return ok
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

var (
	addrsMgr *AddrMgr
)

const (
	extenddb   = "extenddb"
	zeroGasFee = "zerogasfee"
	slashAddr  = "slashaddress"
)

func InitAddrMgr(datadir string) {
	log.Info("init address extend data")
	addrsMgr = NewAddrMgr(datadir)
	if addrsMgr.store == nil {
		return
	}
	// load zero gas fee address
	if err := addrsMgr.store.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(zeroGasFee))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			log.Info("adding", "address", common.BytesToAddress(k))
			addrsMgr.zeroAddrs[common.BytesToAddress(k)] = struct{}{}
			return nil
		})
		b = tx.Bucket([]byte(slashAddr))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			log.Info("adding", "address", common.BytesToAddress(k))
			addrsMgr.slashAddrs[common.BytesToAddress(k)] = struct{}{}
			return nil
		})
		return nil
	}); err != nil {
		log.Error("init address extend data", "error", err)
		return
	}
}

func Close() {

	if addrsMgr.store == nil {
		return
	}

	defer addrsMgr.store.Close()

	log.Info("close address extend data ...")
	if len(addrsMgr.zeroAddrs) > 0 {
		if err := addrsMgr.store.Update(func(tx *bbolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(zeroGasFee))
			if err != nil {
				return err
			}
			for addr := range addrsMgr.zeroAddrs {
				b.Put(addr.Bytes(), []byte{})
			}
			b, err = tx.CreateBucketIfNotExists([]byte(slashAddr))
			if err != nil {
				return err
			}
			for addr := range addrsMgr.slashAddrs {
				b.Put(addr.Bytes(), []byte{})
			}
			return nil
		}); err != nil {
			log.Error("update address extend data", "error", err)
			return
		}
	}
	log.Info("close address extend data success")
}

func AddZeroFeeAddress(addr common.Address) {
	addrsMgr.AddZeroAddr(addr)
}

func RemoveZeroFeeAddress(addr common.Address) {
	addrsMgr.RemoveZeroAddr(addr)
}

func ContainsZeroFeeAddress(addr common.Address) bool {
	return addrsMgr.ContainZeroAddr(addr)
}

func ListZeroFeeAddress() []common.Address {
	addrs := make([]common.Address, 0, len(addrsMgr.zeroAddrs))
	for addr := range addrsMgr.zeroAddrs {
		addrs = append(addrs, addr)
	}
	return addrs
}

func AddSlashAddress(addr common.Address) {
	addrsMgr.AddSlashAddr(addr)
}

func RemoveSlashAddress(addr common.Address) {
	addrsMgr.RemoveSlashAddr(addr)
}

func ContainsSlashAddress(addr common.Address) bool {
	return addrsMgr.ContainSlashAddr(addr)
}

func ListSlashAddress() []common.Address {
	addrs := make([]common.Address, 0, len(addrsMgr.slashAddrs))
	for addr := range addrsMgr.slashAddrs {
		addrs = append(addrs, addr)
	}
	return addrs
}
