package rawdb

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	contractMetadataPrefix  = []byte("metadata")
	ContractVerityPrefix    = []byte("contract/info/")
	ContractVerityTmpPrefix = []byte("contract/tmp/")
	ContractMethodPrefix    = []byte("method/")
)

/*
metadata => types.ValidateContractMetadata
contract/tmp/<address>  ->  status
contract/info/<address> -> info

method/<MethodID> -> method name
*/

func WriteValidateContractMetadata(ctx context.Context, db kv.Writer, data *types.ValidateContractMetadata) error {
	bytesRes, err := data.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, contractMetadataPrefix, bytesRes, &kv.WriteOption{Table: share.ValidateContractTbl})
}

func ReadValidateContractMetadata(ctx context.Context, db kv.Reader) (acc *types.ValidateContractMetadata, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, contractMetadataPrefix, &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.ValidateContractMetadata{}
	err = acc.Unmarshal(bytesRes)
	return
}

func WriteValidateContractStatus(ctx context.Context, db kv.Writer, address common.Address, status *big.Int) error {
	return db.Put(ctx, append(ContractVerityTmpPrefix, address.Bytes()...), status.Bytes(), &kv.WriteOption{Table: share.ValidateContractTbl})
}

func ReadValidateContractStatus(ctx context.Context, db kv.Reader, address common.Address) (status *big.Int, err error) {
	rs, err := db.Get(ctx, append(ContractVerityTmpPrefix, address.Bytes()...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	status = &big.Int{}
	status.SetBytes(rs)
	return
}

func WriteValidateContract(ctx context.Context, db kv.Writer, address common.Address, data *types.ContractVerity) error {
	bytesRes, err := data.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, append(ContractVerityPrefix, address.Bytes()...), bytesRes, &kv.WriteOption{Table: share.ValidateContractTbl})
}

func ReadValidateContract(ctx context.Context, db kv.Reader, address common.Address) (data *types.ContractVerity, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(ContractVerityPrefix, address.Bytes()...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return nil, err
	}
	data = &types.ContractVerity{}
	err = data.Unmarshal(bytesRes)
	return
}

func WriteMethodName(ctx context.Context, db kv.Writer, methodID, methodName string) error {
	return db.Put(ctx, append(ContractMethodPrefix, []byte(methodID)...), []byte(methodName), &kv.WriteOption{Table: share.ValidateContractTbl})
}

func ReadMethodName(ctx context.Context, db kv.Reader, methodID string) (data string, err error) {
	rs, err := db.Get(ctx, append(ContractMethodPrefix, []byte(methodID)...), &kv.ReadOption{Table: share.ValidateContractTbl})
	if err != nil {
		return "", err
	}
	data = string(rs)
	return
}
