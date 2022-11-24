package job

import (
	"encoding/json"
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestTxJob(t *testing.T) {
	hash := common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d")
	txJob := NewSyncTxJob(hash, testRpc)
	txJob.Execute()

	to := common.HexToAddress("0x07861819f3d9773088f67e5572bd645b2e5c15ef")
	assert.Equal(t, txJob.TransactionData.Hash, hash)
	assert.Equal(t, txJob.TransactionData.BlockNum, field.NewInt(2090368))
	assert.Nil(t, txJob.TransactionData.GasTipCap)
	assert.Nil(t, txJob.TransactionData.GasFeeCap)
	assert.Equal(t, txJob.TransactionData.GasPrice, field.NewInt(1500))
	assert.Equal(t, txJob.TransactionData.Nonce, field.NewInt(95))
	assert.Equal(t, txJob.TransactionData.Gas, field.NewInt(29836))
	assert.Equal(t, txJob.TransactionData.From, common.HexToAddress("0x3c10ec535d1a8cba60536a963cc62a1df855e71c"))
	assert.Equal(t, txJob.TransactionData.To, &to)
	assert.Equal(t, txJob.TransactionData.Value, field.NewInt(0))
	assert.Equal(t, txJob.TransactionData.Data, hexutil.Bytes(hexutil.MustDecode("0xa9059cbb0000000000000000000000008ec529c63f174996c5cf360081d94bac07a8615e00000000000000000000000000000000000000000000000000000002540be400")))

	bytesRes, err := json.Marshal(txJob.TransactionData)
	assert.NoError(t, err)
	t.Log(string(bytesRes))
}
