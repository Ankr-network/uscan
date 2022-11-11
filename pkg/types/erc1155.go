package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Erc1155Transfer struct {
	TransactionHash common.Hash
	BlockNumber     *big.Int
	Contract        common.Address
	Method          string
	From            common.Address
	To              common.Address
	TokenID         big.Int
	Quantity        big.Int
}

func (b *Erc1155Transfer) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Erc1155Transfer) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
