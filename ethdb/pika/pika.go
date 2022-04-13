package pika

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/go-redis/redis/v8"
	"github.com/sunvim/utils/tools"
)

type PikaDatabase struct {
	ctx             context.Context
	fn              string // filename for reporting
	db              *redis.Client
	getTimer        metrics.Timer // Timer for measuring the database get request counts and latencies
	putTimer        metrics.Timer // Timer for measuring the database put request counts and latencies
	delTimer        metrics.Timer // Timer for measuring the database delete request counts and latencies
	missMeter       metrics.Meter // Meter for measuring the missed database get requests
	readMeter       metrics.Meter // Meter for measuring the database get request data usage
	writeMeter      metrics.Meter // Meter for measuring the database put request data usage
	batchPutTimer   metrics.Timer
	batchWriteTimer metrics.Timer
	batchWriteMeter metrics.Meter

	quitLock sync.Mutex // Mutex protecting the quit channel access

	log log.Logger // Contextual logger tracking the database path
}

var _ ethdb.KeyValueStore = &PikaDatabase{}

const NoExpired = 0

// NewPikaDatabase returns a Pika wrapped object.
func NewPikaDatabase() (*PikaDatabase, error) {
	logger := log.New("database: pika")

	ret := &PikaDatabase{
		fn:  "pika",
		db:  redis.NewClient(&redis.Options{DB: 0, Addr: "127.0.0.1:9221"}),
		log: logger,
		ctx: context.Background(),
	}

	return ret, nil
}

func (p *PikaDatabase) Stat(property string) (string, error) {
	return "", nil
}

func (p *PikaDatabase) Meter(prefix string) {
	if !metrics.Enabled {
		return
	}
	// Initialize all the metrics collector at the requested prefix
	p.getTimer = metrics.NewRegisteredTimer(prefix+"user/gets", nil)
	p.putTimer = metrics.NewRegisteredTimer(prefix+"user/puts", nil)
	p.delTimer = metrics.NewRegisteredTimer(prefix+"user/dels", nil)
	p.missMeter = metrics.NewRegisteredMeter(prefix+"user/misses", nil)
	p.readMeter = metrics.NewRegisteredMeter(prefix+"user/reads", nil)
	p.writeMeter = metrics.NewRegisteredMeter(prefix+"user/writes", nil)
	p.batchPutTimer = metrics.NewRegisteredTimer(prefix+"user/batchPuts", nil)
	p.batchWriteTimer = metrics.NewRegisteredTimer(prefix+"user/batchWriteTimes", nil)
	p.batchWriteMeter = metrics.NewRegisteredMeter(prefix+"user/batchWrites", nil)
}

// Has retrieves if a key is present in the key-value data store.
func (p *PikaDatabase) Has(key []byte) (bool, error) {
	cmd := p.db.Exists(p.ctx, tools.BytesToStringFast(key))
	return (cmd.Val() > 0), cmd.Err()
}

// Get retrieves the given key if it's present in the key-value data store.
func (p *PikaDatabase) Get(key []byte) ([]byte, error) {
	return p.db.Get(p.ctx, tools.BytesToStringFast(key)).Bytes()
}

// Put inserts the given value into the key-value data store.
func (p *PikaDatabase) Put(key []byte, value []byte) error {
	return p.db.Set(p.ctx, tools.BytesToStringFast(key), value, NoExpired).Err()
}

// Delete removes the key from the key-value data store.
func (p *PikaDatabase) Delete(key []byte) error {
	return p.db.Del(p.ctx, tools.BytesToStringFast(key)).Err()
}

// Compact flattens the underlying data store for the given key range. In essence,
// deleted and overwritten versions are discarded, and the data is rearranged to
// reduce the cost of operations needed to access them.
//
// A nil start is treated as a key before all keys in the data store; a nil limit
// is treated as a key after all keys in the data store. If both is nil then it
// will compact entire data store.
func (p *PikaDatabase) Compact(start []byte, limit []byte) error {
	return nil
}

func (p *PikaDatabase) Close() error {
	return p.db.Close()
}

type PikaBatch struct {
	db   *PikaDatabase
	b    map[string][]byte
	size int
}

var _ ethdb.Batch = &PikaBatch{}

// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called.
func (p *PikaDatabase) NewBatch() ethdb.Batch {
	return &PikaBatch{db: p, b: make(map[string][]byte)}
}

// Put inserts the given value into the key-value data store.
func (p *PikaBatch) Put(key []byte, value []byte) error {
	if p.db.batchPutTimer != nil {
		defer p.db.batchPutTimer.UpdateSince(time.Now())
	}

	p.b[tools.BytesToStringFast(key)] = common.CopyBytes(value)
	p.size += len(value)

	return nil
}

// Delete removes the key from the key-value data store.
func (p *PikaBatch) Delete(key []byte) error {
	p.db.db.Del(p.db.ctx, tools.BytesToStringFast(key))
	p.size += len(key)
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (p *PikaBatch) ValueSize() int {
	return p.size
}

// Write flushes any accumulated data to disk.
func (p *PikaBatch) Write() error {

	if p.db.batchWriteTimer != nil {
		defer p.db.batchWriteTimer.UpdateSince(time.Now())
	}

	if p.db.batchWriteMeter != nil {
		p.db.batchWriteMeter.Mark(int64(p.size))
	}

	for k, v := range p.b {
		p.db.db.Set(p.db.ctx, k, v, NoExpired)
	}

	p.size = 0
	p.b = make(map[string][]byte)

	return nil
}

// Reset resets the batch for reuse.
func (p *PikaBatch) Reset() {
	p.b = make(map[string][]byte)
	p.size = 0
}

// Replay replays the batch contents.
func (p *PikaBatch) Replay(w ethdb.KeyValueWriter) error {
	for k, v := range p.b {
		p.db.db.Set(p.db.ctx, k, v, NoExpired)
	}
	return nil
}

type Item struct {
	key string
	val []byte
}

type PikaIterator struct {
	first  bool
	prefix string
	db     *PikaDatabase
	cursor uint64
	b      []string
	item   *Item
	err    error
}

var _ ethdb.Iterator = &PikaIterator{}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (p *PikaDatabase) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	return &PikaIterator{first: true, db: p, b: make([]string, 0), cursor: 0, prefix: tools.BytesToStringFast(append(prefix, start...))}
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (i *PikaIterator) Next() bool {
	if len(i.b) == 0 && i.cursor == 0 && !i.first {
		return false
	}

	if len(i.b) != 0 {
		i.item = &Item{
			key: i.b[0],
		}
		bs, _ := i.db.db.Get(i.db.ctx, i.b[0]).Bytes()
		i.item.val = bs
		i.b = i.b[1:]
		return true
	}

	// take data from remote database
	res := i.db.db.Scan(i.db.ctx, i.cursor, i.prefix, 10)
	keys, cursor := res.Val()
	if len(keys) == 0 {
		return false
	}

	i.cursor = cursor
	i.item = &Item{
		key: keys[0],
	}
	bs, _ := i.db.db.Get(i.db.ctx, keys[0]).Bytes()
	i.item.val = bs
	i.first = false
	i.b = append(i.b, keys[1:]...)

	return true
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (i *PikaIterator) Error() error {
	return i.err
}

// Key returns the key of the current key/value pair, or nil if done. The caller
// should not modify the contents of the returned slice, and its contents may
// change on the next call to Next.
func (i *PikaIterator) Key() []byte {
	return tools.StringToBytes(i.item.key)
}

// Value returns the value of the current key/value pair, or nil if done. The
// caller should not modify the contents of the returned slice, and its contents
// may change on the next call to Next.
func (i *PikaIterator) Value() []byte {
	return i.item.val
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (i *PikaIterator) Release() {
	i.b = nil
	i.item = nil
}
