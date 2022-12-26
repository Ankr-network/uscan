package service

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

func GetAccountInfo(address common.Address) (*types.AccountResp, error) {
	account, err := store.GetAccount(address)
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
		Erc20:            account.Erc20,
		Erc721:           account.Erc721,
		Erc1155:          account.Erc1155,
	}
	if account.Creator != (common.Address{}) {
		creator := account.Creator.Hex()
		resp.Creator = &creator
	}
	if account.TxHash != (common.Hash{}) {
		txHash := account.TxHash.Hex()
		resp.TxHash = &txHash
	}
	c, err := store.GetContract(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	if c != nil {
		if c.DeployedCode != nil {
			code := hexutil.Encode(c.DeployedCode)
			resp.Code = &code
		}
	}
	return resp, nil
}

func GetAccountTxs(pager *types.Pager, address common.Address) (map[string]interface{}, error) {
	total, err := store.GetAccountTxTotal(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	txsResp := make([]*types.ListTransactionResp, 0)

	resp := map[string]interface{}{
		"items": txsResp,
		"total": 0,
	}
	if total == nil {
		return resp, nil
	}

	txs, err := store.ListAccountTxs(address, total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, err
	}
	rts := make(map[string]*types.Rt, 0)
	for _, tx := range txs {
		rt, err := store.GetRt(tx.Hash)
		if err != nil {
			return nil, err
		}
		rts[tx.Hash.String()] = rt
	}
	addresses := make(map[string]common.Address)
	methodIDs := make([]string, 0)
	for _, tx := range txs {
		t := &types.ListTransactionResp{
			Hash:        tx.Hash.Hex(),
			Method:      tx.Method.String(),
			BlockNumber: DecodeBig(tx.BlockNum.String()).String(),
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			Gas:         rts[tx.Hash.Hex()].GasUsed.StringPointer(),
			GasPrice:    tx.GasPrice.StringPointer(),
			Value:       tx.Value.StringPointer(),
			CreatedTime: tx.TimeStamp.ToUint64(),
		}

		txsResp = append(txsResp, t)
		addresses[tx.From.String()] = tx.From
		if tx.To != nil {
			addresses[tx.To.String()] = *tx.To
		}
		if tx.Method.String() == "0x60806040" {
			contractAddress := rts[tx.Hash.Hex()].ContractAddress
			addresses[rts[tx.Hash.Hex()].ContractAddress.String()] = *contractAddress
		}
		if tx.Method.String() != "0x" && tx.Method.String() != "0x60806040" {
			mid := strings.Split(tx.Method.String(), "0x")
			if len(mid) == 2 {
				methodIDs = append(methodIDs, mid[1])
			}
		}
	}
	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, err
	}
	contracts, err := GetAccountContracts(addresses)
	if err != nil {
		return nil, err
	}
	methodNames, err := GetMethodNames(methodIDs)
	if err != nil {
		return nil, err
	}
	for _, t := range txsResp {
		if t.Method == "0x60806040" {
			t.To = rts[t.Hash].ContractAddress.Hex()
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
		}
		if to, ok := contracts[t.To]; ok {
			if to.DeployedCode != nil {
				t.ToContract = true
			}
		}
		if t.Method == "0x" {
			t.Method = "Transfer"
		}
		if t.Method != "Transfer" && t.Method != "0x60806040" {
			if mn, ok := methodNames[t.Method]; ok {
				md := strings.Split(mn, "(")
				if len(md) >= 1 {
					t.Method = strings.Title(md[0])
				}
			}
		}
	}

	resp = map[string]interface{}{
		"items": txsResp,
		"total": total.ToUint64(),
	}
	return resp, nil
}

func GetAccountTotal(address common.Address) (map[string]uint64, error) {
	otherTotal := map[string]uint64{
		"internalTotal": 0,
		"erc20Total":    0,
		"erc721Total":   0,
		"erc1155Total":  0,
		"txTotal":       0,
	}
	txTotal, err := store.GetAccountTxTotal(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	itxTotal, err := store.GetAccountITxTotal(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc20Total, err := store.GetAccountErc20Total(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc721Total, err := store.GetAccountErc721Total(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	erc1155Total, err := store.GetAccountErc1155Total(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}

	if txTotal != nil {
		otherTotal["txTotal"] = txTotal.ToUint64()
	}

	if itxTotal != nil {
		otherTotal["internalTotal"] = itxTotal.ToUint64()
	}
	if erc20Total != nil {
		otherTotal["erc20Total"] = erc20Total.ToUint64()
	}
	if erc721Total != nil {
		otherTotal["erc721Total"] = erc721Total.ToUint64()
	}
	if erc1155Total != nil {
		otherTotal["erc1155Total"] = erc1155Total.ToUint64()
	}
	return otherTotal, nil
}

func GetAccountItxs(pager *types.Pager, address common.Address) ([]*types.InternalTxResp, uint64, error) {
	total, err := store.GetAccountITxTotal(address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.InternalTxResp, 0)
	if total == nil {
		return resp, 0, nil
	}

	txs, err := store.ListAccountITxs(address, total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
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
	return resp, total.ToUint64(), nil
}

func GetAccountErc20Txns(pager *types.Pager, address common.Address) ([]*types.Erc20TxResp, uint64, error) {
	total, err := store.GetAccountErc20Total(address)
	if err != nil && err != kv.NotFound {
		log.Infof("GetAccountErc20Txns, GetAccountErc20Total:%s", err)
		return nil, 0, err
	}
	resp := make([]*types.Erc20TxResp, 0)
	if total == nil {
		return resp, 0, nil
	}

	txs, err := store.ListAccountErc20Txs(address, total, pager.Offset, pager.Limit)
	if err != nil {
		log.Infof("GetAccountErc20Txns, ListAccountErc20Txs:%s", err)
		return nil, 0, err
	}
	addresses := make(map[string]common.Address)
	methodIDs := make([]string, 0)
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
		if t.Method != "0x" && t.Method != "0x60806040" {
			mid := strings.Split(t.Method, "0x")
			if len(mid) == 2 {
				methodIDs = append(methodIDs, mid[1])
			}
		}
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, 0, err
	}
	contracts, err := GetAccountContracts(addresses)
	if err != nil {
		return nil, 0, err
	}
	methodNames, err := GetMethodNames(methodIDs)
	if err != nil {
		return nil, 0, err
	}
	for _, t := range resp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
		}
		if from, ok := contracts[t.From]; ok {
			if from.DeployedCode != nil {
				t.FromContract = true
			}
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
		}
		if to, ok := contracts[t.To]; ok {
			if to.DeployedCode != nil {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
			t.ContractDecimals = c.Decimals.ToUint64()
		}
		if t.Method == "0x" {
			t.Method = "Transfer"
		}
		if t.Method != "Transfer" && t.Method != "0x60806040" {
			if mn, ok := methodNames[t.Method]; ok {
				md := strings.Split(mn, "(")
				if len(md) >= 1 {
					t.Method = strings.Title(md[0])
				}
			}
		}
	}

	return resp, total.ToUint64(), nil
}

func GetAccountErc721Txs(pager *types.Pager, address common.Address) ([]*types.Erc721TxResp, uint64, error) {
	total, err := store.GetAccountErc721Total(address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.Erc721TxResp, 0)
	if total == nil {
		return resp, 0, nil
	}

	txs, err := store.ListAccountErc721Txs(address, total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
	}
	addresses := make(map[string]common.Address)
	methodIDs := make([]string, 0)
	for _, tx := range txs {
		t := &types.Erc721TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Contract:        tx.Contract.String(),
			Method:          hexutil.Bytes(tx.Method).String(),
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenId.String(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
		if t.Method != "0x" && t.Method != "0x60806040" {
			mid := strings.Split(t.Method, "0x")
			if len(mid) == 2 {
				methodIDs = append(methodIDs, mid[1])
			}
		}
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, 0, err
	}
	contracts, err := GetAccountContracts(addresses)
	if err != nil {
		return nil, 0, err
	}
	methodNames, err := GetMethodNames(methodIDs)
	if err != nil {
		return nil, 0, err
	}
	for _, t := range resp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
		}
		if from, ok := contracts[t.From]; ok {
			if from.DeployedCode != nil {
				t.FromContract = true
			}
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
		}
		if to, ok := contracts[t.To]; ok {
			if to.DeployedCode != nil {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
			t.ContractDecimals = c.Decimals.ToUint64()
		}
		if t.Method == "0x" {
			t.Method = "Transfer"
		}
		if t.Method != "Transfer" && t.Method != "0x60806040" {
			if mn, ok := methodNames[t.Method]; ok {
				md := strings.Split(mn, "(")
				if len(md) >= 1 {
					t.Method = strings.Title(md[0])
				}
			}
		}
	}

	return resp, total.ToUint64(), nil
}

func GetAccountErc1155Txs(pager *types.Pager, address common.Address) ([]*types.Erc1155TxResp, uint64, error) {
	total, err := store.GetAccountErc1155Total(address)
	if err != nil && err != kv.NotFound {
		return nil, 0, err
	}
	resp := make([]*types.Erc1155TxResp, 0)
	if total == nil {
		return resp, 0, nil
	}
	txs, err := store.ListAccountErc1155Txs(address, total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
	}

	addresses := make(map[string]common.Address)
	methodIDs := make([]string, 0)
	for _, tx := range txs {
		t := &types.Erc1155TxResp{
			TransactionHash: tx.TransactionHash.String(),
			BlockHash:       tx.TransactionHash.String(),
			BlockNumber:     DecodeBig(tx.BlockNumber.String()).String(),
			Contract:        tx.Contract.String(),
			Method:          hexutil.Bytes(tx.Method).String(),
			From:            tx.From.Hex(),
			To:              tx.To.Hex(),
			TokenID:         tx.TokenID.String(),
			Value:           tx.Quantity.String(),
			CreatedTime:     tx.TimeStamp.ToUint64(),
		}
		resp = append(resp, t)

		addresses[tx.From.String()] = tx.From
		addresses[tx.To.String()] = tx.To
		addresses[tx.Contract.String()] = tx.Contract
		if t.Method != "0x" && t.Method != "0x60806040" {
			mid := strings.Split(t.Method, "0x")
			if len(mid) == 2 {
				methodIDs = append(methodIDs, mid[1])
			}
		}
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, 0, err
	}
	contracts, err := GetAccountContracts(addresses)
	if err != nil {
		return nil, 0, err
	}
	methodNames, err := GetMethodNames(methodIDs)
	if err != nil {
		return nil, 0, err
	}
	for _, t := range resp {
		if from, ok := accounts[t.From]; ok {
			t.FromName = from.Name
			t.FromSymbol = from.Symbol
		}
		if from, ok := contracts[t.From]; ok {
			if from.DeployedCode != nil {
				t.FromContract = true
			}
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
		}
		if to, ok := contracts[t.To]; ok {
			if to.DeployedCode != nil {
				t.ToContract = true
			}
		}
		if c, ok := accounts[t.Contract]; ok {
			t.ContractName = c.Name
			t.ContractSymbol = c.Symbol
			t.ContractDecimals = c.Decimals.ToUint64()
		}
		if t.Method == "0x" {
			t.Method = "Transfer"
		}
		if t.Method != "Transfer" && t.Method != "0x60806040" {
			if mn, ok := methodNames[t.Method]; ok {
				md := strings.Split(mn, "(")
				if len(md) >= 1 {
					t.Method = strings.Title(md[0])
				}
			}
		}
	}

	return resp, total.ToUint64(), nil
}
