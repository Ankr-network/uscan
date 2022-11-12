package service

import (
	store "github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

func ListTxs(pager *types.Pager) error {
	num, err := store.GetTxNum(nil)
	if err != nil {
		return err
	}
	begin, end := ParseBlockPage(num, pager.Offset, pager.Limit)

	// TODO get tx from begin to end :/all/tx/<index> => <txhash>  /tx/<txhash> /rt/<txhash>

	return nil
}

func GetTx(tx string) error {
	txHash := common.HexToHash(tx)
	txs, err := store.GetTx(nil, txHash)
	if err != nil {
		return err
	}
	rts, err := store.GetRt(nil, txHash)
	if err != nil {
		return err
	}

	return nil
}
