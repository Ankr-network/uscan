package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Erc721Transfer struct {
	TransactionHash common.Hash
	BlockNumber     *field.BigInt
	Contract        common.Address
	Method          []byte
	From            common.Address
	To              common.Address
	TokenId         field.BigInt
}

func (b *Erc721Transfer) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Erc721Transfer) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
