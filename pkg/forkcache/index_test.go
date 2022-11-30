package forkcache

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"testing"
)

func TestIndex(t *testing.T) {
	db := mdbx.NewMdbx("./uscandb")
	blockNumber := field.NewInt(3232254)
	val, _ := rawdb.ReadBlock(context.Background(), db, blockNumber)
	t.Log(val.Number)
}
