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

package cdb

import (
	"bytes"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

type DbIter struct {
	db                    *MDB
	first                 bool
	curkey, prefix, start []byte
	key, val              []byte
	err                   error
	opts                  *ethdb.Option
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *DbIter) Next() bool {
	err := i.db.env.View(func(txn *mdbx.Txn) error {
		var (
			key, val []byte
			err      error
		)

		dbi, _ := txn.OpenDBI(i.opts.Name, mdbx.Create|mdbx.DupSort, nil, nil)
		c, err := txn.OpenCursor(dbi)
		if err != nil {
			return err
		}
		defer c.Close()
		switch {
		case i.first && i.prefix == nil && i.start == nil:
			i.first = false
			key, val, err = c.Get(nil, nil, mdbx.First)
		case i.first && i.start != nil:
			i.first = false
			i.curkey = append(i.prefix, i.start...)
			key, val, err = c.Get(i.curkey, nil, mdbx.Set)
			if key == nil {
				key, val, err = c.Get(i.curkey, nil, mdbx.Last)
			}
		case i.first && (i.prefix != nil && i.start == nil):
			i.first = false
			i.curkey = append(i.prefix, i.start...)
			key, val, err = c.Get(i.curkey, nil, mdbx.First)
		case !i.first:
			key, val, err = c.Get(i.curkey, nil, mdbx.Set)
		}
		if key != nil && bytes.HasPrefix(key, i.prefix) {
			i.key = key
			i.val = val
			i.curkey, _, _ = c.Get(key, nil, mdbx.Next)
			return nil
		} else {
			i.key = nil
			i.val = nil
			return ErrNotFound
		}

	})
	if err != nil {
		return false
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
	i.first = true
	i.key, i.val, i.err, i.curkey = nil, nil, nil, nil
}
