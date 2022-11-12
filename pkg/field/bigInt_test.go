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
}
