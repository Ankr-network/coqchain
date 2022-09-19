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
	"testing"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
)

func newPebble() *Database {
	db, _ := pebble.Open("", &pebble.Options{FS: vfs.NewMem()})
	return &Database{kv: db}
}

func TestIter(t *testing.T) {
	p := newPebble()
	tbl := []struct {
		Key []byte
		Val []byte
		Act []byte
	}{
		{Key: []byte("k1"), Val: []byte("v1"), Act: nil},
		{Key: []byte("k2"), Val: []byte("v2"), Act: nil},
		{Key: []byte("k21"), Val: []byte("v21"), Act: nil},
		{Key: []byte("k31"), Val: []byte("v31"), Act: nil},
		{Key: []byte("k32"), Val: []byte("v32"), Act: nil},
	}

	for _, v := range tbl {
		p.Put(v.Key, v.Val)
	}

	iter := p.NewIterator(nil, nil)

	for iter.Next() {
		exist := false
		for _, v := range tbl {
			if bytes.Equal(v.Key, iter.Key()) {
				exist = bytes.Equal(v.Val, iter.Value())
				break
			}
		}
		if !exist {
			t.Fail()
		}
	}
}
