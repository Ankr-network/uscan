package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type TraceTx struct {
	Res    string
	LogNum int
}

func (b *TraceTx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *TraceTx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

func (b *TraceTx) UnmarshalJSON(bin []byte) error {
	return nil
}

type TraceTx2 struct {
	Res string
}

func (b *TraceTx2) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *TraceTx2) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

func (b *TraceTx2) UnmarshalJSON(bin []byte) error {
	return nil
}

type CallFrame struct {
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to,omitempty"`
	Value   *big.Int       `json:"value,omitempty"`
	Gas     *big.Int       `json:"gas"`
	GasUsed *big.Int       `json:"gasUsed"`
	Input   []byte         `json:"input"`
	Output  []byte         `json:"output,omitempty"`
	Error   string         `json:"error,omitempty"`
	Calls   []CallFrame    `json:"calls,omitempty"`
}
