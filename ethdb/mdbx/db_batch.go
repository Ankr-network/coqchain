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
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/ethdb"
)

type keyvalue struct {
	key    []byte
	value  []byte
	opts   *ethdb.Option
	delete bool
}

type DbBatch struct {
	db     *DbImpl
	writes []keyvalue
	size   int
}

// Put inserts the given value into the key-value data store.
func (b *DbBatch) Put(key []byte, value []byte, opts *ethdb.Option) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), common.CopyBytes(value), opts, false})
	b.size += len(key) + len(value)
	return nil
}

// Delete removes the key from the key-value data store.
func (b *DbBatch) Delete(key []byte, opts *ethdb.Option) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), nil, opts, true})
	b.size += len(key)
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (b *DbBatch) ValueSize() int {
	return b.size
}

// Write flushes any accumulated data to disk.
func (b *DbBatch) Write() error {
	var err error
	for _, keyvalue := range b.writes {
		if keyvalue.delete {
			if err = b.db.Delete(keyvalue.key, keyvalue.opts); err != nil {
				return err
			}
			continue
		}
		if err = b.db.Put(keyvalue.key, keyvalue.value, keyvalue.opts); err != nil {
			return err
		}
	}
	return nil
}

// Reset resets the batch for reuse.
func (b *DbBatch) Reset() {
	b.size = 0
	b.writes = b.writes[:0]
}

// Replay replays the batch contents.
func (b *DbBatch) Replay(w ethdb.KeyValueWriter, opts *ethdb.Option) error {
	for _, keyvalue := range b.writes {
		if keyvalue.delete {
			if err := w.Delete(keyvalue.key, opts); err != nil {
				return err
			}
			continue
		}
		if err := w.Put(keyvalue.key, keyvalue.value, opts); err != nil {
			return err
		}
	}
	return nil
}
