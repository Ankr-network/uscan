package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Erc721Transfer struct {
	TransactionHash common.Hash
	BlockNumber     *big.Int
	Contract        common.Address
	Method          string
	From            common.Address
	To              common.Address
	TokenId         big.Int
}

func (b *Erc721Transfer) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Erc721Transfer) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
