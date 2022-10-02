package cdb

import (
	"bytes"
	"errors"
	"os"
	"sync"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/c2h5oh/datasize"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

type MDB struct {
	mu   sync.RWMutex
	path string
	opts *Option
	env  *mdbx.Env
}

func NewMDB(path string, opts *Option) (*MDB, error) {
	var err error
	m := &MDB{}

	m.env, err = mdbx.NewEnv()
	if err != nil {
		return nil, err
	}

	if opts == nil {
		opts = defaultOption
		m.opts = opts
	}

	err = m.env.SetOption(mdbx.OptMaxDB, opts.MaxDB)
	if err != nil {
		return nil, err
	}
	const pageSize = int(4 * datasize.KB)
	err = m.env.SetGeometry(-1, -1, 64*1024*pageSize, int(2*datasize.GB), -1, pageSize)

	if err != nil {
		return nil, err
	}
	err = m.env.SetOption(mdbx.OptMaxReaders, opts.MaxReaders)
	if err != nil {
		return nil, err
	}
	const MAX_AUGMENT_LIMIT = 0x7fffFFFF
	err = m.env.SetOption(mdbx.OptRpAugmentLimit, MAX_AUGMENT_LIMIT)
	if err != nil {
		return nil, err
	}
	if opts.SyncPeriod != 0 {
		m.env.SetSyncPeriod(opts.SyncPeriod)
	}
	os.MkdirAll(path, 0755)
	err = m.env.Open(path, opts.Flags, 0664)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// init dbi
	m.env.Update(func(txn *mdbx.Txn) error {
		for _, v := range ethdb.Buckets {
			txn.OpenDBI(v, mdbx.Create|mdbx.DupSort, nil, nil)
		}
		return nil
	})

	return m, nil
}

// Has retrieves if a key is present in the key-value data store.
func (m *MDB) Has(key []byte, opts *ethdb.Option) (bool, error) {
	var (
		rs  bool
		err error
	)

	err = m.env.View(func(txn *mdbx.Txn) error {
		dbi, _ := txn.OpenDBI(opts.Name, mdbx.Create|mdbx.DupSort, nil, nil)
		_, err = txn.Get(dbi, key)
		if err != nil {
			if mdbx.IsNotFound(err) {
				rs = false
				return nil
			}
			return err
		}
		rs = true
		return nil
	})

	return rs, err
}

// Get retrieves the given key if it's present in the key-value data store.
func (m *MDB) Get(key []byte, opts *ethdb.Option) ([]byte, error) {
	var (
		rs  []byte
		err error
	)

	err = m.env.View(func(txn *mdbx.Txn) error {
		dbi, _ := txn.OpenDBI(opts.Name, mdbx.Create|mdbx.DupSort, nil, nil)
		rs, err = txn.Get(dbi, key)
		if err != nil {
			if mdbx.IsNotFound(err) {
				return nil
			}
			return err
		}
		return nil
	})

	return rs, err
}

// Put inserts the given value into the key-value data store.
func (m *MDB) Put(key []byte, value []byte, opts *ethdb.Option) error {

	var err error
	err = m.env.Update(func(txn *mdbx.Txn) error {
		dbi, _ := txn.OpenDBI(opts.Name, mdbx.Create|mdbx.DupSort, nil, nil)
		err = txn.Put(dbi, key, value, opts.Flags)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Delete removes the key from the key-value data store.
func (m *MDB) Delete(key []byte, opts *ethdb.Option) error {

	var err error
	err = m.env.Update(func(txn *mdbx.Txn) error {
		dbi, _ := txn.OpenDBI(opts.Name, mdbx.Create|mdbx.DupSort, nil, nil)
		err = txn.Del(dbi, key, nil)
		if err != nil && errors.Is(err, ErrNotFound) {
			return err
		}
		return nil
	})
	return err
}

// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called.
func (m *MDB) NewBatch() ethdb.Batch {
	return &DbBatch{
		db:   m,
		size: 0,
	}
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (m *MDB) NewIterator(prefix []byte, start []byte, opts *ethdb.Option) ethdb.Iterator {
	return &DbIter{
		first: true,
		prefix: func() []byte {
			if bytes.Equal(prefix, []byte("")) {
				return nil
			}
			return prefix
		}(),
		start: func() []byte {
			if bytes.Equal(start, []byte("")) {
				return nil
			}
			return start
		}(),
		curkey: append(prefix, start...),
		db:     m,
		opts:   opts,
	}
}

// Stat returns a particular internal stat of the database.
func (m *MDB) Stat(property string, opts *ethdb.Option) (string, error) {
	return "", nil
}

// Compact flattens the underlying data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (m *MDB) Compact(start []byte, limit []byte, opts *ethdb.Option) error {
	return nil
}

func (m *MDB) Close() error {
	m.env.Close()
	return nil
}
