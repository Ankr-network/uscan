package service

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ListTxs(pager *types.Pager) ([]*types.ListTransactionResp, string, error) {
	num, err := store.ReadTxTotal(context.Background(), nil)
	if err != nil {
		return nil, "0", err
	}
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Tx, 0)
	for {
		tx, err := store.ReadTxByIndex(context.Background(), nil, p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

	resp := make([]*types.ListTransactionResp, 0)
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
		resp = append(resp, t)
		addresses[tx.From.String()] = tx.From
		if tx.To != nil {
			addresses[tx.To.String()] = *tx.To
		}
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
	}
	return resp, num.String(), nil
}

var TokenTopics = []string{
	"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	"0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62",
	"0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb",
}

func GetTxBase(tx string) (*types.TransactionBaseResp, error) {
	txHash := common.HexToHash(tx)
	txData, err := store.ReadTx(context.Background(), nil, txHash)
	if err != nil {
		return nil, err
	}
	rtData, err := store.ReadRt(context.Background(), nil, txHash)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	resp := &types.TransactionBaseResp{
		Hash: txData.Hash.Hex(),
		//Nonce:   0,
		//GasUsed: txData.Gas,
		//GasLimit: 0,
		Status: 3,
	}
	if rtData != nil {
		resp.Status = rtData.Status.ToUint64()
		resp.GasUsed = rtData.GasUsed.String()
	}

	if txData.BlockNum != nil {
		block, _ := store.ReadBlock(context.Background(), nil, txData.BlockNum)
		resp.GasLimit = block.GasLimit.String()
		text, err := block.Nonce.MarshalText()
		if err != nil {
			return nil, err
		}
		resp.Nonce = hexutil.Encode(text)
	}
	return resp, nil
}

func GetTx(tx string) (*types.TxResp, error) {
	txHash := common.HexToHash(tx)
	txData, err := store.ReadTx(context.Background(), nil, txHash)
	if err != nil {
		return nil, err
	}
	rtData, err := store.ReadRt(context.Background(), nil, txHash)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	var blockNumber *string
	if txData.BlockNum != nil {
		num := DecodeBig(txData.BlockNum.String()).String()
		blockNumber = &num
	}
	method, err := txData.Method.MarshalText()
	if err != nil {
		return nil, err
	}
	resp := &types.TxResp{
		Hash:                 txData.Hash.String(),
		Method:               string(method),
		BlockHash:            txData.BlockNum.StringPointer(),
		BlockNumber:          blockNumber,
		From:                 txData.From.Hex(),
		To:                   txData.To.Hex(),
		Gas:                  txData.Gas.String(),
		GasPrice:             txData.GasPrice.String(),
		Value:                txData.Value.String(),
		MaxFeePerGas:         nil,
		MaxPriorityFeePerGas: nil,
		Input:                txData.Data.String(),
		Nonce:                txData.Nonce.String(),
		TransactionIndex:     nil,
		Type:                 nil,
		ChainID:              nil,
		V:                    txData.V.String(),
		R:                    txData.R.String(),
		S:                    txData.S.String(),
		TokensTransferred:    nil,
		MethodName:           "",
	}

	resp.Status = 3
	if rtData != nil {
		var rtCA *string
		if rtData.ContractAddress != nil {
			ca := rtData.ContractAddress.String()
			rtCA = &ca
		}
		bloomByte, err := rtData.Bloom.MarshalText()
		if err != nil {

		}
		bloom := string(bloomByte)
		rtResp := types.RtResp{
			ContractAddress:   rtCA,
			CumulativeGasUsed: rtData.CumulativeGasUsed.StringPointer(),
			EffectiveGasPrice: rtData.EffectiveGasPrice.StringPointer(),
			GasUsed:           rtData.GasUsed.String(),
			LogsBloom:         &bloom,
			Root:              rtData.PostState.String(),
			Status:            rtData.Status.ToUint64(),
			ErrorReturn:       rtData.ReturnErr,
		}
		resp.RtResp = rtResp
	}

	if txData.BlockNum != nil {
		block, _ := store.ReadBlock(context.Background(), nil, txData.BlockNum)
		resp.BaseFeePerGas = block.BaseFee.StringPointer()
		resp.GasLimit = block.GasLimit.String()
	}

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
		//resp.FromCode = hexutil.Encode(from.Code)
	}
	if to, ok := accounts[txData.To.Hex()]; ok {
		resp.FromName = to.Name
		resp.FromSymbol = to.Symbol
		//resp.FromCode = hexutil.Encode(to.Code)
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
				Address: log.Address.Hex(),
				Topics:  topics,
				Data:    log.Data.String(),
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
		// TODO
		//mid := strings.Split(*resp.Method, "0x")
		//if len(mid) == 2 {
		//	mname, err := store.GetMethod(mid[1])
		//	if err != nil && err != gorm.ErrRecordNotFound {
		//		return nil, err
		//	}
		//	resp.MethodName = mname.FuncName
		//}
	}

	return resp, nil
}

func GetAccounts(addresses map[string]common.Address) (map[string]*types.Account, error) {
	accounts := make(map[string]*types.Account, 0)
	for _, address := range addresses {
		account, err := store.ReadAccount(context.Background(), nil, address)
		if err != nil {
			return nil, err
		}
		accounts[account.Owner.String()] = account
	}
	return accounts, nil
}
