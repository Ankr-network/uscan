package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Contract struct {
	Owner                 common.Address `rlp:"-"`
	ByteCode              []byte
	ByteCodeHash          common.Hash
	ConstructorArguements []byte
	DeployedCode          []byte
}

func (b *Contract) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Contract) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
