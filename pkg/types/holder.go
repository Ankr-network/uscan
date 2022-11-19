package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
)

type Holder struct {
	Addr     common.Address
	Quantity field.BigInt
}

func ByteToHolder(bin []byte) (*Holder, error) {
	if len(bin) != 52 {
		return nil, ErrorInvalidByte
	}
	h := &Holder{}
	h.Quantity.SetBytes(bin[:32])
	h.Addr.SetBytes(bin[32:])
	return h, nil
}

func (h Holder) ToBytes() []byte {
	return append(common.BytesToHash(h.Quantity.Bytes()).Bytes(), h.Addr.Bytes()...)
}

type Inventory struct {
	Addr    common.Address
	TokenID field.BigInt
}

func ByteToInventory(bin []byte) (*Inventory, error) {
	if len(bin) != 52 {
		return nil, ErrorInvalidByte
	}
	h := &Inventory{}
	h.TokenID.SetBytes(bin[:32])
	h.Addr.SetBytes(bin[32:])
	return h, nil
}

func (h Inventory) ToBytes() []byte {
	return append(common.BytesToHash(h.TokenID.Bytes()).Bytes(), h.Addr.Bytes()...)
}
