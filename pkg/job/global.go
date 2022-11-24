package job

import "github.com/Ankr-network/uscan/pkg/workpool"

var (
	DebugJobChan workpool.Dispathcher
	TxJobChan    workpool.Dispathcher
)

func GlobalInit(work int) {
	DebugJobChan = workpool.NewDispathcher(work)
	TxJobChan = workpool.NewDispathcher(work * 3)
}
