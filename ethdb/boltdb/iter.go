package boltdb

import (
	"bytes"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/utils"
	"go.etcd.io/bbolt"
)

type Iter struct {
	db            *BoltDB
	prefix, start []byte
	curkey        []byte
	key, val      []byte
	err           error
	first         bool
	opts          *ethdb.Option
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *Iter) Next() bool {
	var (
		key, val []byte
		nextkey  []byte
		rs       bool
	)
	i.db.db.View(func(tx *bbolt.Tx) error {
		if !i.first && i.curkey == nil {
			rs = false
			i.key = nil
			i.val = nil
			return nil
		}
		i.first = false
		c := tx.Bucket(utils.S2B(i.opts.Name)).Cursor()
		key, val = c.Seek(i.curkey)
		if key != nil && bytes.HasPrefix(key, i.prefix) {
			i.key = make([]byte, len(key))
			copy(i.key, key)
			i.val = make([]byte, len(val))
			copy(i.val, val)
			nextkey, _ = c.Next()
			if nextkey != nil {
				i.curkey = make([]byte, len(nextkey))
				copy(i.curkey, nextkey)
			} else {
				i.curkey = nil
			}
			rs = true
		} else {
			rs = false
		}
		return nil
	})
	return rs
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (i *Iter) Error() error {
	return i.err
}

// Key returns the key of the current key/value pair, or nil if done. The caller
// should not modify the contents of the returned slice, and its contents may
// change on the next call to Next.
func (i *Iter) Key() []byte {
	return i.key
}

// Value returns the value of the current key/value pair, or nil if done. The
// caller should not modify the contents of the returned slice, and its contents
// may change on the next call to Next.
func (i *Iter) Value() []byte {
	return i.val
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (i *Iter) Release() {
	i.key, i.val, i.curkey, i.err, i.prefix = nil, nil, nil, nil, nil
}
