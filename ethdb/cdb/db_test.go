package cdb

import (
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/ethdb/dbtest"
)

func TestMemoryDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			return openTestDb()
		})
	})
}
