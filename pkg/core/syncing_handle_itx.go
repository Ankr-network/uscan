package core

import (
	"context"
	"errors"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	accountItxTotalMap = utils.NewCache()
)

func (n *blockHandle) writeITx(ctx context.Context, itxmap map[common.Hash][]*types.InternalTx) (err error) {
	var itxTotal *field.BigInt

	for k, itxs := range itxmap {
		itxTotal, err = rawdb.ReadITxTotal(ctx, n.db, k)
		if errors.Is(err, kv.NotFound) {
			itxTotal = field.NewInt(0)
		} else {
			log.Errorf("get itx total: %v", err)
			return err
		}

		for _, v := range itxs {
			v.TimeStamp = n.blockData.TimeStamp
			if err = rawdb.WriteITx(ctx, n.db, k, itxTotal.Add(field.NewInt(1)), v); err != nil {
				log.Errorf("write itx(%s): %v", k.Hex(), err)
				return err
			}
			key := &types.InternalTxKey{
				TransactionHash: v.TransactionHash,
				Index:           *itxTotal,
			}
			if v.From != (common.Address{}) {
				if err = n.writeAccountItx(ctx, v.From, key); err != nil {
					log.Errorf("write account(from: %s) Itx: %v", v.From.Hex(), err)
				}
			}

			if v.To != (common.Address{}) {
				if err = n.writeAccountItx(ctx, v.To, key); err != nil {
					log.Errorf("write account(to: %s) Itx: %v", v.To.Hex(), err)
				}
			}
		}
		if err = rawdb.WriteItxTotal(ctx, n.db, k, itxTotal); err != nil {
			log.Errorf("write itx total: %v", err)
			return err
		}
	}
	return nil
}

func (n *blockHandle) writeAccountItx(ctx context.Context, addr common.Address, data *types.InternalTxKey) (err error) {
	var total = &field.BigInt{}
	if bytesRes, ok := accountItxTotalMap.Get(addr); ok {
		total.SetBytes(bytesRes.([]byte))
	} else {
		total, err = rawdb.ReadAccountITxTotal(ctx, n.db, addr)
		if err != nil {
			if errors.Is(err, kv.NotFound) {
				total = field.NewInt(0)
				err = nil
			} else {
				log.Errorf("get account itx total: %v", err)
				return err
			}
		}
	}
	total.Add(field.NewInt(1))
	err = rawdb.WriteAccountITxIndex(ctx, n.db, addr, total, data)
	if err != nil {
		log.Errorf("write account itx : %v", err)
		return err
	}

	err = rawdb.WriteAccountITxTotal(ctx, n.db, addr, total)
	if err == nil {
		accountItxTotalMap.Add(addr.Bytes(), total.Bytes())
	}
	return
}
