package service

import (
	"github.com/Ankr-network/uscan/pkg/types"
)

func ListTxs(pager *types.Pager) error {
	// TODO get txs /all/tx/total
	//num := 0

	//begin, end := ParsePage(num, pager.Offset, pager.Limit)
	//NumToHex(fmt.Sprint(begin))
	//NumToHex(fmt.Sprint(end))

	// TODO get tx from begin to end :/all/tx/<index> => <txhash>  /tx/<txhash> /rt/<txhash>

	return nil
}

func GetTx(tx string) error {
	// TODO get /tx/<txhash> /rt/<txhash>
	return nil
}
