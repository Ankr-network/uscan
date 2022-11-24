package mdbx

import (
	"context"
	"os"
	"testing"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/stretchr/testify/assert"
)

func TestMdbx(t *testing.T) {
	var (
		path  = "./db"
		db    = NewMdbx(path)
		ctx   = context.Background()
		key   = []byte("/key")
		value = []byte("value")
		err   error
	)
	os.RemoveAll(path)

	ctx, err = db.BeginTx(context.Background())
	assert.NoError(t, err)

	err = db.Put(ctx, key, value, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)

	val, err := db.Get(ctx, key, &kv.ReadOption{Table: schemas[1]})
	assert.NoError(t, err)
	assert.Equal(t, val, value)

	err = db.Del(ctx, key, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)

	val, err = db.Get(ctx, key, &kv.ReadOption{Table: schemas[1]})
	assert.Error(t, err)
	assert.EqualError(t, err, kv.NotFound.Error())

	db.RollBack(ctx)

	err = db.Put(context.Background(), key, value, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)
	db.Commit(ctx)

	val, err = db.Get(context.Background(), key, &kv.ReadOption{Table: schemas[1]})
	assert.NoError(t, err)
	assert.Equal(t, val, value)

	err = db.Del(context.Background(), key, &kv.WriteOption{Table: schemas[1]})
	assert.NoError(t, err)

	val, err = db.Get(context.Background(), key, &kv.ReadOption{Table: schemas[1]})
	assert.Error(t, err)
	assert.EqualError(t, err, kv.NotFound.Error())
}
