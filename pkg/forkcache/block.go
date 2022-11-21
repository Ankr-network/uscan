package forkcache

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/types"
)

var blocks = make(map[*field.BigInt]types.Block)
