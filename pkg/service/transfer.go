package service

import (
	"context"
	"fmt"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/pkg/response"
	store "github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ListErc20Txs(pager *types.Pager) ([]*types.Erc20TxResp, string, error) {
	resp := make([]*types.Erc20TxResp, 0)
	num, err := store.ReadErc20Total(context.Background(), mdbx.DB)
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc20Transfer, 0)
	for {
		tx, err := store.ReadErc20Transfer(context.Background(), mdbx.DB, p)
		if err != nil {
			return nil, "0", err
		}
		txs = append(txs, tx)
		if p.String() == end.String() {
			break
		}
		p.Add(field.NewInt(-1))
	}

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
		return nil, "0", err
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
	return resp, DecodeBig(total).String(), nil
}

func ListErc721Txs(pager *types.Pager) ([]*types.Erc721TxResp, string, error) {
	resp := make([]*types.Erc721TxResp, 0)
	num, err := store.ReadErc721Total(context.Background(), mdbx.DB)
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc721Transfer, 0)
	for {
		tx, err := store.ReadErc721Transfer(context.Background(), mdbx.DB, p)
		if err != nil {
			return nil, "0", err
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
		return nil, "0", err
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
	return resp, DecodeBig(total).String(), nil
}

func ListErc1155Txs(pager *types.Pager) ([]*types.Erc1155TxResp, string, error) {
	resp := make([]*types.Erc1155TxResp, 0)
	num, err := store.ReadErc1155Total(context.Background(), mdbx.DB)
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}
	total := num.String()
	begin, end := ParsePage(num, pager.Offset, pager.Limit)
	p := begin
	txs := make([]*types.Erc1155Transfer, 0)
	for {
		tx, err := store.ReadErc1155Transfer(context.Background(), mdbx.DB, p)
		if err != nil {
			return nil, "0", err
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
		return nil, "0", err
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
	return resp, DecodeBig(total).String(), nil
}

func GetTraceTx(hash common.Hash) (*types.TraceTxResp, error) {
	t, err := store.ReadTraceTx(context.Background(), mdbx.DB, hash)
	if err != nil {
		if err == kv.NotFound {
			return nil, response.ErrRecordNotFind
		}
		return nil, err
	}
	resp := &types.TraceTxResp{
		Res:    t.Res,
		LogNum: t.LogNum.String(),
	}
	return resp, nil
}

func GetTraceTx2(hash common.Hash) (*types.TraceTx2Resp, error) {
	t, err := store.ReadTraceTx2(context.Background(), mdbx.DB, hash)
	if err != nil {
		if err == kv.NotFound {
			return nil, response.ErrRecordNotFind
		}
		return nil, err
	}
	resp := &types.TraceTx2Resp{
		Res: t.Res,
	}
	return resp, nil
}

func GetTokenType(address common.Address) (interface{}, error) {
	erc20Count, err := store.ReadAccountErc20Total(context.Background(), mdbx.DB, address)
	if err != nil {
		return nil, err
	}
	erc721Count, err := store.ReadAccountErc721Total(context.Background(), mdbx.DB, address)
	if err != nil {
		return nil, err
	}
	erc1155Count, err := store.ReadAccountErc20Total(context.Background(), mdbx.DB, address)
	if err != nil {
		return nil, err
	}
	return map[string]string{"erc20": erc20Count.String(), "erc721": erc721Count.String(), "erc1155": erc1155Count.String()}, nil
}

func ListTokenTransfers(address common.Address, typ string, pager *types.Pager) (map[string]interface{}, error) {
	var items interface{}
	var total string
	var err error
	switch typ {
	case "erc20":
		items, total, err = GetAccountErc20Txns(pager, address)
		if err != nil {
			return nil, err
		}
	case "erc721":
		items, total, err = GetAccountErc721Txs(pager, address)
		if err != nil {
			return nil, err
		}
	case "erc1155":
		items, total, err = GetAccountErc1155Txs(pager, address)
		if err != nil {
			return nil, err
		}
	}
	resp := map[string]interface{}{"items": items, "total": total}
	return resp, nil
}
func ListErc20Holders(pager *types.Pager, address common.Address) ([]*types.HolderResp, string, error) {
	resp := make([]*types.HolderResp, 0)
	holders, err := store.GetErc20Holder(context.Background(), mdbx.DB, address, uint64(pager.Offset), uint64(pager.Limit))
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}

	for _, holder := range holders {
		resp = append(resp, &types.HolderResp{
			Address:  holder.Addr.String(),
			Quantity: holder.Quantity.String(),
		})
	}

	if len(holders) > 0 {
		count, err := store.GetErc20HolderCount(context.Background(), mdbx.DB, address)
		if err != nil {
			return nil, "0", err
		}
		return resp, fmt.Sprint(count), nil
	}
	return resp, "0", nil
}

func ListErc721Holders(pager *types.Pager, address common.Address) ([]*types.HolderResp, string, error) {
	resp := make([]*types.HolderResp, 0)
	holders, err := store.GetErc721Holder(context.Background(), mdbx.DB, address, uint64(pager.Offset), uint64(pager.Limit))
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}

	for _, holder := range holders {
		resp = append(resp, &types.HolderResp{
			Address:  holder.Addr.String(),
			Quantity: holder.Quantity.String(),
		})
	}
	if len(holders) > 0 {
		count, err := store.GetErc721HolderCount(context.Background(), mdbx.DB, address)
		if err != nil {
			return nil, "0", err
		}
		return resp, fmt.Sprint(count), nil
	}
	return resp, "0", nil
}

func ListErc1155Holders(pager *types.Pager, address common.Address) ([]*types.HolderResp, string, error) {
	resp := make([]*types.HolderResp, 0)
	holders, err := store.GetErc1155Holder(context.Background(), mdbx.DB, address, uint64(pager.Offset), uint64(pager.Limit))
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}

	for _, holder := range holders {
		resp = append(resp, &types.HolderResp{
			Address:  holder.Addr.String(),
			Quantity: holder.Quantity.String(),
		})
	}
	if len(holders) > 0 {
		count, err := store.GetErc1155HolderCount(context.Background(), mdbx.DB, address)
		if err != nil {
			return nil, "0", err
		}
		return resp, fmt.Sprint(count), nil
	}
	return resp, "0", nil
}

func ListTokenHolders(typ string, pager *types.Pager, address common.Address) (map[string]interface{}, error) {
	var items interface{}
	var total string
	var err error
	switch typ {
	case "erc20":
		items, total, err = ListErc20Holders(pager, address)
		if err != nil {
			return nil, err
		}
	case "erc721":
		items, total, err = ListErc721Holders(pager, address)
		if err != nil {
			return nil, err
		}
	case "erc1155":
		items, total, err = ListErc1155Holders(pager, address)
		if err != nil {
			return nil, err
		}
	}
	return map[string]interface{}{"items": items, "total": total}, nil
}

func ListInventory(typ string, pager *types.Pager, address common.Address) (map[string]interface{}, error) {
	var items interface{}
	var total string
	var err error
	switch typ {
	case "erc721":
		items, total, err = ListErc721Inventory(pager, address)
		if err != nil {
			return nil, err
		}
	case "erc1155":
		items, total, err = ListErc1155Inventory(pager, address)
		if err != nil {
			return nil, err
		}
	}
	resp := map[string]interface{}{"items": items, "total": total}
	return resp, nil
}

func ListErc721Inventory(pager *types.Pager, address common.Address) ([]*types.InventoryResp, string, error) {
	resp := make([]*types.InventoryResp, 0)
	holders, err := store.GetErc721Inventory(context.Background(), mdbx.DB, address, uint64(pager.Offset), uint64(pager.Limit))
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}

	for _, holder := range holders {
		resp = append(resp, &types.InventoryResp{
			Address: holder.Addr.String(),
			TokenID: holder.TokenID.ToUint64(),
		})
	}
	if len(holders) > 0 {
		count, err := store.GetErc721InventoryCount(context.Background(), mdbx.DB, address)
		if err != nil {
			return nil, "0", err
		}
		return resp, fmt.Sprint(count), nil
	}
	return resp, "0", nil
}

func ListErc1155Inventory(pager *types.Pager, address common.Address) ([]uint64, string, error) {
	resp := make([]uint64, 0)
	tokenIDs, err := store.GetErc1155Inventory(context.Background(), mdbx.DB, address, uint64(pager.Offset), uint64(pager.Limit))
	if err != nil {
		if err == kv.NotFound {
			return resp, "0", nil
		}
		return nil, "0", err
	}

	for _, tokenID := range tokenIDs {
		resp = append(resp, tokenID.ToUint64())
	}
	if len(tokenIDs) > 0 {
		count, err := store.GetErc1155InventoryCount(context.Background(), mdbx.DB, address)
		if err != nil {
			return nil, "0", err
		}
		return resp, fmt.Sprint(count), nil
	}
	return resp, "0", nil
}
