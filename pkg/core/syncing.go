package core

import (
	"context"
	"time"

	"github.com/Ankr-network/uscan/pkg/contract"
	"github.com/Ankr-network/uscan/pkg/job"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/workpool"
	"github.com/ethereum/go-ethereum/common"
)

type Sync struct {
	client         rpcclient.RpcClient
	contractClient contract.Contractor
	jobChan        workpool.Dispathcher
	storeChan      chan *job.SyncJob
}

func NewSync(
	client rpcclient.RpcClient,
	contractClient contract.Contractor,
	chanSize uint64,
) *Sync {
	s := &Sync{
		client:         client,
		contractClient: contractClient,
		jobChan:        workpool.NewDispathcher(int(chanSize)),
		storeChan:      make(chan *job.SyncJob, chanSize*2),
	}
	go s.storeEvent()
	return s
}

func (n *Sync) Execute(ctx context.Context, defaultBlock uint64) {
	var (
		begin, lastBlock, end uint64
	)

	begin = n.getBeginBlock()
	for lastBlock = range n.client.GetLatestBlockNumber(ctx) {
		log.Infof("receive block: %d", lastBlock)
		end = lastBlock - 1
		if begin > end {
			continue
		}

		if end >= begin {
			log.Infof("from %d to %d", begin, end)
			for ; begin <= end; begin++ {
				serveJob := job.NewSyncJob(begin, n.client)
				n.jobChan.AddJob(serveJob)
				n.storeChan <- serveJob
			}
		}
	}

}

func (n *Sync) getBeginBlock() uint64 {
	return 0
}

func (n *Sync) storeEvent() {
	var blockNum uint64
	for job := range n.storeChan {
		for {
			blockNum = job.Block
			if job.Completed {
				if err := n.handleEventData(job.Block, job.BlockData, job.TransactionDatas, job.ReceiptDatas, job.ContractOrMemberData, job.InternalTxs, job.CallFrames); err != nil {
					log.Errorf("handle event data: %d", job.Block)
					goto end
				} else {
					// n.toGetDebugLog(job.TransactionDatas)
				}
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
end:
	log.Fatal("handle failed: %d", blockNum)
}

func (n *Sync) handleEventData(block uint64, blockData *types.Block, transactionData []*types.Tx, receiptData []*types.Rt, contractOrMemberData map[common.Address]*types.Account, internalTxs []*types.InternalTx, callFrames map[common.Hash]*types.CallFrame) (err error) {
	return nil
}

func (n *Sync) CountTransfer(ctx context.Context, block uint64, erc20_transfers, erc721_transfers, erc1155_transfers int) (err error) {
	return nil
}

func (n *Sync) toGetDebugLog(txes []*types.Tx) {
	// for _, tx := range txes {
	// 	if strings.HasPrefix(tx.Input, "0x") && len(tx.Input) > 2 {
	// 		job.DebugJobChan <- job.NewSyncDebugJob(field.Hash(common.HexToHash(tx.Hash)), n.client, n.dbRepo)
	// 	}
	// }
}
