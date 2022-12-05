package fulldb

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
)

var (
	erc20TotalKey   = []byte("/erc20/total")
	erc721TotalKey  = []byte("/erc721/total")
	erc1155TotalKey = []byte("/erc1155/total")

	erc20IndexPrefix   = []byte("/erc20/")
	erc721IndexPrefix  = []byte("/erc721/")
	erc1155IndexPrefix = []byte("/erc1155/")
)

/*
table: transfers

/erc20/total => total
/erc721/total => total
/erc1155/total => total

/erc20/<index> => erc20 transfer info
/erc721/<index> => erc1155 transfer info
/erc1155/<index> => erc1155 transfer info

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
