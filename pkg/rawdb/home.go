package rawdb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	homeKey      = []byte("/home")
	syncingKey   = []byte("/syncing")
	homeData     *types.Home
	syncingBlock *field.BigInt
)

/*
table: home

/home => home
/syncing => block number
*/

func ReadHome(ctx context.Context, db kv.Getter) (home *types.Home, err error) {
	if homeData != nil {
		return homeData, nil
	}
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, homeKey, &kv.ReadOption{Table: share.HomeTbl})
	if err != nil {
		return
	}
	home = &types.Home{}
	err = home.Unmarshal(bytesRes)
	if err == nil {
		home.BlockNumber = syncingBlock
		homeData = home
	}
	return
}

func WriteHome(ctx context.Context, db kv.Putter, home *types.Home) (err error) {
	var bytesRes []byte
	bytesRes, err = home.Marshal()
	if err != nil {
		return
	}
	err = db.Put(ctx, homeKey, bytesRes, &kv.WriteOption{Table: share.HomeTbl})
	homeData = home
	return
}

func ReadSyncingBlock(ctx context.Context, db kv.Getter) (bk *field.BigInt, err error) {
	if syncingBlock != nil {
		return syncingBlock, nil
	}
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, syncingKey, &kv.ReadOption{Table: share.HomeTbl})
	if err != nil {
		return
	}
	bk = &field.BigInt{}
	bk.SetBytes(bytesRes)
	syncingBlock = bk
	return
}

func WriteSyncingBlock(ctx context.Context, db kv.Putter, bk *field.BigInt) (err error) {
	err = db.Put(ctx, homeKey, bk.Bytes(), &kv.WriteOption{Table: share.HomeTbl})
	syncingBlock = bk
	return
}
