package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestHomeNilMarshal(t *testing.T) {
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
	assert.Equal(t, len(out.Blocks), 0)
	assert.Equal(t, len(out.Txs), 0)
}

func TestHomeMarshal(t *testing.T) {

	h := Home{
		AddressTotal: *field.NewInt(11),
		Blocks: []*BkSim{
			{
				Number: *field.NewInt(1),
			}, {
				Number: *field.NewInt(2),
			}, {
				Number: *field.NewInt(3),
			},
		},
		Txs: []*TxSim{
			{
				Hash: common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
			}, {
				Hash: common.HexToHash("0x866fe28eb38d737da9a10a5dcfad3ce0b1fd517cf853609cffb002166abd55d7"),
			},
		},
		DateTxs: map[string]*field.BigInt{
			"20221011": field.NewInt(11),
			"20221012": field.NewInt(12),
		},
	}
	res, err := h.Marshal()
	assert.NoError(t, err)
	t.Log(res)

	out := &Home{}
	err = out.Unmarshal(res)
	assert.NoError(t, err)
	assert.Equal(t, len(out.Blocks), 3)
	assert.Equal(t, len(out.Txs), 2)
	assert.Equal(t, out.Blocks[0].Number, *field.NewInt(1))
	assert.Equal(t, out.Blocks[1].Number, *field.NewInt(2))
	assert.Equal(t, out.Blocks[2].Number, *field.NewInt(3))

	assert.Equal(t, out.Txs[0].Hash, common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"))
	assert.Equal(t, out.Txs[1].Hash, common.HexToHash("0x866fe28eb38d737da9a10a5dcfad3ce0b1fd517cf853609cffb002166abd55d7"))
	assert.Equal(t, out.DateTxs["20221011"].String(), field.NewInt(11).String())
}
