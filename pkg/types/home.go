package types

import (
	"encoding/json"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type BkSim struct {
	Number            field.BigInt   `json:"number"`
	Timestamp         field.BigInt   `json:"timestamp"`
	Miner             common.Address `json:"miner"`
	GasUsed           field.BigInt   `json:"gas_used"`
	TransactionsTotal field.BigInt   `json:"transactions_total"`
}

type TxSim struct {
	Hash     common.Hash    `json:"hash"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	GasPrice field.BigInt   `json:"gas_price"`
	Gas      field.BigInt   `json:"gas"`
}

type Home struct {
	BlockNumber  field.BigInt `rlp:"-"`
	TxTotal      field.BigInt
	AddressTotal field.BigInt
	Erc20Total   field.BigInt
	Erc721Total  field.BigInt
	Erc1155Total field.BigInt
	Blocks       []*BkSim
	Txs          []*TxSim
	DateTxs      map[string]*field.BigInt `rlp:"-"` // example : 20221023 => 0x2
	DateTxsByte  []byte
}

func (b *Home) Marshal() ([]byte, error) {
	b.DateTxsByte, _ = json.Marshal(b.DateTxs)
	return rlp.EncodeToBytes(b)
}

func (b *Home) Unmarshal(bin []byte) (err error) {

	err = rlp.DecodeBytes(bin, &b)
	if err == nil {
		json.Unmarshal(b.DateTxsByte, &b.DateTxs)
		b.DateTxsByte = []byte{}
	}
	return
}
