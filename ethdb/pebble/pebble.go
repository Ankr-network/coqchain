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
	"bytes"
	"sync"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/cockroachdb/pebble"
)

type Database struct {
	kv *pebble.DB
	wo *pebble.WriteOptions
	ro *pebble.IterOptions
}

const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
)

func New(file string, caches int64, namespace string, readonly bool) (*Database, error) {

	c := pebble.NewCache(caches * MiB)

	opts := &pebble.Options{
		BytesPerSync: 4 * MiB,
		MaxOpenFiles: 8 * 1024,
		Cache:        c,
	}

	kv, err := pebble.Open(file, opts)
	if err != nil {
		panic(err)
	}

	ldb := &Database{
		kv: kv,
		wo: &pebble.WriteOptions{},
		ro: &pebble.IterOptions{},
	}

	return ldb, nil
}

// Has retrieves if a key is present in the key-value data store.
func (d *Database) Has(key []byte) (bool, error) {
	val, err := d.Get(key)
	if err == pebble.ErrNotFound {
		err = nil
	}
	if err == nil && len(val) != 0 {
		return true, nil
	}
	return false, err
}

// Get retrieves the given key if it's present in the key-value data store.
func (d *Database) Get(key []byte) ([]byte, error) {
	val, closer, err := d.kv.Get(key)
	if err == nil {
		closer.Close()
	}
	return val, err
}

// HasAncient returns an indicator whether the specified data exists in the
// ancient store.
func (d *Database) HasAncient(kind string, number uint64) (bool, error) {
	return false, nil
}

// Ancient retrieves an ancient binary blob from the append-only immutable files.
func (d *Database) Ancient(kind string, number uint64) ([]byte, error) {
	return nil, nil
}

// AncientRange retrieves multiple items in sequence, starting from the index 'start'.
// It will return
//   - at most 'count' items,
//   - at least 1 item (even if exceeding the maxBytes), but will otherwise
//     return as many items as fit into maxBytes.
func (d *Database) AncientRange(kind string, start uint64, count uint64, maxBytes uint64) ([][]byte, error) {
	return nil, nil
}

// Ancients returns the ancient item numbers in the ancient store.
func (d *Database) Ancients() (uint64, error) {
	return 0, nil
}

// AncientSize returns the ancient size of the specified category.
func (d *Database) AncientSize(kind string) (uint64, error) {
	return 0, nil
}

// ReadAncients runs the given read operation while ensuring that no writes take place
// on the underlying freezer.
func (d *Database) ReadAncients(fn func(ethdb.AncientReader) error) (err error) {
	return nil
}

// Put inserts the given value into the key-value data store.
func (d *Database) Put(key []byte, value []byte) error {
	return d.kv.Set(key, value, d.wo)
}

// Delete removes the key from the key-value data store.
func (d *Database) Delete(key []byte) error {
	return d.kv.Delete(key, d.wo)
}

// ModifyAncients runs a write operation on the ancient store.
// If the function returns an error, any changes to the underlying store are reverted.
// The integer return value is the total size of the written data.
func (d *Database) ModifyAncients(_ func(ethdb.AncientWriteOp) error) (int64, error) {
	return 0, nil
}

// TruncateAncients discards all but the first n ancient data from the ancient store.
func (d *Database) TruncateAncients(n uint64) error {
	return nil
}

// Sync flushes all in-memory ancient store data to disk.
func (d *Database) Sync() error {
	return nil
}

// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called.
func (d *Database) NewBatch() ethdb.Batch {
	return &batch{
		db:     d.kv,
		writes: make([]keyvalue, 0),
	}
}

var (
	bufpool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (d *Database) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	if prefix != nil || start != nil {
		lowBound, upBound := bytesPrefix(prefix, start)
		return &Iter{kvi: d.kv.NewIter(&pebble.IterOptions{
			LowerBound: lowBound,
			UpperBound: upBound,
		})}
	} else {
		return &Iter{kvi: d.kv.NewIter(&pebble.IterOptions{})}
	}
}

// Stat returns a particular internal stat of the database.
func (d *Database) Stat(property string) (string, error) {
	return "", nil
}

// Compact flattens the underlyingb data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (d *Database) Compact(start []byte, limit []byte) error {
	return d.kv.Compact(start, limit, false)
}

func (d *Database) Close() error {
	return d.kv.Close()
}
