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

package mdbx

import (
	"github.com/Ankr-network/coqchain/ethdb"
)

type DbIter struct {
	db       *DbImpl
	prefix   []byte
	nextKey  []byte
	key, val []byte
	err      error
	opts     *ethdb.Option
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *DbIter) Next() bool {
	if i.err != nil {
		return false
	}
	kvtx, err := i.db.chaindb.BeginRw(i.db.ctx)
	if err != nil {
		return false
	}
	c, err := kvtx.RwCursor(i.opts.Name)
	if err != nil {
		return false
	}
	if i.nextKey == nil {
		k, v, err := c.Seek(i.prefix)
		if err != nil {
			return false
		}
		i.key, i.val = k, v
		i.nextKey, _, i.err = c.Next()
	} else {
		i.key, i.val, err = c.Seek(i.nextKey)
		if err != nil {
			return false
		}
	}
	return true
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (i *DbIter) Error() error {
	return nil
}

// Key returns the key of the current key/value pair, or nil if done. The caller
// should not modify the contents of the returned slice, and its contents may
// change on the next call to Next.
func (i *DbIter) Key() []byte {
	return i.key
}

// Value returns the value of the current key/value pair, or nil if done. The
// caller should not modify the contents of the returned slice, and its contents
// may change on the next call to Next.
func (i *DbIter) Value() []byte {
	return i.val
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (i *DbIter) Release() {
	i.key, i.val, i.err, i.prefix = nil, nil, nil, nil
}
