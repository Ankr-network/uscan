package job

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTracerJob(t *testing.T) {
	hash := common.HexToHash("0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d")
	job := NewSyncTracerJob(1691329, hash, testRpc)
	job.Execute()

	assert.True(t, job.Status)
	assert.Equal(t, job.Status, "")
	t.Log(job.CallFrame.JsonToString())
	t.Log(job.ContractOrMemberData)
	for i, v := range job.InternalTxs {
		t.Log("internaltx(", i, "): ", v)
	}
}
