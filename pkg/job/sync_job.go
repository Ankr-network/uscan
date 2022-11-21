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
	InternalTxs          map[common.Hash][]*types.InternalTx
	ContractOrMemberData map[common.Address]*types.Account
	ContractInfoMap      map[common.Address]*types.Contract
	ProxyContracts       map[common.Address]common.Address
}

func NewSyncJob(block uint64, client rpcclient.RpcClient) *SyncJob {
	return &SyncJob{
		Block:                block,
		client:               client,
		ContractOrMemberData: make(map[common.Address]*types.Account),
		ContractInfoMap:      make(map[common.Address]*types.Contract),
		ProxyContracts:       make(map[common.Address]common.Address),
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
		e.InternalTxs = make(map[common.Hash][]*types.InternalTx)

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
					v.txJob.TransactionData.TimeStamp = e.BlockData.TimeStamp
					v.rtJob.ReceiptData.ReturnErr = v.tracerJob.Error
					e.TransactionDatas = append(e.TransactionDatas, v.txJob.TransactionData)
					e.ReceiptDatas = append(e.ReceiptDatas, v.rtJob.ReceiptData)
					e.CallFrames[v.tracerJob.tx] = v.tracerJob.CallFrame
					e.InternalTxs[v.tracerJob.tx] = v.tracerJob.InternalTxs
					e.mergeContractOrMember(v.tracerJob.ContractOrMemberData)
					e.mergeContractOrMember(v.txJob.ContractOrMemberData)
					e.mergeContract(v.tracerJob.ContractInfoMap)
					e.mergeProxyContract(v.tracerJob.ProxyContract)
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
				e.ContractOrMemberData[k].Balance = *v
			}
			break
		}
	}

	e.Completed = true
}

func (e *SyncJob) mergeContractOrMember(data map[common.Address]*types.Account) {
	for k, v := range data {
		if _, ok := e.ContractOrMemberData[k]; ok {
			if v.Creator != (common.Address{}) {
				e.ContractOrMemberData[k].Creator = v.Creator
			}
			if v.TxHash != (common.Hash{}) {
				e.ContractOrMemberData[k].TxHash = v.TxHash
			}
		} else {
			e.ContractOrMemberData[k] = v
		}
	}
}

func (e *SyncJob) mergeContract(data map[common.Address]*types.Contract) {
	for k, v := range data {
		if _, ok := e.ContractInfoMap[k]; !ok {
			e.ContractInfoMap[k] = v
		}
	}
}

func (e *SyncJob) mergeProxyContract(data map[common.Address]common.Address) {
	for k, v := range data {
		e.ProxyContracts[k] = v
	}
}

type Jobs struct {
	txJob     *SyncTxJob
	rtJob     *SyncRtJob
	tracerJob *SyncTracerJob
}
