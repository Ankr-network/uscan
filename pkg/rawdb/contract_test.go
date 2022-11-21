package rawdb

import (
	"context"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/pkg/types"
	"testing"
)

func TestWriteValidateContractMetadata(t *testing.T) {
	db := mdbx.NewMdbx("/Users/johnson/goWork/Ankr-network/uscan/uscandb")
	data := &types.ValidateContractMetadata{
		CompilerVersions: []*types.CompilerVersion{{Name: "v0.4.11+commit.68ef5810", FileName: "solc-static-linux-v0.4.11"}},
		LicenseTypes:     []*types.LicenseType{},
		CompilerTypes:    []string{},
	}
	if err := WriteValidateContractMetadata(context.Background(), db, data); err != nil {
		t.Errorf("err: %s", err)
	}
	acc, err := ReadValidateContractMetadata(context.Background(), db)
	if err != nil {
		t.Errorf("err: %s", err)
	}
	t.Log(data)
	t.Log(acc)
}
