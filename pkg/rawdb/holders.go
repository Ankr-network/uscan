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
	// key = value
	/erc20/<contract>/<address> => <amount>

	/erc721/<contract>/<address> => <amount>  # total nft
	/erc721/<contract>/<address>/<tokenId> => <Quantity>

	/erc1155/<contract>/<address> => <amount>  # total nft
	/erc1155/<contract>/<address>/<tokenId> => <Quantity>
*/

/*
	// key = > sort
	// holder
	/erc20/<contract> => value + address
	/erc721/<contract> => value + address
	/erc1155/<contract> => value + address

	// inventory
	/erc721/<contract>/tokenId => tokenId + address
	/erc1155/<contract>/tokenId => tokenId
*/

var (
	erc20HolderPrefix   = []byte("/erc20/")
	erc721HolderPrefix  = []byte("/erc721/")
	erc1155HolderPrefix = []byte("/erc1155/")
)

/// ------------------- erc20 holder ---------------------
func DelErc20HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc20HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc20Holder(ctx context.Context, db kv.Sorter, contract common.Address, page, pageSize uint64) (holders []*types.Holder, err error) {
	var key = append(erc20HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, page, pageSize, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return nil, err
	}
	holders = make([]*types.Holder, len(res))
	for i, v := range res {
		holders[i], err = types.ByteToHolder(v)
		if err != nil {
			return nil, err
		}
	}
	return
}

func WriteErc20HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc20HolderKey(contract, holder.Addr)
	err = db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
	if err == nil {
		var (
			sortKey = append(erc20HolderPrefix, contract.Bytes()...)
		)
		return db.SPut(ctx, sortKey, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
	}
	return
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

/// ------------------- erc721 holder ---------------------
func DelErc721HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc721HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc721Holder(ctx context.Context, db kv.Sorter, contract common.Address, page, pageSize uint64) (holders []*types.Holder, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, page, pageSize, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return nil, err
	}
	holders = make([]*types.Holder, len(res))
	for i, v := range res {
		holders[i], err = types.ByteToHolder(v)
		if err != nil {
			return nil, err
		}
	}
	return
}

// func DelErc721InventoryTokenId(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
// 	var (
// 		key = append(erc721HolderPrefix, contract.Bytes()...)
// 	)
// 	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
// }

// func GetErc721Inventory(ctx context.Context, db kv.Sorter, contract common.Address, page, pageSize uint64) (holders []*types.Holder, err error) {
// 	var key = append(erc721HolderPrefix, contract.Bytes()...)
// 	var res [][]byte
// 	res, err = db.SGet(ctx, key, page, pageSize, &kv.ReadOption{Table: share.HolderSortTabl})
// 	if err != nil {
// 		return nil, err
// 	}
// 	holders = make([]*types.Holder, len(res))
// 	for i, v := range res {
// 		holders[i], err = types.ByteToHolder(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return
// }

func WriteErc721HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc721HolderKey(contract, holder.Addr)
	err = db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
	if err == nil {
		var (
			sortKey = append(erc721HolderPrefix, contract.Bytes()...)
		)
		return db.SPut(ctx, sortKey, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
	}
	return
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

func WriteErc721HolderTokenIdQuantity(ctx context.Context, db kv.Writer, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc721TokenIdHolderKey(contract, addr, tokenId)
	if quantity.Cmp(field.NewInt(0)) == 0 {
		return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
	} else {
		return db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
	}
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

/// ------------------- erc1155 holder ---------------------
func DelErc1155HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc1155HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc1155Holder(ctx context.Context, db kv.Sorter, contract common.Address, page, pageSize uint64) (holders []*types.Holder, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, page, pageSize, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return nil, err
	}
	holders = make([]*types.Holder, len(res))
	for i, v := range res {
		holders[i], err = types.ByteToHolder(v)
		if err != nil {
			return nil, err
		}
	}
	return
}

func WriteErc1155HolderAmount(ctx context.Context, db kv.Database, contract common.Address, holder *types.Holder) (err error) {
	var key = getErc1155HolderKey(contract, holder.Addr)
	err = db.Put(ctx, key, holder.Quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
	if err == nil {
		var (
			sortKey = append(erc1155HolderPrefix, contract.Bytes()...)
		)
		return db.SPut(ctx, sortKey, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
	}
	return
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

func WriteErc1155HolderTokenIdQuantity(ctx context.Context, db kv.Writer, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc1155TokenIdHolderKey(contract, addr, tokenId)
	if quantity.Cmp(field.NewInt(0)) == 0 {
		return db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
	} else {
		return db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
	}
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
