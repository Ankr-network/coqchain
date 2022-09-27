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

package kv

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"github.com/sunvim/wal"
	"go.etcd.io/bbolt"
)

type Database struct {
	log     *wal.Log
	global  *bbolt.DB
	actdb   *bbolt.DB
	recents *lru.ARCCache // cache recent db
}

func New(path string) (*Database, error) {
	d := &Database{}
	walpath := fmt.Sprintf("%s/LOG", path)
	l, err := wal.Open(walpath, nil)
	if err != nil {
		return nil, nil
	}
	d.log = l
	return d, nil
}

// Has retrieves if a key is present in the key-value data store.
func (d *Database) Has(key []byte) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Get retrieves the given key if it's present in the key-value data store.
func (d *Database) Get(key []byte) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// Put inserts the given value into the key-value data store.
func (d *Database) Put(key []byte, value []byte) error {
	data := &KeyValue{
		Key: key,
		Val: value,
	}
	d.log.Write(data.Marshal())
	return nil
}

// Delete removes the key from the key-value data store.
func (d *Database) Delete(key []byte) error {
	panic("not implemented") // TODO: Implement
}

// Stat returns a particular internal stat of the database.
func (d *Database) Stat(property string) (string, error) {
	panic("not implemented") // TODO: Implement
}

// Compact flattens the underlying data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (d *Database) Compact(start []byte, limit []byte) error {
	return nil
}

func (d *Database) Close() error {
	var err error
	e := d.log.Close()
	if e != nil {
		err = e
	}
	e = d.actdb.Close()
	if e != nil {
		err = e
	}
	e = d.global.Close()
	if e != nil {
		err = e
	}
	for _, key := range d.recents.Keys() {
		v, _ := d.recents.Get(key)
		e = v.(*bbolt.DB).Close()
		if e != nil {
			err = e
		}
	}
	return err
}
