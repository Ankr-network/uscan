package rawdb

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	txKey      = "/tx/"
	rtKey      = "/rt/"
	txNumKey   = "/all/tx/total"
	txIndexKey = "/all/tx/"
)

// /tx/<txhash>          => tx info
// /rt/<txhash>          => rt info
// /all/tx/total => total tx number bigInt转byte
// /all/tx/<index> => <txhash>

func getByteKey(keyPre string, txHash common.Hash) []byte {
	tx := []byte(keyPre)
	tx = append(tx, txHash.Bytes()...)
	return tx
}

func getTxIndex(num *big.Int) []byte {
	tx := []byte(txIndexKey)
	tx = append(tx, num.Bytes()...)
	return tx
}

func GetTxNum(db kv.Getter) (*big.Int, error) {
	data, err := db.Get([]byte(txNumKey), &kv.ReadOption{
		Table: share.TxTbl,
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

func GetTx(db kv.Getter, txHash common.Hash) (*types.Tx, error) {
	data, err := db.Get(getByteKey(txKey, txHash), &kv.ReadOption{
		Table: share.TxTbl,
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &types.Tx{}
	if err := res.Unmarshal(data); err != nil {
		return nil, err
	}
	return res, nil
}

func GetRt(db kv.Getter, txHash common.Hash) (*types.Rt, error) {
	data, err := db.Get(getByteKey(rtKey, txHash), &kv.ReadOption{
		Table: share.TxTbl,
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &types.Rt{}
	if err := res.Unmarshal(data); err != nil {
		return nil, err
	}
	return res, nil
}

func GetTxIndex(db kv.Getter, index *big.Int) (*common.Hash, error) {
	data, err := db.Get(getTxIndex(index), &kv.ReadOption{
		Table: share.TxTbl,
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := common.Hash{}
	res.SetBytes(data)
	return &res, nil
}
