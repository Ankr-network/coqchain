package boltdb

import (
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/utils"
	"go.etcd.io/bbolt"
)

type keyvalue struct {
	key    []byte
	value  []byte
	delete bool
	opts   *ethdb.Option
}

type Batch struct {
	db     *BoltDB
	writes []keyvalue
	size   int
}

// Put inserts the given value into the key-value data store.
func (b *Batch) Put(key []byte, value []byte, opts *ethdb.Option) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), common.CopyBytes(value), false, opts})
	b.size += len(key) + len(value)
	return nil
}

// Delete removes the key from the key-value data store.
func (b *Batch) Delete(key []byte, opts *ethdb.Option) error {
	b.writes = append(b.writes, keyvalue{common.CopyBytes(key), nil, true, opts})
	b.size += len(key)
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (b *Batch) ValueSize() int {
	return b.size
}

// Write flushes any accumulated data to disk.
func (b *Batch) Write() error {
	var err error
	err = b.db.db.Batch(func(tx *bbolt.Tx) error {
		for _, keyvalue := range b.writes {
			bt := tx.Bucket(utils.S2B(keyvalue.opts.Name))
			if keyvalue.delete {
				if err = bt.Delete(keyvalue.key); err != nil {
					return err
				}
				continue
			}
			if err = bt.Put(keyvalue.key, keyvalue.value); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// Reset resets the batch for reuse.
func (b *Batch) Reset() {
	b.size = 0
	b.writes = b.writes[:0]
}

// Replay replays the batch contents.
func (b *Batch) Replay(w ethdb.KeyValueWriter, opts *ethdb.Option) error {
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
