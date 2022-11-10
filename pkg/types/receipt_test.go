package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestRtSerialize(t *testing.T) {
	b := &Rt{
		Type:              1,
		PostState:         []byte{111, 22, 1, 231, 31},
		Status:            11,
		CumulativeGasUsed: big.NewInt(11),
		Bloom:             [types.BloomByteLength]byte{11, 22, 22, 21, 245, 221, 2, 3, 4, 52, 3},
		Logs: []*Log{
			{
				Address: common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc"),
				Topics:  []common.Hash{common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"), common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528")},
				Data:    []byte{1, 12, 3, 4, 2, 12, 24, 32, 1},
			},
		},

		TxHash:          common.HexToHash(""),
		ContractAddress: &common.Address{},
		GasUsed:         big.NewInt(232),

		EffectiveGasPrice: big.NewInt(111),
	}

	res, err := b.Marshal()

	assert.NoError(t, err)
	t.Log(hexutil.Bytes(res).String())

	out := Rt{}

	err = out.Unmarshal(res)
	assert.NoError(t, err)

	assert.Equal(t, out.Type, b.Type)
	assert.Equal(t, out.PostState, b.PostState)
	assert.Equal(t, out.Status, b.Status)
	assert.Equal(t, out.CumulativeGasUsed, b.CumulativeGasUsed)
	assert.Equal(t, out.Bloom, b.Bloom)
	assert.Equal(t, out.TxHash, b.TxHash)
	assert.Equal(t, out.ContractAddress, b.ContractAddress)
	assert.Equal(t, out.GasUsed, b.GasUsed)
	assert.Equal(t, out.EffectiveGasPrice, b.EffectiveGasPrice)
	assert.Equal(t, len(out.Logs), len(b.Logs))
}

var testRt = []byte(`{"blockHash":"0x93dc24ad05a3e73bbec6d52f0991a0c8ef6560ca9257bc8f28b08ec4d5b1f643","blockNumber":"0x1fe580","contractAddress":null,"cumulativeGasUsed":"0x31c0","effectiveGasPrice":"0x5dc","from":"0x3c10ec535d1a8cba60536a963cc62a1df855e71c","gasUsed":"0x31c0","logs":[{"address":"0x07861819f3d9773088f67e5572bd645b2e5c15ef","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000003c10ec535d1a8cba60536a963cc62a1df855e71c","0x0000000000000000000000008ec529c63f174996c5cf360081d94bac07a8615e"],"data":"0x00000000000000000000000000000000000000000000000000000002540be400","blockNumber":"0x1fe580","transactionHash":"0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d","transactionIndex":"0x0","blockHash":"0x93dc24ad05a3e73bbec6d52f0991a0c8ef6560ca9257bc8f28b08ec4d5b1f643","logIndex":"0x0","removed":false}],"logsBloom":"0x00000000000000000020000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000020000000000000000008000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000210000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000020002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010080000","status":"0x1","to":"0x07861819f3d9773088f67e5572bd645b2e5c15ef","transactionHash":"0x9aaa0c4a421d8cd3e52765475acccb23a6dd388d0be384b00bb73fc7e8db796d","transactionIndex":"0x0","type":"0x0"}`)

func TestRtJson(t *testing.T) {
	out := &Rt{}
	err := json.Unmarshal(testRt, out)
	assert.NoError(t, err)
	t.Log(out)
}
