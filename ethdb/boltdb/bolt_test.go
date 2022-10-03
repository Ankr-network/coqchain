package boltdb

import (
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/ethdb/dbtest"
)

func TestBoltDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			db, err := NewBoltDB(t.TempDir())
			if err != nil {
				t.Error(err)
			}
			return db
		})
	})
}
