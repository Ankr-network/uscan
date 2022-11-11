package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

type Log struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
}

type Rt struct {
	TxHash            common.Hash     `json:"transactionHash"  rlp:"-"`
	Type              *field.BigInt   `json:"type,omitempty"`
	PostState         hexutil.Bytes   `json:"root"`
	Status            *field.BigInt   `json:"status"`
	CumulativeGasUsed *field.BigInt   `json:"cumulativeGasUsed"`
	Bloom             Bloom           `json:"logsBloom"`
	Logs              []*Log          `json:"logs"`
	ContractAddress   *common.Address `json:"contractAddress"`
	GasUsed           *field.BigInt   `json:"gasUsed"`
	EffectiveGasPrice *field.BigInt   `json:"effectiveGasPrice"`
	ExistInternalTx   bool
	ReturnErr         string
}

func (b *Rt) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Rt) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
