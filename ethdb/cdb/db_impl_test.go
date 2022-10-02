package cdb

import (
	"os"
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/stretchr/testify/assert"
)

const testdb = "testdb"

func openTestDb() *MDB {
	os.RemoveAll("testdb")
	db, err := NewMDB(testdb, nil)
	if err != nil {
		panic(err)
	}
	return db
}
func TestMdbxPut(t *testing.T) {
	db := openTestDb()
	key := []byte("hello")
	val := []byte("val")
	db.Put(key, val, ethdb.StateOption)
	v, err := db.Get(key, ethdb.StateOption)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, val, v, "should be equal")
}
