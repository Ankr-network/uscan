package job

import (
	"context"
	"math/big"
	"time"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SyncJob struct {
	Completed            bool
	Block                uint64
	client               rpcclient.RpcClient
	BlockData            *types.Block
	TransactionDatas     []*types.Tx
	ReceiptDatas         []*types.Rt
	CallFrames           map[common.Hash]*types.CallFrame
	InternalTxs          []*types.InternalTx
	ContractOrMemberData map[common.Address]*types.Account
}

func NewSyncJob(block uint64, client rpcclient.RpcClient) *SyncJob {
	return &SyncJob{
		Block:                block,
		client:               client,
		ContractOrMemberData: make(map[common.Address]*types.Account),
	}
}

func (e *SyncJob) Execute() {
	var (
		err error
		ctx = context.Background()
	)

	// get block data
	for {
		e.BlockData, err = e.client.GetBlockByNumber(ctx, (*hexutil.Big)(big.NewInt(int64(e.Block))).String())
		if err != nil {
			log.Errorf("get block(%d) data failed: %v", e.Block, err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	// get balance of miner
	if e.BlockData.Coinbase != (common.Address{}) {
		e.ContractOrMemberData[e.BlockData.Coinbase] = &types.Account{
			Owner: e.BlockData.Coinbase,
		}
	}

	if len(e.BlockData.Transactions) != 0 {
		e.TransactionDatas = make([]*types.Tx, 0, len(e.BlockData.Transactions))
		e.ReceiptDatas = make([]*types.Rt, 0, len(e.BlockData.Transactions))
		e.CallFrames = make(map[common.Hash]*types.CallFrame, len(e.BlockData.Transactions))
		e.InternalTxs = make([]*types.InternalTx, 0, len(e.BlockData.Transactions))

		data := make([]*Jobs, len(e.BlockData.Transactions))
		for i, tx := range e.BlockData.Transactions {
			jobs := &Jobs{
				txJob:     NewSyncTxJob(tx, e.client),
				rtJob:     NewSyncRtJob(tx, e.client),
				tracerJob: NewSyncTracerJob(e.Block, tx, e.client),
			}
			data[i] = jobs
			TxJobChan.AddJob(jobs.txJob)
			TxJobChan.AddJob(jobs.rtJob)
			TxJobChan.AddJob(jobs.tracerJob)
		}

		for _, v := range data {
			for {
				if v.txJob.Completed && v.rtJob.Completed && v.tracerJob.Completed {
					if len(v.tracerJob.InternalTxs) > 0 {
						v.rtJob.ReceiptData.ExistInternalTx = true
					}
					v.rtJob.ReceiptData.ReturnErr = v.tracerJob.Error
					e.TransactionDatas = append(e.TransactionDatas, v.txJob.TransactionData)
					e.ReceiptDatas = append(e.ReceiptDatas, v.rtJob.ReceiptData)
					e.CallFrames[v.tracerJob.tx] = v.tracerJob.CallFrame
					e.InternalTxs = append(e.InternalTxs, v.tracerJob.InternalTxs...)
					e.ContractOrMemberData = e.mergeContractOrMember(e.ContractOrMemberData, v.tracerJob.ContractOrMemberData)
					e.ContractOrMemberData = e.mergeContractOrMember(e.ContractOrMemberData, v.txJob.ContractOrMemberData)
					break
				} else {
					time.Sleep(time.Millisecond * 500)
				}
			}
		}
	}
	addresses := make([]common.Address, 0, len(e.ContractOrMemberData))
	for k := range e.ContractOrMemberData {
		if k == (common.Address{}) {
			delete(e.ContractOrMemberData, k)
			continue
		}
		addresses = append(addresses, k)
	}

	for {
		if balanceMap, err := e.client.GetBalances(ctx, addresses, hexutil.EncodeUint64(e.Block)); err != nil {
			time.Sleep(time.Second)
		} else {
			for k, v := range balanceMap {
				e.ContractOrMemberData[k].Balance = v
			}
			break
		}
	}

	e.Completed = true
}

func (e *SyncJob) mergeContractOrMember(data map[common.Address]*types.Account, data2 map[common.Address]*types.Account) map[common.Address]*types.Account {
	for k, v := range data2 {
		if _, ok := data[k]; ok {
			if len(v.Code) != 0 {
				data[k].Code = v.Code
			}
			if v.Creator != nil {
				data[k].Creator = v.Creator
			}
			if v.TxHash != nil {
				data[k].TxHash = v.TxHash
			}
		} else {
			data[k] = v
		}
	}
	return data
}

type Jobs struct {
	txJob     *SyncTxJob
	rtJob     *SyncRtJob
	tracerJob *SyncTracerJob
}
