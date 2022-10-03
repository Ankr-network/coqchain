package boltdb

import (
	"bytes"
	"testing"

	"github.com/Ankr-network/coqchain/ethdb"
)

func openTestdb(t *testing.T) *BoltDB {
	db, err := NewBoltDB(t.TempDir())
	if err != nil {
		panic(err)
	}
	return db
}

func TestIter(t *testing.T) {
	db := openTestdb(t)
	tests := []struct {
		content map[string]string
		prefix  string
		start   string
		order   []string
	}{
		{map[string]string{"key": "val"}, "k", "", []string{"key"}},
	}

	for i, tt := range tests {
		for key, val := range tt.content {
			if err := db.Put([]byte(key), []byte(val), ethdb.GlobalDataOption); err != nil {
				t.Fatalf("test %d: failed to insert item %s:%s into database: %v", i, key, val, err)
			}
		}
		// Iterate over the database with the given configs and verify the results
		it, idx := db.NewIterator([]byte(tt.prefix), []byte(tt.start), ethdb.GlobalDataOption), 0
		for it.Next() {
			if len(tt.order) <= idx {
				t.Errorf("test %d: prefix=%q more items than expected: checking idx=%d (key %q), expecting len=%d", i, tt.prefix, idx, it.Key(), len(tt.order))
				break
			}
			if !bytes.Equal(it.Key(), []byte(tt.order[idx])) {
				t.Errorf("test %d: item %d: key mismatch: have %s, want %s", i, idx, string(it.Key()), tt.order[idx])
			}
			if !bytes.Equal(it.Value(), []byte(tt.content[tt.order[idx]])) {
				t.Errorf("test %d: item %d: value mismatch: have %s, want %s", i, idx, string(it.Value()), tt.content[tt.order[idx]])
			}
			idx++
		}
		if err := it.Error(); err != nil {
			t.Errorf("test %d: iteration failed: %v", i, err)
		}
		if idx != len(tt.order) {
			t.Errorf("test %d: iteration terminated prematurely: have %d, want %d", i, idx, len(tt.order))
		}
		db.Close()
	}

}
