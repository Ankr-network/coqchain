package cdb

import (
	"errors"
	"os"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/log"
	"github.com/c2h5oh/datasize"
	"github.com/go-stack/stack"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

type MDB struct {
	path    string
	opts    *Option
	env     *mdbx.Env
	buckets map[string]mdbx.DBI
}

func NewMDB(path string, opts *Option) (*MDB, error) {
	var err error
	m := &MDB{
		buckets: make(map[string]mdbx.DBI, 32),
	}

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
	os.MkdirAll(path, 0744)
	err = m.env.Open(path, opts.Flags, 0664)
	if err != nil {
		return nil, err
	}
	// init bucket
	var dbi mdbx.DBI
	m.buckets[ethdb.StorageStateFmt] = dbi
	err = m.env.Update(func(txn *mdbx.Txn) error {
		var dbi mdbx.DBI
		for bucketName := range m.buckets {
			dbi, err = txn.CreateDBI(bucketName)
			if err != nil {
				log.Error("NewMDB", "err", err, "trace", stack.Trace().String())
				return err
			}
			m.buckets[bucketName] = dbi
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Has retrieves if a key is present in the key-value data store.
func (m *MDB) Has(key []byte, opts *ethdb.Option) (bool, error) {
	var (
		rs  bool
		err error
	)

	err = m.env.View(func(txn *mdbx.Txn) error {
		_, err = txn.Get(m.buckets[opts.Name], key)
		if err != nil {
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
		rs, err = txn.Get(m.buckets[opts.Name], key)
		if err != nil && errors.Is(err, ErrNotFound) {
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
		err = txn.Put(m.buckets[opts.Name], key, value, opts.Flags)
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
		err = txn.Del(m.buckets[opts.Name], key, nil)
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
	for dbiName := range m.buckets {
		m.env.CloseDBI(m.buckets[dbiName])
	}
	println("close mdb start")
	m.env.Close()
	println("close mdb end")
	return nil
}
