package rawdb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	blockKey []byte = []byte("/block/")
)

/*
table: blocks

/block/<block num> => block info
*/

func ReadBlock(ctx context.Context, db kv.Getter, blockNum *field.BigInt) (bk *types.Block, err error) {
	var (
		key      = append(blockKey, blockNum.Bytes()...)
		bytesRes []byte
	)

	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.BlockTbl})
	if err != nil {
		return
	}

	bk = &types.Block{}
	err = bk.Unmarshal(bytesRes)
	if err == nil {
		bk.Number = blockNum
	}
	return
}

func WriteBlock(ctx context.Context, db kv.Putter, blockNum *field.BigInt, bk *types.Block) (err error) {
	var (
		key      = append(blockKey, blockNum.Bytes()...)
		bytesRes []byte
	)

	bytesRes, err = bk.Marshal()
	if err != nil {
		return
	}

	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.BlockTbl})
}
