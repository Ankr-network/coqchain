package boltdb

import (
	"bytes"

	"github.com/Ankr-network/coqchain/utils"
	"go.etcd.io/bbolt"
)

type Iter struct {
	db              *BoltDB
	prefix, nextkey []byte
	key, val        []byte
	err             error
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *Iter) Next() bool {
	var err error
	err = i.db.db.View(func(tx *bbolt.Tx) error {
		if !bytes.HasPrefix(i.nextkey, i.prefix) {
			return ErrNotFound
		}
		c := tx.Bucket(utils.S2B(defaultBucket)).Cursor()
		i.key, i.val = c.Seek(i.nextkey)
		if i.key != nil && bytes.HasPrefix(i.key, i.prefix) {
			i.nextkey, _ = c.Next()
		} else {
			return ErrNotFound
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
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
	i.key, i.val, i.nextkey, i.err, i.prefix = nil, nil, nil, nil, nil
}
