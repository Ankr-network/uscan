package fulldb

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	contractKeyPrefix   = []byte("/contract/")
	proxyContractPrefix = []byte("/proxy/")
)

/*
/contract/<address> => contract info

/proxy/<proxy contract> => logic address
*/

// ----------------- contract info -----------------
func ReadContract(ctx context.Context, db kv.Reader, addr common.Address) (acc *types.Contract, err error) {
	var bytesRes []byte

	bytesRes, err = db.Get(ctx, append(contractKeyPrefix, addr.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.Contract{}
	err = acc.Unmarshal(bytesRes)
	if err == nil {
		acc.Owner = addr
	}
	return
}

func WriteContract(ctx context.Context, db kv.Writer, addr common.Address, data *types.Contract) error {
	bytesRes, err := data.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, append(contractKeyPrefix, addr.Bytes()...), bytesRes, &kv.WriteOption{Table: share.AccountsTbl})
}

// -------------------- proxy contract -------------
func WriteProxyContract(ctx context.Context, db kv.Writer, proxy common.Address, logic common.Address) error {
	return db.Put(ctx, append(proxyContractPrefix, proxy.Bytes()...), logic.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadProxyContract(ctx context.Context, db kv.Reader, proxy common.Address) (logic common.Address, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(proxyContractPrefix, proxy.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err == nil {
		logic.SetBytes(bytesRes)
	}
	return
}
