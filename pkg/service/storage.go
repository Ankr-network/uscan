package service

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var store Storage

type Store struct {
	ctx context.Context
	St  *storage.StorageImpl
}

func NewStore(st *storage.StorageImpl) {
	store = &Store{
		ctx: context.Background(),
		St:  st,
	}
	WriteMetadata()
}

type Storage interface {
	GetBlock(blockNum *field.BigInt) (*types.Block, error)
	ListBlockTxs(total, blockNum *field.BigInt, offset, limit int64) ([]*types.Tx, error)
	ListBlocks(total *field.BigInt, offset, limit int64) ([]*types.Block, error)
	GetBlockTotal() (bk *field.BigInt, err error)

	GetAccount(address common.Address) (acc *types.Account, err error)
	GetContract(address common.Address) (*types.Contract, error)
	GetAccountTxTotal(address common.Address) (total *field.BigInt, err error)
	GetAccountErc20Total(address common.Address) (total *field.BigInt, err error)
	GetAccountErc721Total(address common.Address) (total *field.BigInt, err error)
	GetAccountErc1155Total(address common.Address) (total *field.BigInt, err error)
	GetAccountITxTotal(address common.Address) (total *field.BigInt, err error)

	ListAccountTxs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Tx, error)
	ListAccountITxs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.InternalTx, error)
	ListAccountErc20Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc20Transfer, error)
	ListAccountErc721Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc721Transfer, error)
	ListAccountErc1155Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc1155Transfer, error)

	GetHome() (home *types.Home, err error)

	GetTx(txhash common.Hash) (data *types.Tx, err error)
	GetRt(txhash common.Hash) (data *types.Rt, err error)
	GetTxTotal() (total *field.BigInt, err error)
	ListTxs(total *field.BigInt, offset, limit int64) ([]*types.Tx, error)

	GetErc20Total() (total *field.BigInt, err error)
	GetErc721Total() (total *field.BigInt, err error)
	GetErc1155Total() (total *field.BigInt, err error)
	ListErc20Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc20Transfer, error)
	ListErc721Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc721Transfer, error)
	ListErc1155Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc1155Transfer, error)

	ListErc20Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error)
	GetErc20HolderCount(address common.Address) (count uint64, err error)
	ListErc721Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error)
	GetErc721HolderCount(address common.Address) (count uint64, err error)
	ListErc1155Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error)
	GetErc1155HolderCount(address common.Address) (count uint64, err error)

	GetMethodName(methodID string) (string, error)

	ListErc721Inventories(address common.Address, offset, limit int64) ([]*types.Inventory, error)
	GetErc721InventoryCount(address common.Address) (count uint64, err error)
	ListErc1155Inventories(address common.Address, offset, limit int64) ([]*field.BigInt, error)
	GetErc1155InventoryCount(address common.Address) (count uint64, err error)

	WriteValidateContractMetadata(metadata *types.ValidateContractMetadata) error
	GetValidateContractMetadata() (data *types.ValidateContractMetadata, err error)
	GetValidateContract(address common.Address) (data *types.ContractVerity, err error)
	WriteValidateContractStatus(address common.Address, status *big.Int) error
	GetValidateContractStatus(address common.Address) (status *big.Int, err error)
	WriteMethodName(id, name string) error
	WriteValidateContract(address common.Address, data *types.ContractVerity) error
	GetProxyContract(address common.Address) (logic common.Address, err error)

	GetErc20ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc20Transfer, total *field.BigInt, err error)
	GetErc721ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc721Transfer, total *field.BigInt, err error)
	GetErc1155ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc1155Transfer, total *field.BigInt, err error)

	ReadTraceTx2(address common.Hash) (res *types.TraceTx2, err error)
}

func (s *Store) GetBlock(blockNum *field.BigInt) (*types.Block, error) {
	return s.St.ReadBlock(s.ctx, blockNum)
}

func (s *Store) ListBlockTxs(total, blockNum *field.BigInt, offset, limit int64) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)

	begin, end := ParsePage(total, offset, limit)
	p := begin
	for {

		tx, err := s.St.ReadBlockTxByIndex(s.ctx, blockNum, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}

func (s *Store) ListBlocks(total *field.BigInt, offset, limit int64) ([]*types.Block, error) {
	blocks := make([]*types.Block, 0)
	if total.ToUint64() == 0 {
		return blocks, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin
	for {
		block, err := s.St.ReadBlock(s.ctx, p)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	return blocks, nil
}

func (s *Store) GetBlockTotal() (bk *field.BigInt, err error) {
	return s.St.ReadSyncingBlock(s.ctx)
}

func (s *Store) GetAccount(address common.Address) (*types.Account, error) {
	return s.St.ReadAccount(s.ctx, address)
}

func (s *Store) GetContract(address common.Address) (*types.Contract, error) {
	return s.St.ReadContract(s.ctx, address)
}

func (s *Store) GetAccountTxTotal(address common.Address) (total *field.BigInt, err error) {
	return s.St.ReadAccountTxTotal(s.ctx, address)
}

func (s *Store) ListAccountTxs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		tx, err := s.St.ReadAccountTxByIndex(s.ctx, address, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}
func (s *Store) GetAccountErc20Total(address common.Address) (total *field.BigInt, err error) {
	return s.St.ReadAccountErc20Total(s.ctx, address)
}
func (s *Store) GetAccountErc721Total(address common.Address) (total *field.BigInt, err error) {
	return s.St.ReadAccountErc721Total(s.ctx, address)
}
func (s *Store) GetAccountErc1155Total(address common.Address) (total *field.BigInt, err error) {
	return s.St.ReadAccountErc1155Total(s.ctx, address)
}
func (s *Store) GetAccountITxTotal(address common.Address) (total *field.BigInt, err error) {
	return s.St.ReadAccountITxTotal(s.ctx, address)
}

func (s *Store) ListAccountITxs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.InternalTx, error) {
	txs := make([]*types.InternalTx, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	log.Infof("ListAccountITxs ParsePage, total:%d, begin:%d, end:%d", total.ToUint64(), begin.ToUint64(), end.ToUint64())
	p := begin
	for {
		tx, err := s.St.ReadAccountITxByIndex(s.ctx, address, p)
		if err != nil {
			log.Errorf("ListAccountITxs ReadAccountITxByIndex error. err:%s, p:%d, total:%d, begin:%d, end:%d", err, p.ToUint64(), total.ToUint64(), begin.ToUint64(), end.ToUint64())
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}

func (s *Store) ListAccountErc20Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc20Transfer, error) {
	txs := make([]*types.Erc20Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin
	for {
		tx, err := s.St.ReadAccountErc20ByIndex(s.ctx, address, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}
func (s *Store) ListAccountErc721Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc721Transfer, error) {
	txs := make([]*types.Erc721Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin
	for {
		tx, err := s.St.ReadAccountErc721ByIndex(s.ctx, address, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}
func (s *Store) ListAccountErc1155Txs(address common.Address, total *field.BigInt, offset, limit int64) ([]*types.Erc1155Transfer, error) {
	txs := make([]*types.Erc1155Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin
	for {
		tx, err := s.St.ReadAccountErc1155ByIndex(s.ctx, address, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}

func (s *Store) GetHome() (home *types.Home, err error) {
	return s.St.ReadHome(s.ctx)
}

func (s *Store) GetTx(txhash common.Hash) (data *types.Tx, err error) {
	return s.St.ReadTx(s.ctx, txhash)
}
func (s *Store) GetRt(txhash common.Hash) (data *types.Rt, err error) {
	return s.St.ReadRt(s.ctx, txhash)
}

func (s *Store) GetTxTotal() (total *field.BigInt, err error) {
	return s.St.ReadTxTotal(s.ctx)
}

func (s *Store) ListTxs(total *field.BigInt, offset, limit int64) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		tx, err := s.St.ReadTxByIndex(s.ctx, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}

func (s *Store) GetErc20Total() (total *field.BigInt, err error) {
	return s.St.ReadErc20Total(s.ctx)
}
func (s *Store) GetErc721Total() (total *field.BigInt, err error) {
	return s.St.ReadErc721Total(s.ctx)
}

func (s *Store) GetErc1155Total() (total *field.BigInt, err error) {
	return s.St.ReadErc1155Total(s.ctx)
}

func (s *Store) ListErc20Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc20Transfer, error) {
	txs := make([]*types.Erc20Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		tx, err := s.St.ReadErc20Transfer(s.ctx, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}
func (s *Store) ListErc721Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc721Transfer, error) {
	txs := make([]*types.Erc721Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		tx, err := s.St.ReadErc721Transfer(s.ctx, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}
func (s *Store) ListErc1155Transfers(total *field.BigInt, offset, limit int64) ([]*types.Erc1155Transfer, error) {
	txs := make([]*types.Erc1155Transfer, 0)
	if total.ToUint64() == 0 {
		return txs, nil
	}
	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		tx, err := s.St.ReadErc1155Transfer(s.ctx, p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return txs, nil
}

func (s *Store) ListErc20Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error) {
	return s.St.GetErc20Holder(s.ctx, address, uint64(offset), uint64(limit))
}

func (s *Store) GetErc20HolderCount(address common.Address) (count uint64, err error) {
	return s.St.GetErc20HolderCount(s.ctx, address)
}

func (s *Store) ListErc721Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error) {
	return s.St.GetErc721Holder(s.ctx, address, uint64(offset), uint64(limit))
}

func (s *Store) GetErc721HolderCount(address common.Address) (count uint64, err error) {
	return s.St.GetErc721HolderCount(s.ctx, address)
}
func (s *Store) ListErc1155Holders(address common.Address, offset, limit int64) (holders []*types.Holder, err error) {
	return s.St.GetErc1155Holder(s.ctx, address, uint64(offset), uint64(limit))
}

func (s *Store) GetErc1155HolderCount(address common.Address) (count uint64, err error) {
	return s.St.GetErc1155HolderCount(s.ctx, address)
}

func (s *Store) GetMethodName(methodID string) (string, error) {
	return s.St.ReadMethodName(s.ctx, methodID, "")
}

func (s *Store) ListErc721Inventories(address common.Address, offset, limit int64) ([]*types.Inventory, error) {
	return s.St.GetErc721Inventory(s.ctx, address, uint64(offset), uint64(limit))
}

func (s *Store) GetErc721InventoryCount(address common.Address) (count uint64, err error) {
	return s.St.GetErc721InventoryCount(s.ctx, address)
}

func (s *Store) ListErc1155Inventories(address common.Address, offset, limit int64) ([]*field.BigInt, error) {
	return s.St.GetErc1155Inventory(s.ctx, address, uint64(offset), uint64(limit))
}

func (s *Store) GetErc1155InventoryCount(address common.Address) (count uint64, err error) {
	return s.St.GetErc1155InventoryCount(s.ctx, address)
}

func (s *Store) WriteValidateContractMetadata(metadata *types.ValidateContractMetadata) error {
	return s.St.WriteValidateContractMetadata(s.ctx, metadata)
}

func (s *Store) GetValidateContractMetadata() (data *types.ValidateContractMetadata, err error) {
	return s.St.ReadValidateContractMetadata(s.ctx)
}

func (s *Store) GetValidateContract(address common.Address) (data *types.ContractVerity, err error) {
	return s.St.ReadValidateContract(s.ctx, address)
}

func (s *Store) WriteValidateContractStatus(address common.Address, status *big.Int) error {
	return s.St.WriteValidateContractStatus(s.ctx, address, status)
}

func (s *Store) GetValidateContractStatus(address common.Address) (status *big.Int, err error) {
	return s.St.ReadValidateContractStatus(s.ctx, address)
}

func (s *Store) WriteMethodName(id, name string) error {
	return s.St.WriteMethodName(s.ctx, id, name)
}

func (s *Store) WriteValidateContract(address common.Address, data *types.ContractVerity) error {
	return s.St.WriteValidateContract(s.ctx, address, data)
}

func (s *Store) GetProxyContract(address common.Address) (logic common.Address, err error) {
	return s.St.ReadProxyContract(s.ctx, address)
}

func (s *Store) GetErc20ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc20Transfer, total *field.BigInt, err error) {
	return s.St.GetErc20ContractTransfer(s.ctx, contract, offset, limit)
}

func (s *Store) GetErc721ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc721Transfer, total *field.BigInt, err error) {
	return s.St.GetErc721ContractTransfer(s.ctx, contract, offset, limit)
}

func (s *Store) GetErc1155ContractTransfer(contract common.Address, offset, limit int64) (data []*types.Erc1155Transfer, total *field.BigInt, err error) {
	return s.St.GetErc1155ContractTransfer(s.ctx, contract, offset, limit)
}

func (s *Store) ReadTraceTx2(address common.Hash) (res *types.TraceTx2, err error) {
	return s.St.ReadTraceTx2(s.ctx, address)
}
