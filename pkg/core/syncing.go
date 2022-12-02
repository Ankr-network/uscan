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
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/pkg/workpool"
	"github.com/Ankr-network/uscan/share"
	"github.com/spf13/viper"
)

type Jobs struct {
	Main *job.SyncJob
	Fork *job.SyncJob
}

type Sync struct {
	client         rpcclient.RpcClient
	contractClient contract.Contractor
	db             kv.Database
	forkDb         kv.Database
	jobChan        workpool.Dispathcher
	storeChan      chan *Jobs
}

func NewSync(
	client rpcclient.RpcClient,
	contractClient contract.Contractor,
	db kv.Database,
	forkDB kv.Database,
	chanSize uint64,
) *Sync {
	s := &Sync{
		client:         client,
		contractClient: contractClient,
		db:             db,
		forkDb:         forkDB,
		jobChan:        workpool.NewDispathcher(int(chanSize)),
		storeChan:      make(chan *Jobs, chanSize*2),
	}
	job.GlobalInit(int(chanSize))
	go s.storeEvent()
	return s
}

func (n *Sync) Execute(ctx context.Context) {
	var (
		begin, lastBlock, end, forkBlockNumber, forkStart uint64
	)

	forkBlockNumber = viper.GetUint64(share.ForkBlockNum)
	begin = n.getBeginBlock()

	go func() {
		for latestBlockNumber := range n.client.GetLatestBlockNumber(ctx) {
			lastBlock = latestBlockNumber
			log.Infof("receive block: %d", lastBlock)
		}
	}()

	for {
		if begin <= lastBlock {
			var mainJob, forkJob *job.SyncJob
			end = lastBlock
			if forkStart > 0 {
				forkJob = job.NewSyncJob(begin, n.client)
				if forkStart <= begin-forkBlockNumber {
					mainJob = job.NewSyncJob(forkStart, n.client)
					forkStart++
				}
			} else {
				if begin <= end-forkBlockNumber {
					mainJob = job.NewSyncJob(begin, n.client)
				} else {
					forkJob = job.NewSyncJob(begin, n.client)
					forkStart = begin

				}
			}
			if mainJob != nil {
				n.jobChan.AddJob(mainJob)
			}
			if forkJob != nil {
				n.jobChan.AddJob(forkJob)
			}
			n.storeChan <- &Jobs{
				Main: mainJob,
				Fork: forkJob,
			}
		} else {
			time.Sleep(time.Microsecond * 100)
		}
	}

}

func (n *Sync) getBeginBlock() uint64 {
	syncingBlock, err := fulldb.ReadSyncingBlock(context.Background(), n.db)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			syncingBlock = field.NewInt(0)
		} else {
			log.Fatalf("get syncing block err: %v", err)
		}
	}
	return syncingBlock.ToUint64() + 1
}

func (n *Sync) storeEvent() {
	var (
		blockNum uint64
		err      error
	)
	for j := range n.storeChan {
		for {
			if ((j.Fork == nil) || (j.Fork != nil && j.Fork.Completed)) &&
				((j.Main == nil) || (j.Main != nil && j.Main.Completed)) {
				blockNum, err = n.handleJobs(j)
				if err != nil {
					goto end
				}
				break
			} else {
				time.Sleep(time.Microsecond * 100)
			}
		}
	}
end:
	log.Fatalf("handle failed: %d", blockNum)
}

func (n *Sync) handleJobs(jobs *Jobs) (blockNum uint64, err error) {
	var (
		ctxMain, ctxFork context.Context
		errMain, errFork error
	)
	ctxMain, errMain = n.db.BeginTx(context.Background())
	if errMain != nil {
		return blockNum, errMain
	}
	ctxFork, errFork = n.forkDb.BeginTx(context.Background())
	if errFork != nil {
		return blockNum, errFork
	}

	defer func() {
		if errMain == nil && errFork == nil {
			n.db.Commit(ctxMain)
			n.forkDb.Commit(ctxFork)
			log.Infof("write block complete: %d", blockNum)
		} else {
			n.db.RollBack(ctxMain)
			n.forkDb.RollBack(ctxFork)
		}
	}()

	// handle contract data
	if jobs.Fork != nil {
		blockNum = jobs.Fork.Block
		if errFork = newBlockHandle(
			jobs.Fork.BlockData,
			jobs.Fork.TransactionDatas,
			jobs.Fork.ReceiptDatas,
			jobs.Fork.ContractOrMemberData,
			jobs.Fork.ContractInfoMap,
			jobs.Fork.ProxyContracts,
			jobs.Fork.InternalTxs,
			jobs.Fork.CallFrames,
			n.contractClient,
			n.db,
		).handleContractData(ctxFork); errFork != nil {
			log.Errorf("handle contract data from fork: %d", blockNum)
			return blockNum, errFork
		}
	} else if jobs.Main != nil {
		blockNum = jobs.Main.Block
		if errMain = newBlockHandle(
			jobs.Main.BlockData,
			jobs.Main.TransactionDatas,
			jobs.Main.ReceiptDatas,
			jobs.Main.ContractOrMemberData,
			jobs.Main.ContractInfoMap,
			jobs.Main.ProxyContracts,
			jobs.Main.InternalTxs,
			jobs.Main.CallFrames,
			n.contractClient,
			n.db,
		).handleContractData(ctxMain); errMain != nil {
			log.Errorf("handle contract data from main: %d", blockNum)
			return blockNum, errMain
		}
	}

	// handle main job
	if jobs.Main != nil {
		blockNum = jobs.Main.Block
		if errMain = newBlockHandle(
			jobs.Main.BlockData,
			jobs.Main.TransactionDatas,
			jobs.Main.ReceiptDatas,
			jobs.Main.ContractOrMemberData,
			jobs.Main.ContractInfoMap,
			jobs.Main.ProxyContracts,
			jobs.Main.InternalTxs,
			jobs.Main.CallFrames,
			n.contractClient,
			n.db,
		).handleMain(ctxMain); errMain != nil {
			log.Errorf("handle main event data: %d", blockNum)
			return blockNum, errMain
		}

		if errMain = newBlockHandle(
			jobs.Main.BlockData,
			jobs.Main.TransactionDatas,
			jobs.Main.ReceiptDatas,
			jobs.Main.ContractOrMemberData,
			jobs.Main.ContractInfoMap,
			jobs.Main.ProxyContracts,
			jobs.Main.InternalTxs,
			jobs.Main.CallFrames,
			n.contractClient,
			n.forkDb,
		).handleDeleteFork(ctxMain); errMain != nil {
			log.Errorf("handle delete fork event data: %d", blockNum)
			return blockNum, errMain
		}
	}

	// handle fork job
	if jobs.Fork != nil {
		blockNum = jobs.Fork.Block
		if errFork = newBlockHandle(
			jobs.Fork.BlockData,
			jobs.Fork.TransactionDatas,
			jobs.Fork.ReceiptDatas,
			jobs.Fork.ContractOrMemberData,
			jobs.Fork.ContractInfoMap,
			jobs.Fork.ProxyContracts,
			jobs.Fork.InternalTxs,
			jobs.Fork.CallFrames,
			n.contractClient,
			n.forkDb,
		).handleFork(ctxFork); errFork != nil {
			log.Errorf("handle fork event data: %d", blockNum)
			return blockNum, errFork
		}
	}

	return blockNum, nil
}

func (n *Sync) toGetDebugLog(txes []*types.Tx) {
	for _, tx := range txes {
		if len(tx.Data) > 0 {
			job.DebugJobChan.AddJob(job.NewSyncDebugJob(tx.Hash, n.client, n.db))
		}
	}
}
