package job

import (
	"context"
	"testing"

	"github.com/Ankr-network/uscan/pkg/kv/memorydb"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var testHash = "0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d"

func TestDebugJob(t *testing.T) {
	txhash := common.HexToHash(testHash)
	cache := memorydb.NewMemoryDb()
	job := NewSyncDebugJob(txhash, testRpc, cache)
	job.Execute()

	res, err := rawdb.ReadTraceTx(context.Background(), cache, txhash)
	assert.NoError(t, err)

	t.Log(res)
}
