package job

import (
	"context"
	"time"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

type SyncTxJob struct {
	Completed            bool
	tx                   common.Hash
	client               rpcclient.RpcClient
	TransactionData      *types.Tx
	ContractOrMemberData map[common.Address]*types.Account
}

func NewSyncTxJob(tx common.Hash,
	client rpcclient.RpcClient) *SyncTxJob {
	return &SyncTxJob{
		tx:                   tx,
		client:               client,
		ContractOrMemberData: make(map[common.Address]*types.Account),
	}
}

func (e *SyncTxJob) Execute() {
	var err error
	for {
		e.TransactionData, err = e.client.GetTransactionByHash(context.Background(), e.tx)
		if err != nil {
			log.Errorf("get transaction(%s) data failed: %v", e.tx.Hex(), err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	if len(e.TransactionData.Data) > 0 {
		e.TransactionData.Method = e.TransactionData.Data[:4]
	}

	if e.TransactionData.To == nil {
		e.TransactionData.To = &common.Address{}
	}

	e.ContractOrMemberData[e.TransactionData.From] = &types.Account{
		Owner:       e.TransactionData.From,
		BlockNumber: e.TransactionData.BlockNum,
	}

	e.Completed = true
}
