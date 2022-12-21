package service

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func ListTxs(pager *types.Pager) ([]*types.ListTransactionResp, uint64, error) {
	total, err := store.GetTxTotal()
	if err != nil {
		return nil, 0, err
	}

	txs, err := store.ListTxs(total, pager.Offset, pager.Limit)
	if err != nil {
		return nil, 0, err
	}
	rts := make(map[string]*types.Rt, 0)
	for _, tx := range txs {
		rt, err := store.GetRt(tx.Hash)
		if err != nil {
			return nil, 0, err
		}
		rts[tx.Hash.String()] = rt
	}

	resp := make([]*types.ListTransactionResp, 0)
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
		resp = append(resp, t)
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
		return nil, 0, err
	}
	methodNames, err := GetMethodNames(methodIDs)
	if err != nil {
		return nil, 0, err
	}
	for _, t := range resp {
		if t.Method == "0x60806040" {
			t.To = rts[t.Hash].ContractAddress.Hex()
		}
		if to, ok := accounts[t.To]; ok {
			t.ToName = to.Name
			t.ToSymbol = to.Symbol
			if to.Erc20 || to.Erc721 || to.Erc1155 {
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
	return resp, total.ToUint64(), nil
}

var TokenTopics = []string{
	"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	"0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62",
	"0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb",
}

func GetTxBase(tx string) (*types.TransactionBaseResp, error) {
	txHash := common.HexToHash(tx)
	txData, err := store.GetTx(txHash)
	if err != nil {
		return nil, err
	}
	rtData, err := store.GetRt(txHash)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	resp := &types.TransactionBaseResp{
		Hash:  txData.Hash.Hex(),
		Nonce: txData.Nonce.String(),
		//GasUsed: txData.Gas,
		GasLimit: txData.Gas.String(),
		Status:   3,
	}
	if rtData != nil {
		resp.Status = rtData.Status.ToUint64()
		resp.GasUsed = rtData.GasUsed.String()
	}

	//x gas => gas limit (前)
	//rt gasUsed => (后)
	return resp, nil
}

func GetTx(tx string) (*types.TxResp, error) {
	txHash := common.HexToHash(tx)
	txData, err := store.GetTx(txHash)
	if err != nil {
		if err != nil {
			return nil, response.ErrRecordNotFind
		}
		return nil, err
	}
	rtData, err := store.GetRt(txHash)
	if err != nil && err != kv.NotFound {
		return nil, err
	}

	resp := &types.TxResp{
		Hash:                 txData.Hash.String(),
		Method:               txData.Method.String(),
		BlockHash:            txData.BlockNum.String(),
		BlockNumber:          DecodeBig(txData.BlockNum.String()).String(),
		From:                 txData.From.Hex(),
		To:                   txData.To.Hex(),
		Gas:                  txData.Gas.String(),
		GasPrice:             txData.GasPrice.String(),
		Value:                txData.Value.String(),
		MaxFeePerGas:         txData.GasFeeCap.StringPointer(),
		MaxPriorityFeePerGas: txData.GasTipCap.StringPointer(),
		Input:                txData.Data.String(),
		Nonce:                txData.Nonce.String(),
		V:                    txData.V.String(),
		R:                    txData.R.String(),
		S:                    txData.S.String(),
		CreatedTime:          txData.TimeStamp.ToUint64(),
	}

	resp.Status = 3
	if rtData != nil {
		rtResp := types.RtResp{
			CumulativeGasUsed: rtData.CumulativeGasUsed.StringPointer(),
			EffectiveGasPrice: rtData.EffectiveGasPrice.StringPointer(),
			GasUsed:           rtData.GasUsed.String(),
			Root:              rtData.PostState.String(),
			Status:            rtData.Status.ToUint64(),
			ErrorReturn:       rtData.ReturnErr,
		}
		if rtData.ContractAddress != nil {
			ca := rtData.ContractAddress.String()
			rtResp.ContractAddress = &ca
		}
		bloomByte, err := rtData.Bloom.MarshalText()
		if err != nil {
			return nil, err
		}
		bloom := string(bloomByte)
		rtResp.LogsBloom = &bloom

		resp.RtResp = rtResp
	}

	block, err := store.GetBlock(&txData.BlockNum)
	if err != nil {
		return nil, err
	}
	resp.BaseFeePerGas = block.BaseFee.StringPointer()
	resp.GasLimit = block.GasLimit.String()

	addresses := make(map[string]common.Address)
	addresses[txData.From.String()] = txData.From
	if txData.To != nil {
		addresses[txData.To.String()] = *txData.To
	}
	if rtData != nil && rtData.ContractAddress != nil {
		addresses[rtData.ContractAddress.String()] = *rtData.ContractAddress
	}

	accounts, err := GetAccounts(addresses)
	if err != nil {
		return nil, err
	}
	if from, ok := accounts[txData.From.Hex()]; ok {
		resp.FromName = from.Name
		resp.FromSymbol = from.Symbol
		if from.Erc20 || from.Erc721 || from.Erc1155 {
			resp.FromContract = true
		}
	}
	if to, ok := accounts[txData.To.Hex()]; ok {
		resp.ToName = to.Name
		resp.ToSymbol = to.Symbol
		if to.Erc20 || to.Erc721 || to.Erc1155 {
			resp.ToContract = true
		}
	}
	if rtData != nil && rtData.ContractAddress != nil {
		if ca, ok := accounts[rtData.ContractAddress.Hex()]; ok {
			resp.ContractAddressName = ca.Name
			resp.ContractAddressSymbol = ca.Symbol
		}
	}
	// event log
	resp.TotalLogs = 0
	resp.Logs = make([]*types.RtLogResp, 0)
	resp.TokensTransferred = make([]*types.TokensTransferred, 0)
	if rtData != nil {
		resp.TotalLogs = len(rtData.Logs)
		resp.Logs = make([]*types.RtLogResp, resp.TotalLogs)

		addresses = make(map[string]common.Address)
		for i, log := range rtData.Logs {
			topics := make([]string, 0)
			for _, topic := range log.Topics {
				topics = append(topics, topic.Hex())
			}
			resp.Logs[i] = &types.RtLogResp{
				Address:  log.Address.Hex(),
				Topics:   topics,
				Data:     log.Data.String(),
				LogIndex: log.LogIndex.ToUint64(),
			}
			addresses[log.Address.Hex()] = log.Address
			for _, topic := range TokenTopics {
				if len(log.Topics) > 0 {
					if log.Topics[0].Hex() == topic {
						var from, fromHex, to, toHex string
						if len(log.Topics) > 1 {
							from = common.HexToAddress(log.Topics[1].Hex()).String()
							fromHex = log.Topics[1].Hex()
						}
						if len(log.Topics) > 2 {
							to = common.HexToAddress(log.Topics[2].Hex()).String()
							toHex = log.Topics[2].Hex()
						}
						resp.TokensTransferred = append(resp.TokensTransferred, &types.TokensTransferred{
							From:    from,
							FromHex: fromHex,
							To:      to,
							ToHex:   toHex,
							Address: log.Address.Hex(),
						})
					}
				}
			}
		}
		accounts, err := GetAccounts(addresses)
		if err != nil {
			return nil, err
		}
		for _, transferred := range resp.TokensTransferred {
			if _, ok := accounts[transferred.Address]; ok {
				transferred.AddressName = accounts[transferred.Address].Name
				transferred.AddressSymbol = accounts[transferred.Address].Symbol
			}
		}
	}

	if resp.Method != "" {
		mid := strings.Split(resp.Method, "0x")
		if len(mid) == 2 {
			methodName, err := store.GetMethodName(mid[1])
			if err != nil && err != kv.NotFound {
				return nil, err
			}
			resp.MethodName = methodName
		}
	}

	return resp, nil
}

func GetAccounts(addresses map[string]common.Address) (map[string]*types.Account, error) {
	accounts := make(map[string]*types.Account, 0)
	for _, address := range addresses {
		account, err := store.GetAccount(address)
		if err != nil {
			if err == kv.NotFound {
				continue
			}
			return nil, err
		}
		accounts[account.Owner.String()] = account
	}
	return accounts, nil
}

func GetMethodNames(methodIDs []string) (methodNames map[string]string, err error) {
	methodNames = make(map[string]string, 0)
	for _, methodID := range methodIDs {
		methodName, err := store.GetMethodName(methodID)
		if err != nil {
			if err == kv.NotFound {
				continue
			}
			return nil, err
		}
		methodNames["0x"+methodID] = methodName
	}
	return
}
