package service

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ListErc20Txs(pager *types.Pager) ([]*types.Erc20TxResp, string, error) {
	num, err := store.ReadErc20Total(context.Background(), nil)
	if err != nil {

	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc20Transfer, 0)
	for {
		tx, err := store.ReadErc20Transfer(context.Background(), nil, p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	resp := make([]*types.Erc20TxResp, 0)
	addresses := make(map[string]common.Address)

	for _, tx := range txs {
		var blockNumber string
		if tx.BlockNumber != nil {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc20TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			ContractName:    "",
			ContractSymbol:  "",
			Method:          tx.Method,
			From:            tx.From.Hex(),
			//FromName:        "",
			//FromSymbol:      "",
			//FromCode:        "",
			To: tx.To.Hex(),
			//ToName:          "",
			//ToSymbol:        "",
			//ToCode:          "",
			Value:       tx.Amount.String(),
			CreatedTime: 0, // TODO
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, "0", err
	}
	for _, t := range resp {
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
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}
	return resp, total, nil
}

func ListErc721Txs(pager *types.Pager) ([]*types.Erc721TxResp, string, error) {
	num, err := store.ReadErc721Total(context.Background(), nil)
	if err != nil {

	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc721Transfer, 0)
	for {
		tx, err := store.ReadErc721Transfer(context.Background(), nil, p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	resp := make([]*types.Erc721TxResp, 0)
	addresses := make(map[string]common.Address)

	for _, tx := range txs {
		var blockNumber string
		if tx.BlockNumber != nil {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc721TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			Method:          tx.Method,
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenId.ToUint64(),
			CreatedTime:     0, // TODO
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, "0", err
	}
	for _, t := range resp {
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
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}
	return resp, total, nil
}

func ListErc1155Txs(pager *types.Pager) ([]*types.Erc1155TxResp, string, error) {
	num, err := store.ReadErc1155Total(context.Background(), nil)
	if err != nil {

	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc1155Transfer, 0)
	for {
		tx, err := store.ReadErc1155Transfer(context.Background(), nil, p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	resp := make([]*types.Erc1155TxResp, 0)
	addresses := make(map[string]common.Address)

	for _, tx := range txs {
		var blockNumber string
		if tx.BlockNumber != nil {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc1155TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			Method:          tx.Method,
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenID.ToUint64(),
			Value:           tx.Quantity.String(),
			CreatedTime:     0, // TODO
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, "0", err
	}
	for _, t := range resp {
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
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}
	return resp, total, nil
}

func GetTraceTx(hash common.Hash) (*types.TraceTxResp, error) {
	t, err := store.ReadTraceTx(context.Background(), nil, hash)
	if err != nil {

	}
	resp := &types.TraceTxResp{
		Res:    t.Res,
		LogNum: t.LogNum.String(),
	}
	return resp, nil
}

func GetTraceTx2(hash common.Hash) (*types.TraceTx2Resp, error) {
	t, err := store.ReadTraceTx2(context.Background(), nil, hash)
	if err != nil {

	}
	resp := &types.TraceTx2Resp{
		Res: t.Res,
	}
	return resp, nil
}

func GetTokenType(address common.Address) (interface{}, error) {
	erc20Count, err := store.ReadAccountErc20Total(context.Background(), nil, address)
	if err != nil {
		return nil, err
	}
	erc721Count, err := store.ReadAccountErc721Total(context.Background(), nil, address)
	if err != nil {
		return nil, err
	}
	erc1155Count, err := store.ReadAccountErc20Total(context.Background(), nil, address)
	if err != nil {
		return nil, err
	}
	return map[string]string{"erc20": erc20Count.String(), "erc721": erc721Count.String(), "erc1155": erc1155Count.String()}, nil
}

func ListTokenTransfers(typ string, pager *types.Pager) (map[string]interface{}, error) {
	var items interface{}
	var total string
	var err error
	switch typ {
	case "erc20":
		items, total, err = ListErc20Txs(pager)
		if err != nil {
			return nil, err
		}
	case "erc721":
		items, total, err = ListErc721Txs(pager)
		if err != nil {
			return nil, err
		}
	case "erc1155":
		items, total, err = ListErc1155Txs(pager)
		if err != nil {
			return nil, err
		}
	}
	resp := map[string]interface{}{"items": items, "total": total}
	return resp, nil
}
