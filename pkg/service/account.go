package service

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func GetAccountInfo(address string) (*types.AccountResp, error) {
	account, err := store.ReadAccount(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil {
		if err == kv.NotFound {
			return nil, response.ErrRecordNotFind
		}
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
	if account.Creator != (common.Address{}) {
		creator := account.Creator.Hex()
		resp.Creator = &creator
	}
	if account.TxHash != (common.Hash{}) {
		txHash := account.TxHash.Hex()
		resp.TxHash = &txHash
	}
	c, err := store.ReadContract(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	if c != nil {
		if c.ByteCode != nil {
			code := hexutil.Encode(c.ByteCode)
			resp.Code = &code
		}
	}
	return resp, nil
}

func GetAccountTxs(pager *types.Pager, address string) (map[string]interface{}, error) {
	num, err := store.ReadAccountTxTotal(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	txsResp := make([]*types.ListTransactionResp, 0)
	otherTotal := map[string]int64{
		"internalTotal": 0,
		"erc20Total":    0,
		"erc721Total":   0,
		"erc1155Total":  0,
	}
	resp := map[string]interface{}{
		"items":      txsResp,
		"total":      0,
		"otherTotal": otherTotal,
	}
	if num == nil {
		return resp, nil
	}
	total := num.String()

	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Tx, 0)
	for {
		tx, err := store.ReadAccountTxByIndex(context.Background(), mdbx.DB, common.HexToAddress(address), p)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

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
			if from.Erc20 || from.Erc721 || from.Erc1155 {
				t.FromContract = true
			}
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
				t.ToContract = true
			}
		}
	}

	itxTotal, err := store.ReadAccountITxTotal(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc20Total, err := store.ReadAccountErc20Total(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc721Total, err := store.ReadAccountErc721Total(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc1155Total, err := store.ReadAccountErc1155Total(context.Background(), mdbx.DB, common.HexToAddress(address))
	if err != nil && err != kv.NotFound {
		return nil, err
	}

	if itxTotal != nil {
		otherTotal["internalTotal"] = (*big.Int)(itxTotal).Int64()
	}
	if erc20Total != nil {
		otherTotal["erc20Total"] = (*big.Int)(erc20Total).Int64()
	}
	if erc721Total != nil {
		otherTotal["erc721Total"] = (*big.Int)(erc721Total).Int64()
	}
	if erc1155Total != nil {
		otherTotal["erc1155Total"] = (*big.Int)(erc1155Total).Int64()
	}
	resp = map[string]interface{}{
		"items":      txsResp,
		"total":      DecodeBig(total).Int64(),
		"otherTotal": otherTotal,
	}
	return resp, nil
}

func GetAccountItxs(pager *types.Pager, address string) ([]*types.InternalTxResp, int64, error) {
	num, err := store.ReadITxTotal(context.Background(), mdbx.DB, common.HexToHash(address))
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.InternalTxResp, 0)
	if num == nil {
		return resp, 0, nil
	}
	total := num.String()

	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.InternalTx, 0)
	for {
		tx, err := store.ReadAccountITxByIndex(context.Background(), mdbx.DB, common.HexToAddress(address), p)
		if err != nil {
			return nil, 0, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	resp = make([]*types.InternalTxResp, len(txs))
	for i, tx := range txs {
		resp[i] = &types.InternalTxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Status:          tx.Status,
			CallType:        tx.CallType,
			Depth:           tx.Depth,
			From:            tx.From.String(),
			To:              tx.To.String(),
			Amount:          tx.Amount.String(),
			GasLimit:        tx.GasLimit.String(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
	}
	return resp, DecodeBig(total).Int64(), nil
}

func GetAccountErc20Txns(pager *types.Pager, address common.Address) ([]*types.Erc20TxResp, int64, error) {
	num, err := store.ReadAccountErc20Total(context.Background(), mdbx.DB, address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.Erc20TxResp, 0)
	if num == nil {
		return resp, 0, nil
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc20Transfer, 0)
	for {
		tx, err := store.ReadAccountErc20ByIndex(context.Background(), mdbx.DB, address, p)
		if err != nil {
			return nil, 0, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	addresses := make(map[string]common.Address)
	for _, tx := range txs {
		t := &types.Erc20TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Contract:        tx.Contract.String(),
			Method:          hexutil.Bytes(tx.Method).String(),
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			Value:           tx.Amount.String(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
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
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, DecodeBig(total).Int64(), nil
}

func GetAccountErc721Txs(pager *types.Pager, address common.Address) ([]*types.Erc721TxResp, int64, error) {
	num, err := store.ReadAccountErc721Total(context.Background(), mdbx.DB, address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.Erc721TxResp, 0)
	if num == nil {
		return resp, 0, nil
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc721Transfer, 0)
	for {
		tx, err := store.ReadAccountErc721ByIndex(context.Background(), mdbx.DB, address, p)
		if err != nil {
			return nil, 0, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	addresses := make(map[string]common.Address)
	for _, tx := range txs {
		t := &types.Erc721TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Contract:        tx.Contract.String(),
			Method:          hexutil.Bytes(tx.Method).String(),
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenId.ToUint64(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
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
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, DecodeBig(total).Int64(), nil
}

func GetAccountErc1155Txs(pager *types.Pager, address common.Address) ([]*types.Erc1155TxResp, int64, error) {
	num, err := store.ReadAccountErc1155Total(context.Background(), mdbx.DB, address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.Erc1155TxResp, 0)
	if num == nil {
		return resp, 0, nil
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc1155Transfer, 0)
	for {
		tx, err := store.ReadAccountErc1155ByIndex(context.Background(), mdbx.DB, address, p)
		if err != nil {
			return nil, 0, err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	addresses := make(map[string]common.Address)

	for _, tx := range txs {
		t := &types.Erc1155TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Contract:        tx.Contract.String(),
			Method:          hexutil.Bytes(tx.Method).String(),
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenID.ToUint64(),
			Value:           tx.Quantity.String(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
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
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
		}
	}

	return resp, DecodeBig(total).Int64(), nil
}
