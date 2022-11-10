package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestInternalTx(t *testing.T) {
	to := common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc")
	b := &InternalTx{
		TransactionHash: common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
		BlockNumber:     field.NewInt(1111),
		Status:          true,
		CallType:        "create2",
		Depth:           "0_1",
		From:            common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc"),
		To:              &to,
		Amount:          field.NewInt(1111),
		GasLimit:        field.NewInt(112321),
	}

	res, err := b.Marshal()

	assert.NoError(t, err)
	t.Log(hexutil.Bytes(res).String())

	out := InternalTx{}

	err = out.Unmarshal(res)
	assert.NoError(t, err)

	assert.Equal(t, out.TransactionHash, b.TransactionHash)
	assert.Equal(t, out.BlockNumber, b.BlockNumber)
	assert.Equal(t, out.Status, b.Status)
	assert.Equal(t, out.CallType, b.CallType)
	assert.Equal(t, out.CallType, b.CallType)
	assert.Equal(t, out.Depth, b.Depth)
	assert.Equal(t, out.From, b.From)
	assert.Equal(t, out.To, b.To)
	assert.Equal(t, out.Amount, b.Amount)
	assert.Equal(t, out.GasLimit, b.GasLimit)

}
