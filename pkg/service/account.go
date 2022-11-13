package service

import (
	"context"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func GetAccountInfo(address string) (*types.AccountResp, error) {
	account, err := store.ReadAccount(context.Background(), nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	hexutil.Encode(account.Code)
	resp := &types.AccountResp{
		Owner:            account.Owner.String(),
		Balance:          account.Balance.String(),
		BlockNumber:      account.BlockNumber.StringPointer(),
		Name:             account.Name,
		Symbol:           account.Symbol,
		TokenTotalSupply: account.TokenTotalSupply.StringPointer(),
		NftTotalSupply:   account.NftTotalSupply.StringPointer(),
		Decimals:         account.Decimals.ToUint64(),
		//CreatedTime:      account., TODO
	}
	var creator string
	if account.Creator != nil {
		creator = account.Creator.Hex()
		resp.Creator = &creator
	}
	var txHash string
	if account.TxHash != nil {
		txHash = account.TxHash.Hex()
		resp.TxHash = &txHash
	}

	var code string
	if account.Creator != nil {
		code = hexutil.Encode(account.Code)
		resp.Code = &code
	}
	return resp, nil
}

func GetAccountTxs(pager *types.Pager, address string) (map[string]interface{}, error) {
	num, err := store.ReadAccountTxTotal(context.Background(), nil, common.HexToAddress(address))
	if err != nil {

	}
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Tx, 0)
	for {
		tx, err := store.ReadAccountTxByIndex(context.Background(), nil, common.HexToAddress(address), p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.Cmp(end) == 0 {
			break
		}
		p = BigIntReduce(p, 1)
	}
	txsResp := make([]*types.ListTransactionResp, 0)
	addresses := make(map[string]common.Address)
	for _, tx := range txs {
		var blockNumber *string
		if tx.BlockNum != nil {
			num := DecodeBig(tx.BlockNum.String()).String()
			blockNumber = &num
		}
		t := &types.ListTransactionResp{
			Hash:        tx.Hash.Hex(),
			Method:      nil, // TODO
			BlockHash:   tx.BlockNum.StringPointer(),
			BlockNumber: blockNumber,
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			ToName:      "",
			ToSymbol:    "",
			ToCode:      "",
			Gas:         tx.Gas.StringPointer(),
			GasPrice:    tx.GasPrice.StringPointer(),
			Value:       tx.Value.StringPointer(),
			CreatedTime: 0, // TODO
		}
		txsResp = append(txsResp, t)
		addresses[tx.From.String()] = tx.From
		if tx.To != nil {
			addresses[tx.To.String()] = *tx.To
		}
	}
	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, "0", err
	}
	for _, t := range txsResp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
			t.FromCode = hexutil.Encode(from.Code)
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			t.FromCode = hexutil.Encode(to.Code)
		}
	}

	itxTotal, err := store.ReadAccountITxTotal(context.Background(), nil, common.HexToAddress(address))
	erc20Total, err := store.ReadAccountErc20Total(context.Background(), nil, common.HexToAddress(address))
	erc721Total, err := store.ReadAccountErc721Total(context.Background(), nil, common.HexToAddress(address))
	erc1155Total, err := store.ReadAccountErc1155Total(context.Background(), nil, common.HexToAddress(address))
	otherTotal := map[string]string{
		"internalTotal": itxTotal.String(),
		"erc20Total":    erc20Total.String(),
		"erc721Total":   erc721Total.String(),
		"erc1155Total":  erc1155Total.String(),
	}
	resp := map[string]interface{}{
		"items":      txsResp,
		"total":      num.String(),
		"otherTotal": otherTotal,
	}
	return resp, nil
}

func GetAccountItxs(pager *types.Pager, address string) {
	num, err := store.ReadITxTotal(context.Background(), nil, common.HexToHash(address))
	if err != nil {
		return nil, 0, err
	}
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.InternalTx, 0)
	for {
		tx, err := store.ReadAccountITxByIndex(context.Background(), nil, common.HexToAddress(address), p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.Cmp(end) == 0 {
			break
		}
		p = BigIntReduce(p, 1)
	}

	resp := make([]*types.InternalTxResp, len(txs))
	for i, tx := range txs {
		resp[i] = tx.ToInternalTxResp()
	}
	return resp, total, nil
}

func GetAccountErc20Txns(pager *types.Pager, address string) {
	num, err := store.ReadAccountErc20Total(context.Background(), nil, common.HexToHash(address))
	if err != nil {
		return nil, 0, err
	}
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc20Transfer, 0)
	for {
		tx, err := store.ReadAccountErc20Index(context.Background(), nil, common.HexToAddress(address), p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.Cmp(end) == 0 {
			break
		}
		p = BigIntReduce(p, 1)
	}
}
