package core

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/contract"
	"github.com/Ankr-network/uscan/pkg/contract/eip"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/forkdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

func (n *blockHandle) writeForkTxAndRtLog(ctx context.Context, transactionData []*types.Tx, receiptData []*types.Rt, deleteMap map[string][][]byte, indexMap, totalMap, accountTotalMap, erc20TotalMap, erc721TotalMap, erc1155TotalMap, erc20ContractTotalMap, erc721ContractTotalMap, erc1155ContractTotalMap map[string]*field.BigInt) (err error) {

	for i, v := range transactionData {
		err = forkdb.WriteBlockIndex(ctx, n.db, n.blockData.Number, field.NewInt(int64(i)), v.Hash)
		if err != nil {
			log.Errorf("write fork block index(%d): %v", i, err)
			return err
		}
		deleteMap[share.ForkBlockTbl] = append(deleteMap[share.ForkBlockTbl], append(append([]byte("/fork/block/"), n.blockData.Number.Bytes()...), append([]byte("/"), field.NewInt(int64(i)).Bytes()...)...))

		if err = n.writeForkTxAndRt(ctx, v, receiptData[i], deleteMap, indexMap, totalMap, accountTotalMap); err != nil {
			log.Errorf("writeForkTxAndRt tx(%s): %v", v.Hash.Hex(), err)
			return err
		}

		var (
			erc20Transfer         *eip.Erc20Transfer
			erc721Transfer        *eip.Ieip721Transfer
			erc1155TransferSignle *eip.Ieip1155TransferSingle
			erc1155TransferBatch  *eip.Ieip1155TransferBatch
		)
		for _, rtLog := range receiptData[i].Logs {
			if len(rtLog.Topics) >= 3 {
				switch rtLog.Topics[0] {
				case contract.TransferEventTopic:
					if len(rtLog.Data) > 0 {
						erc20Transfer, err = n.contractClient.Erc20Transfer(rtLog.Address.Hex(), rtLog.ToEthLog())
						if err == nil {
							if err = n.writeForkErc20Transfer(ctx, &types.Erc20Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc20Transfer.From,
								To:              erc20Transfer.To,
								Amount:          (field.BigInt)(*erc20Transfer.Value),
								TimeStamp:       n.blockData.TimeStamp,
							}, deleteMap, indexMap, totalMap, erc20TotalMap, erc20ContractTotalMap); err != nil {
								log.Errorf("write fork erc20Transfer: %v", err)
								return err
							}
						}
					} else {
						erc721Transfer, err = n.contractClient.Erc721Transfer(rtLog.Address.Hex(), rtLog.ToEthLog())
						if err == nil {
							if err = n.writeForkErc721Transfer(ctx, &types.Erc721Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc721Transfer.From,
								To:              erc721Transfer.To,
								TokenId:         (field.BigInt)(*erc721Transfer.TokenId),
								TimeStamp:       n.blockData.TimeStamp,
							}, deleteMap, indexMap, totalMap, erc721TotalMap, erc721ContractTotalMap); err != nil {
								log.Errorf("write fork erc721Transfer: %v", err)
								return err
							}
						}
					}

				case contract.TransferSingleEventTopic:
					erc1155TransferSignle, err = n.contractClient.Erc1155TransferSingle(rtLog.Address.Hex(), rtLog.ToEthLog())
					if err == nil {
						if err = n.writeForkErc1155Transfer(ctx, &types.Erc1155Transfer{
							TransactionHash: v.Hash,
							BlockNumber:     v.BlockNum,
							Contract:        rtLog.Address,
							Method:          v.Method,
							From:            erc1155TransferSignle.From,
							To:              erc1155TransferSignle.To,
							TokenID:         (field.BigInt)(*erc1155TransferSignle.Id),
							Quantity:        (field.BigInt)(*erc1155TransferSignle.Value),
							TimeStamp:       n.blockData.TimeStamp,
						}, deleteMap, indexMap, totalMap, erc1155TotalMap, erc1155ContractTotalMap); err != nil {
							log.Errorf("write fork erc1155Transfer single: %v", err)
							return err
						}
					}
				case contract.TransferBatchEventTopic:
					erc1155TransferBatch, err = n.contractClient.Erc1155TransferBatch(rtLog.Address.Hex(), rtLog.ToEthLog())
					if err == nil {
						for i := range erc1155TransferBatch.Ids {
							if err = n.writeForkErc1155Transfer(ctx, &types.Erc1155Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc1155TransferBatch.From,
								To:              erc1155TransferBatch.To,
								TokenID:         (field.BigInt)(*erc1155TransferBatch.Ids[i]),
								Quantity:        (field.BigInt)(*erc1155TransferBatch.Values[i]),
								TimeStamp:       n.blockData.TimeStamp,
							}, deleteMap, indexMap, totalMap, erc1155TotalMap, erc1155ContractTotalMap); err != nil {
								log.Errorf("write fork erc1155Transfer batch: %v", err)
								return err
							}
						}
					}
				}
			}
		}

		if err = n.updateForkErc20TrasferTotal(ctx, totalMap, erc20TotalMap); err != nil {
			log.Errorf("update fork erc20 transfer total: %v", err)
			return err
		}
		if err = n.updateForkErc721TrasferTotal(ctx, totalMap, erc721TotalMap); err != nil {
			log.Errorf("update fork erc721 transfer total: %v", err)
			return err
		}
		if err = n.updateForkErc1155TrasferTotal(ctx, totalMap, erc1155TotalMap); err != nil {
			log.Errorf("update fork erc1155 transfer total: %v", err)
			return err
		}
	}

	return n.writeForkTxTotal(ctx, totalMap, accountTotalMap)
}

func (n *blockHandle) writeForkTraceTx2(ctx context.Context, callFrames map[common.Hash]*types.CallFrame, deleteMap map[string][][]byte) (err error) {
	for k, v := range callFrames {
		if err = forkdb.WriteTraceTx2(ctx, n.db, k, &types.TraceTx2{
			Res: v.JsonToString(),
		}); err != nil {
			log.Errorf("write fork trace tx2: %v", err)
			return err
		}
		deleteMap[share.ForkTraceLogTbl] = append(deleteMap[share.ForkTraceLogTbl], append([]byte("/fork/tracetx2/"), k.Bytes()...))
	}
	return nil
}
