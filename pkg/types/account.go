package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	Erc20Flag   byte = 0x01
	Erc721Flag  byte = 0x02
	Erc1155Flag byte = 0x04
)

type Account struct {
	BlockNumber *big.Int
	Owner       common.Address `rlp:"-"`
	Balance     *big.Int

	Erc20            bool `rlp:"-"`
	Erc721           bool `rlp:"-"`
	Erc1155          bool `rlp:"-"`
	ErcFlag          byte
	Creator          *common.Address
	TxHash           *common.Hash
	Code             []byte
	Name, Symbol     string
	TokenTotalSupply *big.Int
	NftTotalSupply   *big.Int
	Decimals         *big.Int // erc20 decimals
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
