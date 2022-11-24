package rpcclient

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/assert"
)

var testClient RpcClient

func TestMain(m *testing.M) {
	// testClient = NewRpcClient("ws://103.23.44.29:28546")
	testClient = NewRpcClient([]string{"wss://testnet.ankr.com/ws"}, 0)

	m.Run()
}

func TestBlockNumber(t *testing.T) {
	numberChan := testClient.GetLatestBlockNumber(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	for {
		select {
		case num := <-numberChan:
			t.Log(num)
		case <-ctx.Done():
			return
		}
	}
}

func TestGetBlockByNumber(t *testing.T) {
	bk, err := testClient.GetBlockByNumber(context.Background(), (*hexutil.Big)(big.NewInt(int64(1584363))).String())
	assert.NoError(t, err)
	bytesRes, err := json.Marshal(bk)
	assert.NoError(t, err)
	t.Log(string(bytesRes))

}

func TestGetTransactionByHash(t *testing.T) {
	hashes := []common.Hash{
		common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d"),
		common.HexToHash("0x8c1420f491679d36cb09912774e6a47fa6bfbd8bb225a33e151d0e1522e7b5a2"),
	}
	rts, err := testClient.GetTransactionsByHash(context.Background(), hashes)
	assert.NoError(t, err)
	for i, v := range rts {
		bytesRes, err := json.Marshal(v)
		assert.NoError(t, err)
		assert.Equal(t, hashes[i], v.Hash)
		t.Log(string(bytesRes))
	}
}

func TestGetTransactionReceipt(t *testing.T) {
	hashes := []common.Hash{
		common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d"),
		common.HexToHash("0x8c1420f491679d36cb09912774e6a47fa6bfbd8bb225a33e151d0e1522e7b5a2"),
	}
	rts, err := testClient.GetTransactionReceiptsByHash(context.Background(), hashes)
	assert.NoError(t, err)
	for i, v := range rts {
		bytesRes, err := json.Marshal(v)
		assert.NoError(t, err)
		assert.Equal(t, hashes[i], v.TxHash)
		t.Log(string(bytesRes))
	}
}

func TestGetCode(t *testing.T) {

	code, err := testClient.GetCode(context.Background(), common.HexToAddress("0xd8a4ff865c120e287e0953551462e0d6084b9f04"), "latest")
	assert.NoError(t, err)
	t.Log(code)
}

func TestGetBalance(t *testing.T) {
	bal, err := testClient.GetBalance(context.Background(), common.HexToAddress("0x20cD8eB93c50BDAc35d6A526f499c0104958e3F6"), "latest")
	assert.NoError(t, err)
	t.Log(bal.String())
}

func TestGetBalances(t *testing.T) {
	data, err := testClient.GetBalances(
		context.Background(),
		[]common.Address{
			common.HexToAddress("0x20cD8eB93c50BDAc35d6A526f499c0104958e3F6"),
			common.HexToAddress("0x5fd8b97F8D8DA84813F583C42d40D1e5A4DA9A17"),
			common.HexToAddress("0x6A82Fa5Cf82Ef724d1F1955fECc27DDd0758132E"),
		},
		hexutil.EncodeUint64(2659050),
	)
	assert.NoError(t, err)

	for k, v := range data {
		t.Log(v)
		t.Log(k, " : ", v.String())
	}
}

func TestGetTracerCall(t *testing.T) {
	cf, err := testClient.GetTracerCall(context.Background(), common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d"))
	assert.NoError(t, err)
	t.Log(cf.JsonToString())
}
func TestGetTracerLog(t *testing.T) {
	cf, err := testClient.GetTracerLog(context.Background(), common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d"))
	assert.NoError(t, err)
	t.Log(cf.JsonToString())
}
