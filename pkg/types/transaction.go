package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Tx struct {
	Hash      common.Hash     `json:"hash"  rlp:"-"`
	BlockNum  *big.Int        `json:"blockNumber"`
	GasTipCap *big.Int        `json:"maxPriorityFeePerGas"`
	GasFeeCap *big.Int        `json:"maxFeePerGas"`
	GasPrice  *big.Int        `json:"gasPrice"`
	Nonce     *big.Int        `json:"nonce"`
	Gas       *big.Int        `json:"gas"`
	From      common.Address  `json:"from"`
	To        *common.Address `json:"to"` // nil means contract creation
	Value     *big.Int        `json:"value"`
	Data      []byte          `json:"input"`

	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (b *Tx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Tx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

func (b *Tx) UnmarshalJSON(bin []byte) error {
	return nil
}
