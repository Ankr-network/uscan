package service

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"unicode"
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
	home, err := store.ReadHome(context.Background(), nil)
	if err != nil {

	}
	//DateTxs
	dateTxs := GetDateTxs(home.DateTxs)

	// Blocks
	blocks := make([]*types.HomeBlock, 0)
	var totalTxs int
	var beginTime, endTime int
	for i, block := range home.Blocks {
		blocks = append(blocks, &types.HomeBlock{
			Number:            block.Number.String(),
			Timestamp:         uint64(block.Timestamp),
			Miner:             block.Miner.String(),
			GasUsed:           block.GasUsed.String(),
			TransactionsTotal: uint64(block.TransactionsTotal),
		})
		totalTxs += block.TransactionsTotal
		if i == 0 {
			endTime = block.Timestamp
		}
		if i == len(blocks)-1 {
			beginTime = block.Timestamp
		}
	}
	t := endTime - beginTime

	// Txs
	txs := make([]*types.HomeTx, 0)
	for _, tx := range home.Txs {
		txs = append(txs, &types.HomeTx{
			Hash:        tx.Hash.Hex(),
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			GasPrice:    tx.GasPrice.StringPointer(),
			Gas:         tx.Gas.StringPointer(),
			CreatedTime: 0, // TODO 少时间
		})
	}

	// metrics
	resp := make(map[string]interface{})
	resp["dateTxs"] = home.DateTxs
	resp["metrics"] = GetHomeMetrics(home, dateTxs, totalTxs, t)
	resp["blocks"] = blocks
	resp["txs"] = txs
	return resp, nil
}

func GetHomeMetrics(home *types.Home, dateTxs []map[string]string, totalTxs, t int) map[string]interface{} {
	metrics := make(map[string]interface{})
	metrics["address"] = home.AddressTotal.String()
	metrics["tx"] = home.TxTotal.String()
	metrics["block"] = home.BlockNumber.String()
	metrics["avgBlockTime"] = 3
	metrics["dailyTx"] = 0
	if len(dateTxs) > 0 {
		metrics["dailyTx"] = dateTxs[len(dateTxs)-1]
	}
	if t == 0 {
		metrics["tps"] = 0
	} else {
		metrics["tps"] = totalTxs / t
	}
	// TODO 舍弃
	metrics["diff"] = 0
	metrics["erc20"] = 0
	metrics["erc721"] = 0
	return metrics
}

func GetDateTxs(dateTxs map[string]*field.BigInt) []map[string]string {
	keys := make([]string, 0)
	for k := range dateTxs {
		keys = append(keys, k)
	}
	resp := make([]map[string]string, 0)
	for _, key := range keys {
		resp = append(resp, map[string]string{"date": key, "txCount": dateTxs[key].String()})
	}
	return resp
}

func Search(f *types.SearchFilter) (map[string]interface{}, error) {
	resp := map[string]interface{}{}
	if f.Keyword == "" {
		resp["type"] = searchNull
		return resp, nil
	}
	switch f.Type {
	case allFilters:
		address := common.IsHexAddress(f.Keyword)
		if address {
			// common.HexToAddress(f.Keyword).Hex()
			account, err := store.ReadAccount(context.Background(), nil, common.HexToAddress(f.Keyword))
			if err != nil {
				return nil, err
			}
			if account != nil {
				resp["type"] = searchAddress
				return resp, nil
			}
		}
		number := IsNumber(f.Keyword)
		if number {
			n := new(big.Int)
			n, ok := n.SetString(f.Keyword, 10)
			if !ok {
				return nil, errors.New("parse block num error")
			}
			num := field.BigInt(*n)
			block, err := store.ReadBlock(context.Background(), nil, &num)
			if err != nil && err != kv.NotFound {
				return nil, err
			}
			if block != nil {
				resp["type"] = searchBlock
				return resp, nil
			}
		}
		transaction, err := store.ReadTx(context.Background(), nil, common.HexToHash(f.Keyword))
		if err != nil && err != kv.NotFound {
			return nil, err
		}
		if transaction != nil {
			resp["type"] = searchTxnHash
			return resp, nil
		}
		//accounts, err := store.GetAccountsByNameOrSymbol(f.Keyword)
		//if err != nil && err != gorm.ErrRecordNotFound {
		//	return nil, err
		//}
		//if len(accounts) > 0 {
		//	resp["type"] = searchName
		//	return resp, nil
		//}
		resp["type"] = searchNull
		return resp, nil
	default:
		return nil, response.ErrInvalidParameter
	}
}

func IsNumber(number string) bool {
	for _, s := range []rune(number) {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

func GetBlock(blockNum string) (*types.BlockResp, error) {
	n := new(big.Int)
	n, ok := n.SetString(blockNum, 10)
	if !ok {
		return nil, errors.New("parse block num error")
	}
	num := field.BigInt(*n)
	block, err := store.ReadBlock(context.Background(), nil, &num)
	if err != nil {
		return nil, err
	}
	nonce, err := block.Nonce.MarshalText()
	if err != nil {
		return nil, err
	}
	bloom, err := block.Bloom.MarshalText()
	if err != nil {
		return nil, err
	}
	txs := []string{}
	for _, transaction := range block.Transactions {
		txs = append(txs, transaction.String())
	}
	resp := &types.BlockResp{
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
		Timestamp:         block.TimeStamp.ToUint64(),
		TotalDifficulty:   block.TotalDifficulty.ToUint64(),
		Transactions:      txs,
		TransactionsTotal: block.TransactionTotal.ToUint64(),
		//TransactionsRoot:  block,
	}
	return resp, nil
}

func ListBlocks(pager *types.Pager) ([]*types.Block, string, error) {
	home, err := store.ReadHome(context.Background(), nil)
	if err != nil {
		return nil, "0", err
	}
	total := home.BlockNumber.String()
	blocks := make([]*types.Block, 0)
	if total == "" {
		return blocks, "0", nil
	}
	begin, end := ParsePage(home.BlockNumber, pager.Offset, pager.Limit)
	p := begin
	for {
		block, err := store.ReadBlock(context.Background(), nil, p)
		if err != nil {
			return nil, "0", err
		}
		blocks = append(blocks, block)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	return blocks, total, nil
}

func ListFullFieldBlocks(pager *types.Pager) ([]*types.ListBlockResp, string, error) {
	blocks, total, err := ListBlocks(pager)
	if err != nil {
		return nil, "0", err
	}
	resp := make([]*types.ListBlockResp, len(blocks))
	for i, block := range blocks {
		resp[i] = &types.ListBlockResp{
			GasLimit:          block.GasLimit.String(),
			GasUsed:           block.GasUsed.String(),
			Miner:             block.Coinbase.String(),
			Number:            block.Number.String(),
			Timestamp:         block.TimeStamp.ToUint64(),
			TransactionsTotal: block.TransactionTotal.ToUint64(),
		}
	}
	return resp, total, nil
}

func ParsePage(num *field.BigInt, offset, limit int64) (*field.BigInt, *field.BigInt) {
	if uint64(offset) >= num.ToUint64() {
		offset = 0
	}
	num.Add(field.NewInt(-offset))
	beginHex := num.String()

	num.Add(field.NewInt(-(limit - 1)))
	endHex := num.String()

	begin := field.BigInt(*DecodeBig(beginHex))
	end := field.BigInt(*DecodeBig(endHex))

	if end.ToUint64() <= 0 {
		end = *(field.NewInt(1))
	}
	return &begin, &end
}

func DecodeBig(num string) *big.Int {
	res, _ := hexutil.DecodeBig(num)
	return res
}

func EncodeBig(num *big.Int) string {
	return hexutil.EncodeBig(num)
}
