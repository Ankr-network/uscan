package field

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

	a := bi.String()

	bi.Add(NewInt(111))
	fmt.Println(bi.ToUint64())
	res, _ := hexutil.DecodeBig(a)
	fmt.Println(res.String())

	assert.Equal(t, bi.ToUint64(), uint64(333))
	biRes := bi.Add(NewInt(2))
	assert.Equal(t, biRes.ToUint64(), uint64(335))

	biRes = bi.Sub(NewInt(3))
	assert.Equal(t, biRes.ToUint64(), uint64(332))
	assert.Equal(t, bi.ToUint64(), uint64(332))
}
