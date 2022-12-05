package storage

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Storage interface {
	ReadAccount(ctx context.Context, addr common.Address) (acc *types.Account, err error)
	ReadAccountTxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error)
	ReadAccountTxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (hash common.Hash, err error)
	ReadAccountTxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (tx *types.Tx, err error)
	ReadAccountITxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error)
	ReadAccountITxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.InternalTxKey, err error)
	ReadAccountITxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (itx *types.InternalTx, err error)
	ReadAccountErc20Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error)
	ReadAccountErc20Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc20TransferIndex *field.BigInt, err error)
	ReadAccountErc20ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc20Transfer, err error)
	ReadAccountErc721Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error)
	ReadAccountErc721Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc721TransferIndex *field.BigInt, err error)
	ReadAccountErc721ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc721Transfer, err error)
	ReadAccountErc1155Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error)
	ReadAccountErc1155Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc1155TransferIndex *field.BigInt, err error)
	ReadAccountErc1155ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc1155Transfer, err error)

	ReadBlock(ctx context.Context, blockNum *field.BigInt) (bk *types.Block, err error)
	ReadBlockIndex(ctx context.Context, blockNum *field.BigInt, index *field.BigInt) (txHash common.Hash, err error)
	ReadBlockTxByIndex(ctx context.Context, blockNum *field.BigInt, index *field.BigInt) (tx *types.Tx, err error)

	WriteValidateContractMetadata(ctx context.Context, data *types.ValidateContractMetadata) error
	ReadValidateContractMetadata(ctx context.Context) (acc *types.ValidateContractMetadata, err error)
	WriteValidateContractStatus(ctx context.Context, address common.Address, status *big.Int) error
	ReadValidateContractStatus(ctx context.Context, address common.Address) (status *big.Int, err error)
	WriteValidateContract(ctx context.Context, address common.Address, data *types.ContractVerity) error
	ReadValidateContract(ctx context.Context, address common.Address) (data *types.ContractVerity, err error)
	WriteMethodName(ctx context.Context, methodID, methodName string) error
	ReadMethodName(ctx context.Context, methodID, methodName string) (data string, err error)

	ReadContract(ctx context.Context, addr common.Address) (acc *types.Contract, err error)
	ReadProxyContract(ctx context.Context, proxy common.Address) (logic common.Address, err error)

	GetErc20Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error)
	GetErc20HolderCount(ctx context.Context, contract common.Address) (count uint64, err error)
	ReadErc20HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error)
	GetErc721Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error)
	GetErc721HolderCount(ctx context.Context, contract common.Address) (count uint64, err error)
	GetErc721Inventory(ctx context.Context, contract common.Address, offset, limit uint64) (inventorys []*types.Inventory, err error)
	GetErc721InventoryCount(ctx context.Context, contract common.Address) (count uint64, err error)
	ReadErc721HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error)
	ReadErc721HolderTokenIdQuantity(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error)
	GetErc1155Inventory(ctx context.Context, contract common.Address, offset, limit uint64) (inventorys []*field.BigInt, err error)
	GetErc1155InventoryCount(ctx context.Context, contract common.Address) (count uint64, err error)
	GetErc1155Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error)
	GetErc1155HolderCount(ctx context.Context, contract common.Address) (count uint64, err error)
	ReadErc1155HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error)
	ReadErc1155HolderTokenIdQuantity(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error)

	ReadHome(ctx context.Context) (home *types.Home, err error)
	ReadSyncingBlock(ctx context.Context) (bk *field.BigInt, err error)

	ReadITx(ctx context.Context, hash common.Hash, index *field.BigInt) (data *types.InternalTx, err error)
	ReadITxTotal(ctx context.Context, hash common.Hash) (total *field.BigInt, err error)

	ReadTraceTx(ctx context.Context, hash common.Hash) (res *types.TraceTx, err error)
	ReadTraceTx2(ctx context.Context, hash common.Hash) (res *types.TraceTx2, err error)

	ReadTx(ctx context.Context, hash common.Hash) (data *types.Tx, err error)
	ReadTxByIndex(ctx context.Context, index *field.BigInt) (data *types.Tx, err error)
	ReadTxTotal(ctx context.Context) (total *field.BigInt, err error)
	ReadRt(ctx context.Context, hash common.Hash) (data *types.Rt, err error)

	ReadErc20Total(ctx context.Context) (total *field.BigInt, err error)
	ReadErc721Total(ctx context.Context) (total *field.BigInt, err error)
	ReadErc1155Total(ctx context.Context) (total *field.BigInt, err error)
	ReadErc20Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc20Transfer, err error)
	ReadErc721Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc721Transfer, err error)
	ReadErc1155Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc1155Transfer, err error)
}
