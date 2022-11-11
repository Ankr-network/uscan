package job

import (
	"context"
	"time"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

type SyncDebugJob struct {
	txhash common.Hash
	client rpcclient.RpcClient
	retry  int
}

func NewSyncDebugJob(
	txhash common.Hash,
	client rpcclient.RpcClient,
) *SyncDebugJob {
	return &SyncDebugJob{
		txhash: txhash,
		client: client,
	}
}

func (e *SyncDebugJob) Execute() {
	var (
		ctx = context.Background()
		res *types.ExecutionResult
		err error
	)

	for {
		res, err = e.client.GetTracerLog(ctx, e.txhash)
		if err != nil {
			if e.retry >= share.Retry {
				return
			}
			e.retry++
			log.Errorf("get trancerlogs(%s) data failed: %v", e.txhash.Hex(), err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	if len(res.StructLogs) == 0 {
		return
	}

	logNum := len(res.StructLogs)
	if logNum > 1000 {
		res.StructLogs = res.StructLogs[:1000]
	}
}
