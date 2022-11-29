package forkcache

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

/*
table: transactions

/fork/iTx/<txhash>/total => total
/fork/iTx/<txhash>/<index> => internal tx info
*/

var (
	iTxPrefix = []byte("/fork/iTx/")
)

func WriteITx(ctx context.Context, db kv.Writer, hash common.Hash, index *field.BigInt, data *types.InternalTx) (err error) {
	var (
		key      = GetITxKey(hash, index)
		bytesRes []byte
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.ForkTxTbl})
}

func ReadITx(ctx context.Context, db kv.Reader, hash common.Hash, index *field.BigInt) (data *types.InternalTx, err error) {
	var (
		key      = GetITxKey(hash, index)
		bytesRes []byte
	)
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		return
	}
	data = &types.InternalTx{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.TransactionHash = hash
	}
	return
}

func WriteItxTotal(ctx context.Context, db kv.Writer, hash common.Hash, total *field.BigInt) (err error) {
	return db.Put(ctx, GetITxTotalKey(hash), total.Bytes(), &kv.WriteOption{Table: share.ForkTxTbl})
}

func ReadITxTotal(ctx context.Context, db kv.Reader, hash common.Hash) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, GetITxTotalKey(hash), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func GetITxKey(hash common.Hash, index *field.BigInt) []byte {
	return append(append(iTxPrefix, hash.Bytes()...), append([]byte("/"), index.Bytes()...)...)
}

func GetITxTotalKey(hash common.Hash) []byte {
	return append(append(iTxPrefix, hash.Bytes()...), []byte("/total")...)
}
