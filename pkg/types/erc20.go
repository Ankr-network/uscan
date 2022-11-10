package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Erc20Transfer struct {
	TransactionHash common.Hash
	BlockNumber     *big.Int
	Contract        common.Address
	Method          string
	From            common.Address
	To              common.Address
	Amount          big.Int
}

func (b *Erc20Transfer) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Erc20Transfer) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
