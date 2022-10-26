package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {

	a := &Account{
		Erc20:   true,
		Balance: 35621,
		Code:    []byte("8066234511"),
	}

	bs, err := a.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	b := &Account{}

	b.Unmarshal(bs)

	assert.Equal(t, a.Erc20, b.Erc20, "ERC20")
	assert.Equal(t, a.Balance, b.Balance, "Balance")
	assert.Equal(t, a.Code, b.Code, "Code")

}
