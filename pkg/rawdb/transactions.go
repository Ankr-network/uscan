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
	txKey      []byte = []byte("/tx/")
	rtKey      []byte = []byte("/rt/")
	txNumKey   []byte = []byte("/all/tx/total")
	txIndexKey []byte = []byte("/all/tx/")
)

/*
table : transactions

/tx/<txhash>          => tx info
/rt/<txhash>          => rt info
/all/tx/total => total
/all/tx/<index> => <txhash>
*/
func WriteTx(ctx context.Context, db kv.Putter, hash common.Hash, data *types.Tx) (err error) {
	var (
		key      = append(txKey, hash.Bytes()...)
		bytesRes []byte
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.TxTbl})
}

func ReadTx(ctx context.Context, db kv.Getter, hash common.Hash) (data *types.Tx, err error) {
	var (
		key      = append(txKey, hash.Bytes()...)
		bytesRes []byte
	)

	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		return
	}
	data = &types.Tx{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.Hash = hash
	}
	return
}

func WriteTxIndex(ctx context.Context, db kv.Putter, index *field.BigInt, hash common.Hash) error {
	return db.Put(ctx, append(txIndexKey, index.Bytes()...), hash.Bytes(), &kv.WriteOption{Table: share.TxTbl})
}

func ReadTxByIndex(ctx context.Context, db kv.Getter, index *field.BigInt) (data *types.Tx, err error) {
	var hashByte []byte
	hashByte, err = db.Get(ctx, append(txIndexKey, index.Bytes()...), &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		return
	}
	hash := common.BytesToHash(hashByte)
	return ReadTx(ctx, db, hash)
}

func WriteRt(ctx context.Context, db kv.Putter, hash common.Hash, data *types.Rt) (err error) {
	var (
		key      = append(rtKey, hash.Bytes()...)
		bytesRes []byte
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.TxTbl})
}

func ReadRt(ctx context.Context, db kv.Getter, hash common.Hash) (data *types.Rt, err error) {
	var (
		key      = append(rtKey, hash.Bytes()...)
		bytesRes []byte
	)
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		return
	}
	data = &types.Rt{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.TxHash = hash
	}
	return
}
