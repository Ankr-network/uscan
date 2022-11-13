package field

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

var ErrInvalidBigInt = errors.New("invalid BigInt")

type BigInt big.Int

func NewInt(x int64) *BigInt {
	return (*BigInt)(big.NewInt(x))
}

func (b *BigInt) String() string {
	if b == nil {
		return ""
	}
	return hexutil.EncodeBig((*big.Int)(b))
}

func (b *BigInt) StringPointer() *string {
	if b == nil {
		return nil
	}
	res := hexutil.EncodeBig((*big.Int)(b))
	return &res
}

func (b *BigInt) Bytes() []byte {
	return b.Bytes()
}

func (b *BigInt) Add(num *BigInt) {
	(*big.Int)(b).Add((*big.Int)(b), (*big.Int)(num))
}

func (b *BigInt) SetBytes(bin []byte) {
	(*big.Int)(b).SetBytes(bin)
}

func (b *BigInt) ToUint64() uint64 {
	return (*big.Int)(b).Uint64()
}

func (b *BigInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", hexutil.EncodeBig((*big.Int)(b)))), nil
}

func (b *BigInt) UnmarshalJSON(bs []byte) error {
	input := string(bytes.Trim(bs, "\""))
	if input == "0x" {
		return nil
	}
	bi, err := hexutil.DecodeBig(input)
	if err != nil {
		return err
	}
	(*big.Int)(b).SetBytes(bi.Bytes())

	return nil
}

func (b *BigInt) DecodeRLP(s *rlp.Stream) (err error) {
	res := &big.Int{}
	err = s.Decode(res)
	if err == nil {
		(*big.Int)(b).SetBytes(res.Bytes())
	}
	return err
}

func (b *BigInt) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, (*big.Int)(b))
}
