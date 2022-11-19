package memorydb

import (
	"context"
	"testing"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/stretchr/testify/assert"
)

func TestMemoryDb(t *testing.T) {
	db := NewMemoryDb()

	var (
		table = "test"
		key   = []byte("key")
		val   = []byte("value")
	)

	err := db.Put(context.Background(), key, val, &kv.WriteOption{Table: table})
	assert.NoError(t, err)
	exists, err := db.Has(context.Background(), key, &kv.ReadOption{Table: table})
	assert.NoError(t, err)
	assert.True(t, exists)

	v, err := db.Get(context.Background(), key, &kv.ReadOption{Table: table})
	assert.NoError(t, err)
	assert.Equal(t, v, val)

	exists, err = db.Has(context.Background(), key, &kv.ReadOption{Table: "table"})
	assert.NoError(t, err)
	assert.False(t, exists)

	v, err = db.Get(context.Background(), key, &kv.ReadOption{Table: "table"})
	assert.Error(t, err)
	assert.Equal(t, err, kv.NotFound)
	assert.Nil(t, v)

	err = db.Del(context.Background(), key, &kv.WriteOption{Table: table})
	assert.NoError(t, err)

	exists, err = db.Has(context.Background(), key, &kv.ReadOption{Table: table})
	assert.NoError(t, err)
	assert.False(t, exists)

}
