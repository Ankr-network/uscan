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
)

type Jobs struct {
	Main *job.SyncJob
	Fork *job.SyncJob
}

type Sync struct {
	client         rpcclient.RpcClient
	contractClient contract.Contractor
	forkNum        int64
	db             kv.Database
	forkDb         kv.Database
	jobChan        workpool.Dispathcher
	storeChan      chan *Jobs
}

func NewSync(
	client rpcclient.RpcClient,
	contractClient contract.Contractor,
	forkNum int64,
	db kv.Database,
	forkDB kv.Database,
	chanSize uint64,
) *Sync {
	s := &Sync{
		client:         client,
		contractClient: contractClient,
		forkNum:        forkNum,
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
		begin, lastBlock, end, forkStart uint64
	)

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
				if int64(forkStart) <= int64(begin)-n.forkNum {
					mainJob = job.NewSyncJob(forkStart, n.client)
					forkStart++
				}
			} else {
				if begin <= end-uint64(n.forkNum) {
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
			begin++
		} else {
			time.Sleep(time.Millisecond * 100)
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
		err error
	)
	for j := range n.storeChan {
		for {
			if ((j.Fork == nil) || (j.Fork != nil && j.Fork.Completed)) &&
				((j.Main == nil) || (j.Main != nil && j.Main.Completed)) {
				err = n.handleJobs(j)
				if err != nil {
					goto end
				}
				break
			} else {
				time.Sleep(time.Millisecond * 100)
			}
		}
	}
end:
	log.Fatalf("handle failed")
}

func (n *Sync) handleJobs(jobs *Jobs) (err error) {
	var (
		ctxMain, ctxFork       context.Context
		errMain, errFork       error
		mainHandle, forkHandle *blockHandle
	)
	ctxMain, errMain = n.db.BeginTx(context.Background())
	if errMain != nil {
		return errMain
	}
	ctxFork, errFork = n.forkDb.BeginTx(context.Background())
	if errFork != nil {
		return errFork
	}

	defer func() {
		if errMain == nil && errFork == nil {
			n.db.Commit(ctxMain)
			n.forkDb.Commit(ctxFork)
		} else {
			n.db.RollBack(ctxMain)
			n.forkDb.RollBack(ctxFork)
		}
	}()

	if jobs.Fork != nil {
		log.Infof("handle fork block: %s", jobs.Fork.BlockData.Number.String())
		forkHandle = newBlockHandle(
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
		)
		if errFork = forkHandle.handleFork(ctxFork); errFork != nil {
			log.Errorf("handle fork event data: %d", jobs.Fork.BlockData.Number.String())
			return errFork
		}
		delForkBlockNumber := jobs.Fork.BlockData.Number.Sub(field.NewInt(n.forkNum))

		if errFork = forkHandle.handleDeleteFork(ctxFork, delForkBlockNumber); errFork != nil {
			log.Errorf("delete fork  data: %d", jobs.Fork.BlockData.Number.String())
			return errFork
		}

		// errFork = forkHandle.handleContractData(ctxFork)
		// if errFork != nil {
		// 	log.Errorf("handle contract data from fork: %d", blockNum)
		// 	return blockNum, errFork
		// }
	}

	if jobs.Main != nil {
		log.Infof("handle block: %s", jobs.Main.BlockData.Number.String())
		mainHandle = newBlockHandle(
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
		)
		if errMain = mainHandle.handleMain(ctxMain); errMain != nil {
			log.Errorf("handle main data: %d", jobs.Main.BlockData.Number.String())
			return errMain
		}
	}

	//  if jobs.Main != nil {
	// 	blockNum = jobs.Main.Block
	// 	if errMain = newBlockHandle(
	// 		jobs.Main.BlockData,
	// 		jobs.Main.TransactionDatas,
	// 		jobs.Main.ReceiptDatas,
	// 		jobs.Main.ContractOrMemberData,
	// 		jobs.Main.ContractInfoMap,
	// 		jobs.Main.ProxyContracts,
	// 		jobs.Main.InternalTxs,
	// 		jobs.Main.CallFrames,
	// 		n.contractClient,
	// 		n.db,
	// 	).handleContractData(ctxMain); errMain != nil {
	// 		log.Errorf("handle contract data from main: %d", blockNum)
	// 		return blockNum, errMain
	// 	}
	// }

	// handle main job
	if jobs.Main != nil {
		// blockNum = jobs.Main.Block
		// if errMain = newBlockHandle(
		// 	jobs.Main.BlockData,
		// 	jobs.Main.TransactionDatas,
		// 	jobs.Main.ReceiptDatas,
		// 	jobs.Main.ContractOrMemberData,
		// 	jobs.Main.ContractInfoMap,
		// 	jobs.Main.ProxyContracts,
		// 	jobs.Main.InternalTxs,
		// 	jobs.Main.CallFrames,
		// 	n.contractClient,
		// 	n.db,
		// ).handleMain(ctxMain); errMain != nil {
		// 	log.Errorf("handle main event data: %d", blockNum)
		// 	return blockNum, errMain
		// }

		// if errMain = newBlockHandle(
		// 	jobs.Main.BlockData,
		// 	jobs.Main.TransactionDatas,
		// 	jobs.Main.ReceiptDatas,
		// 	jobs.Main.ContractOrMemberData,
		// 	jobs.Main.ContractInfoMap,
		// 	jobs.Main.ProxyContracts,
		// 	jobs.Main.InternalTxs,
		// 	jobs.Main.CallFrames,
		// 	n.contractClient,
		// 	n.forkDb,
		// ).handleDeleteFork(ctxFork); errMain != nil {
		// 	log.Errorf("handle delete fork event data: %d", blockNum)
		// 	return blockNum, errMain
		// }
	}

	// handle fork job
	if jobs.Fork != nil {
		// blockNum = jobs.Fork.Block
		// if errFork = newBlockHandle(
		// 	jobs.Fork.BlockData,
		// 	jobs.Fork.TransactionDatas,
		// 	jobs.Fork.ReceiptDatas,
		// 	jobs.Fork.ContractOrMemberData,
		// 	jobs.Fork.ContractInfoMap,
		// 	jobs.Fork.ProxyContracts,
		// 	jobs.Fork.InternalTxs,
		// 	jobs.Fork.CallFrames,
		// 	n.contractClient,
		// 	n.forkDb,
		// ).handleFork(ctxFork); errFork != nil {
		// 	log.Errorf("handle fork event data: %d", blockNum)
		// 	return blockNum, errFork
		// }
	}

	if forkHandle != nil {
		if errFork = forkHandle.handleContractData(ctxFork); errFork != nil {
			log.Errorf("handle contract data from fork: %d", forkHandle.blockData.Number.String())
			return errFork
		}
	} else if mainHandle != nil {
		if errFork = mainHandle.handleContractData(ctxMain); errFork != nil {
			log.Errorf("handle contract data from fork: %d", mainHandle.blockData.Number.String())
			return errFork
		}
	}

	return nil
}

func (n *Sync) toGetDebugLog(txes []*types.Tx) {
	for _, tx := range txes {
		if len(tx.Data) > 0 {
			job.DebugJobChan.AddJob(job.NewSyncDebugJob(tx.Hash, n.client, n.db))
		}
	}
}
