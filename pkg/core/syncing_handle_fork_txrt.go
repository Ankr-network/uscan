package core

import (
	"context"
	"errors"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/forkcache"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

func (n *blockHandle) writeForkTxAndRt(ctx context.Context, tx *types.Tx, rt *types.Rt) (err error) {
	if txTotal == nil {
		txTotal, err = forkcache.ReadTxTotal(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				txTotal = field.NewInt(0)
			} else {
				log.Errorf("get fork tx total: %v", err)
				return err
			}
		}
	}

	if err = forkcache.WriteTx(ctx, n.db, tx.Hash, tx); err != nil {
		log.Errorf("write fork tx(%s): %v", tx.Hash.Hex(), err)
		return err
	}

	if err = forkcache.WriteTxIndex(ctx, n.db, txTotal.Add(field.NewInt(1)), tx.Hash); err != nil {
		log.Errorf("write fork tx(%s) index: %v", tx.Hash.Hex(), err)
		return err
	}

	if err = forkcache.WriteRt(ctx, n.db, tx.Hash, rt); err != nil {
		log.Errorf("write fork rt: %v", err)
		return err
	}

	if tx.From != (common.Address{}) {
		if err = n.writeForkAccountTx(ctx, tx.From, tx.Hash); err != nil {
			log.Errorf("write fork account(%s) tx: %v", tx.From, err)
			return err
		}
	}

	if tx.To != nil && tx.To.Hex() != (common.Address{}).Hex() {
		if err = n.writeForkAccountTx(ctx, *tx.To, tx.Hash); err != nil {
			log.Errorf("write fork account(%s) tx: %v", tx.To.Hex(), err)
			return err
		}
	}

	return nil
}

func (n *blockHandle) writeForkAccountTx(ctx context.Context, addr common.Address, hash common.Hash) (err error) {
	var total = &field.BigInt{}
	if bytesRes, ok := accountTxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = forkcache.ReadAccountTxTotal(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			}
		}
	}
	total.Add(field.NewInt(1))
	err = forkcache.WriteAccountTxIndex(ctx, n.db, addr, total, hash)
	if err != nil {
		log.Errorf("write fork account(%s) tx(%s) index: %v", addr.Hex(), hash.Hex(), err)
		return err
	}
	err = forkcache.WriteAccountTxTotal(ctx, n.db, addr, total)
	if err == nil {
		accountItxTotalMap.Add(addr, total.Bytes())
	}
	return
}

func (n *blockHandle) writeForkTxTotal(ctx context.Context) (err error) {
	if txTotal != nil {
		err = forkcache.WriteTxTotal(ctx, n.db, txTotal)
	}
	return
}
