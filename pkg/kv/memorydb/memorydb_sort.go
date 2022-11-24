package memorydb

import (
	"bytes"
	"context"

	"github.com/Ankr-network/uscan/pkg/kv"
)

var _ kv.Sorter = (*Database)(nil)

func (db *Database) SPut(ctx context.Context, key, val []byte, opts *kv.WriteOption) error {
	_, ok := db.dbList[opts.Table]
	if !ok {
		db.dbList[opts.Table] = make([][]byte, 0, 1)
	}
	db.dbList[opts.Table] = append(db.dbList[opts.Table], val)
	return nil
}

func (db *Database) SDel(ctx context.Context, key, val []byte, opts *kv.WriteOption) error {
	_, ok := db.dbList[opts.Table]
	if ok {
		var vals = make([][]byte, 0, len(db.dbList[opts.Table]))
		for _, v := range db.dbList[opts.Table] {
			if !bytes.Equal(v, val) {
				vals = append(vals, v)
			}
		}
		db.dbList[opts.Table] = vals
	}
	return nil
}

func (db *Database) SCount(ctx context.Context, key []byte, opts *kv.ReadOption) (uint64, error) {
	_, ok := db.dbList[opts.Table]
	if ok {
		return uint64(len(db.dbList[opts.Table])), nil
	}
	return uint64(0), nil
}

func (db *Database) SGet(ctx context.Context, key []byte, page, pageSize uint64, opts *kv.ReadOption) ([][]byte, error) {
	return nil, nil
}
