package field

import (
	"testing"
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
