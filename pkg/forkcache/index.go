package forkcache

import (
	"context"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/share"
)

/*
table: index

/fork/<address>/tx/index => index
/fork/<address>/itx/index => index
/fork/<address>/erc20/index => index
/fork/<address>/erc721/index => index
/fork/<address>/erc1155/index => index
/fork/iTx/<txhash>/index => index
/fork/all/tx/index => index
/fork/erc20/index => index
/fork/erc721/index => index
/fork/erc1155/index => index
/fork/erc20/<contract>/<index> => <erc20 total index>
/fork/erc721/<contract>/<index> => <erc721 total index>
/fork/erc1155/<contract>/<index> => <erc1155 total index>
*/

func ReadAddressTxIndex(ctx context.Context, db kv.Reader, addr common.Address) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/index")...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteAddressTxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/index")...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadAddressITxIndex(ctx context.Context, db kv.Reader, addr common.Address) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/index")...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteAddressITxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/index")...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadAddressErc20Index(ctx context.Context, db kv.Reader, addr common.Address) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/index")...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteAddressErc20Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/index")...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadAddressErc721Index(ctx context.Context, db kv.Reader, addr common.Address) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/index")...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteAddressErc721Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/index")...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadAddressErc1155Index(ctx context.Context, db kv.Reader, addr common.Address) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/index")...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteAddressErc1155Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/index")...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadITxIndex(ctx context.Context, db kv.Reader, hash common.Hash) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append(iTxPrefix, hash.Bytes()...), append([]byte("/index"))...), &kv.ReadOption{Table: share.ForkIndexTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteITxIndex(ctx context.Context, db kv.Writer, hash common.Hash, index *field.BigInt) (err error) {
	return db.Put(ctx, append(append(iTxPrefix, hash.Bytes()...), append([]byte("/index"))...), index.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
}

func ReadTxTotalIndex(ctx context.Context, db kv.Reader) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, []byte("/fork/all/tx/index"), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		return
	}
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteTxTotalIndex(ctx context.Context, db kv.Writer, index *field.BigInt) error {
	return db.Put(ctx, []byte("/fork/all/tx/index"), index.Bytes(), &kv.WriteOption{Table: share.ForkTxTbl})
}

func ReadErc20Index(ctx context.Context, db kv.Reader) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc20IndexPrefix, append([]byte("/index"))...), &kv.ReadOption{Table: share.ForkTxTbl})
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteErc20Index(ctx context.Context, db kv.Writer, index *field.BigInt) error {
	return db.Put(ctx, append(erc20IndexPrefix, append([]byte("/index"))...), index.Bytes(), &kv.WriteOption{Table: share.ForkTxTbl})
}

func ReadErc721Index(ctx context.Context, db kv.Reader) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc721IndexPrefix, append([]byte("/index"))...), &kv.ReadOption{Table: share.ForkTxTbl})
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteErc721Index(ctx context.Context, db kv.Writer, index *field.BigInt) error {
	return db.Put(ctx, append(erc721IndexPrefix, append([]byte("/index"))...), index.Bytes(), &kv.WriteOption{Table: share.ForkTxTbl})
}

func ReadErc1155Index(ctx context.Context, db kv.Reader) (index *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(erc1155IndexPrefix, append([]byte("/index"))...), &kv.ReadOption{Table: share.ForkTxTbl})
	index = &field.BigInt{}
	index.SetBytes(bytesRes)
	return
}

func WriteErc1155Index(ctx context.Context, db kv.Writer, index *field.BigInt) error {
	return db.Put(ctx, append(erc1155IndexPrefix, append([]byte("/index"))...), index.Bytes(), &kv.WriteOption{Table: share.ForkTxTbl})
}
