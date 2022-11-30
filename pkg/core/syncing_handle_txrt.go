package core

import (
	"context"
	"errors"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	txTotal           *field.BigInt
	accountTxTotalMap = utils.NewCache()
)

func (n *blockHandle) writeTxAndRt(ctx context.Context, tx *types.Tx, rt *types.Rt) (err error) {
	if txTotal == nil {
		txTotal, err = fulldb.ReadTxTotal(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				txTotal = field.NewInt(0)
			} else {
				log.Errorf("get tx total: %v", err)
				return err
			}
		}
	}

	if err = fulldb.WriteTx(ctx, n.db, tx.Hash, tx); err != nil {
		log.Errorf("write tx(%s): %v", tx.Hash.Hex(), err)
		return err
	}

	if err = fulldb.WriteTxIndex(ctx, n.db, txTotal.Add(field.NewInt(1)), tx.Hash); err != nil {
		log.Errorf("write tx(%s) index: %v", tx.Hash.Hex(), err)
		return err
	}

	if err = fulldb.WriteRt(ctx, n.db, tx.Hash, rt); err != nil {
		log.Errorf("write rt: %v", err)
		return err
	}

	if tx.From != (common.Address{}) {
		if err = n.writeAccountTx(ctx, tx.From, tx.Hash); err != nil {
			log.Errorf("write account(%s) tx: %v", tx.From, err)
			return err
		}
	}

	if tx.To != nil && tx.To.Hex() != (common.Address{}).Hex() {
		if err = n.writeAccountTx(ctx, *tx.To, tx.Hash); err != nil {
			log.Errorf("write account(%s) tx: %v", tx.To.Hex(), err)
			return err
		}
	}

	return nil
}

func (n *blockHandle) writeAccountTx(ctx context.Context, addr common.Address, hash common.Hash) (err error) {
	var total = &field.BigInt{}
	if bytesRes, ok := accountTxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = fulldb.ReadAccountTxTotal(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			}
		}
	}
	total.Add(field.NewInt(1))
	err = fulldb.WriteAccountTxIndex(ctx, n.db, addr, total, hash)
	if err != nil {
		log.Errorf("write account(%s) tx(%s) index: %v", addr.Hex(), hash.Hex(), err)
		return err
	}
	err = fulldb.WriteAccountTxTotal(ctx, n.db, addr, total)
	if err == nil {
		accountItxTotalMap.Add(addr, total.Bytes())
	}
	return
}

func (n *blockHandle) writeTxTotal(ctx context.Context) (err error) {
	if txTotal != nil {
		err = fulldb.WriteTxTotal(ctx, n.db, txTotal)
	}
	return
}
