package rawdb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	homeKey    = []byte("/home")
	syncingKey = []byte("/syncing")
)

/*
table: home

/home => home
/syncing => block number
*/

func ReadHome(ctx context.Context, db kv.Getter) (home *types.Home, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, homeKey, &kv.ReadOption{Table: share.HomeTbl})
	if err != nil {
		return
	}
	home = &types.Home{}
	err = home.Unmarshal(bytesRes)
	if err == nil && home.DateTxs == nil {
		home.DateTxs = make(map[string]*field.BigInt)
	}
	return
}

func WriteHome(ctx context.Context, db kv.Putter, home *types.Home) (err error) {
	var bytesRes []byte
	bytesRes, err = home.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, homeKey, bytesRes, &kv.WriteOption{Table: share.HomeTbl})
}

func ReadSyncingBlock(ctx context.Context, db kv.Getter) (bk *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, syncingKey, &kv.ReadOption{Table: share.HomeTbl})

	if err != nil {
		return
	}
	bk = &field.BigInt{}
	bk.SetBytes(bytesRes)
	return
}

func WriteSyncingBlock(ctx context.Context, db kv.Putter, bk *field.BigInt) (err error) {
	return db.Put(ctx, syncingKey, bk.Bytes(), &kv.WriteOption{Table: share.HomeTbl})
}
