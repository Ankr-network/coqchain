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
	"context"
	"path/filepath"
	"time"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/mdbx"
)

type DbImpl struct {
	chaindb Database
	dir     string
}

func NewMDBXDB(path string) *DbImpl {
	db := &DbImpl{dir: filepath.Join(path, "chaindb")}
	chaindb, err := mdbx.NewMDBX(nil).Label(kv.ChainDB).WithTableCfg(func(defaultBuckets kv.TableCfg) kv.TableCfg {
		return kv.ChaindataTablesCfg
	}).
		SyncPeriod(10 * time.Second).Path(db.dir).Open()
	if err != nil {
		panic(err)
	}
	chainRw, err := chaindb.BeginRw(context.Background())
	if err != nil {
		panic(err)
	}

	db.chaindb = WrapIntoTxDB(chainRw)

	return db
}

func (d *DbImpl) Path() string {
	return d.dir
}

// Has retrieves if a key is present in the key-value data store.
func (d *DbImpl) Has(key []byte, opts *ethdb.Option) (bool, error) {
	return d.chaindb.Has(opts.Name, key)
}

// Get retrieves the given key if it's present in the key-value data store.
func (d *DbImpl) Get(key []byte, opts *ethdb.Option) ([]byte, error) {
	return d.chaindb.Get(opts.Name, key)
}

// Put inserts the given value into the key-value data store.
func (d *DbImpl) Put(key []byte, value []byte, opts *ethdb.Option) error {
	return d.chaindb.Put(opts.Name, key, value)
}

// Delete removes the key from the key-value data store.
func (d *DbImpl) Delete(key []byte, opts *ethdb.Option) error {
	return d.chaindb.Delete(opts.Name, key)
}

// Sync flushes all in-memory ancient store data to disk.
func (d *DbImpl) Sync() error {
	return nil
}

// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called.
func (d *DbImpl) NewBatch() ethdb.Batch {
	return &DbBatch{
		db: d,
	}
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (d *DbImpl) NewIterator(prefix []byte, start []byte, opts *ethdb.Option) ethdb.Iterator {
	return &DbIter{db: d, prefix: append(prefix, start...), opts: opts}
}

// Stat returns a particular internal stat of the database.
func (d *DbImpl) Stat(property string, opts *ethdb.Option) (string, error) {
	return "", nil
}

// Compact flattens the underlying data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (d *DbImpl) Compact(start []byte, limit []byte, opts *ethdb.Option) error {
	return nil
}

func (d *DbImpl) Close() error {
	d.chaindb.RwKV().Close()
	d.chaindb.Close()
	return nil
}
