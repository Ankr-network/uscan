package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigInt(t *testing.T) {
	var b = &BigInt{}
	t.Log(b.String())
}

func TestBytes(t *testing.T) {
	arr := []byte{1, 2, 33, 123}

	in := BigInt{}
	in.SetBytes(arr)

	t.Log(in.String())
}

func TestAddNum(t *testing.T) {
	bi := NewInt(111)
	bi.Add(NewInt(111))
	bi.Add(NewInt(111))

	assert.Equal(t, bi.ToUint64(), uint64(333))
	res := bi.Add(NewInt(2))
	assert.Equal(t, res.ToUint64(), uint64(335))

	res = bi.Sub(NewInt(3))
	assert.Equal(t, res.ToUint64(), uint64(332))
	assert.Equal(t, bi.ToUint64(), uint64(332))
}
