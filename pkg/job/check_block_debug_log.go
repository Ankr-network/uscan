package job

// import (
// 	"context"
// 	"math/big"
// 	"time"

// 	"github.com/Ankr-network/chainscan_syncing/pkg/field"
// 	"github.com/Ankr-network/uscan/pkg/model"
// 	"github.com/Ankr-network/uscan/pkg/repository"
// 	"github.com/Ankr-network/uscan/pkg/rpcclient"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/rs/zerolog/log"
// )

// type CheckBlockDebugLog struct {
// 	Block    uint64
// 	client   rpcclient.RpcClient
// 	dbRepo   repository.DBer
// 	Complete bool
// }

// func NewCheckBlockDebugLog(block uint64,
// 	client rpcclient.RpcClient,
// 	dbRepo repository.DBer,
// ) *CheckBlockDebugLog {
// 	return &CheckBlockDebugLog{
// 		Block:  block,
// 		client: client,
// 		dbRepo: dbRepo,
// 	}
// }

// func (e *CheckBlockDebugLog) Execute() {
// 	var (
// 		err       error
// 		tx        = e.dbRepo.GetDB().Begin()
// 		ctx       = e.dbRepo.ContextWithTx(context.Background(), tx)
// 		blockData *rpcclient.Block
// 		logRes    *rpcclient.ExecutionResult
// 	)
// 	log.Debug().Uint64("block", e.Block).Msg("rpc call start")

// 	defer func() {
// 		if err == nil {
// 			tx.Commit()
// 		} else {
// 			log.Error().Err(err).Msg("callback because of error")
// 			tx.Callback()
// 		}
// 	}()
// 	// get block data
// 	for {
// 		blockData, err = e.client.GetBlockByNumber(ctx, (*hexutil.Big)(big.NewInt(int64(e.Block))).String())
// 		if err != nil {
// 			log.Error().Err(err).Uint64("block", e.Block).Msg("get block data failed")
// 			time.Sleep(time.Second)
// 		} else {
// 			break
// 		}
// 	}

// 	for _, v := range blockData.Transactions {
// 		count, err := e.dbRepo.Count(ctx, &model.TraceTx{}, "transaction_hash = ?", []interface{}{v})
// 		if err != nil {
// 			log.Error().Err(err).Uint64("block", e.Block).Str("tx", v).Msg("get count")
// 			return
// 		}
// 		if count == 0 {
// 			for {
// 				logRes, err = e.client.GetTracerLog(ctx, v)
// 				if err != nil {
// 					log.Error().Err(err).Str("tx", v).Msg("get trancerlogs data failed")
// 					time.Sleep(time.Second)
// 				} else {
// 					break
// 				}
// 			}
// 			if len(logRes.StructLogs) == 0 {
// 				continue
// 			} else {
// 				err = e.dbRepo.Save(ctx, model.TraceTx{}, &model.TraceTx{
// 					TransactionHash: field.Hash(common.HexToHash(v)),
// 					Res:             logRes.StructLogs.JsonToString(),
// 				})
// 				if err != nil {
// 					log.Error().Err(err).Str("tx", v).Msg("save to table")
// 					return
// 				}
// 			}
// 		}
// 	}
// 	e.Complete = true
// }
