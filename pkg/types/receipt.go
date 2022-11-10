package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Log struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    []byte         `json:"data"`
}

type Rt struct {
	TxHash            common.Hash     `json:"transactionHash"  rlp:"-"`
	Type              uint8           `json:"type,omitempty"`
	PostState         []byte          `json:"root"`
	Status            uint64          `json:"status"`
	CumulativeGasUsed *big.Int        `json:"cumulativeGasUsed"`
	Bloom             Bloom           `json:"logsBloom"        `
	Logs              []*Log          `json:"logs"             `
	ContractAddress   *common.Address `json:"contractAddress"`
	GasUsed           *big.Int        `json:"gasUsed"`
	EffectiveGasPrice *big.Int        `json:"effectiveGasPrice"`
}

func (b *Rt) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Rt) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

func (b *Rt) UnmarshalJSON(bin []byte) error {
	return nil
}
