package job

import (
	"context"
	"math/big"
	"time"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type CheckBlockDebugLog struct {
	Block    uint64
	client   rpcclient.RpcClient
	Complete bool
	db       kv.Database
}

func NewCheckBlockDebugLog(block uint64,
	client rpcclient.RpcClient,
	db kv.Database,
) *CheckBlockDebugLog {
	return &CheckBlockDebugLog{
		Block:  block,
		client: client,
		db:     db,
	}
}

func (e *CheckBlockDebugLog) Execute() {
	var (
		ctx, err  = e.db.BeginTx(context.Background())
		blockData *types.Block
		logRes    *types.ExecutionResult
	)
	if err != nil {
		log.Fatalf("db begin: %v", err)
	}
	log.Infof("rpc call start: %d", e.Block)

	defer func() {
		if err == nil {
			e.db.Commit(ctx)
		} else {
			log.Errorf("callback because of error: %v", err)
			e.db.RollBack(ctx)
		}
	}()
	// get block data
	for {
		blockData, err = e.client.GetBlockByNumber(ctx, (*hexutil.Big)(big.NewInt(int64(e.Block))).String())
		if err != nil {
			log.Errorf("get block(%d) data failed: %v", e.Block, err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	for _, v := range blockData.Transactions {
		for {
			logRes, err = e.client.GetTracerLog(ctx, v)
			if err != nil {
				log.Errorf("get trancerlogs data failed: %v; tx: %s", err, v.Hex())
				time.Sleep(time.Second)
			} else {
				break
			}
		}
		if len(logRes.StructLogs) == 0 {
			continue
		} else {
			logNum := len(logRes.StructLogs)
			if logNum > 1000 {
				logRes.StructLogs = logRes.StructLogs[:1000]
			}
			err = fulldb.WriteTraceTx(context.Background(), e.db, v, &types.TraceTx{
				Res:    logRes.JsonToString(),
				LogNum: *field.NewInt(int64(logNum)),
			})
			if err != nil {
				log.Errorf("write Trace tx : %v", err)
				return
			}
		}

	}
	e.Complete = true
}
