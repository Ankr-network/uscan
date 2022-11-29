package core

import (
	"context"
	"errors"
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/forkcache"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

func (n *blockHandle) checkForkNewAddr(ctx context.Context) (*field.BigInt, error) {
	var newAddrTotal = field.NewInt(0)
	for k, v := range n.contractOrMemberData {
		account, err := forkcache.ReadAccount(ctx, n.db, k)
		if err != nil {
			if !errors.Is(err, kv.NotFound) {
				log.Errorf("read fork account(%s): %v", k.Hex(), err)
				return nil, err
			}
			newAddrTotal.Add(field.NewInt(1))
			account = &types.Account{}
		}

		n.contractOrMemberData[k] = n.mergeForkAccount(account, v)
	}

	return newAddrTotal, nil
}

func (n *blockHandle) mergeForkAccount(beforeAcc *types.Account, afterAcc *types.Account) *types.Account {
	if beforeAcc.BlockNumber.String() == "0x0" {
		beforeAcc.BlockNumber = afterAcc.BlockNumber
	}
	beforeAcc.Balance = afterAcc.Balance
	if afterAcc.Erc20 {
		beforeAcc.Erc20 = true
	}

	if afterAcc.Erc721 {
		beforeAcc.Erc721 = true
	}
	if afterAcc.Erc1155 {
		beforeAcc.Erc1155 = true
	}
	if beforeAcc.Creator == (common.Address{}) {
		beforeAcc.Creator = afterAcc.Creator
	}

	if beforeAcc.TxHash == (common.Hash{}) {
		beforeAcc.TxHash = afterAcc.TxHash
	}

	if beforeAcc.Name == "" {
		beforeAcc.Name = afterAcc.Name
	}

	if beforeAcc.Symbol == "" {
		beforeAcc.Symbol = afterAcc.Symbol
	}
	if afterAcc.TokenTotalSupply.String() != "0x0" {
		beforeAcc.TokenTotalSupply = afterAcc.TokenTotalSupply
	}
	if afterAcc.NftTotalSupply.String() != "0x0" {
		beforeAcc.NftTotalSupply = afterAcc.NftTotalSupply
	}
	return beforeAcc
}

func (n *blockHandle) readForkAccount(ctx context.Context, addr common.Address) (*types.Account, error) {
	acc, ok := n.contractOrMemberData[addr]
	if ok {
		return acc, nil
	}
	account, err := forkcache.ReadAccount(ctx, n.db, addr)
	if err != nil {
		if !errors.Is(err, kv.NotFound) {
			log.Errorf("read fork account(%s): %v", addr.Hex(), err)
			return nil, err
		}
		account = &types.Account{}
	}
	n.contractOrMemberData[addr] = account
	return account, nil
}

func (n *blockHandle) updateForkAccounts(ctx context.Context) (err error) {
	for k, v := range n.contractOrMemberData {
		if err = forkcache.WriteAccount(ctx, n.db, k, v); err != nil {
			log.Errorf("write fork account(%s): %v", k.Hex(), err)
			return err
		}
	}
	return nil
}
