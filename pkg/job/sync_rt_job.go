package job

import (
	"context"
	"time"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

type SyncRtJob struct {
	Completed   bool
	tx          common.Hash
	client      rpcclient.RpcClient
	ReceiptData *types.Rt
}

func NewSyncRtJob(tx common.Hash,
	client rpcclient.RpcClient,
) *SyncRtJob {
	return &SyncRtJob{
		tx:     tx,
		client: client,
	}
}

func (e *SyncRtJob) Execute() {
	var err error
	for {
		e.ReceiptData, err = e.client.GetTransactionReceiptByHash(context.Background(), e.tx)
		if err != nil {
			log.Errorf("get transaction(%s) data failed: %v", e.tx.Hex(), err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	if e.ReceiptData.ContractAddress == nil {
		e.ReceiptData.ContractAddress = &common.Address{}
	}
	e.Completed = true
}
