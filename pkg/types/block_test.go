package types

import (
	"encoding/json"
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestBlockSerialize(t *testing.T) {
	b := &Block{
		ParentHash:       common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
		UncleHash:        common.HexToHash("0x6cca2193971f5813123795a259f026b4387ecf00dd134eb7a44e59885983f3f4"),
		Coinbase:         common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc"),
		Root:             common.HexToHash("0x866fe28eb38d737da9a10a5dcfad3ce0b1fd517cf853609cffb002166abd55d7"),
		TxHash:           common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
		ReceiptHash:      common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
		Bloom:            [types.BloomByteLength]byte{11, 22, 22, 21, 245, 221, 2, 3, 4, 52, 3},
		Difficulty:       field.NewInt(111),
		Number:           field.NewInt(111),
		Hash:             common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528"),
		GasLimit:         field.NewInt(1112),
		GasUsed:          field.NewInt(213),
		Time:             field.NewInt(1231),
		Extra:            []byte("adsdada291s92sada"),
		MixDigest:        common.HexToHash("0xd883010b00846765746888676f312e31392e31856c696e7578"),
		Nonce:            [8]byte{12, 3, 4, 33, 1},
		BaseFee:          field.NewInt(111111),
		Size:             field.NewInt(22222),
		TotalDifficulty:  field.NewInt(4213521321413),
		Transactions:     []common.Hash{common.HexToHash("0x5977673bbb2382d0192dd79e96b2dfb9e3f7752374c5680ffb175767da10d8e9"), common.HexToHash("0x90e4625f5b3eb1cdfd6037d97530084b8cd3afc2ba16b5f48d2848fec84620a3"), common.HexToHash("0x74c48de18ff97958910dee2b65dca536334582d96bdc6f49dafce33bfc6004f7")},
		TransactionTotal: field.NewInt(1111),
	}

	res, err := b.Marshal()
	assert.NoError(t, err)

	out := &Block{}
	err = out.Unmarshal(res)
	assert.NoError(t, err)

	assert.Equal(t, out.ParentHash, b.ParentHash)
	assert.Equal(t, out.UncleHash, b.UncleHash)
	assert.Equal(t, out.Coinbase, b.Coinbase)
	assert.Equal(t, out.Root, b.Root)
	assert.Equal(t, out.TxHash, b.TxHash)
	assert.Equal(t, out.ReceiptHash, b.ReceiptHash)
	assert.Equal(t, out.Bloom, b.Bloom)
	assert.Equal(t, out.Difficulty, b.Difficulty)
	assert.Equal(t, out.Hash, b.Hash)
	assert.Equal(t, out.GasLimit, b.GasLimit)
	assert.Equal(t, out.Time, b.Time)
	assert.Equal(t, out.Extra, b.Extra)
	assert.Equal(t, out.MixDigest, b.MixDigest)
	assert.Equal(t, out.Nonce, b.Nonce)
	assert.Equal(t, out.BaseFee, b.BaseFee)
	assert.Equal(t, out.Size, b.Size)
	assert.Equal(t, out.TotalDifficulty, b.TotalDifficulty)
	assert.Equal(t, out.TransactionTotal, b.TransactionTotal)
}

var testBlock = []byte(`{"baseFeePerGas":"0x","difficulty":"0x2","extraData":"0xd883010a0d846765746888676f312e31382e32856c696e757800000000000000a7b58a705aae18821c766200af2808640af36197c5d000fba9c1d121708d303f1eb9b8963f5c5fcba5212bd0ede185a723b35ddd3c78150d0072b4ae138f6c0f00","gasLimit":"0x6619de","gasUsed":"0x0","hash":"0xcc2708783820f52ca48637cc6d37b91b5ec62feac809f8e154231503d6c001a4","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x16a","parentHash":"0x6aa7ec70ebaf19c604b6412bd0e08d574d167f29c6852d41c43f857fa9e3dcd4","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x261","stateRoot":"0xdb9e3559eef6ebecda75033eb7b0482bfc8ba2a51e97b8c404521fbe4b4bf84b","timestamp":"0x62f319f7","totalDifficulty":"0x2d5","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}`)

func TestBlockJson(t *testing.T) {
	out := &Block{}

	err := json.Unmarshal(testBlock, out)
	// err := out.UnmarshalJSON(testBlock)
	assert.NoError(t, err)
	t.Log(out)

	byteRes, err := json.Marshal(out)
	assert.NoError(t, err)
	t.Log(string(byteRes))
}
