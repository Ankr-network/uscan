package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type BkSim struct {
	Number            field.BigInt   `json:"number"`
	Timestamp         int            `json:"timestamp"`
	Miner             common.Address `json:"miner"`
	GasUsed           field.BigInt   `json:"gas_used"`
	TransactionsTotal int            `json:"transactions_total"`
}

type TxSim struct {
	Hash     common.Hash    `json:"hash"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	GasPrice *field.BigInt  `json:"gas_price"`
	Gas      *field.BigInt  `json:"gas"`
}

type Home struct {
	BlockNumber  *field.BigInt `rlp:"-"`
	TxTotal      *field.BigInt
	AddressTotal *field.BigInt
	Blocks       []*BkSim
	Txs          []*TxSim
	DateTxs      map[string]*field.BigInt // example : 20221023 => 0x2
}

func (b *Home) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Home) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
