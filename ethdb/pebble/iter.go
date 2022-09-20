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

package pebble

import (
	"github.com/cockroachdb/pebble"
)

type Iter struct {
	kvi   *pebble.Iterator
	index uint64
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *Iter) Next() bool {
	if i.index == 0 {
		i.kvi.First()
		i.index++
		return i.kvi.Valid()
	}
	i.kvi.Next()
	return i.kvi.Valid()
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (i *Iter) Error() error {
	return i.kvi.Error()
}

// Key returns the key of the current key/value pair, or nil if done. The caller
// should not modify the contents of the returned slice, and its contents may
// change on the next call to Next.
func (i *Iter) Key() []byte {
	return i.kvi.Key()
}

// Value returns the value of the current key/value pair, or nil if done. The
// caller should not modify the contents of the returned slice, and its contents
// may change on the next call to Next.
func (i *Iter) Value() []byte {
	return i.kvi.Value()
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (i *Iter) Release() {
	i.index = 0
	i.kvi.Close()
}

func bytesPrefix(prefix []byte, start []byte) ([]byte, []byte) {
	if prefix == nil && start == nil {
		return nil, nil
	}
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	prefix = append(prefix, start...)
	return prefix, limit
}
