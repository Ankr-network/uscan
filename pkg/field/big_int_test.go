package field

import "testing"

func TestBigInt(t *testing.T) {
	var b = &BigInt{}
	t.Log(b.String())
}
