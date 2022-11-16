package rawdb

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

/iTx/<txhash>/total => total
/iTx/<txhash>/<index> => internal tx info
*/

var (
	iTxPrefix []byte = []byte("/iTx/")
)

func WriteITx(ctx context.Context, db kv.Putter, hash common.Hash, index *field.BigInt, data *types.InternalTx) (err error) {
	var (
		key      = GetITxKey(hash, index)
		bytesRes []byte
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.TxTbl})
}

func ReadITx(ctx context.Context, db kv.Getter, hash common.Hash, index *field.BigInt) (data *types.InternalTx, err error) {
	var (
		key      = GetITxKey(hash, index)
		bytesRes []byte
	)
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.TxTbl})
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

func WriteItxTotal(ctx context.Context, db kv.Putter, hash common.Hash, total *field.BigInt) (err error) {
	return db.Put(ctx, GetITxTotalKey(hash), total.Bytes(), &kv.WriteOption{Table: share.TxTbl})
}

func ReadITxTotal(ctx context.Context, db kv.Getter, hash common.Hash) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, GetITxTotalKey(hash), &kv.ReadOption{Table: share.TxTbl})
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
