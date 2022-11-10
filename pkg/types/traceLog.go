package types

import (
	"encoding/json"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

type TraceTx struct {
	Res    string
	LogNum int
}

func (b *TraceTx) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *TraceTx) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

type TraceTx2 struct {
	Res string
}

func (b *TraceTx2) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *TraceTx2) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

type CallFrame struct {
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to,omitempty"`
	Value   *field.BigInt  `json:"value,omitempty"`
	Gas     *field.BigInt  `json:"gas"`
	GasUsed *field.BigInt  `json:"gasUsed"`
	Input   hexutil.Bytes  `json:"input"`
	Output  hexutil.Bytes  `json:"output,omitempty"`
	Error   string         `json:"error,omitempty"`
	Calls   []CallFrame    `json:"calls,omitempty"`
}

func (c *CallFrame) JsonToString() string {
	res, _ := json.Marshal(c)
	return string(res)
}

type ExecutionResult struct {
	Gas         uint64     `json:"gas"`
	Failed      bool       `json:"failed"`
	ReturnValue string     `json:"returnValue"`
	StructLogs  StructLogs `json:"structLogs"`
}

type StructLogRes struct {
	Pc      uint64      `json:"pc"`
	Op      string      `json:"op"`
	Gas     uint64      `json:"gas"`
	GasCost uint64      `json:"gasCost"`
	Depth   int         `json:"depth"`
	Error   interface{} `json:"error,omitempty"`
}
type StructLogs []StructLogRes

func (c *ExecutionResult) JsonToString() string {
	res, _ := json.Marshal(c)
	return string(res)
}

func (c *StructLogs) JsonToString() string {
	res, _ := json.Marshal(c)
	return string(res)
}
