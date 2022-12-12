package storage

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	contractMetadataPrefix  = []byte("metadata")
	ContractVerityPrefix    = []byte("contract/info/")
	ContractVerityTmpPrefix = []byte("contract/tmp/")
	ContractMethodPrefix    = []byte("method/")

	contractKeyPrefix   = []byte("/contract/")
	proxyContractPrefix = []byte("/proxy/")

	erc20HolderPrefix   = []byte("/erc20/")
	erc721HolderPrefix  = []byte("/erc721/")
	erc1155HolderPrefix = []byte("/erc1155/")
)

var _ Storage = (*StorageImpl)(nil)

type StorageImpl struct {
	ForkDB *mdbx.MdbxDB
	FullDB *mdbx.MdbxDB
}

var schemas = []string{}

func NewStorage(path string) *StorageImpl {
	return &StorageImpl{
		ForkDB: mdbx.NewMdbx(path+"/fork", []string{
			share.ForkHomeTbl,
			share.ForkAccountsTbl,
			share.ForkTxTbl,
			share.ForkBlockTbl,
			share.ForkTraceLogTbl,
			share.ForkTransferTbl,
			share.ForkIndexTbl,
		}, []string{}),
		FullDB: mdbx.NewMdbx(path, []string{
			share.AccountsTbl,
			share.HomeTbl,
			share.TxTbl,
			share.BlockTbl,
			share.TraceLogTbl,
			share.TransferTbl,
			share.HolderTbl,
			share.ValidateContractTbl,
		}, []string{
			share.HolderSortTabl,
			share.InventorySortTabl,
		}),
	}
}

func (s *StorageImpl) ReadAccount(ctx context.Context, addr common.Address) (acc *types.Account, err error) {
	var bytesRes []byte
	accFork := &types.Account{}
	accFull := &types.Account{}
	acc = &types.Account{}

	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/info/"), addr.Bytes()...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			err = nil
		} else {
			return
		}
	} else {
		err = accFork.Unmarshal(bytesRes)
		if err == nil {
			accFork.Owner = addr
		}
	}

	bytesRes, err = s.FullDB.Get(ctx, append([]byte("/info/"), addr.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			err = nil
		} else {
			return
		}
	} else {
		err = accFull.Unmarshal(bytesRes)
		if err == nil {
			accFull.Owner = addr
		}
	}

	if accFork.Owner.String() != "0x0000000000000000000000000000000000000000" && accFull.Owner.String() != "0x0000000000000000000000000000000000000000" {
		acc = &types.Account{
			Owner:            accFork.Owner,
			Erc20:            accFork.Erc20,
			Erc721:           accFork.Erc721,
			Erc1155:          accFork.Erc1155,
			ErcFlag:          accFork.ErcFlag,
			BlockNumber:      accFork.BlockNumber,
			Balance:          accFork.Balance,
			Name:             accFork.Name,
			Symbol:           accFork.Symbol,
			TokenTotalSupply: accFull.TokenTotalSupply,
			NftTotalSupply:   accFull.NftTotalSupply,
			Decimals:         accFork.Decimals,
			Creator:          accFork.Creator,
			TxHash:           accFork.TxHash,
			Retry:            accFork.Retry,
		}
	} else if accFork.Owner.String() != "0x0000000000000000000000000000000000000000" {
		acc = accFork
	} else {
		acc = accFull
	}
	return
}

func (s *StorageImpl) ReadAccountTxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountTxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (hash common.Hash, err error) {
	i := &field.BigInt{}
	hash = common.Hash{}

	totalAll, err := s.ReadAccountTxTotal(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountTxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressTxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountTxIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountTxIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountTxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (tx *types.Tx, err error) {
	i := &field.BigInt{}
	tx = &types.Tx{}

	totalAll, err := s.ReadAccountTxTotal(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountTxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressTxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(0)
				err = nil
			} else {
				return
			}
		}
		totalAll.Add(totalFork)
		return forkdb.ReadAccountTxByIndex(ctx, s.ForkDB, addr, i.Add(totalAll).Sub(index))
	} else {
		return fulldb.ReadAccountTxByIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountITxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountITxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.InternalTxKey, err error) {
	i := &field.BigInt{}
	data = &types.InternalTxKey{}

	totalAll, err := s.ReadAccountITxTotal(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountITxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressITxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountITxIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountITxIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountITxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (itx *types.InternalTx, err error) {
	i := &field.BigInt{}
	itx = &types.InternalTx{}

	totalAll, err := s.ReadAccountITxTotal(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountITxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressITxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountITxByIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountITxByIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc20Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc20Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc20TransferIndex *field.BigInt, err error) {
	i := &field.BigInt{}
	erc20TransferIndex = &field.BigInt{}

	totalAll, err := s.ReadAccountErc20Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc20Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc20Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc20Index(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc20Index(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc20ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc20Transfer{}

	totalAll, err := s.ReadAccountErc20Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc20Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc20Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc20ByIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc20ByIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc721Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc721Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc721TransferIndex *field.BigInt, err error) {
	i := &field.BigInt{}
	erc721TransferIndex = &field.BigInt{}

	totalAll, err := s.ReadAccountErc721Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc721Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc721Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc721Index(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc721Index(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc721ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc721Transfer{}

	totalAll, err := s.ReadAccountErc721Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc721Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc721Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc721ByIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc721ByIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc1155Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc1155Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc1155TransferIndex *field.BigInt, err error) {
	i := &field.BigInt{}
	erc1155TransferIndex = &field.BigInt{}

	totalAll, err := s.ReadAccountErc1155Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc1155Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc1155Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc1155Index(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc1155Index(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadAccountErc1155ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc1155Transfer{}

	totalAll, err := s.ReadAccountErc1155Total(ctx, addr)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadAccountErc1155Total(ctx, s.ForkDB, addr)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadAddressErc1155Index(ctx, s.ForkDB, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadAccountErc1155ByIndex(ctx, s.ForkDB, addr, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadAccountErc1155ByIndex(ctx, s.FullDB, addr, index)
	}
}

func (s *StorageImpl) ReadBlock(ctx context.Context, blockNum *field.BigInt) (bk *types.Block, err error) {
	var bytesRes []byte

	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/block/"), blockNum.Bytes()...), &kv.ReadOption{Table: share.ForkBlockTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, append([]byte("/block/"), blockNum.Bytes()...), &kv.ReadOption{Table: share.BlockTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	bk = &types.Block{}
	err = bk.Unmarshal(bytesRes)
	if err == nil {
		bk.Number.SetBytes(blockNum.Bytes())
	}
	return
}

func (s *StorageImpl) ReadBlockIndex(ctx context.Context, blockNum *field.BigInt, index *field.BigInt) (txHash common.Hash, err error) {
	var bytesRes []byte

	forkKey := make([]byte, 0, len([]byte("/fork/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
	forkKey = append(forkKey, []byte("/fork/block/")...)
	forkKey = append(forkKey, blockNum.Bytes()...)
	forkKey = append(forkKey, byte('/'))
	forkKey = append(forkKey, index.Bytes()...)
	bytesRes, err = s.ForkDB.Get(ctx, forkKey, &kv.ReadOption{Table: share.ForkBlockTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			fullKey := make([]byte, 0, len([]byte("/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
			fullKey = append(fullKey, []byte("/block/")...)
			fullKey = append(fullKey, blockNum.Bytes()...)
			fullKey = append(fullKey, byte('/'))
			fullKey = append(fullKey, index.Bytes()...)
			bytesRes, err = s.FullDB.Get(ctx, fullKey, &kv.ReadOption{Table: share.BlockTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	txHash.SetBytes(bytesRes)

	return
}

func (s *StorageImpl) ReadBlockTxByIndex(ctx context.Context, blockNum *field.BigInt, index *field.BigInt) (tx *types.Tx, err error) {
	var bytesRes []byte

	forkKey := make([]byte, 0, len([]byte("/fork/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
	forkKey = append(forkKey, []byte("/fork/block/")...)
	forkKey = append(forkKey, blockNum.Bytes()...)
	forkKey = append(forkKey, byte('/'))
	forkKey = append(forkKey, index.Bytes()...)

	bytesRes, err = s.ForkDB.Get(ctx, forkKey, &kv.ReadOption{Table: share.ForkBlockTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			fullKey := make([]byte, 0, len([]byte("/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
			fullKey = append(fullKey, []byte("/block/")...)
			fullKey = append(fullKey, blockNum.Bytes()...)
			fullKey = append(fullKey, byte('/'))
			fullKey = append(fullKey, index.Bytes()...)
			bytesRes, err = s.FullDB.Get(ctx, fullKey, &kv.ReadOption{Table: share.BlockTbl})
			if err != nil {
				return
			}
			hash := common.BytesToHash(bytesRes)
			return fulldb.ReadTx(ctx, s.FullDB, hash)
		} else {
			return
		}
	} else {
		hash := common.BytesToHash(bytesRes)
		return forkdb.ReadTx(ctx, s.ForkDB, hash)
	}
}

func (s *StorageImpl) WriteValidateContractMetadata(ctx context.Context, data *types.ValidateContractMetadata) error {
	bytesRes, err := data.Marshal()
	if err != nil {
		return err
	}
	return s.FullDB.Put(ctx, contractMetadataPrefix, bytesRes, &kv.WriteOption{Table: share.ValidateContractTbl})
}

func (s *StorageImpl) ReadValidateContractMetadata(ctx context.Context) (acc *types.ValidateContractMetadata, err error) {
	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, contractMetadataPrefix, &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.ValidateContractMetadata{}
	err = acc.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) WriteValidateContractStatus(ctx context.Context, address common.Address, status *big.Int) error {
	return s.FullDB.Put(ctx, append(ContractVerityTmpPrefix, address.Bytes()...), status.Bytes(), &kv.WriteOption{Table: share.ValidateContractTbl})
}

func (s *StorageImpl) ReadValidateContractStatus(ctx context.Context, address common.Address) (status *big.Int, err error) {
	rs, err := s.FullDB.Get(ctx, append(ContractVerityTmpPrefix, address.Bytes()...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	status = &big.Int{}
	status.SetBytes(rs)
	return
}

func (s *StorageImpl) WriteValidateContract(ctx context.Context, address common.Address, data *types.ContractVerity) error {
	bytesRes, err := data.Marshal()
	if err != nil {
		return err
	}
	return s.FullDB.Put(ctx, append(ContractVerityPrefix, address.Bytes()...), bytesRes, &kv.WriteOption{Table: share.ValidateContractTbl})
}

func (s *StorageImpl) ReadValidateContract(ctx context.Context, address common.Address) (data *types.ContractVerity, err error) {
	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, append(ContractVerityPrefix, address.Bytes()...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	data = &types.ContractVerity{}
	err = data.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) WriteMethodName(ctx context.Context, methodID, methodName string) error {
	return s.FullDB.Put(ctx, append(ContractMethodPrefix, []byte(methodID)...), []byte(methodName), &kv.WriteOption{Table: share.ValidateContractTbl})
}

func (s *StorageImpl) ReadMethodName(ctx context.Context, methodID, methodName string) (data string, err error) {
	rs, err := s.FullDB.Get(ctx, append(ContractMethodPrefix, []byte(methodID)...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return "", err
	}
	data = string(rs)
	return
}

func (s *StorageImpl) ReadContract(ctx context.Context, addr common.Address) (acc *types.Contract, err error) {
	var bytesRes []byte

	bytesRes, err = s.FullDB.Get(ctx, append(contractKeyPrefix, addr.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.Contract{}
	err = acc.Unmarshal(bytesRes)
	if err == nil {
		acc.Owner = addr
	}
	return
}

func (s *StorageImpl) ReadProxyContract(ctx context.Context, proxy common.Address) (logic common.Address, err error) {
	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, append(proxyContractPrefix, proxy.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err == nil {
		logic.SetBytes(bytesRes)
	}
	return
}

func (s *StorageImpl) GetErc20Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc20HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = s.FullDB.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func (s *StorageImpl) GetErc20HolderCount(ctx context.Context, contract common.Address) (count uint64, err error) {
	var key = append(erc20HolderPrefix, contract.Bytes()...)
	count, err = s.FullDB.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func (s *StorageImpl) ReadErc20HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	key := make([]byte, 0, len(erc20HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)

	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) GetErc721Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = s.FullDB.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func (s *StorageImpl) GetErc721HolderCount(ctx context.Context, contract common.Address) (count uint64, err error) {
	var key = append(erc721HolderPrefix, contract.Bytes()...)
	count, err = s.FullDB.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func (s *StorageImpl) GetErc721Inventory(ctx context.Context, contract common.Address, offset, limit uint64) (inventorys []*types.Inventory, err error) {
	var key = append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	var res [][]byte
	res, err = s.FullDB.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.InventorySortTabl})
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

func (s *StorageImpl) GetErc721InventoryCount(ctx context.Context, contract common.Address) (count uint64, err error) {
	var key = append(append(erc721HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	count, err = s.FullDB.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func (s *StorageImpl) ReadErc721HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	key := make([]byte, 0, len(erc721HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)

	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadErc721HolderTokenIdQuantity(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error) {
	key := make([]byte, 0, len(erc721HolderPrefix)+common.AddressLength*2+2+len(tokenId.Bytes()))
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	key = append(key, '/')
	key = append(key, tokenId.Bytes()...)

	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return
	}
	quantity = &field.BigInt{}
	quantity.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) GetErc1155Inventory(ctx context.Context, contract common.Address, offset, limit uint64) (inventorys []*field.BigInt, err error) {
	var key = append(append(erc1155HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	var res [][]byte
	res, err = s.FullDB.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.InventorySortTabl})
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

func (s *StorageImpl) GetErc1155InventoryCount(ctx context.Context, contract common.Address) (count uint64, err error) {
	var key = append(append(erc1155HolderPrefix, contract.Bytes()...), []byte("/tokenId")...)
	count, err = s.FullDB.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func (s *StorageImpl) GetErc1155Holder(ctx context.Context, contract common.Address, offset, limit uint64) (holders []*types.Holder, err error) {
	var key = append(erc1155HolderPrefix, contract.Bytes()...)
	var res [][]byte
	res, err = s.FullDB.SGet(ctx, key, offset, limit, &kv.ReadOption{Table: share.HolderSortTabl})
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

func (s *StorageImpl) GetErc1155HolderCount(ctx context.Context, contract common.Address) (count uint64, err error) {
	var key = append(erc1155HolderPrefix, contract.Bytes()...)
	count, err = s.FullDB.SCount(ctx, key, &kv.ReadOption{Table: share.HolderSortTabl})
	if err != nil {
		return 0, err
	}
	return
}

func (s *StorageImpl) ReadErc1155HolderAmount(ctx context.Context, contract common.Address, addr common.Address) (amount *field.BigInt, err error) {
	key := make([]byte, 0, len(erc1155HolderPrefix)+common.AddressLength*2+1)
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)

	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return nil, err
	}
	amount = &field.BigInt{}
	amount.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadErc1155HolderTokenIdQuantity(ctx context.Context, contract common.Address, addr common.Address, tokenId *field.BigInt) (quantity *field.BigInt, err error) {
	key := make([]byte, 0, len(erc1155HolderPrefix)+common.AddressLength*2+2+len(tokenId.Bytes()))
	key = append(key, contract.Bytes()...)
	key = append(key, '/')
	key = append(key, addr.Bytes()...)
	key = append(key, '/')
	key = append(key, tokenId.Bytes()...)

	var bytesRes []byte
	bytesRes, err = s.FullDB.Get(ctx, key, &kv.ReadOption{Table: share.HolderTbl})
	if err != nil {
		return
	}
	quantity = &field.BigInt{}
	quantity.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadHome(ctx context.Context) (home *types.Home, err error) {
	var bytesRes []byte
	blocks := make([]*types.BkSim, 0, 10)
	txs := make([]*types.TxSim, 0, 10)
	homeFork := &types.Home{}
	homeFull := &types.Home{}
	home = &types.Home{}

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/home"), &kv.ReadOption{Table: share.ForkHomeTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			err = nil
		} else {
			return
		}
	} else {
		err = homeFork.Unmarshal(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, []byte("/home"), &kv.ReadOption{Table: share.HomeTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			err = nil
		} else {
			return
		}
	} else {
		err = homeFull.Unmarshal(bytesRes)
	}

	if len(homeFork.Blocks) == 0 {
		blocks = homeFull.Blocks
	} else if len(homeFork.Blocks) < 10 {
		blocks = append(blocks, append(homeFull.Blocks, homeFork.Blocks...)...)
	} else {
		blocks = homeFork.Blocks
	}
	if len(blocks) > 10 {
		blocks = blocks[(len(home.Blocks) - 10):]
	}

	if len(homeFork.Txs) == 0 {
		txs = homeFull.Txs
	} else if len(homeFork.Txs) < 10 {
		txs = append(txs, append(homeFull.Txs, homeFork.Txs...)...)
	} else {
		txs = homeFork.Txs
	}
	if len(home.Txs) > 10 {
		home.Txs = home.Txs[(len(home.Txs) - 10):]
	}

	home = &types.Home{
		BlockNumber:  homeFork.BlockNumber,
		TxTotal:      *homeFull.TxTotal.Add(&homeFork.TxTotal),
		AddressTotal: *homeFull.AddressTotal.Add(&homeFork.AddressTotal),
		Erc20Total:   *homeFull.Erc20Total.Add(&homeFork.Erc20Total),
		Erc721Total:  *homeFull.Erc721Total.Add(&homeFork.Erc721Total),
		Erc1155Total: *homeFull.Erc1155Total.Add(&homeFork.Erc1155Total),
		Blocks:       blocks,
		Txs:          txs,
		DateTxs:      homeFull.DateTxs,
		DateTxsByte:  homeFull.DateTxsByte,
	}
	return
}

func (s *StorageImpl) ReadSyncingBlock(ctx context.Context) (bk *field.BigInt, err error) {
	var bytesRes []byte

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/syncing"), &kv.ReadOption{Table: share.ForkHomeTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, []byte("/syncing"), &kv.ReadOption{Table: share.HomeTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	bk = &field.BigInt{}
	bk.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadITx(ctx context.Context, hash common.Hash, index *field.BigInt) (data *types.InternalTx, err error) {
	i := &field.BigInt{}
	data = &types.InternalTx{}

	totalAll, err := s.ReadITxTotal(ctx, hash)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadITxTotal(ctx, s.ForkDB, hash)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadITxIndex(ctx, s.ForkDB, hash)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadITx(ctx, s.ForkDB, hash, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadITx(ctx, s.FullDB, hash, index)
	}
}

func (s *StorageImpl) ReadITxTotal(ctx context.Context, hash common.Hash) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/iTx/"), hash.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/iTx/"), hash.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadTraceTx(ctx context.Context, hash common.Hash) (res *types.TraceTx, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/tracetx/"), hash.Bytes()...), &kv.ReadOption{Table: share.ForkTraceLogTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, append([]byte("/tracetx/"), hash.Bytes()...), &kv.ReadOption{Table: share.TraceLogTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	res = &types.TraceTx{}
	err = res.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) ReadTraceTx2(ctx context.Context, hash common.Hash) (res *types.TraceTx2, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/tracetx2/"), hash.Bytes()...), &kv.ReadOption{Table: share.ForkTraceLogTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, append([]byte("/tracetx2/"), hash.Bytes()...), &kv.ReadOption{Table: share.TraceLogTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	res = &types.TraceTx2{}
	err = res.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) ReadTx(ctx context.Context, hash common.Hash) (data *types.Tx, err error) {
	var bytesRes []byte

	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/tx/"), hash.Bytes()...), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, append([]byte("/tx/"), hash.Bytes()...), &kv.ReadOption{Table: share.TxTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	data = &types.Tx{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.Hash = hash
	}
	return
}

func (s *StorageImpl) ReadTxByIndex(ctx context.Context, index *field.BigInt) (data *types.Tx, err error) {
	i := &field.BigInt{}
	data = &types.Tx{}

	totalAll, err := s.ReadTxTotal(ctx)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadTxTotal(ctx, s.ForkDB)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadTxTotalIndex(ctx, s.ForkDB)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadTxByIndex(ctx, s.ForkDB, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadTxByIndex(ctx, s.FullDB, index)
	}
}

func (s *StorageImpl) ReadTxTotal(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/all/tx/total"), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, []byte("/all/tx/total"), &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadRt(ctx context.Context, hash common.Hash) (data *types.Rt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/rt/"), hash.Bytes()...), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, append([]byte("/rt/"), hash.Bytes()...), &kv.ReadOption{Table: share.TxTbl})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	data = &types.Rt{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.TxHash = hash
	}
	return
}

func (s *StorageImpl) ReadErc20Total(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc20/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc20/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc721Total(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc721/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc721/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc1155Total(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc1155/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc1155/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc20Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc20Transfer{}

	totalAll, err := s.ReadErc20Total(ctx)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc20Total(ctx, s.ForkDB)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc20Index(ctx, s.ForkDB)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc20Transfer(ctx, s.ForkDB, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc20Transfer(ctx, s.FullDB, index)
	}
}

func (s *StorageImpl) ReadErc721Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc721Transfer{}

	totalAll, err := s.ReadErc721Total(ctx)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc721Total(ctx, s.ForkDB)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc721Index(ctx, s.ForkDB)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc721Transfer(ctx, s.ForkDB, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc721Transfer(ctx, s.FullDB, index)
	}
}

func (s *StorageImpl) ReadErc1155Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	i := &field.BigInt{}
	data = &types.Erc1155Transfer{}

	totalAll, err := s.ReadErc1155Total(ctx)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc1155Total(ctx, s.ForkDB)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc1155Index(ctx, s.ForkDB)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc1155Transfer(ctx, s.ForkDB, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc1155Transfer(ctx, s.FullDB, index)
	}
}

func (s *StorageImpl) ReadErc20ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc20/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc20/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc721ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc721/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc721/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc1155ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	totalFork := &field.BigInt{}
	totalFull := &field.BigInt{}
	total = &field.BigInt{}

	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc1155/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFork.SetBytes(bytesRes)
	}

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc1155/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFull = field.NewInt(0)
			err = nil
		} else {
			return
		}
	} else {
		totalFull.SetBytes(bytesRes)
	}
	total.Add(totalFork).Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc20ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	i := &field.BigInt{}
	data = &field.BigInt{}

	totalAll, err := s.ReadErc20ContractTotal(ctx, contract)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc20ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc20ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc20ContractTransfer(ctx, s.ForkDB, contract, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc20ContractTransfer(ctx, s.FullDB, contract, index)
	}
}

func (s *StorageImpl) ReadErc721ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	i := &field.BigInt{}
	data = &field.BigInt{}

	totalAll, err := s.ReadErc721ContractTotal(ctx, contract)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc721ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc721ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc721ContractTransfer(ctx, s.ForkDB, contract, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc721ContractTransfer(ctx, s.FullDB, contract, index)
	}
}

func (s *StorageImpl) ReadErc1155ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	i := &field.BigInt{}
	data = &field.BigInt{}

	totalAll, err := s.ReadErc1155ContractTotal(ctx, contract)
	if err != nil {
		return
	}

	totalFork, err := forkdb.ReadErc1155ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			totalFork = field.NewInt(0)
			err = nil
		} else {
			return
		}
	}

	if totalAll.Sub(totalFork).Cmp(index) == -1 {
		i, err = forkdb.ReadErc1155ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				i = field.NewInt(1)
				err = nil
			} else {
				return
			}
		}
		return forkdb.ReadErc1155ContractTransfer(ctx, s.ForkDB, contract, i.Add(totalFork).Sub(field.NewInt(1)))
	} else {
		return fulldb.ReadErc1155ContractTransfer(ctx, s.FullDB, contract, index)
	}
}

func (s *StorageImpl) GetErc20ContractTransfer(ctx context.Context, contract common.Address, offset, limit int64) (data []*types.Erc20Transfer, total *field.BigInt, err error) {
	transfer := &types.Erc20Transfer{}
	index := &field.BigInt{}
	data = make([]*types.Erc20Transfer, 0)
	total, err = s.ReadErc20ContractTotal(ctx, contract)
	if err != nil {
		if err == kv.NotFound {
			return data, field.NewInt(0), nil
		}
		return nil, field.NewInt(0), err
	}
	if total.Cmp(field.NewInt(0)) == 0 {
		return data, field.NewInt(0), nil
	}

	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		index, err = s.ReadErc20ContractTransfer(ctx, contract, p)
		if err != nil {
			return nil, total, err
		}
		transfer, err = s.ReadErc20Transfer(ctx, index)
		if err != nil {
			return nil, total, err
		}
		data = append(data, transfer)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return data, total, nil
}

func (s *StorageImpl) GetErc721ContractTransfer(ctx context.Context, contract common.Address, offset, limit int64) (data []*types.Erc721Transfer, total *field.BigInt, err error) {
	transfer := &types.Erc721Transfer{}
	index := &field.BigInt{}
	data = make([]*types.Erc721Transfer, 0)
	total, err = s.ReadErc721ContractTotal(ctx, contract)
	if err != nil {
		if err == kv.NotFound {
			return data, field.NewInt(0), nil
		}
		return nil, field.NewInt(0), err
	}
	if total.Cmp(field.NewInt(0)) == 0 {
		return data, field.NewInt(0), nil
	}

	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		index, err = s.ReadErc721ContractTransfer(ctx, contract, p)
		if err != nil {
			return nil, total, err
		}
		transfer, err = s.ReadErc721Transfer(ctx, index)
		if err != nil {
			return nil, total, err
		}
		data = append(data, transfer)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return data, total, nil
}

func (s *StorageImpl) GetErc1155ContractTransfer(ctx context.Context, contract common.Address, offset, limit int64) (data []*types.Erc1155Transfer, total *field.BigInt, err error) {
	transfer := &types.Erc1155Transfer{}
	index := &field.BigInt{}
	data = make([]*types.Erc1155Transfer, 0)
	total, err = s.ReadErc1155ContractTotal(ctx, contract)
	if err != nil {
		if err == kv.NotFound {
			return data, field.NewInt(0), nil
		}
		return nil, field.NewInt(0), err
	}
	if total.Cmp(field.NewInt(0)) == 0 {
		return data, field.NewInt(0), nil
	}

	begin, end := ParsePage(total, offset, limit)
	p := begin

	for {
		index, err = s.ReadErc1155ContractTransfer(ctx, contract, p)
		if err != nil {
			return nil, total, err
		}
		transfer, err = s.ReadErc1155Transfer(ctx, index)
		if err != nil {
			return nil, total, err
		}
		data = append(data, transfer)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	return data, total, nil
}

func ParsePage(num *field.BigInt, offset, limit int64) (*field.BigInt, *field.BigInt) {
	if uint64(offset) >= num.ToUint64() {
		offset = 0
	}

	n := field.BigInt(*DecodeBig(num.String()))

	n.Add(field.NewInt(-offset))
	beginHex := n.String()

	n.Add(field.NewInt(-(limit - 1)))
	endHex := n.String()
	if n.Cmp(field.NewInt(0)) <= 0 {
		endHex = "0x1"
	}

	begin := field.BigInt(*DecodeBig(beginHex))
	end := field.BigInt(*DecodeBig(endHex))

	return &begin, &end
}

func DecodeBig(num string) *big.Int {
	res, _ := hexutil.DecodeBig(num)
	return res
}
