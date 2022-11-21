package core

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rawdb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

func (n *blockHandle) writeContract(ctx context.Context, data map[common.Address]*types.Contract) (err error) {
	for k, v := range data {
		if err = rawdb.WriteContract(ctx, n.db, k, v); err != nil {
			log.Errorf("write contract(%s): %v ", k, err)
			return err
		}
	}
	return nil
}

func (n *blockHandle) writeProxyContract(ctx context.Context, data map[common.Address]common.Address) (err error) {
	for k, v := range data {
		if err = rawdb.WriteProxyContract(ctx, n.db, k, v); err != nil {
			log.Errorf("write proxy contract(%s): %v ", k, err)
			return err
		}
	}
	return nil
}
