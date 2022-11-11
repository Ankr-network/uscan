package job

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var testHash = "0x3ae4bb27f9e98cdf8d7eb831c9407d8001d754ed5c4b501ec9e9d1624f4b8866"

func TestDebugJob(t *testing.T) {
	txhash := common.HexToHash(testHash)
	job := NewSyncDebugJob(txhash, testRpc)
	job.Execute()
}
