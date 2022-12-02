package service

import (
	"errors"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sort"
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
	home, err := store.GetHome()
	if err != nil {
		return nil, err
	}
	//DateTxs
	dateTxs := GetDateTxs(home.DateTxs)

	// Blocks
	l := len(home.Blocks)
	blocks := make([]*types.HomeBlock, 0)
	var totalTxs uint64
	var beginTime, endTime uint64

	for i := l - 1; i >= 0; i-- {
		blocks = append(blocks, &types.HomeBlock{
			Number:            home.Blocks[i].Number.String(),
			Timestamp:         home.Blocks[i].Timestamp.ToUint64(),
			Miner:             home.Blocks[i].Miner.String(),
			GasUsed:           home.Blocks[i].GasUsed.String(),
			TransactionsTotal: home.Blocks[i].TransactionsTotal.ToUint64(),
		})
		totalTxs += home.Blocks[i].TransactionsTotal.ToUint64()
		if i == 0 {
			beginTime = home.Blocks[i].Timestamp.ToUint64()
		}
		if i == l-1 {
			endTime = home.Blocks[i].Timestamp.ToUint64()
		}
	}

	t := endTime - beginTime

	// Txs
	txs := make([]*types.HomeTx, 0)
	for i := len(home.Txs) - 1; i >= 0; i-- {
		txs = append(txs, &types.HomeTx{
			Hash:        home.Txs[i].Hash.Hex(),
			From:        home.Txs[i].From.Hex(),
			To:          home.Txs[i].To.Hex(),
			GasPrice:    home.Txs[i].GasPrice.StringPointer(),
			Gas:         home.Txs[i].Gas.StringPointer(),
			CreatedTime: home.Txs[i].Timestamp.ToUint64(),
		})
	}
	// metrics
	resp := make(map[string]interface{})
	resp["dateTxs"] = dateTxs
	resp["metrics"] = GetHomeMetrics(home, dateTxs, totalTxs, t)
	resp["blocks"] = blocks
	resp["txs"] = txs
	return resp, nil
}

func GetHomeMetrics(home *types.Home, dateTxs []map[string]string, totalTxs, t uint64) map[string]interface{} {
	metrics := make(map[string]interface{})
	metrics["address"] = home.AddressTotal.String()
	metrics["tx"] = home.TxTotal.String()
	blockNum, err := store.GetBlockTotal()
	if err != nil {
		return nil
	}
	metrics["block"] = blockNum.String()
	metrics["avgBlockTime"] = 3
	metrics["dailyTx"] = 0
	if len(dateTxs) > 0 {
		metrics["dailyTx"] = dateTxs[len(dateTxs)-1]["txCount"]
	}
	if t == 0 {
		metrics["tps"] = 0
	} else {
		metrics["tps"] = totalTxs / t
	}

	metrics["diff"] = 0
	metrics["erc20"] = home.Erc20Total.String()
	metrics["erc721"] = home.Erc721Total.String()
	metrics["erc1155"] = home.Erc1155Total.String()
	return metrics
}

func GetDateTxs(dateTxs map[string]*field.BigInt) []map[string]string {
	keys := make([]string, 0)
	for k := range dateTxs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
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
			account, err := store.GetAccount(common.HexToAddress(f.Keyword))
			if err != nil && err != kv.NotFound {
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
			block, err := store.GetBlock(&num)
			if err != nil && err != kv.NotFound {
				return nil, err
			}
			if block != nil {
				resp["type"] = searchBlock
				return resp, nil
			}
		}

		transaction, err := store.GetTx(common.HexToHash(f.Keyword))
		if err != nil && err != kv.NotFound {
			return nil, err
		}
		if transaction != nil {
			resp["type"] = searchTxnHash
			return resp, nil
		}
		//accounts, err := rawdb.GetAccountsByNameOrSymbol(f.Keyword)
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

	block, err := store.GetBlock(&num)
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
	txs := make([]string, 0)
	for _, transaction := range block.Transactions {
		txs = append(txs, transaction.String())
	}

	resp := &types.BlockResp{
		BaseFeePerGas:     block.BaseFee.StringPointer(),
		Difficulty:        block.Difficulty.String(),
		ExtraData:         hexutil.Encode(block.Extra),
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

func ListFullFieldBlocks(pager *types.Pager) ([]*types.ListBlockResp, uint64, error) {
	total, err := store.GetBlockTotal()
	if err != nil {
		return nil, 0, err
	}
	blocks, err := store.ListBlocks(total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
	}
	resp := make([]*types.ListBlockResp, len(blocks))
	for i, block := range blocks {
		resp[i] = &types.ListBlockResp{
			Number:            block.Number.String(),
			Timestamp:         block.TimeStamp.ToUint64(),
			TransactionsTotal: block.TransactionTotal.ToUint64(),
			Miner:             block.Coinbase.String(),
			GasLimit:          block.GasLimit.String(),
			GasUsed:           block.GasUsed.String(),
			BaseFeePerGas:     block.BaseFee.StringPointer(),
		}
	}
	return resp, total.ToUint64(), nil
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
func GetBlockTxs(blockNum string, pager *types.Pager) ([]*types.ListTransactionResp, uint64, error) {
	n := new(big.Int)
	n, ok := n.SetString(blockNum, 10)
	if !ok {
		return nil, 0, errors.New("parse block num error")
	}
	num := field.BigInt(*n)

	block, err := store.GetBlock(&num)
	if err != nil {
		return nil, 0, err
	}

	total := block.TransactionTotal
	txs, err := store.ListBlockTxs(&total, &num, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
	}
	resp := make([]*types.ListTransactionResp, 0)
	addresses := make(map[string]common.Address)
	for _, tx := range txs {
		t := &types.ListTransactionResp{
			Hash:        tx.Hash.Hex(),
			Method:      tx.Method.String(),
			BlockHash:   tx.BlockNum.String(),
			BlockNumber: DecodeBig(tx.BlockNum.String()).String(),
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			Gas:         tx.Gas.StringPointer(),
			GasPrice:    tx.GasPrice.StringPointer(),
			Value:       tx.Value.StringPointer(),
			CreatedTime: tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		if tx.To != nil {
			addresses[tx.To.String()] = *tx.To
		}
	}
	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, 0, err
	}
	for _, t := range resp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
			if from.Erc20 || from.Erc721 || from.Erc1155 {
				t.FromContract = true
			}
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
				t.ToContract = true
			}
		}
	}
	return resp, total.ToUint64(), nil
}
