package rawdb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	blockKey []byte = []byte("/block/")
)

/*
table: blocks

/block/<block num> => block info
/block/<block num>/index => tx hash
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

func ReadBlockIndex(ctx context.Context, db kv.Reader, blockNum *field.BigInt, index *field.BigInt) (txHash common.Hash, err error) {
	var (
		bytesRes []byte
	)

	bytesRes, err = db.Get(ctx, getBlockIndex(blockNum, index), &kv.ReadOption{Table: share.BlockTbl})
	if err != nil {
		return
	}
	txHash.SetBytes(bytesRes)
	return
}

func WriteBlockIndex(ctx context.Context, db kv.Writer, blockNum *field.BigInt, index *field.BigInt, txHash common.Hash) (err error) {
	return db.Put(ctx, getBlockIndex(blockNum, index), txHash.Bytes(), &kv.WriteOption{Table: share.BlockTbl})
}

func getBlockIndex(blockNum *field.BigInt, index *field.BigInt) []byte {
	key := make([]byte, 0, len(blockKey)+len(blockNum.Bytes())+len(index.Bytes())+1)

	key = append(key, blockKey...)
	key = append(key, blockNum.Bytes()...)
	key = append(key, byte('/'))
	key = append(key, index.Bytes()...)
	return key
}
