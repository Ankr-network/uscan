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

func (n *blockHandle) writeForkTxAndRt(ctx context.Context, tx *types.Tx, rt *types.Rt, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
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
	key := []byte("/fork/all/tx/index")
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	if err = forkdb.WriteRt(ctx, n.db, tx.Hash, rt); err != nil {
		log.Errorf("write fork rt: %v", err)
		return err
	}
	deleteMap[share.ForkTxTbl] = append(deleteMap[share.ForkTxTbl], append([]byte("/fork/rt/"), tx.Hash.Bytes()...))

	if tx.From != (common.Address{}) {
		if err = n.writeForkAccountTx(ctx, tx.From, tx.Hash, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(%s) tx: %v", tx.From, err)
			return err
		}
	}

	if tx.To != nil && tx.To.Hex() != (common.Address{}).Hex() {
		if err = n.writeForkAccountTx(ctx, *tx.To, tx.Hash, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(%s) tx: %v", tx.To.Hex(), err)
			return err
		}
	}

	if rt.ContractAddress != nil && rt.ContractAddress.Hex() != (common.Address{}).Hex() {
		if err = n.writeForkAccountTx(ctx, *rt.ContractAddress, tx.Hash, deleteMap, indexMap, totalMap); err != nil {
			log.Errorf("write fork account(%s) tx: %v", rt.ContractAddress.Hex(), err)
			return err
		}
	}

	return nil
}

func (n *blockHandle) writeForkAccountTx(ctx context.Context, addr common.Address, hash common.Hash, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
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
	key := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/index")...)
	if indexMap[string(key)] == nil {
		indexMap[string(key)] = field.NewInt(0)
	}
	indexMap[string(key)].Add(field.NewInt(1))

	err = forkdb.WriteAccountTxTotal(ctx, n.db, addr, total)
	if err == nil {
		forkAccountTxTotalMap.Add(addr, total.Bytes())
	}

	key2 := append(append([]byte("/fork/"), addr.Bytes()...), []byte("/tx/total")...)
	if totalMap[share.ForkAccountsTbl+":"+string(key2)] == nil {
		totalMap[share.ForkAccountsTbl+":"+string(key2)] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+string(key2)].Add(field.NewInt(1))

	return
}

func (n *blockHandle) writeForkTxTotal(ctx context.Context, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) error {
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

		forkTxTotal.Sub(oldTotal)
		key := []byte("/fork/all/tx/total")
		if totalMap[share.ForkTxTbl+":"+string(key)] == nil {
			totalMap[share.ForkTxTbl+":"+string(key)] = field.NewInt(0)
		}
		totalMap[share.ForkTxTbl+":"+string(key)].Add(forkTxTotal)
	}
	return nil
}
