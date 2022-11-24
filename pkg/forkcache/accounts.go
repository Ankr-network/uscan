package forkcache

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	addressKeyPrefix = []byte("/fork/info/")
)

/*
table: accounts

/fork/info/<address> => account info

/fork/<address>/tx/total => num
/fork/<address>/tx/<index> => <txhash>

/fork/<address>/itx/total => num
/fork/<address>/itx/<index> => InternalTxKey{txhash,index}

/fork/<address>/erc20/total => num
/fork/<address>/erc20/<index> => <index>(erc20 transfer index)

/fork/<address>/erc721/total => num
/fork/<address>/erc721/<index> => <index>(erc721 transfer index)

/fork/<address>/erc1155/total => num
/fork/<address>/erc1155/<index> => <index>(erc1155 transfer index)
*/

// ----------------- account info -----------------

func ReadAccount(ctx context.Context, db kv.Reader, addr common.Address) (acc *types.Account, err error) {
	var bytesRes []byte

	bytesRes, err = db.Get(ctx, append(addressKeyPrefix, addr.Bytes()...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.Account{}
	err = acc.Unmarshal(bytesRes)
	if err == nil {
		acc.Owner = addr
	}
	return
}

func WriteAccount(ctx context.Context, db kv.Database, addr common.Address, acc *types.Account) error {
	bytesRes, err := acc.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, append(addressKeyPrefix, addr.Bytes()...), bytesRes, &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func DeleteAccount(ctx context.Context, db kv.Database, addr common.Address) error {
	return db.Del(ctx, append(addressKeyPrefix, addr.Bytes()...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

// ----------------- tx ----------------

func WriteAccountTxTotal(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...), total.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountTxTotal(ctx context.Context, db kv.Reader, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteAccountTxTotal(ctx context.Context, db kv.Writer, addr common.Address) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func WriteAccountTxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt, hash common.Hash) error {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), hash.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountTxIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (hash common.Hash, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	hash.SetBytes(bytesRes)
	return
}

func DeleteAccountTxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) error {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountTxByIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (tx *types.Tx, err error) {
	var hashByte []byte
	hashByte, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	hash := common.BytesToHash(hashByte)
	return ReadTx(ctx, db, hash)
}

// ------------ internal tx -------------

func WriteAccountITxTotal(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/total")...), total.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountITxTotal(ctx context.Context, db kv.Reader, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteAccountITxTotal(ctx context.Context, db kv.Writer, addr common.Address) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/total")...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func WriteAccountITxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt, data *types.InternalTxKey) (err error) {
	var bytesRes []byte
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), bytesRes, &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountITxIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (data *types.InternalTxKey, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	data = &types.InternalTxKey{}
	err = data.Unmarshal(bytesRes)
	return
}

func DeleteAccountITxIndex(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountITxByIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (itx *types.InternalTx, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	data := &types.InternalTxKey{}
	err = data.Unmarshal(bytesRes)
	return ReadITx(ctx, db, data.TransactionHash, &data.Index)
}

//  ---------------- erc20 transfer ---------------

func WriteAccountErc20Total(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...), total.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc20Total(ctx context.Context, db kv.Reader, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteAccountErc20Total(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func WriteAccountErc20Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt, erc20TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), erc20TransferIndex.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc20Index(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (erc20TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc20TransferIndex = &field.BigInt{}
	erc20TransferIndex.SetBytes(bytesRes)
	return
}

func DeleteAccountErc20Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc20ByIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc20TransferIndex := &field.BigInt{}
	erc20TransferIndex.SetBytes(bytesRes)

	return ReadErc20Transfer(ctx, db, erc20TransferIndex)
}

//  ---------------- erc721 transfer ---------------

func WriteAccountErc721Total(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...), total.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc721Total(ctx context.Context, db kv.Reader, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteAccountErc721Total(ctx context.Context, db kv.Writer, addr common.Address) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func WriteAccountErc721Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt, erc721TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), erc721TransferIndex.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc721Index(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (erc721TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc721TransferIndex = &field.BigInt{}
	erc721TransferIndex.SetBytes(bytesRes)
	return
}

func DeleteAccountErc721Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc721ByIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc721TransferIndex := &field.BigInt{}
	erc721TransferIndex.SetBytes(bytesRes)

	return ReadErc721Transfer(ctx, db, erc721TransferIndex)
}

//  ---------------- erc115 transfer ---------------

func WriteAccountErc1155Total(ctx context.Context, db kv.Writer, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...), total.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc1155Total(ctx context.Context, db kv.Reader, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func DeleteAccountErc1155Total(ctx context.Context, db kv.Writer, addr common.Address) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func WriteAccountErc1155Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt, erc1155TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), erc1155TransferIndex.Bytes(), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc1155Index(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (erc1155TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc1155TransferIndex = &field.BigInt{}
	erc1155TransferIndex.SetBytes(bytesRes)
	return
}

func DeleteAccountErc1155Index(ctx context.Context, db kv.Writer, addr common.Address, index *field.BigInt) (err error) {
	return db.Del(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), &kv.WriteOption{Table: share.ForkAccountsTbl})
}

func ReadAccountErc1155ByIndex(ctx context.Context, db kv.Reader, addr common.Address, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}

	erc1155TransferIndex := &field.BigInt{}
	erc1155TransferIndex.SetBytes(bytesRes)

	return ReadErc1155Transfer(ctx, db, erc1155TransferIndex)
}
