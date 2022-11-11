package job

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/rpcclient"
)

var (
	testRpc rpcclient.RpcClient
)

func TestMain(m *testing.M) {
	testRpc = rpcclient.NewRpcClient([]string{"wss://testnet.ankr.com/ws"})
	m.Run()
}
