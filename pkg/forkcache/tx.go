package forkcache

import (
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

var txs = make(map[common.Hash]*types.Tx)
