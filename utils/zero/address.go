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
	"bufio"
	"os"
	"sync"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/log"
	"gopkg.in/urfave/cli.v1"
)

type ZeroAddrs struct {
	lock    sync.Mutex
	addrs   map[common.Address]struct{}
	rmAddrs map[common.Address]struct{}
	db      *os.File
	len     int
}

func NewZeroAddrs(ctx *cli.Context) *ZeroAddrs {
	db, err := os.OpenFile(ctx.GlobalString("datadir")+"/zero_addrs", os.O_CREATE|os.O_RDWR, 0600)
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

	if addr == (common.Address{}) {
		return
	}

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
	scanner := bufio.NewScanner(defaultZeroAddrs.db)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		defaultZeroAddrs.Add(common.HexToAddress(line))
	}
}

func Close() {
	log.Info("close zero fee address db ...")

	wr := bufio.NewWriter(defaultZeroAddrs.db)
	for addr := range defaultZeroAddrs.addrs {
		wr.WriteString(addr.Hex())
		wr.WriteString("\n")
	}
	wr.Flush()

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
