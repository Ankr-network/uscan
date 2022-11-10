package types

import (
	"encoding/json"
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestTxSerialize(t *testing.T) {
	to := common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc")
	b := &Tx{
		BlockNum:  field.NewInt(11),
		GasTipCap: field.NewInt(11),
		GasFeeCap: field.NewInt(11),
		GasPrice:  field.NewInt(11),
		Nonce:     field.NewInt(11123123),
		Gas:       field.NewInt(123421),
		To:        &to,
		Value:     field.NewInt(12312313),
		Data:      []byte{11, 22, 212, 4, 54, 213, 2, 41, 41, 54},

		// Signature values
		V: field.NewInt(111),
		R: field.NewInt(111),
		S: field.NewInt(111),
	}

	res, err := b.Marshal()

	assert.NoError(t, err)
	t.Log(hexutil.Bytes(res).String())

	out := Tx{}

	err = out.Unmarshal(res)
	assert.NoError(t, err)

	assert.Equal(t, out.BlockNum, b.BlockNum)
	assert.Equal(t, out.GasTipCap, b.GasTipCap)
	assert.Equal(t, out.GasFeeCap, b.GasFeeCap)
	assert.Equal(t, out.GasPrice, b.GasPrice)
	assert.Equal(t, out.Nonce, b.Nonce)
	assert.Equal(t, out.Gas, b.Gas)
	assert.Equal(t, out.To, b.To)
	assert.Equal(t, out.Value, b.Value)
	assert.Equal(t, out.V, b.V)
	assert.Equal(t, out.R, b.R)
	assert.Equal(t, out.S, b.S)
}

var testTx = []byte(`{"blockHash":"0x93dc24ad05a3e73bbec6d52f0991a0c8ef6560ca9257bc8f28b08ec4d5b1f643","blockNumber":"0x1fe580","from":"0x3c10ec535d1a8cba60536a963cc62a1df855e71c","gas":"0x748c","gasPrice":"0x5dc","hash":"0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d","input":"0xa9059cbb0000000000000000000000008ec529c63f174996c5cf360081d94bac07a8615e00000000000000000000000000000000000000000000000000000002540be400","nonce":"0x5f","to":"0x07861819f3d9773088f67e5572bd645b2e5c15ef","transactionIndex":"0x0","value":"0x0","type":"0x0","v":"0x5e7d","r":"0x106dcc2018fbb19b93693348635c1b8f78b008778d8ef28201079ffd7789e569","s":"0x457579ddeed18c9a59ab5acd2d77c0862fc00b3e64d40af173e7e067f82a1ae4"}`)

func TestTxJson(t *testing.T) {
	out := &Tx{}
	err := json.Unmarshal(testTx, out)
	assert.NoError(t, err)
	t.Log(out)

	byteRes, err := json.Marshal(out)
	assert.NoError(t, err)
	t.Log(string(byteRes))
}
