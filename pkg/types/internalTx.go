package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type InternalTx struct {
	TransactionHash common.Hash
	BlockNumber     *field.BigInt
	Status          bool
	CallType        string
	Depth           string
	From            common.Address
	To              *common.Address
	Amount          *field.BigInt
	GasLimit        *field.BigInt
}

func (b *InternalTx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *InternalTx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
