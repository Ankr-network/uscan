package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/stretchr/testify/assert"
)

func TestHomeMarshal(t *testing.T) {
	h := Home{
		AddressTotal: *field.NewInt(11),
		DateTxs:      make(map[string]*field.BigInt),
	}
	res, err := h.Marshal()
	assert.NoError(t, err)
	t.Log(res)

	out := &Home{}
	err = out.Unmarshal(res)
	assert.NoError(t, err)
	t.Log(out)
}
