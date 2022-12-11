package core

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/Ankr-network/uscan/share"
	"math/big"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	forkErc20TrasferTotal   *field.BigInt
	forkErc721TrasferTotal  *field.BigInt
	forkErc1155TrasferTotal *field.BigInt

	forkErc20TrasferAccountTotalMap   = utils.NewCache()
	forkErc721TrasferAccountTotalMap  = utils.NewCache()
	forkErc1155TrasferAccountTotalMap = utils.NewCache()

	forkErc20TransferContractTotalMap   = utils.NewCache()
	forkErc721TransferContractTotalMap  = utils.NewCache()
	forkErc1155TransferContractTotalMap = utils.NewCache()
)

// ------------------- erc20 transfer -----------------
func (n *blockHandle) writeForkErc20Transfer(ctx context.Context, data *types.Erc20Transfer, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	if forkErc20TrasferTotal == nil {
		forkErc20TrasferTotal, err = forkdb.ReadErc20Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				forkErc20TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get fork erc20 transfer total: %v", err)
				return err
			}
		}
	}
	forkErc20TrasferTotal.Add(field.NewInt(1))
	err = forkdb.WriteErc20Transfer(ctx, n.db, forkErc20TrasferTotal, data)
	if err != nil {
		log.Errorf("write fork erc20 transfer: %v", err)
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append([]byte("/fork/erc20/"), forkErc20TrasferTotal.Bytes()...))
	key := []byte("/fork/erc20/index")
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	if data.From != (common.Address{}) {
		if err = n.writeForkAccountErc20TransferIndex(ctx, data.From, forkErc20TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(From: %v) erc20 transfer index:%v", data.From.Hex(), err)
			return err
		}

		//if err = n.writeForkErc20HolderAmount(ctx, data.Contract, data.From, &data.Amount, decrease); err != nil {
		//	log.Errorf("decrease fork account(From: %v) erc20:%v", data.From.Hex(), err)
		//	return err
		//}
	} else {
		if data.To != (common.Address{}) {
			if err = n.updateForkErc20Account(ctx, data.Contract, &data.Amount, increase); err != nil {
				log.Errorf("update fork erc20 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeForkAccountErc20TransferIndex(ctx, data.To, forkErc20TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(to: %v) erc20 transfer index:%v", data.To.Hex(), err)
			return err
		}
		//if err = n.writeForkErc20HolderAmount(ctx, data.Contract, data.To, &data.Amount, increase); err != nil {
		//	log.Errorf("increase fork account(to: %v) erc20:%v", data.From.Hex(), err)
		//	return err
		//}
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateForkErc20Account(ctx, data.Contract, &data.Amount, decrease); err != nil {
				log.Errorf("update fork erc20 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if err = n.writeForkErc20ContractTransferIndex(ctx, data.Contract, forkErc20TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
		log.Errorf("write fork erc20 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeForkErc20ContractTransferIndex(ctx context.Context, contract common.Address, transfer20Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc20TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadErc20ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork erc20 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteErc20ContractTransfer(ctx, n.db, contract, total, transfer20Index); err != nil {
		log.Errorf("write fork erc20 contract transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append(append([]byte("/fork/erc20/"), contract.Bytes()...), append([]byte("/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/erc20/"), contract.Bytes()...), []byte("/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteErc20ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		forkErc20TransferContractTotalMap.Add(contract, total.Bytes())
	}
	key2 := append(append([]byte("/fork/erc20/"), contract.Bytes()...), []byte("/total")...)
	if totalMap[share.ForkTransferTbl+":"+string(key2)] == nil {
		totalMap[share.ForkTransferTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkTransferTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) writeForkAccountErc20TransferIndex(ctx context.Context, addr common.Address, transfer20Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc20TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadAccountErc20Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork account erc20 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteAccountErc20Index(ctx, n.db, addr, total, transfer20Index); err != nil {
		log.Errorf("write fork account erc20 transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkAccountsTbl] = append(deleteMap[share.ForkAccountsTbl], append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteAccountErc20Total(ctx, n.db, addr, total)
	if err == nil {
		forkErc20TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	key2 := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...)
	if totalMap[share.ForkAccountsTbl+":"+string(key2)] == nil {
		totalMap[share.ForkAccountsTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) updateForkErc20Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readForkAccount(ctx, addr)
	if err != nil {
		log.Errorf("read fork account(%s): %v", addr, err)
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

	//if inde == increase {
	//	acc.TokenTotalSupply.Add(value)
	//} else {
	//	acc.TokenTotalSupply.Sub(value)
	//	if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
	//		acc.TokenTotalSupply = *field.NewInt(0)
	//	}
	//}
	acc.TokenTotalSupply = *field.NewInt(0)

	return nil
}

// ------------------- erc721 transfer -----------------
func (n *blockHandle) writeForkErc721Transfer(ctx context.Context, data *types.Erc721Transfer, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	if forkErc721TrasferTotal == nil {
		forkErc721TrasferTotal, err = forkdb.ReadErc721Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				forkErc721TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get fork erc721 transfer total: %v", err)
				return err
			}
		}
	}
	forkErc721TrasferTotal.Add(field.NewInt(1))
	err = forkdb.WriteErc721Transfer(ctx, n.db, forkErc721TrasferTotal, data)
	if err != nil {
		log.Errorf("write fork erc721 transfer: %v", err)
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append([]byte("/fork/erc721/"), forkErc721TrasferTotal.Bytes()...))
	key := []byte("/fork/erc721/index")
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	if data.From != (common.Address{}) {
		if err = n.writeForkAccountErc721TransferIndex(ctx, data.From, forkErc721TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(From: %v) erc721 transfer index:%v", data.From.Hex(), err)
			return err
		}
		//if err = n.writeErc721HolderAmount(ctx, data.Contract, data.From, &data.TokenId, decrease); err != nil {
		//	log.Errorf("decrease account(From: %v) erc721 tokenId:%v", data.From.Hex(), err)
		//	return err
		//}
	} else {
		if data.To != (common.Address{}) {
			if err = n.updateForkErc721Account(ctx, data.Contract, &data.TokenId, increase); err != nil {
				log.Errorf("update fork erc721 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeForkAccountErc721TransferIndex(ctx, data.To, forkErc721TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(to: %v) erc721 transfer index:%v", data.To.Hex(), err)
			return err
		}
		//if err = n.writeErc721HolderAmount(ctx, data.Contract, data.To, &data.TokenId, increase); err != nil {
		//	log.Errorf("increase account(to: %v) erc721 tokenId:%v", data.To.Hex(), err)
		//	return err
		//}
	} else {
		if data.From != (common.Address{}) {
			if err = n.updateForkErc721Account(ctx, data.Contract, &data.TokenId, decrease); err != nil {
				log.Errorf("update fork erc721 account(%s): %v", data.Contract.Hex(), err)
				return err
			}
		}
	}

	if err = n.writeForkErc721ContractTransferIndex(ctx, data.Contract, forkErc721TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
		log.Errorf("write fork erc721 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeForkErc721ContractTransferIndex(ctx context.Context, contract common.Address, transfer721Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc721TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadErc721ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork erc721 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteErc721ContractTransfer(ctx, n.db, contract, total, transfer721Index); err != nil {
		log.Errorf("write fork erc721 contract transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append(append([]byte("/fork/erc721/"), contract.Bytes()...), append([]byte("/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/erc721/"), contract.Bytes()...), []byte("/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteErc721ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		forkErc721TransferContractTotalMap.Add(contract, total.Bytes())
	}
	key2 := append(append([]byte("/fork/erc721/"), contract.Bytes()...), []byte("/total")...)
	if totalMap[share.ForkTransferTbl+":"+string(key2)] == nil {
		totalMap[share.ForkTransferTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkTransferTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) writeForkAccountErc721TransferIndex(ctx context.Context, addr common.Address, transfer721Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc721TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadAccountErc721Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork account erc20 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteAccountErc721Index(ctx, n.db, addr, total, transfer721Index); err != nil {
		log.Errorf("write fork account erc721 transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkAccountsTbl] = append(deleteMap[share.ForkAccountsTbl], append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteAccountErc721Total(ctx, n.db, addr, total)
	if err == nil {
		forkErc721TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	key2 := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...)
	if totalMap[share.ForkAccountsTbl+":"+string(key2)] == nil {
		totalMap[share.ForkAccountsTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) updateForkErc721Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readForkAccount(ctx, addr)
	if err != nil {
		log.Errorf("read fork account(%s): %v", addr, err)
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

	//if inde == increase {
	//	acc.TokenTotalSupply.Add(field.NewInt(1))
	//} else {
	//	acc.TokenTotalSupply.Sub(field.NewInt(1))
	//	if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
	//		acc.TokenTotalSupply = *field.NewInt(0)
	//	}
	//}
	acc.TokenTotalSupply = *field.NewInt(0)

	return nil
}

// ------------------- erc1155 transfer -----------------
func (n *blockHandle) writeForkErc1155Transfer(ctx context.Context, data *types.Erc1155Transfer, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	if forkErc1155TrasferTotal == nil {
		forkErc1155TrasferTotal, err = forkdb.ReadErc1155Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				forkErc1155TrasferTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get fork erc1155 transfer total: %v", err)
				return err
			}
		}
	}
	forkErc1155TrasferTotal.Add(field.NewInt(1))
	err = forkdb.WriteErc1155Transfer(ctx, n.db, forkErc1155TrasferTotal, data)
	if err != nil {
		log.Errorf("write fork erc1155 transfer: %v", err)
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append([]byte("/fork/erc1155/"), forkErc1155TrasferTotal.Bytes()...))
	key := []byte("/fork/erc1155/index")
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	if data.From != (common.Address{}) {
		if err = n.writeForkAccountErc1155TransferIndex(ctx, data.From, forkErc1155TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(From: %v) erc1155 transfer index:%v", data.From.Hex(), err)
			return err
		}
		//if err = n.writeForkErc1155HolderAmount(ctx, data.Contract, data.From, &data.TokenID, &data.Quantity, decrease); err != nil {
		//	log.Errorf("decrease fork account(From: %v) erc1155 tokenId:%v", data.From.Hex(), err)
		//	return err
		//}
	} else {
		if err = n.updateForkErc1155Account(ctx, data.Contract, &data.Quantity, decrease); err != nil {
			log.Errorf("decrease fork erc1155 account(%s): %v", data.Contract.Hex(), err)
			return err
		}
	}

	if data.To != (common.Address{}) {
		if err = n.writeForkAccountErc1155TransferIndex(ctx, data.To, forkErc1155TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(to: %v) erc1155 transfer index:%v", data.To.Hex(), err)
			return err
		}
		//if err = n.writeForkErc1155HolderAmount(ctx, data.Contract, data.To, &data.TokenID, &data.Quantity, increase); err != nil {
		//	log.Errorf("increase fork account(to: %v) erc1155 tokenId:%v", data.From.Hex(), err)
		//	return err
		//}
	} else {
		if err = n.updateForkErc1155Account(ctx, data.Contract, &data.Quantity, increase); err != nil {
			log.Errorf("increase fork erc1155 account(%s): %v", data.Contract.Hex(), err)
			return err
		}
	}

	if err = n.writeForkErc1155ContractTransferIndex(ctx, data.Contract, forkErc1155TrasferTotal, deleteMap, indexMap, totalMap); err != nil {
		log.Errorf("write fork erc1155 contract transfer index:%v", data.Contract.Hex(), err)
		return err
	}

	return nil
}

func (n *blockHandle) writeForkErc1155ContractTransferIndex(ctx context.Context, contract common.Address, transfer1155Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc1155TransferContractTotalMap.Get(contract); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadErc1155ContractTotal(ctx, n.db, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork erc1155 contract transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteErc1155ContractTransfer(ctx, n.db, contract, total, transfer1155Index); err != nil {
		log.Errorf("write fork erc1155 contract transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkTransferTbl] = append(deleteMap[share.ForkTransferTbl], append(append([]byte("/fork/erc1155/"), contract.Bytes()...), append([]byte("/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/erc1155/"), contract.Bytes()...), []byte("/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteErc1155ContractTotal(ctx, n.db, contract, total)
	if err == nil {
		forkErc1155TransferContractTotalMap.Add(contract, total.Bytes())
	}
	key2 := append(append([]byte("/fork/erc1155/"), contract.Bytes()...), []byte("/total")...)
	if totalMap[share.ForkTransferTbl+":"+string(key2)] == nil {
		totalMap[share.ForkTransferTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkTransferTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) writeForkAccountErc1155TransferIndex(ctx context.Context, addr common.Address, transfer1155Index *field.BigInt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if BytesRes, ok := forkErc1155TrasferAccountTotalMap.Get(addr); ok {
		total.SetBytes(BytesRes.([]byte))
	} else {
		total, err = forkdb.ReadAccountErc1155Total(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("read fork account erc1155 transfer: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	if err = forkdb.WriteAccountErc1155Index(ctx, n.db, addr, total, transfer1155Index); err != nil {
		log.Errorf("write fork account erc1155 transfer index: %v", err)
		return err
	}
	deleteMap[share.ForkAccountsTbl] = append(deleteMap[share.ForkAccountsTbl], append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), total.Bytes()...)...))
	key := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteAccountErc1155Total(ctx, n.db, addr, total)
	if err == nil {
		forkErc1155TrasferAccountTotalMap.Add(addr, total.Bytes())
	}
	key2 := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...)
	if totalMap[share.ForkAccountsTbl+":"+string(key2)] == nil {
		totalMap[share.ForkAccountsTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+string(key2)].Add(field.NewInt(1))

	return err
}

func (n *blockHandle) updateForkErc1155Account(ctx context.Context, addr common.Address, value *field.BigInt, inde bool) (err error) {
	acc, err := n.readAccount(ctx, addr)
	if err != nil {
		log.Errorf("read fork account(%s): %v", addr, err)
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

	//if inde == increase {
	//	acc.TokenTotalSupply.Add(value)
	//} else {
	//	acc.TokenTotalSupply.Sub(value)
	//	if acc.TokenTotalSupply.Cmp(field.NewInt(0)) < 0 {
	//		acc.TokenTotalSupply = *field.NewInt(0)
	//	}
	//}
	acc.TokenTotalSupply = *field.NewInt(0)

	return nil
}

// write total for erc20
func (n *blockHandle) updateForkErc20TrasferTotal(ctx context.Context, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) error {
	if forkErc20TrasferTotal != nil {
		oldTotal, err := forkdb.ReadErc20Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				oldTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc20 fork transfer total: %v", err)
				return err
			}
		}

		err = forkdb.WriteErc20Total(ctx, n.db, forkErc20TrasferTotal)
		if err != nil {
			return err
		}

		forkErc20TrasferTotal.Sub(oldTotal)
		key := []byte("/fork/erc20/total")
		if totalMap[share.ForkTransferTbl+":"+string(key)] == nil {
			totalMap[share.ForkTransferTbl+":"+string(key)] = field.NewInt(0)
		}
		totalMap[share.ForkTransferTbl+":"+string(key)].Add(forkErc20TrasferTotal)
	}
	return nil
}

// write total for erc721
func (n *blockHandle) updateForkErc721TrasferTotal(ctx context.Context, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) error {
	if forkErc721TrasferTotal != nil {
		oldTotal, err := forkdb.ReadErc721Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				oldTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc721 fork transfer total: %v", err)
				return err
			}
		}

		err = forkdb.WriteErc721Total(ctx, n.db, forkErc721TrasferTotal)
		if err != nil {
			return err
		}

		forkErc721TrasferTotal.Sub(oldTotal)
		key := []byte("/fork/erc721/total")
		if totalMap[share.ForkTransferTbl+":"+string(key)] == nil {
			totalMap[share.ForkTransferTbl+":"+string(key)] = field.NewInt(0)
		}
		totalMap[share.ForkTransferTbl+":"+string(key)].Add(forkErc721TrasferTotal)
	}
	return nil
}

// write total for erc155
func (n *blockHandle) updateForkErc1155TrasferTotal(ctx context.Context, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) error {
	if forkErc1155TrasferTotal != nil {
		oldTotal, err := forkdb.ReadErc1155Total(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				oldTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get erc1155 fork transfer total: %v", err)
				return err
			}
		}

		err = forkdb.WriteErc1155Total(ctx, n.db, forkErc1155TrasferTotal)
		if err != nil {
			return err
		}

		forkErc1155TrasferTotal.Sub(oldTotal)
		key := []byte("/fork/erc1155/total")
		if totalMap[share.ForkTransferTbl+":"+string(key)] == nil {
			totalMap[share.ForkTransferTbl+":"+string(key)] = field.NewInt(0)
		}
		totalMap[share.ForkTransferTbl+":"+string(key)].Add(forkErc1155TrasferTotal)
	}
	return nil
}
