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

func (n *blockHandle) writeForkITx(ctx context.Context, itxmap map[common.Hash][]*types.InternalTx) (err error) {
	var itxTotal *field.BigInt

	for k, itxs := range itxmap {
		itxTotal, err = forkcache.ReadITxTotal(ctx, n.db, k)
		if errors.Is(err, kv.NotFound) {
			itxTotal = field.NewInt(0)
		} else {
			log.Errorf("get fork itx total: %v", err)
			return err
		}

		for _, v := range itxs {
			v.TimeStamp = n.blockData.TimeStamp
			if err = forkcache.WriteITx(ctx, n.db, k, itxTotal.Add(field.NewInt(1)), v); err != nil {
				log.Errorf("write fork itx(%s): %v", k.Hex(), err)
				return err
			}
			key := &types.InternalTxKey{
				TransactionHash: v.TransactionHash,
				Index:           *itxTotal,
			}
			if v.From != (common.Address{}) {
				if err = n.writeForkAccountItx(ctx, v.From, key); err != nil {
					log.Errorf("write fork account(from: %s) Itx: %v", v.From.Hex(), err)
				}
			}

			if v.To != (common.Address{}) {
				if err = n.writeForkAccountItx(ctx, v.To, key); err != nil {
					log.Errorf("write fork account(to: %s) Itx: %v", v.To.Hex(), err)
				}
			}
		}
		if err = forkcache.WriteItxTotal(ctx, n.db, k, itxTotal); err != nil {
			log.Errorf("write fork itx total: %v", err)
			return err
		}
	}
	return nil
}

func (n *blockHandle) writeForkAccountItx(ctx context.Context, addr common.Address, data *types.InternalTxKey) (err error) {
	var total = &field.BigInt{}
	if bytesRes, ok := accountItxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = forkcache.ReadAccountITxTotal(ctx, n.db, addr)
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
	err = forkcache.WriteAccountITxIndex(ctx, n.db, addr, total, data)
	if err != nil {
		log.Errorf("write fork account itx : %v", err)
		return err
	}

	err = forkcache.WriteAccountITxTotal(ctx, n.db, addr, total)
	if err == nil {
		accountItxTotalMap.Add(addr, total.Bytes())
	}
	return
}