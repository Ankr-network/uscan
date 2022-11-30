package fulldb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	traceTxPrefix  = []byte("/tracetx/")
	traceTx2Prefix = []byte("/tracetx2/")
)

/*
table: traceLogs

/tracetx/<txhash> => trace tx info
/tracetx2/<txhash> => trace tx2 info
*/

func WriteTraceTx(ctx context.Context, db kv.Writer, hash common.Hash, data *types.TraceTx) (err error) {
	var (
		bytesRes []byte
		key      = append(traceTxPrefix, hash.Bytes()...)
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.TraceLogTbl})
}

func ReadTraceTx(ctx context.Context, db kv.Reader, hash common.Hash) (res *types.TraceTx, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(traceTxPrefix, hash.Bytes()...), &kv.ReadOption{Table: share.TraceLogTbl})
	if err != nil {
		return
	}
	res = &types.TraceTx{}
	err = res.Unmarshal(bytesRes)
	return
}

func WriteTraceTx2(ctx context.Context, db kv.Writer, hash common.Hash, data *types.TraceTx2) (err error) {
	var (
		bytesRes []byte
		key      = append(traceTx2Prefix, hash.Bytes()...)
	)
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, key, bytesRes, &kv.WriteOption{Table: share.TraceLogTbl})
}

func ReadTraceTx2(ctx context.Context, db kv.Reader, hash common.Hash) (res *types.TraceTx2, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(traceTx2Prefix, hash.Bytes()...), &kv.ReadOption{Table: share.TraceLogTbl})
	if err != nil {
		return
	}
	res = &types.TraceTx2{}
	err = res.Unmarshal(bytesRes)
	return
}
