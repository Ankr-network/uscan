package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type InternalTx struct {
	TransactionHash common.Hash
	BlockNumber     *big.Int
	Status          bool
	CallType        string
	Depth           string
	From            common.Address
	To              *common.Address
	Amount          *big.Int
	GasLimit        *big.Int
}

func (b *InternalTx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *InternalTx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
