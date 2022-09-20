package pebble

import (
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/cockroachdb/pebble"
)

type keyvalue struct {
	key    []byte
	value  []byte
	delete bool
}
type batch struct {
	db     *pebble.DB
	writes []keyvalue
	size   int
}

// Put inserts the given value into the batch for later committing.
func (b *batch) Put(key, value []byte) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), common.CopyBytes(value), false})
	b.size += len(key) + len(value)
	return nil
}

// Delete inserts the a key removal into the batch for later committing.
func (b *batch) Delete(key []byte) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), nil, true})
	b.size += len(key)
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (b *batch) ValueSize() int {
	return b.size
}

// Write flushes any accumulated data to disk.
func (b *batch) Write() error {
	for _, keyvalue := range b.writes {
		if keyvalue.delete {
			b.db.Delete(keyvalue.key, &pebble.WriteOptions{})
			continue
		}
		b.db.Set(keyvalue.key, keyvalue.value, &pebble.WriteOptions{})
	}
	return nil
}

// Reset resets the batch for reuse.
func (b *batch) Reset() {
	b.writes = b.writes[:0]
	b.size = 0
}

// Replay replays the batch contents.
func (b *batch) Replay(w ethdb.KeyValueWriter) error {
	for _, keyvalue := range b.writes {
		if keyvalue.delete {
			if err := w.Delete(keyvalue.key); err != nil {
				return err
			}
			continue
		}
		if err := w.Put(keyvalue.key, keyvalue.value); err != nil {
			return err
		}
	}
	return nil
}
