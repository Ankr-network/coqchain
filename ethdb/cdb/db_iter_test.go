package cdb

import (
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
	"github.com/Ankr-network/coqchain/utils"
)

func TestIter(t *testing.T) {
	db := openTestDb()
	ks := map[string]interface{}{
		"1": nil,
		"2": nil,
		"3": nil,
		"4": nil,
		"6": nil,
		"7": nil,
		"8": nil,
		"9": nil,
	}
	for k := range ks {
		db.Put(utils.S2B(k), nil, ethdb.StateOption)
	}

	it := db.NewIterator([]byte(""), []byte("5"), ethdb.StateOption)
	for it.Next() {
		t.Logf("key: %s val: %s \n", it.Key(), it.Value())
	}
}
