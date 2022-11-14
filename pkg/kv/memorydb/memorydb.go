package memorydb

import (
	"context"
	"sync"

	"github.com/Ankr-network/uscan/pkg/kv"
)

var _ kv.Database = (*Database)(nil)

type Database struct {
	db   map[string]map[string][]byte
	lock sync.RWMutex
}

func NewMemoryDb() *Database {
	return &Database{
		db: make(map[string]map[string][]byte),
	}
}

func (db *Database) BeginTx(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
func (db *Database) Commit(context.Context)   {}
func (db *Database) RollBack(context.Context) {}
func (db *Database) Put(ctx context.Context, key, val []byte, opts *kv.WriteOption) error {
	_, ok := db.db[opts.Table]
	if !ok {
		db.db[opts.Table] = make(map[string][]byte)
	}
	db.db[opts.Table][string(key)] = val
	return nil
}

func (db *Database) Has(ctx context.Context, key []byte, opts *kv.ReadOption) (bool, error) {
	data, ok := db.db[opts.Table]
	if ok {
		_, exists := data[string(key)]
		return exists, nil
	}
	return false, nil
}
func (db *Database) Get(ctx context.Context, key []byte, opts *kv.ReadOption) ([]byte, error) {
	data, ok := db.db[opts.Table]
	if ok {
		val, exists := data[string(key)]
		if exists {
			return val, nil
		}
		return nil, kv.NotFound
	}
	return nil, kv.NotFound
}
func (db *Database) Close() error { return nil }
