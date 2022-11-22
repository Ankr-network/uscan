package forkcache

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	traceTxPrefix  = []byte("/fork/tracetx/")
	traceTx2Prefix = []byte("/fork/tracetx2/")
)

/*
table: traceLogs

/fork/tracetx/<txhash> => trace tx info
/fork/tracetx2/<txhash> => trace tx2 info
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

func DeleteTraceTx(ctx context.Context, db kv.Writer, hash common.Hash) (err error) {
	var key = append(traceTxPrefix, hash.Bytes()...)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.TraceLogTbl})
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

func DeleteTraceTx2(ctx context.Context, db kv.Writer, hash common.Hash) (err error) {
	var key = append(traceTx2Prefix, hash.Bytes()...)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.TraceLogTbl})
}
