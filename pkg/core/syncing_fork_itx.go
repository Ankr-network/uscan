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
	forkAccountItxTotalMap = utils.NewCache()
)

func (n *blockHandle) writeForkITx(ctx context.Context, itxmap map[common.Hash][]*types.InternalTx, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var itxTotal *field.BigInt

	for k, itxs := range itxmap {
		itxTotal, err = forkdb.ReadITxTotal(ctx, n.db, k)
		if errors.Is(err, kv.NotFound) {
			itxTotal = field.NewInt(0)
		} else {
			log.Errorf("get fork itx total: %v", err)
			return err
		}

		for _, v := range itxs {
			v.TimeStamp = n.blockData.TimeStamp
			if err = forkdb.WriteITx(ctx, n.db, k, itxTotal.Add(field.NewInt(1)), v); err != nil {
				log.Errorf("write fork itx(%s): %v", k.Hex(), err)
				return err
			}
			deleteMap[share.ForkTxTbl] = append(deleteMap[share.ForkTxTbl], append(append([]byte("/fork/iTx/"), k.Bytes()...), append([]byte("/"), itxTotal.Add(field.NewInt(1)).Bytes()...)...))
			if indexMap["/fork/iTx/"+k.String()+"/index"] == nil {
				indexMap["/fork/iTx/"+k.String()+"/index"] = field.NewInt(0)
			}
			indexMap["/fork/iTx/"+k.String()+"/index"].Add(field.NewInt(1))

			key := &types.InternalTxKey{
				TransactionHash: v.TransactionHash,
				Index:           *itxTotal,
			}
			if v.From != (common.Address{}) {
				if err = n.writeForkAccountItx(ctx, v.From, key, deleteMap, indexMap, totalMap); err != nil {
					log.Errorf("write fork account(from: %s) Itx: %v", v.From.Hex(), err)
				}
			}

			if v.To != (common.Address{}) {
				if err = n.writeForkAccountItx(ctx, v.To, key, deleteMap, indexMap, totalMap); err != nil {
					log.Errorf("write fork account(to: %s) Itx: %v", v.To.Hex(), err)
				}
			}
		}

		if err = forkdb.WriteItxTotal(ctx, n.db, k, itxTotal); err != nil {
			log.Errorf("write fork itx total: %v", err)
			return err
		}

		oldTotal, err := forkdb.ReadITxTotal(ctx, n.db, k)
		if errors.Is(err, kv.NotFound) {
			oldTotal = field.NewInt(0)
		} else {
			log.Errorf("get fork itx total: %v", err)
			return err
		}

		itxTotal.Sub(oldTotal)
		if totalMap[share.ForkTxTbl+":"+"/fork/iTx/"+k.String()+"/total"] == nil {
			totalMap[share.ForkTxTbl+":"+"/fork/iTx/"+k.String()+"/total"] = field.NewInt(0)
		}
		totalMap[share.ForkTxTbl+":"+"/fork/iTx/"+k.String()+"/total"].Add(itxTotal)
	}
	return nil
}

func (n *blockHandle) writeForkAccountItx(ctx context.Context, addr common.Address, data *types.InternalTxKey, deleteMap map[string][][]byte, indexMap map[string]*field.BigInt, totalMap map[string]*field.BigInt) (err error) {
	var total = &field.BigInt{}
	if bytesRes, ok := forkAccountItxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = forkdb.ReadAccountITxTotal(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get fork account itx total: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	err = forkdb.WriteAccountITxIndex(ctx, n.db, addr, total, data)
	if err != nil {
		log.Errorf("write fork account itx : %v", err)
		return err
	}
	deleteMap[share.ForkAccountsTbl] = append(deleteMap[share.ForkAccountsTbl], append(append([]byte("/fork/"), addr.Bytes()...), append([]byte("/itx/"), total.Bytes()...)...))
	if indexMap["/fork/"+addr.String()+"/itx/index"] == nil {
		indexMap["/fork/"+addr.String()+"/itx/index"] = field.NewInt(0)
	}
	indexMap["/fork/"+addr.String()+"/itx/index"].Add(field.NewInt(1))

	err = forkdb.WriteAccountITxTotal(ctx, n.db, addr, total)
	if err == nil {
		forkAccountItxTotalMap.Add(addr, total.Bytes())
	}
	if totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/itx/total"] == nil {
		totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/itx/total"] = field.NewInt(0)
	}
	totalMap[share.ForkAccountsTbl+":"+"/fork/"+addr.String()+"/itx/total"].Add(field.NewInt(1))

	return
}
