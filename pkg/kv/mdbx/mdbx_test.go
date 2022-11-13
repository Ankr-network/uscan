package mdbx

import (
	"context"
	"testing"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/stretchr/testify/assert"
)

func TestMdbx(t *testing.T) {
	var (
		path  = "./"
		db    = NewMdbx(path)
		ctx   = context.Background()
		key   = []byte("/key")
		value = []byte("value")
		err   error
	)

	ctx, err = db.BeginTx(context.Background())
	assert.NoError(t, err)

	err = db.Put(ctx, key, value, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)

	val, err := db.Get(ctx, key, &kv.ReadOption{Table: schemas[1]})
	assert.NoError(t, err)
	assert.Equal(t, val, value)
	db.RollBack(ctx)

	ctx, err = db.BeginTx(context.Background())
	assert.NoError(t, err)

	err = db.Put(ctx, key, value, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)
	db.Commit(ctx)

	val, err = db.Get(context.Background(), key, &kv.ReadOption{Table: schemas[1]})
	t.Log(val)
	t.Log(err)
}
