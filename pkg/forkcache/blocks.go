package forkcache

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	blockKey = []byte("/fork/block/")
)

/*
table: blocks

/fork/block/<block num> => block info
*/

func ReadBlock(ctx context.Context, db kv.Reader, blockNum *field.BigInt) (bk *types.Block, err error) {
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
		bk.Number.SetBytes(blockNum.Bytes())
	}
	return
}

func WriteBlock(ctx context.Context, db kv.Writer, blockNum *field.BigInt, bk *types.Block) (err error) {
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

func DeleteBlock(ctx context.Context, db kv.Writer, blockNum *field.BigInt) (err error) {
	var key = append(blockKey, blockNum.Bytes()...)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.BlockTbl})
}
