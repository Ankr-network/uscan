package job

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTracerJob(t *testing.T) {
	hash := common.HexToHash("0x077dac720003821e7c6ef716cce0b7c312ba72e0fe574ccb0edbf598e07d8ac7")
	job := NewSyncTracerJob(1691329, hash, testRpc)
	job.Execute()

	assert.True(t, job.Status)
	assert.Equal(t, job.Error, "")
	t.Log(job.CallFrame.JsonToString())
	t.Log(job.ContractOrMemberData)
	for i, v := range job.InternalTxs {
		t.Log("internaltx(", i, "): ", v)
	}

	for k, v := range job.ContractOrMemberData {
		t.Log(k.Hex(), ": ", v)
	}
}
