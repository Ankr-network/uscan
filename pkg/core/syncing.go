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
	"github.com/ethereum/go-ethereum/common"
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
) *Sync {
	s := &Sync{
		client:         client,
		contractClient: contractClient,
		db:             db,
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

func (n *Sync) storeEvent() {
	var blockNum uint64
	for job := range n.storeChan {
		for {
			blockNum = job.Block
			if job.Completed {
				if err := n.handleEventData(
					job.BlockData,
					job.TransactionDatas,
					job.ReceiptDatas,
					job.ContractOrMemberData,
					job.ContractInfoMap,
					job.InternalTxs,
					job.CallFrames,
				); err != nil {
					log.Errorf("handle event data: %d", job.Block)
					goto end
				} else {
					n.toGetDebugLog(job.TransactionDatas)
				}
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
end:
	log.Fatalf("handle failed: %d", blockNum)
}

func (n *Sync) handleEventData(
	blockData *types.Block,
	transactionData []*types.Tx,
	receiptData []*types.Rt,
	contractOrMemberData map[common.Address]*types.Account,
	contractInfoMap map[common.Address]*types.Contract,
	internalTxs map[common.Hash][]*types.InternalTx,
	callFrames map[common.Hash]*types.CallFrame,
) error {
	var (
		ctx, err = n.db.BeginTx(context.Background())
	)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			n.db.Commit(ctx)
		} else {
			n.db.RollBack(ctx)
		}
	}()

	err = rawdb.WriteBlock(ctx, n.db, blockData.Number, blockData)
	if err != nil {
		log.Errorf("write block : %v, block: %s", err, blockData.Number.String())
		return err
	}

	contractOrMemberData, err = n.readAccountToMerge(ctx, contractOrMemberData)
	if err != nil {
		log.Errorf("read acccount to merge: %v", err)
		return err
	}

	if len(contractInfoMap) > 0 {
		if err = n.writeContract(ctx, contractInfoMap); err != nil {
			log.Errorf("write contract: %v", err)
			return err
		}
	}

	if len(transactionData) > 0 {
		if err = n.writeTxAndRt(ctx, transactionData, receiptData); err != nil {
			log.Errorf("write tx and rt: %v", err)
			return err
		}

		if err = n.writeITx(ctx, internalTxs); err != nil {
			log.Errorf("write itxs: %v", err)
			return err
		}

		if err = n.writeTraceTx2(ctx, callFrames); err != nil {
			log.Errorf("write callFrames: %v", err)
			return err
		}

		if err = n.writeReceiptDataLog(ctx, contractOrMemberData, receiptData); err != nil {
			log.Errorf("write receipt log: %v", err)
			return err
		}

	}
	// to do

	// all account about block write to kv
	err = n.writeAccount(ctx, contractOrMemberData)
	if err != nil {
		log.Errorf("write account : %v", err)
		return err
	}
	return nil
}

func (n *Sync) writeContract(ctx context.Context, data map[common.Address]*types.Contract) (err error) {
	for k, v := range data {
		if err = rawdb.WriteContract(ctx, n.db, k, v); err != nil {
			log.Errorf("write contract(%s): %v ", k, err)
			return err
		}
	}
	return nil
}

func (n *Sync) writeTxAndRt(ctx context.Context, transactionData []*types.Tx, receiptData []*types.Rt) error {
	txTotal, err := rawdb.ReadTxTotal(ctx, n.db)
	if err != nil {
		if errors.Is(err, kv.NotFound) {
			txTotal = field.NewInt(0)
		} else {
			log.Errorf("get tx total: %v", err)
			return err
		}
	}

	for i, v := range transactionData {
		if err = rawdb.WriteTx(ctx, n.db, v.Hash, v); err != nil {
			log.Errorf("write tx(%s): %v", v.Hash.Hex(), err)
			return err
		}
		if err = rawdb.WriteTxIndex(ctx, n.db, txTotal.Add(field.NewInt(1)), v.Hash); err != nil {
			log.Errorf("write tx(%s) index: %v", v.Hash.Hex(), err)
			return err
		}

		if err = rawdb.WriteRt(ctx, n.db, v.Hash, receiptData[i]); err != nil {
			log.Errorf("write rt: %v", err)
			return err
		}
	}
	return rawdb.WriteTxTotal(ctx, n.db, txTotal)
}

func (n *Sync) writeITx(ctx context.Context, itxmap map[common.Hash][]*types.InternalTx) error {
	for k, itxs := range itxmap {
		itxTotal, err := rawdb.ReadITxTotal(ctx, n.db, k)
		if errors.Is(err, kv.NotFound) {
			itxTotal = field.NewInt(0)
		} else {
			log.Errorf("get itx total: %v", err)
			return err
		}

		for _, v := range itxs {
			if err = rawdb.WriteITx(ctx, n.db, k, itxTotal.Add(field.NewInt(1)), v); err != nil {
				log.Errorf("write itx(%s): %v", k.Hex(), err)
				return err
			}
		}
		if err = rawdb.WriteItxTotal(ctx, n.db, k, itxTotal); err != nil {
			log.Errorf("write itx total: %v", err)
			return err
		}
	}
	return nil
}

func (n *Sync) readAccount(ctx context.Context, addr common.Address, data map[common.Address]*types.Account) (*types.Account, error) {
	acc, ok := data[addr]
	if ok {
		return acc, nil
	}
	account, err := rawdb.ReadAccount(ctx, n.db, addr)
	if err != nil {
		if !errors.Is(err, kv.NotFound) {
			log.Errorf("read account(%s): %v", addr.Hex(), err)
			return nil, err
		}
		account = &types.Account{}
	}
	data[addr] = account
	return account, nil
}

func (n *Sync) writeAccount(ctx context.Context, data map[common.Address]*types.Account) (err error) {
	for k, v := range data {
		if err = rawdb.WriteAccount(ctx, n.db, k, v); err != nil {
			log.Errorf("write account(%s): %v", k.Hex(), err)
			return err
		}
	}
	return nil
}

func (n *Sync) readAccountToMerge(ctx context.Context, data map[common.Address]*types.Account) (map[common.Address]*types.Account, error) {
	var accs = make(map[common.Address]*types.Account, len(data))
	for k, v := range data {
		account, err := rawdb.ReadAccount(ctx, n.db, k)
		if err != nil {
			if !errors.Is(err, kv.NotFound) {
				log.Errorf("read account(%s): %v", k.Hex(), err)
				return nil, err
			}
			account = &types.Account{}
		}

		accs[k] = n.mergeAccount(account, v)
	}

	return accs, nil
}

func (n *Sync) mergeAccount(beforeAcc *types.Account, afterAcc *types.Account) *types.Account {
	if beforeAcc.BlockNumber.String() == "0x0" {
		beforeAcc.BlockNumber = afterAcc.BlockNumber
	}
	beforeAcc.Balance = afterAcc.Balance
	if afterAcc.Erc20 {
		beforeAcc.Erc20 = true
	}

	if afterAcc.Erc721 {
		beforeAcc.Erc721 = true
	}
	if afterAcc.Erc1155 {
		beforeAcc.Erc1155 = true
	}
	if beforeAcc.Creator == (&common.Address{}) {
		beforeAcc.Creator = afterAcc.Creator
	}

	if beforeAcc.TxHash == (&common.Hash{}) {
		beforeAcc.TxHash = afterAcc.TxHash
	}

	if beforeAcc.Name == "" {
		beforeAcc.Name = afterAcc.Name
	}

	if beforeAcc.Symbol == "" {
		beforeAcc.Symbol = afterAcc.Symbol
	}
	if afterAcc.TokenTotalSupply != nil {
		beforeAcc.TokenTotalSupply = afterAcc.TokenTotalSupply
	}
	if afterAcc.NftTotalSupply != nil {
		beforeAcc.NftTotalSupply = afterAcc.NftTotalSupply
	}
	return beforeAcc
}

func (n *Sync) writeTraceTx2(ctx context.Context, callFrames map[common.Hash]*types.CallFrame) (err error) {
	for k, v := range callFrames {
		if err = rawdb.WriteTraceTx2(ctx, n.db, k, &types.TraceTx2{
			Res: v.JsonToString(),
		}); err != nil {
			log.Errorf("write trace tx2: %v", err)
			return err
		}
	}
	return nil
}

func (n *Sync) writeReceiptDataLog(ctx context.Context, accs map[common.Address]*types.Account, rts []*types.Rt) error {
	return nil
}

func (n *Sync) CountTransfer(ctx context.Context, block uint64, erc20_transfers, erc721_transfers, erc1155_transfers int) (err error) {
	return nil
}

func (n *Sync) toGetDebugLog(txes []*types.Tx) {
	for _, tx := range txes {
		if len(tx.Data) > 0 {
			job.DebugJobChan.AddJob(job.NewSyncDebugJob(tx.Hash, n.client, n.db))
		}
	}
}
