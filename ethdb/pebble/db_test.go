package pebble

import (
	"os"
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/ethdb/dbtest"
)

func TestPebble(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			os.RemoveAll("tmp")
			pdb, err := New("tmp", 4096, "", false)
			if err != nil {
				t.Fatal(err)
			}
			return pdb
		})
	})
}
