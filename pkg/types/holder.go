package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
)

type Holder struct {
	Addr   common.Address
	Amount field.BigInt
}

func ByteToHolder(bin []byte) (*Holder, error) {
	if len(bin) != 52 {
		return nil, ErrorInvalidByte
	}
	h := &Holder{}
	h.Amount.SetBytes(bin[:32])
	h.Addr.SetBytes(bin[32:])
	return h, nil
}

func (h Holder) ToBytes() []byte {
	return append(common.BytesToHash(h.Amount.Bytes()).Bytes(), h.Addr.Bytes()...)
}
