package fulldb

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

// / ------------------- erc20 holder ---------------------
func DelErc20HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc20HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc20Holder(ctx context.Context, db kv.Sorter, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc20HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func GetErc20HolderCount(ctx context.Context, db kv.Sorter, contract common.Address) (count uint64, err error) {
	var key = append(erc20HolderPrefix, contract.Bytes()...)
	count, err = db.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
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

// / ------------------- erc721 holder ---------------------
func DelErc721HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc721HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc721Holder(ctx context.Context, db kv.Sorter, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func GetErc721HolderCount(ctx context.Context, db kv.Sorter, contract common.Address) (count uint64, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	count, err = db.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func GetErc721Inventory(ctx context.Context, db kv.Sorter, contract common.Address, offset, limit uint64) (inventorys []*types.Inventory, err error) {
	var key = append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	var res [][]byte
	res, err = db.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.InventorySortTabl})
	if err != nil {
		return nil, err
	}
	inventorys = make([]*types.Inventory, len(res))
	for i, v := range res {
		inventorys[i], err = types.ByteToInventory(v)
		if err != nil {
			return nil, err
		}
	}
	return
}

func GetErc721InventoryCount(ctx context.Context, db kv.Sorter, contract common.Address) (count uint64, err error) {
	var key = append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	count, err = db.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

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

func WriteErc721HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc721TokenIdHolderKey(contract, addr, tokenId)
	inventory := types.Inventory{
		Addr:    addr,
		TokenID: *tokenId,
	}
	if quantity.Cmp(field.NewInt(0)) == 0 {
		err = db.SDel(ctx, append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...), inventory.ToBytes(), &kv.WriteOption{Table: share.InventorySortTabl})
		if err == nil {
			err = db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
		}
	} else {
		err = db.SPut(ctx, append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...), inventory.ToBytes(), &kv.WriteOption{Table: share.InventorySortTabl})
		if err == nil {
			err = db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
		}
	}
	return err
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

// / ------------------- erc1155 holder ---------------------
func DelErc1155HolderAmount(ctx context.Context, db kv.Sorter, contract common.Address, holder *types.Holder) (err error) {
	var (
		key = append(erc1155HolderPrefix, contract.Bytes()...)
	)
	return db.SDel(ctx, key, holder.ToBytes(), &kv.WriteOption{Table: share.HolderSortTabl})
}

func GetErc1155Inventory(ctx context.Context, db kv.Sorter, contract common.Address, offset, limit uint64) (inventorys []*field.BigInt, err error) {
	var key = append(append(erc1155HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	var res [][]byte
	res, err = db.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.InventorySortTabl})
	if err != nil {
		return nil, err
	}
	inventorys = make([]*field.BigInt, len(res))
	for i, v := range res {
		bi := &field.BigInt{}
		bi.SetBytes(v)
		inventorys[i] = bi
	}
	return
}

func GetErc1155InventoryCount(ctx context.Context, db kv.Sorter, contract common.Address) (count uint64, err error) {
	var key = append(append(erc1155HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	count, err = db.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func GetErc1155Holder(ctx context.Context, db kv.Sorter, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc1155HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = db.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func GetErc1155HolderCount(ctx context.Context, db kv.Sorter, contract common.Address) (count uint64, err error) {
	var key = append(erc1155HolderPrefix, contract.Bytes()...)
	count, err = db.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
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

func WriteErc1155HolderTokenIdQuantity(ctx context.Context, db kv.Database, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt) (err error) {
	var key = getErc1155TokenIdHolderKey(contract, addr, tokenId)
	err = db.SPut(ctx, append(append(erc1155HolderPrefix, contract.Bytes()...), []byte("/tokenId")...), tokenId.Bytes(), &kv.WriteOption{Table: share.InventorySortTabl})
	if err == nil {
		if quantity.Cmp(field.NewInt(0)) == 0 {
			err = db.Del(ctx, key, &kv.WriteOption{Table: share.HolderTbl})
		} else {
			err = db.Put(ctx, key, quantity.Bytes(), &kv.WriteOption{Table: share.HolderTbl})
		}
	}
	return err
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
