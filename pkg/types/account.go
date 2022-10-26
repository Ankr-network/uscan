package types

import (
	"github.com/Ankr-network/uscan/pkg/utils"
)

const (
	Erc20Flag   byte = 0x80
	Erc721Flag  byte = 0x40
	Erc1155Flag byte = 0x20
)

type Account struct {
	Erc20   bool
	Erc721  bool
	Erc1155 bool
	Balance uint64
	Code    []byte
}

func (a *Account) Marshal() ([]byte, error) {
	var flag byte
	switch {
	case a.Erc1155:
		flag |= Erc1155Flag
	case a.Erc20:
		flag |= Erc20Flag
	case a.Erc721:
		flag |= Erc721Flag
	}
	bs := make([]byte, 0, 2)
	bs = append(bs, flag)
	bs = append(bs, utils.WrapLen(utils.EncodeVarint(a.Balance))...)
	bs = append(bs, a.Code...)
	return bs, nil
}

func (a *Account) Unmarshal(bin []byte) error {

	if a == nil {
		a = &Account{}
	}

	if bin[0]&Erc20Flag == Erc20Flag {
		a.Erc20 = true
	}
	if bin[0]&Erc721Flag == Erc721Flag {
		a.Erc721 = true
	}
	if bin[0]&Erc1155Flag == Erc1155Flag {
		a.Erc1155 = true
	}

	balLen := int(bin[1])
	a.Balance, _ = utils.DecodeVarint(bin[2 : 2+balLen])

	a.Code = make([]byte, len(bin)-2-balLen)
	copy(a.Code, bin[2+balLen:])

	return nil
}
