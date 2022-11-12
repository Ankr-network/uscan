package rawdb

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"math/big"
)

const (
	blockNumKey = "/syncing/block"
)

func GetBlockNum(db kv.Getter) (*big.Int, error) {
	data, err := db.Get([]byte(blockNumKey), &kv.ReadOption{
		Table: share.ConfigTbl,
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	num := (&big.Int{}).SetBytes(data)
	return num, nil
}

func GetBlock(db kv.Getter, blockNum string) (*types.Block, error) {
	data, err := db.Get([]byte(blockNum), &kv.ReadOption{
		Table: share.BlockTbl,
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &types.Block{}
	if err := res.Unmarshal(data); err != nil {
		return nil, err
	}
	return res, nil
}
