package storage

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

var St *StorageImpl

var _ Storage = (*StorageImpl)(nil)

type StorageImpl struct {
	ForkDB *mdbx.MdbxDB
	FullDB *mdbx.MdbxDB
}

func NewStorage(path string) *StorageImpl {
	mdbx.ForkDB = mdbx.NewMdbx(path + "/fork")
	mdbx.DB = mdbx.NewMdbx(path)
	return &StorageImpl{
		ForkDB: mdbx.ForkDB,
		FullDB: mdbx.DB,
	}
}

func (s *StorageImpl) ReadAccount(ctx context.Context, addr common.Address) (acc *types.Account, err error) {
	var bytesRes []byte

	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/info/"), addr.Bytes()...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return nil, err
	}
	accFork := &types.Account{}
	err = accFork.Unmarshal(bytesRes)
	if err == nil {
		accFork.Owner = addr
	}

	bytesRes, err = s.FullDB.Get(ctx, append([]byte("/info/"), addr.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return nil, err
	}
	accFull := &types.Account{}
	err = accFull.Unmarshal(bytesRes)
	if err == nil {
		accFull.Owner = addr
	}

	if accFork != nil && accFull != nil {
		accFork.TokenTotalSupply = accFull.TokenTotalSupply
		accFork.NftTotalSupply = accFull.NftTotalSupply
		return accFork, err
	} else if accFork != nil {
		return accFork, err
	} else if accFull != nil {
		return accFull, err
	}

	return
}

func (s *StorageImpl) ReadAccountTxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)

	return
}

func (s *StorageImpl) ReadAccountTxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (hash common.Hash, err error) {
	var bytesRes []byte
	var i *field.BigInt
	total, err := forkdb.ReadAccountTxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressTxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/tx/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
	}
	hash.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadAccountTxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (tx *types.Tx, err error) {
	var bytesRes []byte
	var i *field.BigInt
	total, err := forkdb.ReadAccountTxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressTxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
		hash := common.BytesToHash(bytesRes)
		return forkdb.ReadTx(ctx, s.ForkDB, hash)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/tx/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
		hash := common.BytesToHash(bytesRes)
		return fulldb.ReadTx(ctx, s.FullDB, hash)
	}
}

func (s *StorageImpl) ReadAccountITxTotal(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)

	return
}

func (s *StorageImpl) ReadAccountITxIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.InternalTxKey, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountITxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressITxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/itx/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
	}
	data = &types.InternalTxKey{}
	err = data.Unmarshal(bytesRes)

	return

}

func (s *StorageImpl) ReadAccountITxByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (itx *types.InternalTx, err error) {
	var bytesRes []byte
	var i *field.BigInt
	total, err := forkdb.ReadAccountTxTotal(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressITxIndex(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
		data := &types.InternalTxKey{}
		err = data.Unmarshal(bytesRes)
		return forkdb.ReadITx(ctx, s.ForkDB, data.TransactionHash, &data.Index)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/itx/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
		data := &types.InternalTxKey{}
		err = data.Unmarshal(bytesRes)
		return fulldb.ReadITx(ctx, s.FullDB, data.TransactionHash, &data.Index)
	}
}

func (s *StorageImpl) ReadAccountErc20Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc20Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc20TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc20Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc20Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc20/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
	}

	erc20TransferIndex = &field.BigInt{}
	erc20TransferIndex.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadAccountErc20ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc20Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc20Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc20/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
		erc20TransferIndex := &field.BigInt{}
		erc20TransferIndex.SetBytes(bytesRes)

		return forkdb.ReadErc20Transfer(ctx, s.ForkDB, erc20TransferIndex)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc20/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
		erc20TransferIndex := &field.BigInt{}
		erc20TransferIndex.SetBytes(bytesRes)

		return fulldb.ReadErc20Transfer(ctx, s.FullDB, erc20TransferIndex)
	}
}

func (s *StorageImpl) ReadAccountErc721Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc721Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc721TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc721Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc721Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc721/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
	}

	erc721TransferIndex = &field.BigInt{}
	erc721TransferIndex.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadAccountErc721ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc721Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc721Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc721/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
		erc721TransferIndex := &field.BigInt{}
		erc721TransferIndex.SetBytes(bytesRes)

		return forkdb.ReadErc721Transfer(ctx, s.ForkDB, erc721TransferIndex)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc721/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
		erc721TransferIndex := &field.BigInt{}
		erc721TransferIndex.SetBytes(bytesRes)

		return fulldb.ReadErc721Transfer(ctx, s.FullDB, erc721TransferIndex)
	}
}

func (s *StorageImpl) ReadAccountErc1155Total(ctx context.Context, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.ForkAccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadAccountErc1155Index(ctx context.Context, addr common.Address, index *field.BigInt) (erc1155TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc1155Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc1155Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc1155/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
	}

	erc1155TransferIndex = &field.BigInt{}
	erc1155TransferIndex.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadAccountErc1155ByIndex(ctx context.Context, addr common.Address, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadAccountErc1155Total(ctx, s.ForkDB, addr)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadAddressErc1155Index(ctx, s.ForkDB, addr)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/erc1155/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkAccountsTbl})
		if err != nil {
			return
		}
		erc1155TransferIndex := &field.BigInt{}
		erc1155TransferIndex.SetBytes(bytesRes)

		return forkdb.ReadErc1155Transfer(ctx, s.ForkDB, erc1155TransferIndex)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc1155/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
		if err != nil {
			return
		}
		erc1155TransferIndex := &field.BigInt{}
		erc1155TransferIndex.SetBytes(bytesRes)

		return fulldb.ReadErc1155Transfer(ctx, s.FullDB, erc1155TransferIndex)
	}
}

func (s *StorageImpl) ReadBlock(ctx context.Context, blockNum *field.BigInt) (bk *types.Block, err error) {
	var bytesRes []byte

	bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/block/"), blockNum.Bytes()...), &kv.ReadOption{Table: share.ForkBlockTbl})
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			bytesRes, err = s.FullDB.Get(ctx, []byte("/block/"), &kv.ReadOption{Table: share.ForkBlockTbl})
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
	forkKey := make([]byte, 0, len([]byte("/fork/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
	forkKey = append(forkKey, []byte("/fork/block/")...)
	forkKey = append(forkKey, blockNum.Bytes()...)
	forkKey = append(forkKey, byte('/'))
	forkKey = append(forkKey, index.Bytes()...)

	var bytesRes []byte

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
	forkKey := make([]byte, 0, len([]byte("/fork/block/"))+len(blockNum.Bytes())+len(index.Bytes())+1)
	forkKey = append(forkKey, []byte("/fork/block/")...)
	forkKey = append(forkKey, blockNum.Bytes()...)
	forkKey = append(forkKey, byte('/'))
	forkKey = append(forkKey, index.Bytes()...)

	var bytesRes []byte

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
	}
	hash := common.BytesToHash(bytesRes)
	return forkdb.ReadTx(ctx, s.ForkDB, hash)
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

	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/home"), &kv.ReadOption{Table: share.ForkHomeTbl})
	if err != nil {
		return
	}
	homeFork := &types.Home{}
	err = homeFork.Unmarshal(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, []byte("/home"), &kv.ReadOption{Table: share.HomeTbl})
	if err != nil {
		return
	}
	homeFull := &types.Home{}
	err = homeFull.Unmarshal(bytesRes)

	if homeFull != nil && homeFork != nil {
		home = &types.Home{
			BlockNumber:  *homeFull.BlockNumber.Add(&homeFork.BlockNumber),
			TxTotal:      *homeFull.TxTotal.Add(&homeFork.TxTotal),
			AddressTotal: *homeFull.AddressTotal.Add(&homeFork.AddressTotal),
			Erc20Total:   *homeFull.Erc20Total.Add(&homeFork.Erc20Total),
			Erc721Total:  *homeFull.Erc721Total.Add(&homeFork.Erc721Total),
			Erc1155Total: *homeFull.Erc1155Total.Add(&homeFork.Erc1155Total),
			Blocks:       homeFork.Blocks,
			Txs:          homeFork.Txs,
			DateTxs:      homeFull.DateTxs,
			DateTxsByte:  homeFork.DateTxsByte,
		}
	} else if homeFull != nil {
		home = homeFull
	} else if homeFork != nil {
		home = homeFork
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
	var bytesRes []byte
	var i *field.BigInt
	total, err := forkdb.ReadITxTotal(ctx, s.ForkDB, hash)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadITxIndex(ctx, s.ForkDB, hash)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/iTx/"), hash.Bytes()...), append([]byte("/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkTxTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/iTx/"), hash.Bytes()...), append([]byte("/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.TxTbl})
		if err != nil {
			return
		}
	}

	data = &types.InternalTx{}
	err = data.Unmarshal(bytesRes)
	if err == nil {
		data.TransactionHash = hash
	}
	return
}

func (s *StorageImpl) ReadITxTotal(ctx context.Context, hash common.Hash) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/iTx/"), hash.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/iTx/"), hash.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
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
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadTxTotal(ctx, s.ForkDB)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadTxTotalIndex(ctx, s.ForkDB)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/all/tx/"), i.Bytes()...), &kv.ReadOption{Table: share.ForkTxTbl})
		if err != nil {
			return
		}
		hash := common.BytesToHash(bytesRes)
		return forkdb.ReadTx(ctx, s.ForkDB, hash)
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append([]byte("/all/tx/"), diff.Bytes()...), &kv.ReadOption{Table: share.TxTbl})
		if err != nil {
			return
		}
		hash := common.BytesToHash(bytesRes)
		return fulldb.ReadTx(ctx, s.FullDB, hash)
	}
}

func (s *StorageImpl) ReadTxTotal(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/all/tx/total"), &kv.ReadOption{Table: share.ForkTxTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, []byte("/all/tx/total"), &kv.ReadOption{Table: share.TxTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
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
	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc20/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc20/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc721Total(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc721/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc721/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc1155Total(ctx context.Context) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, []byte("/fork/erc1155/total"), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, []byte("/erc1155/total"), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc20Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc20Total(ctx, s.ForkDB)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc20Index(ctx, s.ForkDB)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/erc20/"), i.Bytes()...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append([]byte("/erc20/"), diff.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &types.Erc20Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) ReadErc721Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc721Total(ctx, s.ForkDB)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc721Index(ctx, s.ForkDB)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/erc721/"), i.Bytes()...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append([]byte("/erc721/"), diff.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &types.Erc721Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) ReadErc1155Transfer(ctx context.Context, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc1155Total(ctx, s.ForkDB)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc1155Index(ctx, s.ForkDB)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append([]byte("/fork/erc1155/"), i.Bytes()...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append([]byte("/erc1155/"), diff.Bytes()...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &types.Erc1155Transfer{}
	err = data.Unmarshal(bytesRes)
	return
}

func (s *StorageImpl) ReadErc20ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc20/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc20/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc721ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc721/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc721/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc1155ContractTotal(ctx context.Context, contract common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc1155/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.ForkTransferTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)

	bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc1155/"), contract.Bytes()...), []byte("/total")...), &kv.ReadOption{Table: share.TransferTbl})
	if err != nil {
		return
	}
	totalFull := &field.BigInt{}
	totalFull.SetBytes(bytesRes)

	total.Add(totalFull)
	return
}

func (s *StorageImpl) ReadErc20ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc20ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc20ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc20/"), contract.Bytes()...), append([]byte("/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc20/"), contract.Bytes()...), append([]byte("/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &field.BigInt{}
	data.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadErc721ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc721ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc721ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc721/"), contract.Bytes()...), append([]byte("/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc721/"), contract.Bytes()...), append([]byte("/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &field.BigInt{}
	data.SetBytes(bytesRes)
	return
}

func (s *StorageImpl) ReadErc1155ContractTransfer(ctx context.Context, contract common.Address, index *field.BigInt) (data *field.BigInt, err error) {
	var bytesRes []byte
	var i *field.BigInt

	total, err := forkdb.ReadErc1155ContractTotal(ctx, s.ForkDB, contract)
	if err != nil {
		return
	}

	if total.Cmp(index) >= 0 {
		i, err = forkdb.ReadErc1155ContractIndex(ctx, s.ForkDB, contract)
		if err != nil {
			return
		}
		i.Add(index)
		bytesRes, err = s.ForkDB.Get(ctx, append(append([]byte("/fork/erc1155/"), contract.Bytes()...), append([]byte("/"), i.Bytes()...)...), &kv.ReadOption{Table: share.ForkTransferTbl})
		if err != nil {
			return
		}
	} else {
		i = index
		diff := i.Sub(total)
		bytesRes, err = s.FullDB.Get(ctx, append(append([]byte("/erc1155/"), contract.Bytes()...), append([]byte("/"), diff.Bytes()...)...), &kv.ReadOption{Table: share.TransferTbl})
		if err != nil {
			return
		}
	}

	data = &field.BigInt{}
	data.SetBytes(bytesRes)
	return
}
