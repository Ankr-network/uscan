package fulldb

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"testing"
)

func TestReadBlock(t *testing.T) {
	db := mdbx.NewMdbx("")
	ReadBlock(context.Background(), db, nil)
}
