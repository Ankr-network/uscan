package forkcache

import "github.com/Ankr-network/uscan/pkg/types"

var (
	erc20Transfer   = make(map[string]*types.Erc20Transfer)
	erc721Transfer  = make(map[string]*types.Erc721Transfer)
	erc1155Transfer = make(map[string]*types.Erc1155Transfer)
)
