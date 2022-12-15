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
		if t.Method == "0x" {
			t.Method = ""
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
		//GasLimit: txData.Gas,
		Status: 3,
	}
	if rtData != nil {
		resp.Status = rtData.Status.ToUint64()
		resp.GasUsed = rtData.GasUsed.String()
	}

	block, _ := store.GetBlock(&txData.BlockNum)
	resp.GasLimit = block.GasLimit.String()
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
