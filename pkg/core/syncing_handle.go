package core

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/contract"
	"github.com/Ankr-network/uscan/pkg/contract/eip"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/forkcache"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var deleteMap = make(map[string][][]byte, 0)                            // table => key
var blockDeleteMap = make(map[*field.BigInt]map[string][][]byte, 0)     // block number => table/key
var indexMap = make(map[string]*field.BigInt, 0)                        // key => index
var blockIndexMap = make(map[*field.BigInt]map[string]*field.BigInt, 0) // block number => key/index
var totalMap = make(map[string]*field.BigInt, 0)                        // table:key => total
var blockTotalMap = make(map[*field.BigInt]map[string]*field.BigInt, 0) // block number => table:key/total
var HomeMap = make(map[*field.BigInt]*Home, 0)                          // block number => home

type Home struct {
	TxTotal      field.BigInt
	AddressTotal field.BigInt
	Erc20Total   field.BigInt
	Erc721Total  field.BigInt
	Erc1155Total field.BigInt
}

type blockHandle struct {
	blockData            *types.Block
	transactionData      []*types.Tx
	receiptData          []*types.Rt
	contractOrMemberData map[common.Address]*types.Account
	contractInfoMap      map[common.Address]*types.Contract
	proxyContracts       map[common.Address]common.Address
	internalTxs          map[common.Hash][]*types.InternalTx
	callFrames           map[common.Hash]*types.CallFrame
	contractClient       contract.Contractor
	db                   kv.Database

	newAddrTotal    *field.BigInt
	newErc20Total   *field.BigInt
	newErc721Total  *field.BigInt
	newErc1155Total *field.BigInt
}

func newBlockHandle(
	blockData *types.Block,
	transactionData []*types.Tx,
	receiptData []*types.Rt,
	contractOrMemberData map[common.Address]*types.Account,
	contractInfoMap map[common.Address]*types.Contract,
	proxyContracts map[common.Address]common.Address,
	internalTxs map[common.Hash][]*types.InternalTx,
	callFrames map[common.Hash]*types.CallFrame,
	contractClient contract.Contractor,
	db kv.Database,
) *blockHandle {
	return &blockHandle{
		blockData:            blockData,
		transactionData:      transactionData,
		receiptData:          receiptData,
		contractOrMemberData: contractOrMemberData,
		contractInfoMap:      contractInfoMap,
		proxyContracts:       proxyContracts,
		internalTxs:          internalTxs,
		callFrames:           callFrames,
		contractClient:       contractClient,
		db:                   db,
		newAddrTotal:         field.NewInt(0),
		newErc20Total:        field.NewInt(0),
		newErc721Total:       field.NewInt(0),
		newErc1155Total:      field.NewInt(0),
	}
}

func (n *blockHandle) handleMain(ctx context.Context) (err error) {
	err = rawdb.WriteBlock(ctx, n.db, n.blockData.Number, n.blockData)
	if err != nil {
		log.Errorf("write block : %v, block: %s", err, n.blockData.Number.String())
		return err
	}

	n.newAddrTotal, err = n.checkNewAddr(ctx)
	if err != nil {
		log.Errorf("read acccount to merge: %v", err)
		return err
	}

	//if len(n.contractInfoMap) > 0 {
	//	if err = n.writeContract(ctx, n.contractInfoMap); err != nil {
	//		log.Errorf("write contract: %v", err)
	//		return err
	//	}
	//}
	//if len(n.proxyContracts) > 0 {
	//	if err = n.writeProxyContract(ctx, n.proxyContracts); err != nil {
	//		log.Errorf("write proxy contract: %v", err)
	//		return err
	//	}
	//}

	if len(n.transactionData) > 0 {
		if err = n.writeTxAndRtLog(ctx, n.transactionData, n.receiptData); err != nil {
			log.Errorf("write tx and rt: %v", err)
			return err
		}

		if err = n.writeITx(ctx, n.internalTxs); err != nil {
			log.Errorf("write itxs: %v", err)
			return err
		}

		if err = n.writeTraceTx2(ctx, n.callFrames); err != nil {
			log.Errorf("write callFrames: %v", err)
			return err
		}
	}

	// all account about block write to kv
	if err = n.updateAccounts(ctx); err != nil {
		log.Errorf("write account : %v", err)
		return err
	}

	if err = n.updateHome(ctx); err != nil {
		log.Errorf("write home : %v", err)
		return err
	}

	return nil
}

func (n *blockHandle) handleDeleteFork(ctx context.Context) (err error) {
	err = forkcache.DeleteBlock(ctx, n.db, n.blockData.Number)
	if err != nil {
		log.Errorf("delete fork block : %v, block: %s", err, n.blockData.Number.String())
		return err
	}

	if len(n.transactionData) > 0 {
		if err = n.deleteForkTxAndRtLog(ctx, n.transactionData, n.receiptData); err != nil {
			log.Errorf("delete fork tx and rt: %v", err)
			return err
		}

		if err = n.deleteForkTraceTx2(ctx, n.callFrames); err != nil {
			log.Errorf("delete fork callFrames: %v", err)
			return err
		}
	}

	if err = n.deleteForkHome(ctx); err != nil {
		log.Errorf("delete fork home : %v", err)
		return err
	}

	for k, v := range blockDeleteMap {
		if k == n.blockData.Number {
			for k1, v1 := range v {
				for _, v2 := range v1 {
					err = n.db.Del(ctx, v2, &kv.WriteOption{Table: k1})
					if err != nil {
						return err
					}
				}
				delete(deleteMap, k1)
			}
			delete(blockDeleteMap, k)
		}
	}

	for k, v := range blockIndexMap {
		if k == n.blockData.Number {
			for k1, v1 := range v {
				var i *field.BigInt
				bytesRes, err := n.db.Get(ctx, []byte(k1), &kv.ReadOption{Table: share.ForkIndexTbl})
				if err != nil {
					if errors.Is(err, kv.NotFound) {
						i = field.NewInt(0)
						err = nil
					} else {
						return err
					}
				}
				i.SetBytes(bytesRes)
				i.Add(v1)
				err = n.db.Put(ctx, []byte(k1), i.Bytes(), &kv.WriteOption{Table: share.ForkIndexTbl})
				if err != nil {
					return err
				}
				delete(indexMap, k1)
			}
			delete(blockIndexMap, k)
		}
	}

	for k, v := range blockTotalMap {
		if k == n.blockData.Number {
			for k1, v1 := range v {
				arr := strings.Split(k1, ":")
				tableName := arr[0]
				key := []byte(arr[1])
				var i *field.BigInt
				bytesRes, err := n.db.Get(ctx, key, &kv.ReadOption{Table: tableName})
				if err != nil {
					if errors.Is(err, kv.NotFound) {
						i = field.NewInt(0)
						err = nil
					} else {
						return err
					}
				}
				i.SetBytes(bytesRes)
				i.Sub(v1)
				err = n.db.Put(ctx, key, i.Bytes(), &kv.WriteOption{Table: tableName})
				if err != nil {
					return err
				}
				delete(totalMap, k1)
			}
			delete(blockTotalMap, k)
		}
	}

	return nil
}

func (n *blockHandle) handleFork(ctx context.Context) (err error) {

	err = forkcache.WriteBlock(ctx, n.db, n.blockData.Number, n.blockData)
	if err != nil {
		log.Errorf("write fork block : %v, block: %s", err, n.blockData.Number.String())
		return err
	}

	n.newAddrTotal, err = n.checkForkNewAddr(ctx)
	if err != nil {
		log.Errorf("read fork account to merge: %v", err)
		return err
	}

	//if len(n.contractInfoMap) > 0 {
	//	if err = n.writeContract(ctx, n.contractInfoMap); err != nil {
	//		log.Errorf("write contract: %v", err)
	//		return err
	//	}
	//}
	//if len(n.proxyContracts) > 0 {
	//	if err = n.writeProxyContract(ctx, n.proxyContracts); err != nil {
	//		log.Errorf("write proxy contract: %v", err)
	//		return err
	//	}
	//}

	if len(n.transactionData) > 0 {
		if err = n.writeForkTxAndRtLog(ctx, n.transactionData, n.receiptData); err != nil {
			log.Errorf("write fork tx and rt: %v", err)
			return err
		}

		if err = n.writeForkITx(ctx, n.internalTxs); err != nil {
			log.Errorf("write fork itxs: %v", err)
			return err
		}

		if err = n.writeForkTraceTx2(ctx, n.callFrames); err != nil {
			log.Errorf("write fork callFrames: %v", err)
			return err
		}
	}

	// all account about block write to kv
	if err = n.updateForkAccounts(ctx); err != nil {
		log.Errorf("write fork account : %v", err)
		return err
	}

	if err = n.updateForkHome(ctx); err != nil {
		log.Errorf("write fork home : %v", err)
		return err
	}

	blockDeleteMap[n.blockData.Number] = deleteMap
	blockIndexMap[n.blockData.Number] = indexMap
	blockTotalMap[n.blockData.Number] = totalMap

	return nil
}

func (n *blockHandle) handleContractData(ctx context.Context) (err error) {
	if len(n.contractInfoMap) > 0 {
		if err = n.writeContract(ctx, n.contractInfoMap); err != nil {
			log.Errorf("write contract: %v", err)
			return err
		}
	}
	if len(n.proxyContracts) > 0 {
		if err = n.writeProxyContract(ctx, n.proxyContracts); err != nil {
			log.Errorf("write proxy contract: %v", err)
			return err
		}
	}
	return nil
}

func (n *blockHandle) writeTxAndRtLog(ctx context.Context, transactionData []*types.Tx, receiptData []*types.Rt) (err error) {

	for i, v := range transactionData {
		err = rawdb.WriteBlockIndex(ctx, n.db, n.blockData.Number, field.NewInt(int64(i)), v.Hash)
		if err != nil {
			log.Errorf("write block index(%d): %v", i, err)
			return err
		}
		if err = n.writeTxAndRt(ctx, v, receiptData[i]); err != nil {
			log.Errorf("writeTxAndRt tx(%s): %v", v.Hash.Hex(), err)
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
							if err = n.writeErc20Transfer(ctx, &types.Erc20Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc20Transfer.From,
								To:              erc20Transfer.To,
								Amount:          (field.BigInt)(*erc20Transfer.Value),
								TimeStamp:       n.blockData.TimeStamp,
							}); err != nil {
								log.Errorf("write erc20Transfer: %v", err)
								return err
							}
						}
					} else {
						erc721Transfer, err = n.contractClient.Erc721Transfer(rtLog.Address.Hex(), rtLog.ToEthLog())
						if err == nil {
							if err = n.writeErc721Transfer(ctx, &types.Erc721Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc721Transfer.From,
								To:              erc721Transfer.To,
								TokenId:         (field.BigInt)(*erc721Transfer.TokenId),
								TimeStamp:       n.blockData.TimeStamp,
							}); err != nil {
								log.Errorf("write erc721Transfer: %v", err)
								return err
							}
						}
					}

				case contract.TransferSingleEventTopic:
					erc1155TransferSignle, err = n.contractClient.Erc1155TransferSingle(rtLog.Address.Hex(), rtLog.ToEthLog())
					if err == nil {
						if err = n.writeErc1155Transfer(ctx, &types.Erc1155Transfer{
							TransactionHash: v.Hash,
							BlockNumber:     v.BlockNum,
							Contract:        rtLog.Address,
							Method:          v.Method,
							From:            erc1155TransferSignle.From,
							To:              erc1155TransferSignle.To,
							TokenID:         (field.BigInt)(*erc1155TransferSignle.Id),
							Quantity:        (field.BigInt)(*erc1155TransferSignle.Value),
							TimeStamp:       n.blockData.TimeStamp,
						}); err != nil {
							log.Errorf("write erc1155Transfer single: %v", err)
							return err
						}
					}
				case contract.TransferBatchEventTopic:
					erc1155TransferBatch, err = n.contractClient.Erc1155TransferBatch(rtLog.Address.Hex(), rtLog.ToEthLog())
					if err == nil {
						for i := range erc1155TransferBatch.Ids {
							if err = n.writeErc1155Transfer(ctx, &types.Erc1155Transfer{
								TransactionHash: v.Hash,
								BlockNumber:     v.BlockNum,
								Contract:        rtLog.Address,
								Method:          v.Method,
								From:            erc1155TransferBatch.From,
								To:              erc1155TransferBatch.To,
								TokenID:         (field.BigInt)(*erc1155TransferBatch.Ids[i]),
								Quantity:        (field.BigInt)(*erc1155TransferBatch.Values[i]),
								TimeStamp:       n.blockData.TimeStamp,
							}); err != nil {
								log.Errorf("write erc1155Transfer single: %v", err)
								return err
							}
						}
					}
				}
			}
		}

		if err = n.updateErc20TrasferTotal(ctx); err != nil {
			log.Errorf("update erc20 transfer total: %v", err)
			return err
		}
		if err = n.updateErc721TrasferTotal(ctx); err != nil {
			log.Errorf("update erc721 transfer total: %v", err)
			return err
		}
		if err = n.updateErc1155TrasferTotal(ctx); err != nil {
			log.Errorf("update erc1155 transfer total: %v", err)
			return err
		}
	}

	return n.writeTxTotal(ctx)
}

func (n *blockHandle) writeTraceTx2(ctx context.Context, callFrames map[common.Hash]*types.CallFrame) (err error) {
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
