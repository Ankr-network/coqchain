// Copyright (c) 2022 coqchain team
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

package container

import "math/big"

type IHeaderInTailOut interface {
	Put(item *big.Int) *big.Int
	List() []*big.Int
}

type HeaderInTailOut struct {
	size int
	set  map[*big.Int]struct{}
	buf  []*big.Int
}

func New(size int) IHeaderInTailOut {
	return &HeaderInTailOut{
		size: size,
		set:  make(map[*big.Int]struct{}, size),
		buf:  make([]*big.Int, 0, size),
	}
}

func (hi *HeaderInTailOut) List() []*big.Int {
	rs := make([]*big.Int, 0, len(hi.buf))
	rs = append(rs, hi.buf...)
	return rs
}

func (hi *HeaderInTailOut) Put(item *big.Int) *big.Int {
	bufLen := len(hi.buf)
	if bufLen >= hi.size {

		if _, ok := hi.set[item]; ok {
			return nil
		}

		var rs *big.Int
		rs = hi.buf[0]
		hi.buf = hi.buf[1:]
		delete(hi.set, rs)
		hi.buf = append(hi.buf, item)
		hi.set[item] = struct{}{}
		return rs

	} else {
		if _, ok := hi.set[item]; ok {
			return nil
		}
		hi.set[item] = struct{}{}
		hi.buf = append(hi.buf, item)
	}
	return nil
}
