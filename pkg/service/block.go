package service

import (
	"errors"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	allFilters = 1
	addresses  = 2

	searchNull    = 1
	searchAddress = 2
	searchBlock   = 3
	searchTxnHash = 4
	searchName    = 5
)

func Home() (map[string]interface{}, error) {
	// 面板 趋势图 block列表 tx列表
	page := &types.Pager{
		Offset: 0,
		Limit:  10,
	}
	blocks, _, err := ListBaseFieldBlocks(page)
	if err != nil {
		return nil, err
	}
	ListTxs(page)

	resp := make(map[string]interface{})
	resp["metrics"] = nil
	resp["metrics"] = nil
	resp["blocks"] = blocks
	resp["txs"] = nil
	return resp, nil
}

func Search(f *types.SearchFilter) (map[string]interface{}, error) {
	resp := map[string]interface{}{}
	if f.Keyword == "" {
		resp["type"] = searchNull
		return resp, nil
	}
	switch f.Type {
	case allFilters:
		//address := common.IsHexAddress(f.Keyword)
		//if address {
		//	// common.HexToAddress(f.Keyword).Hex()
		//	account, err := store.GetAccount()
		//	if err != nil && err != gorm.ErrRecordNotFound {
		//		return nil, err
		//	}
		//	if account.ID > 0 {
		//		resp["type"] = searchAddress
		//		return resp, nil
		//	}
		//}
		//number := IsNumber(f.Keyword)
		//if number {
		//	blockNum, _ := strconv.ParseInt(f.Keyword, 10, 64)
		//	block, err := store.GetBlock(field.BigInt(*big.NewInt(blockNum)))
		//	if err != nil && err != gorm.ErrRecordNotFound {
		//		return nil, err
		//	}
		//	if block.ID > 0 {
		//		resp["type"] = searchBlock
		//		return resp, nil
		//	}
		//}
		//transaction, err := store.GetTransaction(common.HexToHash(f.Keyword).Hex())
		//if err != nil && err != gorm.ErrRecordNotFound {
		//	return nil, err
		//}
		//if transaction.ID > 0 {
		//	resp["type"] = searchTxnHash
		//	return resp, nil
		//}
		//accounts, err := store.GetAccountsByNameOrSymbol(f.Keyword)
		//if err != nil && err != gorm.ErrRecordNotFound {
		//	return nil, err
		//}
		//if len(accounts) > 0 {
		//	resp["type"] = searchName
		//	return resp, nil
		//}
		//resp["type"] = searchNull
		return resp, nil
	default:
		return nil, response.ErrInvalidParameter
	}
}

func GetBlock(blockNum string) error {
	_, err := NumToHex(blockNum)
	if err != nil {
		return err
	}
	// get block from db

	return nil
}

func ListBlocks(pager *types.Pager) ([]*types.Block, string, error) {
	num, err := store.GetBlockNum(nil)
	if err != nil {
		return nil, "0", err
	}
	blocks := make([]*types.Block, 0)
	if num == "" {
		return blocks, "0", nil
	}
	numBig := DecodeBig("0" + num)
	begin, end := ParseBlockPage(numBig, pager.Offset, pager.Limit)
	p := begin
	for {
		block, err := store.GetBlock(nil, EncodeBig(p))
		if err != nil {
			return nil, "0", err
		}
		blocks = append(blocks, block)
		if p.Cmp(end) == 0 {
			break
		}
		p = BigIntReduce(p, 1)
	}

	return blocks, numBig.String(), nil
}

func ListBaseFieldBlocks(pager *types.Pager) ([]*types.HomeBlock, string, error) {
	blocks, total, err := ListBlocks(pager)
	if err != nil {
		return nil, "", err
	}
	res := make([]*types.HomeBlock, len(blocks))
	for i, block := range blocks {
		res[i] = &types.HomeBlock{
			Number:            block.Number.String(),
			Timestamp:         block.Time.ToUint64(),
			Miner:             block.Coinbase.String(),
			GasUsed:           block.GasUsed.String(),
			TransactionsTotal: block.TransactionTotal.ToUint64(),
		}
	}
	return res, total, nil
}

func ListFullFieldBlocks(pager *types.Pager) ([]*types.BlockResp, string, error) {
	blocks, total, err := ListBlocks(pager)
	if err != nil {
		return nil, "0", err
	}
	resp := make([]*types.BlockResp, len(blocks))
	for i, block := range blocks {
		nonce, err := block.Nonce.MarshalText()
		if err != nil {
			return nil, "0", err
		}
		bloom, err := block.Bloom.MarshalText()
		if err != nil {
			return nil, "0", err
		}
		txs := []string{}
		for _, transaction := range block.Transactions {
			txs = append(txs, transaction.String())
		}
		resp[i] = &types.BlockResp{
			BaseFeePerGas:     block.BaseFee.StringPointer(),
			Difficulty:        block.Difficulty.String(),
			ExtraData:         string(block.Extra),
			GasLimit:          block.GasLimit.String(),
			GasUsed:           block.GasUsed.String(),
			Hash:              block.Hash.Hex(),
			LogsBloom:         string(bloom),
			Miner:             block.Coinbase.String(),
			MixHash:           block.MixDigest.String(),
			Nonce:             string(nonce),
			Number:            block.Number.String(),
			ParentHash:        block.ParentHash.Hex(),
			ReceiptsRoot:      block.ReceiptHash.Hex(),
			Sha3Uncles:        block.UncleHash.Hex(),
			Size:              block.Size.String(),
			StateRoot:         block.Root.Hex(),
			Timestamp:         block.Time.ToUint64(),
			TotalDifficulty:   block.TotalDifficulty.ToUint64(),
			Transactions:      txs,
			TransactionsTotal: block.TransactionTotal.ToUint64(),
			//TransactionsRoot:  block,
		}
	}
	return resp, total, nil
}

func BigIntReduce(n *big.Int, num int64) *big.Int {
	m := new(big.Int)
	m.SetInt64(-num)
	m.Add(n, m)
	return m
}

func ParseBlockPage(num *big.Int, offset, limit int64) (*big.Int, *big.Int) {
	if offset >= num.Int64() {
		offset = 0
	}
	begin := BigIntReduce(num, offset)
	end := BigIntReduce(begin, limit-1)

	if end.Int64() <= 0 {
		e := new(big.Int)
		end = e.SetInt64(1)
	}
	return begin, end
}

func NumToHex(num string) (string, error) {
	n := new(big.Int)
	n, ok := n.SetString(num, 10)
	if !ok {
		return "", errors.New("parse block num error")
	}
	return hexutil.EncodeBig(n), nil
}

func DecodeBig(num string) *big.Int {
	res, _ := hexutil.DecodeBig(num)
	return res
}

func EncodeBig(num *big.Int) string {
	return hexutil.EncodeBig(num)
}
