package core

import (
	"context"
	"errors"
	"time"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/types"
)

var forkHomeCache *types.Home

func (n *blockHandle) updateForkHome(ctx context.Context) (err error) {
	var home *types.Home
	if forkHomeCache == nil {
		home, err = forkdb.ReadHome(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				log.Infof("read fork home not found")
				home = &types.Home{
					TxTotal:      *field.NewInt(0),
					AddressTotal: *field.NewInt(0),
					Blocks:       make([]*types.BkSim, 0, 1),
					Txs:          make([]*types.TxSim, 0, 10),
					DateTxs:      make(map[string]*field.BigInt),
				}
				err = nil
				forkHomeCache = home
			} else {
				return err
			}
		} else {
			forkHomeCache = home
		}
	} else {
		home = forkHomeCache
	}

	home.BlockNumber.SetBytes(n.blockData.Number.Bytes())

	home.TxTotal.Add(field.NewInt(int64(len(n.transactionData))))
	home.AddressTotal.Add(n.newAddrTotal)
	home.Erc20Total.Add(n.newErc20Total)
	home.Erc721Total.Add(n.newErc721Total)
	home.Erc721Total.Add(n.newErc1155Total)
	home.Blocks = append(home.Blocks, &types.BkSim{
		Number:            *n.blockData.Number,
		Timestamp:         n.blockData.TimeStamp,
		Miner:             n.blockData.Coinbase,
		GasUsed:           n.blockData.GasUsed,
		TransactionsTotal: *field.NewInt(int64(len(n.transactionData))),
	})
	for _, v := range n.transactionData {
		home.Txs = append(home.Txs, &types.TxSim{
			Hash:      v.Hash,
			From:      v.From,
			To:        *v.To,
			GasPrice:  v.GasPrice,
			Gas:       v.Gas,
			Timestamp: v.TimeStamp,
		})
	}

	timeLayout := "20060102"
	timeStr := time.Unix(int64(n.blockData.TimeStamp.ToUint64()), 0).UTC().Format(timeLayout)
	if _, ok := home.DateTxs[timeStr]; ok {
		home.DateTxs[timeStr].Add(field.NewInt(int64(len(n.transactionData))))
	} else {
		home.DateTxs[timeStr] = field.NewInt(int64(len(n.transactionData)))
	}

	if len(home.Blocks) > 10 {
		home.Blocks = home.Blocks[(len(home.Blocks) - 10):]
	}

	if len(home.Txs) > 10 {
		home.Txs = home.Txs[(len(home.Txs) - 10):]
	}

	delete(home.DateTxs, time.Unix(int64(n.blockData.TimeStamp.ToUint64()-(3600*24*14)), 0).UTC().Format(timeLayout))

	if err = forkdb.WriteSyncingBlock(ctx, n.db, n.blockData.Number); err != nil {
		log.Errorf("write fork syncing block: %v", err)
		return err
	}

	HomeMap[n.blockData.Number] = &types.Home{
		BlockNumber:  home.BlockNumber,
		TxTotal:      home.TxTotal,
		AddressTotal: home.AddressTotal,
		Erc20Total:   home.Erc20Total,
		Erc721Total:  home.Erc721Total,
		Erc1155Total: home.Erc1155Total,
		Blocks:       home.Blocks,
		Txs:          home.Txs,
		DateTxs:      home.DateTxs,
		DateTxsByte:  home.DateTxsByte,
	}

	return forkdb.WriteHome(ctx, n.db, home)
}

func (n *blockHandle) deleteForkHome(ctx context.Context, forkBlockNumber *field.BigInt) (err error) {
	newHome := &types.Home{}

	home, err := forkdb.ReadHome(ctx, n.db)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			return nil
		}
		return err
	}
	if forkHome, ok := HomeMap[forkBlockNumber]; ok {
		blocks := make([]*types.BkSim, 0)
		txs := make([]*types.TxSim, 0)

		for _, v := range home.Blocks {
			flag := true
			for _, v1 := range forkHome.Blocks {
				if (&v.Number).Cmp(&v1.Number) == 0 {
					flag = false
					break
				}
			}
			if flag {
				blocks = append(blocks, v)
			}
		}

		for _, v := range home.Txs {
			flag := true
			for _, v1 := range forkHome.Txs {
				if v.Hash.Hex() == v1.Hash.Hex() {
					flag = false
					break
				}
			}
			if flag {
				txs = append(txs, v)
			}
		}

		tx := home.TxTotal.Sub(&forkHome.TxTotal)
		address := home.AddressTotal.Sub(&forkHome.AddressTotal)
		erc20 := home.Erc20Total.Sub(&forkHome.Erc20Total)
		erc721 := home.Erc721Total.Sub(&forkHome.Erc721Total)
		erc1155 := home.Erc1155Total.Sub(&forkHome.Erc1155Total)
		newHome = &types.Home{
			BlockNumber:  home.BlockNumber,
			TxTotal:      *tx,
			AddressTotal: *address,
			Erc20Total:   *erc20,
			Erc721Total:  *erc721,
			Erc1155Total: *erc1155,
			Blocks:       blocks,
			Txs:          txs,
			DateTxs:      home.DateTxs,
			DateTxsByte:  home.DateTxsByte,
		}
		delete(HomeMap, forkBlockNumber)
	}

	return forkdb.WriteHome(ctx, n.db, newHome)
}
