package core

import (
	"context"
	"errors"
	"math/big"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rawdb"
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
)

// ------------------- erc20 transfer -----------------
func (n *blockHandle) writeErc20Transfer(ctx context.Context, data *types.Erc20Transfer) (err error) {
	if erc20TrasferTotal == nil {
		erc20TrasferTotal, err = rawdb.ReadErc20Total(ctx, n.db)
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
	err = rawdb.WriteErc20Transfer(ctx, n.db, erc20TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc20 transfer: %v", err)
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc20TransferIndex(ctx, data.From, erc20TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc20 transfer index:%v", data.From.Hex(), err)
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
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateErc20Account(ctx, data.Contract, &data.Amount, decrease); err != nil {
				log.Errorf("update erc20 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}
	return nil
}

func (n *blockHandle) writeAccountErc20TransferIndex(ctx context.Context, addr common.Address, transfer20Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc20TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = rawdb.ReadAccountErc20Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc2- transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = rawdb.WriteAccountErc20Index(ctx, n.db, addr, total, transfer20Index); err != nil {
		log.Errorf("write account erc20 transfer index: %v", err)
		return err
	}

	err = rawdb.WriteAccountErc20Total(ctx, n.db, addr, total)
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

	if acc.Decimals.String() == "0" {
		var symbol *big.Int
		symbol, err = n.contractClient.GetContractDecimals(addr.Hex())
		if err == nil {
			acc.Decimals.SetBytes(symbol.Bytes())
		} else {
			acc.Retry.Add(field.NewInt(1))
		}
	}

	if inde == increase {
		acc.TokenTotalSupply.Add(value)
	} else {
		acc.TokenTotalSupply.Sub(value)
	}

	return nil
}

// ------------------- erc721 transfer -----------------
func (n *blockHandle) writeErc721Transfer(ctx context.Context, data *types.Erc721Transfer) (err error) {
	if erc721TrasferTotal == nil {
		erc721TrasferTotal, err = rawdb.ReadErc721Total(ctx, n.db)
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
	err = rawdb.WriteErc721Transfer(ctx, n.db, erc721TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc721 transfer: %v", err)
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc721TransferIndex(ctx, data.From, erc721TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc721 transfer index:%v", data.From.Hex(), err)
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
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateErc721Account(ctx, data.Contract, &data.TokenId, decrease); err != nil {
				log.Errorf("update erc721 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}
	return nil
}

func (n *blockHandle) writeAccountErc721TransferIndex(ctx context.Context, addr common.Address, transfer721Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc721TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = rawdb.ReadAccountErc721Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc2- transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = rawdb.WriteAccountErc721Index(ctx, n.db, addr, total, transfer721Index); err != nil {
		log.Errorf("write account erc721 transfer index: %v", err)
		return err
	}

	err = rawdb.WriteAccountErc721Total(ctx, n.db, addr, total)
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

	if acc.Name == "" {
		acc.Name, _ = n.contractClient.GetContractName(addr.Hex())
	}

	if acc.Symbol == "" {
		acc.Symbol, _ = n.contractClient.GetContractSymbol(addr.Hex())
	}

	if inde == increase {
		acc.TokenTotalSupply.Add(field.NewInt(1))
	} else {
		acc.TokenTotalSupply.Sub(field.NewInt(1))
	}

	return nil
}

// ------------------- erc1155 transfer -----------------
func (n *blockHandle) writeErc1155Transfer(ctx context.Context, data *types.Erc1155Transfer) (err error) {
	if erc1155TrasferTotal == nil {
		erc1155TrasferTotal, err = rawdb.ReadErc1155Total(ctx, n.db)
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
	err = rawdb.WriteErc1155Transfer(ctx, n.db, erc1155TrasferTotal, data)
	if err != nil {
		log.Errorf("write erc1155 transfer: %v", err)
	}

	if err = n.updateErc1155Account(ctx, data.Contract); err != nil {
		log.Errorf("update erc1155 account(%s): %v", data.Contract.Hex(), err)
		return err
	}

	if data.From != (common.Address{}) {
		if err = n.writeAccountErc1155TransferIndex(ctx, data.From, erc1155TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc1155 transfer index:%v", data.From.Hex(), err)
			return err
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeAccountErc1155TransferIndex(ctx, data.To, erc1155TrasferTotal); err != nil {
			log.Errorf("write account(to: %v) erc1155 transfer index:%v", data.To.Hex(), err)
			return err
		}
	}
	return nil
}

func (n *blockHandle) writeAccountErc1155TransferIndex(ctx context.Context, addr common.Address, transfer1155Index *field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := erc1155TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = rawdb.ReadAccountErc1155Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read account erc2- transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = rawdb.WriteAccountErc1155Index(ctx, n.db, addr, total, transfer1155Index); err != nil {
		log.Errorf("write account erc1155 transfer index: %v", err)
		return err
	}

	err = rawdb.WriteAccountErc1155Total(ctx, n.db, addr, total)
	if err == nil {
		erc1155TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	return err
}

func (n *blockHandle) updateErc1155Account(ctx context.Context, addr common.Address) (err error) {
	acc, err := n.readAccount(ctx, addr)
	if err != nil {
		log.Errorf("read account(%s): %v", addr, err)
		return err
	}
	if !acc.Erc1155 {
		acc.Erc1155 = true
		n.newErc1155Total.Add(field.NewInt(1))
	}

	if acc.Name == "" {
		acc.Name, _ = n.contractClient.GetContractName(addr.Hex())
	}

	if acc.Symbol == "" {
		acc.Symbol, _ = n.contractClient.GetContractSymbol(addr.Hex())
	}

	return nil
}

// write total for erc20
func (n *blockHandle) updateErc20TrasferTotal(ctx context.Context) (err error) {
	if erc20TrasferTotal != nil {
		err = rawdb.WriteErc20Total(ctx, n.db, erc20TrasferTotal)
	}
	return
}

// write total for erc721
func (n *blockHandle) updateErc721TrasferTotal(ctx context.Context) (err error) {
	if erc721TrasferTotal != nil {
		err = rawdb.WriteErc721Total(ctx, n.db, erc721TrasferTotal)
	}
	return
}

// write total for erc155
func (n *blockHandle) updateErc1155TrasferTotal(ctx context.Context) (err error) {
	if erc1155TrasferTotal != nil {
		err = rawdb.WriteErc1155Total(ctx, n.db, erc1155TrasferTotal)
	}
	return
}
