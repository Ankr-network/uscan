package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestHolder(t *testing.T) {
	h := &Holder{
		Addr:     common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc"),
		Quantity: *field.NewInt(1111),
	}
	bytesRes := h.ToBytes()
	t.Log(bytesRes)
	t.Log(len(bytesRes))

	out, err := ByteToHolder(bytesRes)
	assert.NoError(t, err)
	assert.Equal(t, h.Addr, out.Addr)
	assert.Equal(t, h.Quantity, out.Quantity)
}
