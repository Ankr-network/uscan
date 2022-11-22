package core

import (
	"context"
	"errors"
	"time"

	"github.com/Ankr-network/uscan/pkg/contract"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/job"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/workpool"
)

type Sync struct {
	client         rpcclient.RpcClient
	contractClient contract.Contractor
	db             kv.Database
	jobChan        workpool.Dispathcher
	storeChan      chan *job.SyncJob
}

func NewSync(
	client rpcclient.RpcClient,
	contractClient contract.Contractor,
	db kv.Database,
	chanSize uint64,
	forkBlockNum uint64,
) *Sync {
	s := &Sync{
		client:         client,
		contractClient: contractClient,
		db:             db,
		jobChan:        workpool.NewDispathcher(int(chanSize)),
		storeChan:      make(chan *job.SyncJob, chanSize*2),
	}
	go s.storeEvent(forkBlockNum)
	job.GlobalInit(int(chanSize))
	return s
}

func (n *Sync) Execute(ctx context.Context) {
	var (
		begin, lastBlock, end uint64
	)

	begin = n.getBeginBlock()
	for lastBlock = range n.client.GetLatestBlockNumber(ctx) {
		log.Infof("receive block: %d", lastBlock)
		end = lastBlock
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
	syncingBlock, err := rawdb.ReadSyncingBlock(context.Background(), n.db)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			syncingBlock = field.NewInt(0)
		} else {
			log.Fatalf("get syncing block err: %v", err)
		}
	}
	return syncingBlock.ToUint64() + 1
}

func (n *Sync) storeEvent(forkBlockNum uint64) {
	var blockNum uint64
	for job := range n.storeChan {
		for {
			blockNum = job.Block
			if job.Completed {
				if err := newBlockHandle(
					job.BlockData,
					job.TransactionDatas,
					job.ReceiptDatas,
					job.ContractOrMemberData,
					job.ContractInfoMap,
					job.ProxyContracts,
					job.InternalTxs,
					job.CallFrames,
					n.contractClient,
					n.db,
				).handle(); err != nil {
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
	log.Fatalf("handle failed: %d", blockNum)
}

func (n *Sync) toGetDebugLog(txes []*types.Tx) {
	for _, tx := range txes {
		if len(tx.Data) > 0 {
			job.DebugJobChan.AddJob(job.NewSyncDebugJob(tx.Hash, n.client, n.db))
		}
	}
}
