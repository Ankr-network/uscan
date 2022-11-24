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
	// key = value
	/fork/erc20/<contract>/<address> => <amount>

	/fork/erc721/<contract>/<address> => <amount>  # total nft
	/fork/erc721/<contract>/<address>/<tokenId> => <Quantity>

	/fork/erc1155/<contract>/<address> => <amount>  # total nft
	/fork/erc1155/<contract>/<address>/<tokenId> => <Quantity>
*/

var (
	erc20HolderPrefix   = []byte("/fork/erc20/")
	erc721HolderPrefix  = []byte("/fork/erc721/")
	erc1155HolderPrefix = []byte("/fork/erc1155/")
)

// / ------------------- erc20 holder ---------------------

func WriteErc20HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc20HolderKey(contract, holder.Addr)
	return db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
}

func ReadErc20HolderAmount(ctx context.Context, db kv.Reader, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	var key = getErc20HolderKey(contract, addr)
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func DeleteErc20HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc20HolderKey(contract, holder.Addr)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
}

// / ------------------- erc721 holder ---------------------

func WriteErc721HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc721HolderKey(contract, holder.Addr)
	return db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
}

func ReadErc721HolderAmount(ctx context.Context, db kv.Reader, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	var key = getErc721HolderKey(contract, addr)
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func DeleteErc721HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc721HolderKey(contract, holder.Addr)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
}

func WriteErc721HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc721TokenIdHolderKey(contract, addr, tokenId)
	return db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
}

func ReadErc721HolderTokenIdQuantity(ctx context.Context, db kv.Reader, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error) {
	var key = getErc721TokenIdHolderKey(contract, addr, tokenId)
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return
	}
	quantity = &field.BigInt{}
	quantity.SetBytes(bytesRes)
	return
}

func DeleteErc721HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt) (err error) {
	var key = getErc721TokenIdHolderKey(contract, addr, tokenId)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
}

// / ------------------- erc1155 holder ---------------------

func WriteErc1155HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc1155HolderKey(contract, holder.Addr)
	return db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
}

func ReadErc1155HolderAmount(ctx context.Context, db kv.Reader, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	var key = getErc1155HolderKey(contract, addr)
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func DeleteErc1155HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc1155HolderKey(contract, holder.Addr)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
}

func WriteErc1155HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc1155TokenIdHolderKey(contract, addr, tokenId)
	return db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
}

func ReadErc1155HolderTokenIdQuantity(ctx context.Context, db kv.Reader, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error) {
	var key = getErc1155TokenIdHolderKey(contract, addr, tokenId)
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return
	}
	quantity = &field.BigInt{}
	quantity.SetBytes(bytesRes)
	return
}

func DeleteErc1155HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt) (err error) {
	var key = getErc1155TokenIdHolderKey(contract, addr, tokenId)
	return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
}

// --------------------- get key  -----------------

func getErc20HolderKey(contract common.Address, addr common.Address) []byte {
	key := make([]byte, 0, len(erc20HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	return key
}

func getErc721HolderKey(contract common.Address, addr common.Address) []byte {
	key := make([]byte, 0, len(erc721HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	return key
}

func getErc1155HolderKey(contract common.Address, addr common.Address) []byte {
	key := make([]byte, 0, len(erc1155HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	return key
}

func getErc721TokenIdHolderKey(contract common.Address, addr common.Address, tokenId *field.BigInt) []byte {
	key := make([]byte, 0, len(erc721HolderPrefix)+common.AddressLength*2+2+len(tokenId.Bytes()))
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	key = append(key, '/')
	key = append(key, tokenId.Bytes()...)
	return key
}

func getErc1155TokenIdHolderKey(contract common.Address, addr common.Address, tokenId *field.BigInt) []byte {
	key := make([]byte, 0, len(erc1155HolderPrefix)+common.AddressLength*2+2+len(tokenId.Bytes()))
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	key = append(key, '/')
	key = append(key, tokenId.Bytes()...)
	return key
}
