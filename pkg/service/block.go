package service

import (
	"errors"
	"fmt"
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
	ListBlocks(page)
	ListTxs(page)

	resp := make(map[string]interface{})
	resp["metrics"] = nil
	resp["metrics"] = nil
	resp["blocks"] = nil
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

func ListBlocks(pager *types.Pager) error {
	// TODO get /block/num
	num := 0

	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	NumToHex(fmt.Sprint(begin))
	NumToHex(fmt.Sprint(end))
	// get block

	return nil
}

func ParsePage(num, offset, limit int) (int, int) {
	if offset >= num {
		offset = 0
	}
	begin := num - offset
	end := num - offset - (limit - 1)
	if end <= 0 {
		end = 1
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
