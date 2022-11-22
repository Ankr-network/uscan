package forkcache

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	erc20TotalKey   = []byte("/fork/erc20/total")
	erc721TotalKey  = []byte("/fork/erc721/total")
	erc1155TotalKey = []byte("/fork/erc1155/total")

	erc20IndexPrefix   = []byte("/fork/erc20/")
	erc721IndexPrefix  = []byte("/fork/erc721/")
	erc1155IndexPrefix = []byte("/fork/erc1155/")
)

/*
table: transfers

/fork/erc20/total => total
/fork/erc721/total => total
/fork/erc1155/total => total

/fork/erc20/<index> => erc20 transfer info
/fork/erc721/<index> => erc1155 transfer info
/fork/erc1155/<index> => erc1155 transfer info
*/

func WriteErc20Total(ctx context.Context, db kv.Writer, total *field.BigInt) error {
	return db.Put(ctx, erc20TotalKey, total.Bytes(), &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc20Total(ctx context.Context, db kv.Reader) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, erc20TotalKey, &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteErc20Total(ctx context.Context, db kv.Writer) error {
	return db.Del(ctx, erc20TotalKey, &kv.WriteOption{Table: share.TransferTbl})
}

func WriteErc721Total(ctx context.Context, db kv.Writer, total *field.BigInt) error {
	return db.Put(ctx, erc721TotalKey, total.Bytes(), &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc721Total(ctx context.Context, db kv.Reader) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, erc721TotalKey, &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteErc721Total(ctx context.Context, db kv.Writer) error {
	return db.Del(ctx, erc721TotalKey, &kv.WriteOption{Table: share.TransferTbl})
}

func WriteErc1155Total(ctx context.Context, db kv.Writer, total *field.BigInt) error {
	return db.Put(ctx, erc1155TotalKey, total.Bytes(), &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc1155Total(ctx context.Context, db kv.Reader) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, erc1155TotalKey, &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteErc1155Total(ctx context.Context, db kv.Writer) error {
	return db.Del(ctx, erc1155TotalKey, &kv.WriteOption{Table: share.TransferTbl})
}

func WriteErc20Transfer(ctx context.Context, db kv.Writer, index *field.BigInt, data *types.Erc20Transfer) (err error) {
	var bytesRes []byte
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, append(erc20IndexPrefix, index.Bytes()...), bytesRes, &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc20Transfer(ctx context.Context, db kv.Reader, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc20IndexPrefix, index.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	data = &types.Erc20Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func DeleteErc20Transfer(ctx context.Context, db kv.Writer, index *field.BigInt) (err error) {
	return db.Del(ctx, append(erc20IndexPrefix, index.Bytes()...), &kv.WriteOption{Table: share.TransferTbl})
}

func WriteErc721Transfer(ctx context.Context, db kv.Writer, index *field.BigInt, data *types.Erc721Transfer) (err error) {
	var bytesRes []byte
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, append(erc721IndexPrefix, index.Bytes()...), bytesRes, &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc721Transfer(ctx context.Context, db kv.Reader, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc721IndexPrefix, index.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}

	data = &types.Erc721Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func DeleteErc721Transfer(ctx context.Context, db kv.Writer, index *field.BigInt, data *types.Erc721Transfer) (err error) {
	return db.Del(ctx, append(erc721IndexPrefix, index.Bytes()...), &kv.WriteOption{Table: share.TransferTbl})
}

func WriteErc1155Transfer(ctx context.Context, db kv.Writer, index *field.BigInt, data *types.Erc1155Transfer) (err error) {
	var bytesRes []byte
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, append(erc1155IndexPrefix, index.Bytes()...), bytesRes, &kv.WriteOption{Table: share.TransferTbl})
}

func ReadErc1155Transfer(ctx context.Context, db kv.Reader, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc1155IndexPrefix, index.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}

	data = &types.Erc1155Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func DeleteErc1155Transfer(ctx context.Context, db kv.Writer, index *field.BigInt) (err error) {
	return db.Del(ctx, append(erc1155IndexPrefix, index.Bytes()...), &kv.WriteOption{Table: share.TransferTbl})
}
