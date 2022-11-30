package forkdb

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/fulldb"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"testing"
)

func TestIndex(t *testing.T) {
	db := mdbx.NewMdbx("./uscandb")
	blockNumber := field.NewInt(3232254)
	val, _ := fulldb.ReadBlock(context.Background(), db, blockNumber)
	t.Log(val.Number)
}
