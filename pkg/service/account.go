package service

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
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
	c, err := store.ReadContract(context.Background(), nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	resp := &types.AccountResp{
		Owner:            account.Owner.String(),
		Balance:          account.Balance.String(),
		BlockNumber:      account.BlockNumber.StringPointer(),
		Name:             account.Name,
		Symbol:           account.Symbol,
		TokenTotalSupply: account.TokenTotalSupply.StringPointer(),
		NftTotalSupply:   account.NftTotalSupply.StringPointer(),
		Decimals:         account.Decimals.ToUint64(),
	}
	if account.Creator.Hex() != "" {
		creator := account.Creator.Hex()
		resp.Creator = &creator
	}
	if account.TxHash.Hex() != "" {
		txHash := account.TxHash.Hex()
		resp.TxHash = &txHash
	}

	if c.ByteCode != nil {
		code := hexutil.Encode(c.ByteCode)
		resp.Code = &code
	}
	return resp, nil
}

func GetAccountTxs(pager *types.Pager, address string) (map[string]interface{}, error) {
	num, err := store.ReadAccountTxTotal(context.Background(), nil, common.HexToAddress(address))
	if err != nil {

	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Tx, 0)
	for {
		tx, err := store.ReadAccountTxByIndex(context.Background(), nil, common.HexToAddress(address), p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}
	txsResp := make([]*types.ListTransactionResp, 0)
	addresses := make(map[string]common.Address)
	for _, tx := range txs {
		var blockNumber *string
		if tx.BlockNum.String() != "" {
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
		return nil, err
	}
	for _, t := range txsResp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
			//t.FromCode = hexutil.Encode(from.Code)
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			//t.FromCode = hexutil.Encode(to.Code)
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
		"total":      total,
		"otherTotal": otherTotal,
	}
	return resp, nil
}

func GetAccountItxs(pager *types.Pager, address string) ([]*types.InternalTxResp, string, error) {
	num, err := store.ReadITxTotal(context.Background(), nil, common.HexToHash(address))
	if err != nil {
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.InternalTx, 0)
	for {
		tx, err := store.ReadAccountITxByIndex(context.Background(), nil, common.HexToAddress(address), p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	resp := make([]*types.InternalTxResp, len(txs))
	for i, tx := range txs {
		var blockNumber string
		if tx.BlockNumber.String() != "" {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		var to string
		if tx.To.String() != "" {
			to = tx.To.String()
		}
		resp[i] = &types.InternalTxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Status:          tx.Status,
			CallType:        tx.CallType,
			Depth:           tx.Depth,
			From:            tx.From.String(),
			To:              &to,
			Amount:          tx.Amount.String(),
			GasLimit:        tx.GasLimit.String(),
			//CreatedTime:     tx.,// TODO
		}
	}
	return resp, total, nil
}

func GetAccountErc20Txns(pager *types.Pager, address string) ([]*types.Erc20TxResp, string, error) {
	num, err := store.ReadAccountErc20Total(context.Background(), nil, common.HexToAddress(address))
	if err != nil {
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc20Transfer, 0)
	for {
		tx, err := store.ReadAccountErc20ByIndex(context.Background(), nil, common.HexToAddress(address), p)
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
		if tx.BlockNumber.String() != "" {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc20TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			ContractName:    "",
			ContractSymbol:  "",
			//Method:          tx.Method,
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
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
			//t.FromCode = hexutil.Encode(from.Code)
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			//t.FromCode = hexutil.Encode(to.Code)
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, total, nil
}

func GetAccountErc721Txs(pager *types.Pager, address string) ([]*types.Erc721TxResp, string, error) {
	num, err := store.ReadAccountErc721Total(context.Background(), nil, common.HexToAddress(address))
	if err != nil {
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc721Transfer, 0)
	for {
		tx, err := store.ReadAccountErc721ByIndex(context.Background(), nil, common.HexToAddress(address), p)
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
		if tx.BlockNumber.String() != "" {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc721TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			//Method:          tx.Method,
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			TokenID:     tx.TokenId.ToUint64(),
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
			//t.FromCode = hexutil.Encode(from.Code)
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			//t.FromCode = hexutil.Encode(to.Code)
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, total, nil
}

func GetAccountErc1155Txs(pager *types.Pager, address string) ([]*types.Erc1155TxResp, string, error) {
	num, err := store.ReadAccountErc1155Total(context.Background(), nil, common.HexToAddress(address))
	if err != nil {
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc1155Transfer, 0)
	for {
		tx, err := store.ReadAccountErc1155ByIndex(context.Background(), nil, common.HexToAddress(address), p)
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
		if tx.BlockNumber.String() != "" {
			blockNumber = DecodeBig(tx.BlockNumber.String()).String()
		}
		t := &types.Erc1155TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     blockNumber,
			Contract:        tx.Contract.String(),
			//Method:          tx.Method,
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			TokenID:     tx.TokenID.ToUint64(),
			Value:       tx.Quantity.String(),
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
			//t.FromCode = hexutil.Encode(from.Code)
		}
		if to, ok := accounts[t.To]; ok {
			t.FromName = to.Name
			t.FromSymbol = to.Symbol
			//t.FromCode = hexutil.Encode(to.Code)
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, total, nil
}
