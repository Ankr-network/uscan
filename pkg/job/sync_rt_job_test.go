package job

import (
	"encoding/json"
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestRtJob(t *testing.T) {
	tx := common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d")
	rtJob := NewSyncRtJob(tx, testRpc)
	rtJob.Execute()

	assert.Equal(t, tx, rtJob.ReceiptData.TxHash)
	assert.Equal(t, rtJob.ReceiptData.Type, field.NewInt(0))
	assert.Nil(t, rtJob.ReceiptData.PostState)
	assert.Equal(t, rtJob.ReceiptData.Status, field.NewInt(1))
	assert.Equal(t, rtJob.ReceiptData.CumulativeGasUsed, field.NewInt(12736))
	assert.Nil(t, rtJob.ReceiptData.ContractAddress)
	assert.Equal(t, rtJob.ReceiptData.GasUsed, field.NewInt(12736))
	assert.Equal(t, rtJob.ReceiptData.EffectiveGasPrice, field.NewInt(1500))

	bytesRes, err := json.Marshal(rtJob.ReceiptData)
	assert.NoError(t, err)
	t.Log(string(bytesRes))
}
