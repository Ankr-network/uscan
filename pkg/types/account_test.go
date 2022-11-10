package types

import (
	"testing"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	t.Log(Erc20Flag + (Erc721Flag + Erc1155Flag))
}

func TestAccount(t *testing.T) {

	creator := common.HexToAddress("0x3c10ec535d1a8cba60536a963cc62a1df855e71c")
	txhash := common.HexToHash("0x537e032e5bc31b5e52f5e28c61c5aefd631b438bf5b9c71913c19d022a4ae528")
	a := &Account{
		BlockNumber: field.NewInt(12231),
		Owner:       common.HexToAddress("0x473780deaf4a2ac070bbba936b0cdefe7f267dfc"),
		Balance:     field.NewInt(1111),

		Erc20:            true,
		Erc721:           true,
		Erc1155:          true,
		Creator:          &creator,
		TxHash:           &txhash,
		Code:             []byte{11, 23, 2, 14, 51, 24, 51, 23, 4, 21, 4},
		Name:             "dasdada",
		Symbol:           "Da",
		TokenTotalSupply: field.NewInt(219282312313),
		NftTotalSupply:   field.NewInt(431421),
		Decimals:         field.NewInt(18),
	}
	res, err := a.Marshal()

	assert.NoError(t, err)
	t.Log(hexutil.Bytes(res).String())

	out := Account{}
	err = out.Unmarshal(res)
	assert.NoError(t, err)

	assert.Equal(t, out.BlockNumber, a.BlockNumber)
	assert.Equal(t, out.Balance, a.Balance)
	assert.Equal(t, out.Erc20, a.Erc20)
	assert.Equal(t, out.Erc721, a.Erc721)
	assert.Equal(t, out.Erc1155, a.Erc1155)
	assert.Equal(t, out.Creator, a.Creator)
	assert.Equal(t, out.TxHash, a.TxHash)
	assert.Equal(t, out.Code, a.Code)
	assert.Equal(t, out.Name, a.Name)
	assert.Equal(t, out.Symbol, a.Symbol)
	assert.Equal(t, out.TokenTotalSupply, a.TokenTotalSupply)
	assert.Equal(t, out.NftTotalSupply, a.NftTotalSupply)
	assert.Equal(t, out.Decimals, a.Decimals)
}
