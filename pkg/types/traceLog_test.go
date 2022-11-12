package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/stretchr/testify/assert"
)

func TestTraceTx(t *testing.T) {
	traceTx := TraceTx{
		Res:    "dsadadsada",
		LogNum: field.NewInt(1111),
	}

	val, err := traceTx.Marshal()
	assert.NoError(t, err)

	tt := &TraceTx{}
	err = tt.Unmarshal(val)
	assert.NoError(t, err)
	assert.Equal(t, traceTx.Res, tt.Res)
	assert.Equal(t, traceTx.LogNum, tt.LogNum)
}

func TestTraceTx2(t *testing.T) {
	traceTx2 := TraceTx2{
		Res: "dsadadadada",
	}
	val, err := traceTx2.Marshal()
	assert.NoError(t, err)

	tt := &TraceTx2{}

	err = tt.Unmarshal(val)
	assert.NoError(t, err)
	assert.Equal(t, traceTx2.Res, tt.Res)

}
