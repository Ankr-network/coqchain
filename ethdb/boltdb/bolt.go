package boltdb

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/utils"
	"go.etcd.io/bbolt"
)

type BoltDB struct {
	path string
	db   *bbolt.DB
}

func NewBoltDB(path string) (*BoltDB, error) {

	os.MkdirAll(path, 0755)

	d := &BoltDB{path: filepath.Join(path, "blockchain.dat")}

	opt := &bbolt.Options{
		Timeout:         0,
		NoGrowSync:      false,
		NoSync:          false,
		MmapFlags:       syscall.MAP_POPULATE,
		PageSize:        1 << 16,
		InitialMmapSize: 1 << 31,
		FreelistType:    bbolt.FreelistMapType,
	}

	db, err := bbolt.Open(d.path, 0664, opt)
	if err != nil {
		return nil, err
	}

	// init bucket
	db.Update(func(tx *bbolt.Tx) error {
		for _, v := range ethdb.Buckets {
			tx.CreateBucketIfNotExists(utils.S2B(v))
		}
		return nil
	})

	d.db = db

	return d, nil
}

func (d *BoltDB) Path() string {
	return d.path
}

// Has retrieves if a key is present in the key-value data store.
func (d *BoltDB) Has(key []byte, opts *ethdb.Option) (bool, error) {
	var (
		err error
		rs  bool
	)
	err = d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(opts.Name))
		val := b.Get(key)
		if val == nil {
			rs = false
		}
		rs = true
		return nil
	})
	return rs, err
}

// Get retrieves the given key if it's present in the key-value data store.
func (d *BoltDB) Get(key []byte, opts *ethdb.Option) ([]byte, error) {
	var (
		rs []byte
	)
	d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(opts.Name))
		rs = b.Get(key)
		return nil
	})
	if rs == nil {
		return rs, ErrNotFound
	}
	return rs, nil
}

// Put inserts the given value into the key-value data store.
func (d *BoltDB) Put(key []byte, value []byte, opts *ethdb.Option) error {
	var err error
	err = d.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(opts.Name))
		return b.Put(key, value)
	})
	return err
}

// Delete removes the key from the key-value data store.
func (d *BoltDB) Delete(key []byte, opts *ethdb.Option) error {
	var err error
	err = d.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(utils.S2B(opts.Name))
		return b.Delete(key)
	})
	return err
}

// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called.
func (d *BoltDB) NewBatch() ethdb.Batch {
	return &Batch{
		db: d,
	}
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (d *BoltDB) NewIterator(prefix []byte, start []byte, opts *ethdb.Option) ethdb.Iterator {
	curkey := append(prefix, start...)
	return &Iter{
		prefix:  curkey,
		nextkey: curkey,
		opts:    opts,
		db:      d,
	}
}

// Stat returns a particular internal stat of the database.
func (d *BoltDB) Stat(property string, opts *ethdb.Option) (string, error) {
	return "", nil
}

// Compact flattens the underlying data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (d *BoltDB) Compact(start []byte, limit []byte, opts *ethdb.Option) error {
	return nil
}

func (d *BoltDB) Close() error {
	return d.db.Close()
}
