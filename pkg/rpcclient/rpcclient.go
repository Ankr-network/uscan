package rpcclient

import (
	"context"
	"math/big"
	"time"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ RpcClient = (*manage)(nil)

type rpcGroup struct {
	wsuri     string
	rpcClient *rpc.Client
	client    *ethclient.Client
}
type manage struct {
	chainId           *big.Int
	index             int
	headerChan        chan *ethTypes.Header
	clients           []*rpcGroup
	latestBlockNumber uint64
	latestChan        chan uint64
}

func NewRpcClient(ws []string) *manage {
	clients := make([]*rpcGroup, len(ws))
	for i, v := range ws {
		rpcClient, err := rpc.Dial(v)
		if err != nil {
			log.Fatal("rpc.dial failed ", "ws: ", v, "err: ", err)
		}
		client := ethclient.NewClient(rpcClient)
		clients[i] = &rpcGroup{
			wsuri:     v,
			rpcClient: rpcClient,
			client:    client,
		}
	}

	var (
		chainId         *big.Int
		index           int
		lastBlockNumber uint64
		lastNumber      uint64
		err             error
	)
	for i, v := range clients {
		if chainId == nil {
			chainId, err = v.client.ChainID(context.Background())
			if err != nil {
				log.Fatal("get chainId: ", err)
			}
		}
		lastBlockNumber, err = v.client.BlockNumber(context.Background())
		if err != nil {
			log.Fatal("get blockNumber: ", err)
		}
		if lastBlockNumber > lastNumber {
			lastNumber = lastBlockNumber
			index = i
		}
	}
	if chainId == nil {
		log.Fatal("get chainId")
	}

	r := &manage{
		chainId:           chainId,
		index:             index,
		headerChan:        make(chan *ethTypes.Header, share.MaxChanSize),
		clients:           clients,
		latestBlockNumber: lastNumber,
		latestChan:        make(chan uint64, share.MaxChanSize),
	}

	go r.syncerBlock(context.Background())
	return r
}

func (m *manage) syncerBlock(ctx context.Context) {
	for i, v := range m.clients {
		go func(client *rpcGroup, index int) {
			var (
				ctx        = context.Background()
				headerChan = make(chan *ethTypes.Header, share.MaxChanSize)
			)
			go func() {
				for {
					sub, err := client.client.SubscribeNewHead(ctx, headerChan)
					if err != nil {
						log.Errorf("subscribe(%s) head failed: %+v", client.wsuri, err)
						time.Sleep(time.Second * 3)
						continue
					}
					for err = range sub.Err() {
						log.Errorf("subscribe(%s) err: %+v", client.wsuri, err)
						time.Sleep(time.Second * 10)
						break
					}
				}
			}()

			for head := range headerChan {
				if head.Number.Uint64() > m.latestBlockNumber {
					m.latestBlockNumber = head.Number.Uint64()
					m.latestChan <- m.latestBlockNumber
					m.index = index
				}
			}
		}(v, i)
	}
}

func (r *manage) ChainID(ctx context.Context) uint64 {
	return r.chainId.Uint64()
}

func (r *manage) GetLatestBlockNumber(ctx context.Context) <-chan uint64 {
	return r.latestChan
}

func (r *manage) GetBlockByNumber(ctx context.Context, blockNumber string) (*types.Block, error) {
	result := &types.Block{}
	err := r.clients[r.index].rpcClient.CallContext(ctx, result, "eth_getBlockByNumber", blockNumber, false)
	if err != nil {
		log.Errorf("eth_getBlockByNumber err: %v; endpoint: %s", err, r.clients[r.index].wsuri)
		return nil, err
	}
	return result, nil
}

func (r *manage) GetTransactionsByHash(ctx context.Context, transactionHash []common.Hash) ([]*types.Tx, error) {
	result := make([]*types.Tx, 0, len(transactionHash))
	elem := make([]rpc.BatchElem, 0, len(transactionHash))
	for _, v := range transactionHash {
		res := types.Tx{}
		elem = append(elem, rpc.BatchElem{
			Method: "eth_getTransactionByHash",
			Args:   []interface{}{v.Hex()},
			Result: &res,
		})
		result = append(result, &res)
	}
	newCtx, cancel := context.WithTimeout(ctx, share.HttpTimeout)
	defer cancel()
	err := r.clients[r.index].rpcClient.BatchCallContext(newCtx, elem)
	if err != nil {
		log.Errorf("eth_getTransactionByHash err: %+v; endpoint: %s", err, r.clients[r.index].wsuri)
		return nil, err
	}
	for _, v := range elem {
		if v.Error != nil {
			log.Errorf("eth_getTransactionReceipt failed: %+v; args: %s", err, v.Args)
			return nil, v.Error
		}
	}

	return result, nil
}

func (r *manage) GetTransactionReceiptsByHash(ctx context.Context, transactionHash []common.Hash) ([]*types.Rt, error) {
	result := make([]*types.Rt, 0, len(transactionHash))
	elem := make([]rpc.BatchElem, 0, len(transactionHash))
	for _, v := range transactionHash {
		res := types.Rt{}
		elem = append(elem, rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{v.Hex()},
			Result: &res,
		})
		result = append(result, &res)
	}
	newCtx, cancel := context.WithTimeout(ctx, share.HttpTimeout)
	defer cancel()
	err := r.clients[r.index].rpcClient.BatchCallContext(newCtx, elem)
	if err != nil {
		log.Errorf("eth_getTransactionReceipt err: %+v; endpoint: %s", err, r.clients[r.index].wsuri)
		return nil, err
	}

	for _, v := range elem {
		if v.Error != nil {
			log.Errorf("eth_getTransactionReceipt failed: %+v; args :%s", err, v.Args)
			return nil, v.Error
		}
	}

	return result, nil
}

func (r *manage) GetTransactionByHash(ctx context.Context, transactionHash common.Hash) (res *types.Tx, err error) {
	res = &types.Tx{}
	err = r.clients[r.index].rpcClient.Call(res, "eth_getTransactionByHash", transactionHash.Hex())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *manage) GetTransactionReceiptByHash(ctx context.Context, transactionHash common.Hash) (res *types.Rt, err error) {
	res = &types.Rt{}
	err = r.clients[r.index].rpcClient.Call(res, "eth_getTransactionReceipt", transactionHash.Hex())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *manage) GetCode(ctx context.Context, address common.Address, blockNumber string) (string, error) {
	var res string
	err := r.clients[r.index].rpcClient.CallContext(ctx, &res, "eth_getCode", address, blockNumber)
	if err != nil {
		log.Errorf("eth_getCode err: %+v; endpoints: %s", err, r.clients[r.index].wsuri)
		return "", err
	}
	return res, nil
}

func (r *manage) GetBalance(ctx context.Context, address common.Address, blockNumber string) (*field.BigInt, error) {
	var res field.BigInt
	err := r.clients[r.index].rpcClient.CallContext(ctx, &res, "eth_getBalance", address, blockNumber)
	if err != nil {
		log.Errorf("eth_getBalance err: %+v; endpoint: %s", err, r.clients[r.index].wsuri)
		return nil, err
	}

	return &res, nil
}

func (r *manage) GetBalances(ctx context.Context, addresses []common.Address, blockNumber string) (map[common.Address]*field.BigInt, error) {
	result := make(map[common.Address]*field.BigInt, len(addresses))
	elem := make([]rpc.BatchElem, 0, len(addresses))
	for _, v := range addresses {
		var res = field.BigInt{}
		elem = append(elem, rpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []interface{}{v, blockNumber},
			Result: &res,
		})
		result[v] = &res
	}
	err := r.clients[r.index].rpcClient.BatchCallContext(context.Background(), elem)
	if err != nil {
		log.Errorf("eth_getBalance err: %+v; endpoint:%s", err, r.clients[r.index].wsuri)
		return nil, err
	}

	for _, v := range elem {
		if v.Error != nil {
			log.Errorf("eth_getBalance failed: %+v; args: %s", err, v.Args)
			return nil, v.Error
		}
	}

	return result, nil
}

func (r *manage) Close() {
	for _, v := range r.clients {
		v.client.Close()
	}
}

func (r *manage) GetClient() *ethclient.Client {
	return r.clients[r.index].client
}

func (r *manage) GetTracerCall(ctx context.Context, txhash common.Hash) (*types.CallFrame, error) {
	var err error
	var res = types.CallFrame{}
	err = r.clients[r.index].rpcClient.Call(&res, "debug_traceTransaction", txhash, &TracerConfig{Tracer: "callTracer"})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *manage) GetTracerLog(ctx context.Context, txHash common.Hash) (*types.ExecutionResult, error) {
	var err error
	var res = types.ExecutionResult{}
	err = r.clients[r.index].rpcClient.Call(&res, "debug_traceTransaction", txHash, &TracerConfig{})
	if err != nil {
		return nil, err
	}
	return &res, nil
}
