package core

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/storage/fulldb"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

type contractHandle struct {
	contractInfoMap map[common.Address]*types.Contract
	proxyContracts  map[common.Address]common.Address
	db              kv.Database
}

func newContractHandle(contractInfoMap map[common.Address]*types.Contract,
	proxyContracts map[common.Address]common.Address,
	db kv.Database) *contractHandle {
	return &contractHandle{
		contractInfoMap: contractInfoMap,
		proxyContracts:  proxyContracts,
		db:              db,
	}
}

func (n *contractHandle) handleContractData(ctx context.Context) (err error) {
	if len(n.contractInfoMap) > 0 {
		if err = n.writeContract(ctx); err != nil {
			log.Errorf("write contract: %v", err)
			return err
		}
	}
	if len(n.proxyContracts) > 0 {
		if err = n.writeProxyContract(ctx); err != nil {
			log.Errorf("write proxy contract: %v", err)
			return err
		}
	}
	return nil
}

func (n *contractHandle) writeContract(ctx context.Context) (err error) {
	for k, v := range n.contractInfoMap {
		if err = fulldb.WriteContract(ctx, n.db, k, v); err != nil {
			log.Errorf("write contract(%s): %v ", k, err)
			return err
		}
	}
	return nil
}

func (n *contractHandle) writeProxyContract(ctx context.Context) (err error) {
	for k, v := range n.proxyContracts {
		if err = fulldb.WriteProxyContract(ctx, n.db, k, v); err != nil {
			log.Errorf("write proxy contract(%s): %v ", k, err)
			return err
		}
	}
	return nil
}
