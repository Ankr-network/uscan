package rpcclient

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RpcClient interface {
	ChainID(ctx context.Context) uint64
	GetLatestBlockNumber(ctx context.Context) <-chan uint64
	GetClient() *ethclient.Client
	Close()

	GetBlockByNumber(ctx context.Context, blockNumber string) (*types.Block, error)
	GetTransactionsByHash(ctx context.Context, transactionHash []common.Hash) ([]*types.Tx, error)
	GetTransactionReceiptsByHash(ctx context.Context, transactionHash []common.Hash) ([]*types.Rt, error)
	GetTransactionByHash(ctx context.Context, transactionHash common.Hash) (*types.Tx, error)
	GetTransactionReceiptByHash(ctx context.Context, transactionHash common.Hash) (*types.Rt, error)
	GetCode(ctx context.Context, address common.Address, blockNumber string) (string, error)
	GetBalance(ctx context.Context, address common.Address, blockNumber string) (*field.BigInt, error)
	GetBalances(ctx context.Context, addresses []common.Address, blockNumber string) (map[common.Address]*field.BigInt, error)
	GetTracerCall(ctx context.Context, txhash common.Hash) (*types.CallFrame, error)
	GetTracerLog(ctx context.Context, txHash common.Hash) (*types.ExecutionResult, error)
}
