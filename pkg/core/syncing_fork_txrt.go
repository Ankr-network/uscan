package core

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	forkTxTotal           *field.BigInt
	forkAccountTxTotalMap = utils.NewCache()
)

func (n *blockHandle) writeForkTxAndRt(ctx context.Context, tx *types.Tx, rt *types.Rt) (err error) {
	if forkTxTotal == nil {
		forkTxTotal, err = forkdb.ReadTxTotal(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				forkTxTotal = field.NewInt(0)
			} else {
				log.Errorf("get fork tx total: %v", err)
				return err
			}
		}
	}

	if err = forkdb.WriteTx(ctx, n.db, tx.Hash, tx); err != nil {
		log.Errorf("write fork tx(%s): %v", tx.Hash.Hex(), err)
		return err
	}
	deleteMap[share.ForkTxTbl] = append(deleteMap[share.ForkTxTbl], append([]byte("/fork/tx/"), tx.Hash.Bytes()...))

	if err = forkdb.WriteTxIndex(ctx, n.db, forkTxTotal.Add(field.NewInt(1)), tx.Hash); err != nil {
		log.Errorf("write fork tx(%s) index: %v", tx.Hash.Hex(), err)
		return err
	}
	deleteMap[share.ForkTxTbl] = append(deleteMap[share.ForkTxTbl], append([]byte("/fork/all/tx/"), forkTxTotal.Bytes()...))
	if indexMap["/fork/all/tx/index"] == nil {
		indexMap["/fork/all/tx/index"] = field.NewInt(0)
	}
	indexMap["/fork/all/tx/index"].Add(field.NewInt(1))

	if err = forkdb.WriteRt(ctx, n.db, tx.Hash, rt); err != nil {
		log.Errorf("write fork rt: %v", err)
		return err
	}
	deleteMap[share.ForkTxTbl] = append(deleteMap[share.ForkTxTbl], append([]byte("/fork/rt/"), tx.Hash.Bytes()...))

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
	if bytesRes, ok := forkAccountTxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = forkdb.ReadAccountTxTotal(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			}
		}
	}

	total.Add(field.NewInt(1))
	err = forkdb.WriteAccountTxIndex(ctx, n.db, addr, total, hash)
	if err != nil {
		log.Errorf("write fork account(%s) tx(%s) index: %v", addr.Hex(), hash.Hex(), err)
		return err
	}
	deleteMap[share.ForkAccountsTbl] = append(deleteMap[share.ForkAccountsTbl], append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/tx/"), total.Bytes()...)...))
	if indexMap["/fork/"+addr.String()+"/tx/index"] == nil {
		indexMap["/fork/"+addr.String()+"/tx/index"] = field.NewInt(0)
	}
	indexMap["/fork/"+addr.String()+"/tx/index"].Add(field.NewInt(1))

	err = forkdb.WriteAccountTxTotal(ctx, n.db, addr, total)
	if err == nil {
		forkAccountItxTotalMap.Add(addr, total.Bytes())
	}

	if totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/tx/total"] == nil {
		totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/tx/total"] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/tx/total"].Add(field.NewInt(1))

	return
}

func (n *blockHandle) writeForkTxTotal(ctx context.Context) error {
	if forkTxTotal != nil {
		oldTotal, err := forkdb.ReadTxTotal(ctx, n.db)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				oldTotal = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get fork tx total: %v", err)
				return err
			}
		}

		err = forkdb.WriteTxTotal(ctx, n.db, forkTxTotal)
		if err != nil {
			return err
		}

		newTotal := forkTxTotal
		newTotal.Sub(oldTotal)
		if totalMap[share.ForkTxTbl+":"+"/fork/all/tx/total"] == nil {
			totalMap[share.ForkTxTbl+":"+"/fork/all/tx/total"] = field.NewInt(0)
		}
		totalMap[share.ForkTxTbl+":"+"/fork/all/tx/total"].Add(newTotal)
	}
	return nil
}

func (n *blockHandle) deleteForkTxAndRt(ctx context.Context, tx *types.Tx, rt *types.Rt) (err error) {
	if err = forkdb.DeleteTx(ctx, n.db, tx.Hash); err != nil {
		log.Errorf("delete fork tx(%s): %v", tx.Hash.Hex(), err)
		return err
	}

	if err = forkdb.DeleteRt(ctx, n.db, tx.Hash); err != nil {
		log.Errorf("delete fork rt: %v", err)
		return err
	}

	return nil
}
