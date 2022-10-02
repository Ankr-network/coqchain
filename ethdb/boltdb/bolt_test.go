package boltdb

import (
	"io/ioutil"
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/ethdb/dbtest"
)

func TestBoltDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			f, _ := ioutil.TempFile("", "*")
			db, err := NewBoltDB(f.Name())
			if err != nil {
				t.Error(err)
			}
			return db
		})
	})
}
