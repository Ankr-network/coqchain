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

import "github.com/Ankr-network/coqchain/common"

type ZeroAddrs struct {
	addrs map[common.Address]struct{}
}

func NewZeroAddrs() *ZeroAddrs {
	return &ZeroAddrs{
		addrs: make(map[common.Address]struct{}),
	}
}

func (z *ZeroAddrs) Add(addr common.Address) {
	z.addrs[addr] = struct{}{}
}

func (z *ZeroAddrs) Remove(addr common.Address) {
	delete(z.addrs, addr)
}

func (z *ZeroAddrs) Contains(addr common.Address) bool {
	_, ok := z.addrs[addr]
	return ok
}

var defaultZeroAddrs = NewZeroAddrs()

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
