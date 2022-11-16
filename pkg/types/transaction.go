package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

type Tx struct {
	Hash      common.Hash     `json:"hash"  rlp:"-"`
	BlockNum  *field.BigInt   `json:"blockNumber"`
	GasTipCap *field.BigInt   `json:"maxPriorityFeePerGas"`
	GasFeeCap *field.BigInt   `json:"maxFeePerGas"`
	GasPrice  *field.BigInt   `json:"gasPrice"`
	Nonce     *field.BigInt   `json:"nonce"`
	Gas       *field.BigInt   `json:"gas"`
	From      common.Address  `json:"from"`
	To        *common.Address `json:"to"` // nil means contract creation
	Value     *field.BigInt   `json:"value"`
	Method    hexutil.Bytes   `json:"-"`
	Data      hexutil.Bytes   `json:"input"`
	TimeStamp *field.BigInt   `json:"-"`

	// Signature values
	V *field.BigInt `json:"v"`
	R *field.BigInt `json:"r"`
	S *field.BigInt `json:"s"`
}

func (b *Tx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Tx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
