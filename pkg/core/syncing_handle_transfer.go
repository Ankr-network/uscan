package core

import (
	"context"
	"errors"
	"math/big"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
)

const (
	increase bool = true
	decrease bool = false
)

var (
	erc20TrasferTotal   *field.BigInt
	erc721TrasferTotal  *field.BigInt
	erc1155TrasferTotal *field.BigInt

	erc20TrasferAccountTotalMap   = utils.NewCache()
	erc721TrasferAccountTotalMap  = utils.NewCache()
	erc1155TrasferAccountTotalMap = utils.NewCache()

	erc20TransferContractTotalMap   = utils.NewCache()
	erc721TransferContractTotalMap  = utils.NewCache()
	erc1155TransferContractTotalMap = utils.NewCache()
)

// ------------------- erc20 transfer -----------------
func (n *blockHandle) writeErc20Transfer(ctx context.Context, data *types.Erc20Transfer) (err error) {
	if erc20TrasferTotal == nil {
		erc20TrasferTotal, err = fulldb.ReadErc20Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				erc20TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc20 transfer total: %v", err)
				return err
			}
		}
	}
	erc20TrasferTotal.Add(field.NewInt(1))
	err = fulldb.WriteErc20Transfer(ctx, n.db, erc20TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc20 transfer: %v", err)
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc20TransferIndex(ctx, data.From, erc20TrasferTotal); err != nil {
			log.Errorf("write account(From: %v) erc20 transfer index:%v", data.From.Hex(), err)
			return err
		}

		if err = n.writeErc20HolderAmount(ctx, data.Contract, data.From, &data.Amount, decrease); err != nil {
			log.Errorf("decrease account(From: %v) erc20:%v", data.From.Hex(), err)
			return err
		}
	} else {
		if data.To != (common.Address{}) {
			if err = n.updateErc20Account(ctx, data.Contract, &data.Amount, increase); err != nil {
				log.Errorf("update erc20 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeAccountErc20TransferIndex(ctx, data.To, erc20TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc20 transfer index:%v", data.To.Hex(), err)
			return err
		}
		if err = n.writeErc20HolderAmount(ctx, data.Contract, data.To, &data.Amount, increase); err != nil {
			log.Errorf("increase account(to: %v) erc20:%v", data.From.Hex(), err)
			return err
		}
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateErc20Account(ctx, data.Contract, &data.Amount, decrease); err != nil {
				log.Errorf("update erc20 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if err = n.writeErc20ContractTransferIndex(ctx, data.Contract, erc20TrasferTotal); err != nil {
		log.Errorf("write erc20 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeErc20ContractTransferIndex(ctx context.Context, contract common.Address, transfer20Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc20TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadErc20ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read erc20 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteErc20ContractTransfer(ctx, n.db, contract, total, transfer20Index); err != nil {
		log.Errorf("write erc20 contract transfer index: %v", err)
		return err
	}

	err = fulldb.WriteErc20ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		erc20TransferContractTotalMap.Add(contract, total.Bytes())
	}
	return err
}

func (n *blockHandle) writeAccountErc20TransferIndex(ctx context.Context, addr common.Address, transfer20Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc20TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadAccountErc20Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc20 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteAccountErc20Index(ctx, n.db, addr, total, transfer20Index); err != nil {
		log.Errorf("write account erc20 transfer index: %v", err)
		return err
	}

	err = fulldb.WriteAccountErc20Total(ctx, n.db, addr, total)
	if err == nil {
		erc20TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	return err
}

func (n *blockHandle) updateErc20Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readAccount(ctx, addr)
	if err != nil {
		log.Errorf("read account(%s): %v", addr, err)
		return err
	}
	if !acc.Erc20 {
		acc.Erc20 = true
		n.newErc20Total.Add(field.NewInt(1))
	}

	if acc.Retry.Cmp(field.NewInt(6)) < 0 {
		if acc.Name == "" {
			if acc.Name, err = n.contractClient.GetContractName(addr.Hex()); err != nil {
				acc.Retry.Add(field.NewInt(1))
			}
		}

		if acc.Symbol == "" {
			if acc.Symbol, err = n.contractClient.GetContractSymbol(addr.Hex()); err != nil {
				acc.Retry.Add(field.NewInt(1))
			}
		}

		if acc.Decimals.String() == "0x0" {
			var symbol *big.Int
			symbol, err = n.contractClient.GetContractDecimals(addr.Hex())
			if err == nil {
				acc.Decimals.SetBytes(symbol.Bytes())
			} else {
				acc.Retry.Add(field.NewInt(1))
			}
		}
	}

	if inde == increase {
		acc.TokenTotalSupply.Add(value)
	} else {
		acc.TokenTotalSupply.Sub(value)
		if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
			acc.TokenTotalSupply = *field.NewInt(0)
		}
	}

	return nil
}

func (n *blockHandle) writeErc20HolderAmount(ctx context.Context, contract common.Address, addr common.Address, amount *field.BigInt, inde bool) (err error) {
	var oriAmount *field.BigInt
	oriAmount, err = fulldb.ReadErc20HolderAmount(ctx, n.db, contract, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			oriAmount = &field.BigInt{}
			err = nil
		} else {
			return err
		}
	} else {
		err = fulldb.DelErc20HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
		if err != nil {
			return err
		}
	}
	if inde == increase {
		oriAmount.Add(amount)
	} else {
		oriAmount.Sub(amount)
		if oriAmount.Cmp(field.NewInt(0)) < 0 {
			oriAmount = field.NewInt(0)
		}
	}
	return fulldb.WriteErc20HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
}

// ------------------- erc721 transfer -----------------
func (n *blockHandle) writeErc721Transfer(ctx context.Context, data *types.Erc721Transfer) (err error) {
	if erc721TrasferTotal == nil {
		erc721TrasferTotal, err = fulldb.ReadErc721Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				erc721TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc721 transfer total: %v", err)
				return err
			}
		}
	}
	erc721TrasferTotal.Add(field.NewInt(1))
	err = fulldb.WriteErc721Transfer(ctx, n.db, erc721TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc721 transfer: %v", err)
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc721TransferIndex(ctx, data.From, erc721TrasferTotal); err != nil {
			log.Errorf("write account(From: %v) erc721 transfer index:%v", data.From.Hex(), err)
			return err
		}
		if err = n.writeErc721HolderAmount(ctx, data.Contract, data.From, &data.TokenId, decrease); err != nil {
			log.Errorf("decrease account(From: %v) erc721 tokenId:%v", data.From.Hex(), err)
			return err
		}
	} else {
		if data.To != (common.Address{}) {
			if err = n.updateErc721Account(ctx, data.Contract, &data.TokenId, increase); err != nil {
				log.Errorf("update erc721 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeAccountErc721TransferIndex(ctx, data.To, erc721TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc721 transfer index:%v", data.To.Hex(), err)
			return err
		}
		if err = n.writeErc721HolderAmount(ctx, data.Contract, data.To, &data.TokenId, increase); err != nil {
			log.Errorf("increase account(to: %v) erc721 tokenId:%v", data.To.Hex(), err)
			return err
		}
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateErc721Account(ctx, data.Contract, &data.TokenId, decrease); err != nil {
				log.Errorf("update erc721 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if err = n.writeErc721ContractTransferIndex(ctx, data.Contract, erc721TrasferTotal); err != nil {
		log.Errorf("write erc721 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeErc721ContractTransferIndex(ctx context.Context, contract common.Address, transfer721Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc721TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadErc721ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read erc721 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteErc721ContractTransfer(ctx, n.db, contract, total, transfer721Index); err != nil {
		log.Errorf("write erc721 contract transfer index: %v", err)
		return err
	}

	err = fulldb.WriteErc721ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		erc721TransferContractTotalMap.Add(contract, total.Bytes())
	}
	return err
}

func (n *blockHandle) writeAccountErc721TransferIndex(ctx context.Context, addr common.Address, transfer721Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc721TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadAccountErc721Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc20 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteAccountErc721Index(ctx, n.db, addr, total, transfer721Index); err != nil {
		log.Errorf("write account erc721 transfer index: %v", err)
		return err
	}

	err = fulldb.WriteAccountErc721Total(ctx, n.db, addr, total)
	if err == nil {
		erc721TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	return err
}

func (n *blockHandle) updateErc721Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readAccount(ctx, addr)
	if err != nil {
		log.Errorf("read account(%s): %v", addr, err)
		return err
	}
	if !acc.Erc721 {
		acc.Erc721 = true
		n.newErc721Total.Add(field.NewInt(1))
	}

	if acc.Retry.Cmp(field.NewInt(6)) < 0 {
		if acc.Name == "" {
			acc.Name, err = n.contractClient.GetContractName(addr.Hex())
			if err != nil {
				acc.Retry.Add(field.NewInt(1))
				err = nil
			}
		}

		if acc.Symbol == "" {
			acc.Symbol, err = n.contractClient.GetContractSymbol(addr.Hex())
			if err != nil {
				acc.Retry.Add(field.NewInt(1))
				err = nil
			}
		}
	}

	if inde == increase {
		acc.TokenTotalSupply.Add(field.NewInt(1))
	} else {
		acc.TokenTotalSupply.Sub(field.NewInt(1))
		if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
			acc.TokenTotalSupply = *field.NewInt(0)
		}
	}

	return nil
}

func (n *blockHandle) writeErc721HolderAmount(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt, inde bool) (err error) {
	var oriAmount *field.BigInt
	oriAmount, err = fulldb.ReadErc721HolderAmount(ctx, n.db, contract, addr)
	if err != nil {
		if !errors.Is(err, kv.NotFound) {
			return err
		}
		oriAmount = field.NewInt(0)
	} else {
		err = fulldb.DelErc721HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
		if err != nil {
			return err
		}
	}
	if inde == increase {
		err = fulldb.WriteErc721HolderTokenIdQuantity(ctx, n.db, contract, addr, tokenId, field.NewInt(1))
		if err != nil {
			return err
		}
		oriAmount.Add(field.NewInt(1))
	} else {
		err = fulldb.WriteErc721HolderTokenIdQuantity(ctx, n.db, contract, addr, tokenId, field.NewInt(0))
		if err != nil {
			return err
		}
		oriAmount.Sub(field.NewInt(1))
	}
	if oriAmount.Cmp(field.NewInt(0)) < 0 {
		oriAmount = field.NewInt(0)
	}
	return fulldb.WriteErc721HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
}

// ------------------- erc1155 transfer -----------------
func (n *blockHandle) writeErc1155Transfer(ctx context.Context, data *types.Erc1155Transfer) (err error) {
	if erc1155TrasferTotal == nil {
		erc1155TrasferTotal, err = fulldb.ReadErc1155Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				erc1155TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc1155 transfer total: %v", err)
				return err
			}
		}
	}
	erc1155TrasferTotal.Add(field.NewInt(1))
	err = fulldb.WriteErc1155Transfer(ctx, n.db, erc1155TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc1155 transfer: %v", err)
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc1155TransferIndex(ctx, data.From, erc1155TrasferTotal); err != nil {
			log.Errorf("write account(From: %v) erc1155 transfer index:%v", data.From.Hex(), err)
			return err
		}
		if err = n.writeErc1155HolderAmount(ctx, data.Contract, data.From, &data.TokenID, &data.Quantity, decrease); err != nil {
			log.Errorf("decrease account(From: %v) erc1155 tokenId:%v", data.From.Hex(), err)
			return err
		}
	} else {
		if err = n.updateErc1155Account(ctx, data.Contract, &data.Quantity, decrease); err != nil {
			log.Errorf("decrease erc1155 account(%s): %v", data.Contract.Hex(), err)
			return err
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeAccountErc1155TransferIndex(ctx, data.To, erc1155TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc1155 transfer index:%v", data.To.Hex(), err)
			return err
		}
		if err = n.writeErc1155HolderAmount(ctx, data.Contract, data.To, &data.TokenID, &data.Quantity, increase); err != nil {
			log.Errorf("increase account(to: %v) erc1155 tokenId:%v", data.From.Hex(), err)
			return err
		}
	} else {
		if err = n.updateErc1155Account(ctx, data.Contract, &data.Quantity, increase); err != nil {
			log.Errorf("increase erc1155 account(%s): %v", data.Contract.Hex(), err)
			return err
		}
	}

	if err = n.writeErc1155ContractTransferIndex(ctx, data.Contract, erc1155TrasferTotal); err != nil {
		log.Errorf("write erc1155 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeErc1155ContractTransferIndex(ctx context.Context, contract common.Address, transfer1155Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc1155TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadErc1155ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read erc1155 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteErc1155ContractTransfer(ctx, n.db, contract, total, transfer1155Index); err != nil {
		log.Errorf("write erc1155 contract transfer index: %v", err)
		return err
	}

	err = fulldb.WriteErc1155ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		erc1155TransferContractTotalMap.Add(contract, total.Bytes())
	}
	return err
}

func (n *blockHandle) writeAccountErc1155TransferIndex(ctx context.Context, addr common.Address, transfer1155Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc1155TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = fulldb.ReadAccountErc1155Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc20 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = fulldb.WriteAccountErc1155Index(ctx, n.db, addr, total, transfer1155Index); err != nil {
		log.Errorf("write account erc1155 transfer index: %v", err)
		return err
	}

	err = fulldb.WriteAccountErc1155Total(ctx, n.db, addr, total)
	if err == nil {
		erc1155TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	return err
}

func (n *blockHandle) updateErc1155Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readAccount(ctx, addr)
	if err != nil {
		log.Errorf("read account(%s): %v", addr, err)
		return err
	}
	if !acc.Erc1155 {
		acc.Erc1155 = true
		n.newErc1155Total.Add(field.NewInt(1))
	}

	if acc.Retry.Cmp(field.NewInt(6)) < 0 {
		if acc.Name == "" {
			acc.Name, err = n.contractClient.GetContractName(addr.Hex())
			if err != nil {
				acc.Retry.Add(field.NewInt(1))
				err = nil
			}
		}

		if acc.Symbol == "" {
			acc.Symbol, err = n.contractClient.GetContractSymbol(addr.Hex())
			if err != nil {
				acc.Retry.Add(field.NewInt(1))
				err = nil
			}
		}
	}

	if inde == increase {
		acc.TokenTotalSupply.Add(value)
	} else {
		acc.TokenTotalSupply.Sub(value)
		if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
			acc.TokenTotalSupply = *field.NewInt(0)
		}
	}

	return nil
}

func (n *blockHandle) writeErc1155HolderAmount(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt, quantity *field.BigInt, inde bool) (err error) {
	var oriAmount *field.BigInt
	oriAmount, err = fulldb.ReadErc1155HolderAmount(ctx, n.db, contract, addr)
	if err != nil {
		if !errors.Is(err, kv.NotFound) {
			return err
		}
		oriAmount = field.NewInt(0)
	} else {
		err = fulldb.DelErc721HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
		if err != nil {
			return err
		}
	}

	var oriQuantity *field.BigInt
	oriQuantity, err = fulldb.ReadErc1155HolderTokenIdQuantity(ctx, n.db, contract, addr, tokenId)
	if err != nil {
		if !errors.Is(err, kv.NotFound) {
			return err
		}
		oriQuantity = field.NewInt(0)
	}
	if inde == increase {
		oriQuantity.Add(quantity)
		err = fulldb.WriteErc1155HolderTokenIdQuantity(ctx, n.db, contract, addr, tokenId, oriQuantity)
		if err != nil {
			return err
		}
		oriAmount.Add(oriQuantity)
	} else {
		oriQuantity.Sub(quantity)
		if oriQuantity.Cmp(field.NewInt(0)) < 0 {
			oriQuantity = field.NewInt(0)
		}
		err = fulldb.WriteErc1155HolderTokenIdQuantity(ctx, n.db, contract, addr, tokenId, oriQuantity)
		if err != nil {
			return err
		}
		oriAmount.Sub(quantity)
	}
	if oriAmount.Cmp(field.NewInt(0)) < 0 {
		oriAmount = field.NewInt(0)
	}
	return fulldb.WriteErc1155HolderAmount(ctx, n.db, contract, &types.Holder{Addr: addr, Quantity: *oriAmount})
}

// write total for erc20
func (n *blockHandle) updateErc20TrasferTotal(ctx context.Context) (err error) {
	if erc20TrasferTotal != nil {
		err = fulldb.WriteErc20Total(ctx, n.db, erc20TrasferTotal)
	}
	return
}

// write total for erc721
func (n *blockHandle) updateErc721TrasferTotal(ctx context.Context) (err error) {
	if erc721TrasferTotal != nil {
		err = fulldb.WriteErc721Total(ctx, n.db, erc721TrasferTotal)
	}
	return
}

// write total for erc155
func (n *blockHandle) updateErc1155TrasferTotal(ctx context.Context) (err error) {
	if erc1155TrasferTotal != nil {
		err = fulldb.WriteErc1155Total(ctx, n.db, erc1155TrasferTotal)
	}
	return
}
