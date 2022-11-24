package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	Erc20Flag   byte = 0x01
	Erc721Flag  byte = 0x02
	Erc1155Flag byte = 0x04
)

type Account struct {
	Owner common.Address `rlp:"-"`

	Erc20            bool `rlp:"-"`
	Erc721           bool `rlp:"-"`
	Erc1155          bool `rlp:"-"`
	ErcFlag          byte
	BlockNumber      field.BigInt
	Balance          field.BigInt
	Name, Symbol     string
	TokenTotalSupply field.BigInt
	NftTotalSupply   field.BigInt
	Decimals         field.BigInt // erc20 decimals

	Creator common.Address
	TxHash  common.Hash

	Retry field.BigInt
}

func (b *Account) Marshal() ([]byte, error) {
	b.ErcFlag = 0x0
	if b.Erc20 {
		b.ErcFlag += Erc20Flag
	}

	if b.Erc721 {
		b.ErcFlag += Erc721Flag
	}

	if b.Erc1155 {
		b.ErcFlag += Erc1155Flag
	}
	return rlp.EncodeToBytes(b)
}

func (b *Account) Unmarshal(bin []byte) (err error) {
	err = rlp.DecodeBytes(bin, &b)
	if err != nil {
		return err
	}

	if b.ErcFlag >= Erc1155Flag {
		b.ErcFlag -= Erc1155Flag
		b.Erc1155 = true
	}

	if b.ErcFlag >= Erc721Flag {
		b.ErcFlag -= Erc721Flag
		b.Erc721 = true
	}

	if b.ErcFlag >= Erc20Flag {
		b.ErcFlag -= Erc20Flag
		b.Erc20 = true
	}
	return nil
}
