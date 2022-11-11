package rawdb

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
)

const (
	blockNumKey = "/syncing/block"
)

func GetBlockNum(db kv.Getter) (string, error) {
	data, err := db.Get([]byte(blockNumKey), &kv.ReadOption{
		Table: "config",
	})
	if err != nil {
		return "", err
	}
	if data != nil {
		return string(data), nil
	}
	return "", nil
}

func GetBlock(db kv.Getter, blockNum string) (*types.Block, error) {
	data, err := db.Get([]byte(blockNum), &kv.ReadOption{
		Table: "block",
	})
	if err != nil {
		return nil, err
	}
	res := &types.Block{}
	if err := res.Unmarshal(data); err != nil {
		return nil, err
	}
	return res, nil
}
