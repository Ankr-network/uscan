package types

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Bloom [types.BloomByteLength]byte

// MarshalText encodes b as a hex string with 0x prefix.
func (b Bloom) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

// UnmarshalText b as a hex string with 0x prefix.
func (b *Bloom) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Bloom", input, b[:])
}

type BlockNonce [8]byte

// MarshalText encodes n as a hex string with 0x prefix.
func (n BlockNonce) MarshalText() ([]byte, error) {
	return hexutil.Bytes(n[:]).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *BlockNonce) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BlockNonce", input, n[:])
}

type Block struct {
	Number           *field.BigInt  `json:"number"  rlp:"-"`
	Hash             common.Hash    `json:"hash"`
	ParentHash       common.Hash    `json:"parentHash"`
	UncleHash        common.Hash    `json:"sha3Uncles"`
	Coinbase         common.Address `json:"miner"`
	Root             common.Hash    `json:"stateRoot"`
	TxHash           common.Hash    `json:"transactionsRoot"`
	ReceiptHash      common.Hash    `json:"receiptsRoot"`
	Bloom            Bloom          `json:"logsBloom"`
	Difficulty       field.BigInt   `json:"difficulty"`
	GasLimit         field.BigInt   `json:"gasLimit"`
	GasUsed          field.BigInt   `json:"gasUsed"`
	TimeStamp        field.BigInt   `json:"timestamp"`
	Extra            []byte         `json:"extraData"`
	MixDigest        common.Hash    `json:"mixHash"`
	Nonce            BlockNonce     `json:"nonce"`
	BaseFee          field.BigInt   `json:"baseFeePerGas"`
	Size             field.BigInt   `json:"size"`
	TotalDifficulty  field.BigInt   `json:"totalDifficulty"`
	Transactions     []common.Hash  `json:"transactions" rlp:"-"`
	TransactionTotal field.BigInt   `json:"-"`
}

func (b *Block) Marshal() ([]byte, error) {
	b.TransactionTotal = *field.NewInt(int64(len(b.Transactions)))
	return rlp.EncodeToBytes(b)
}

func (b *Block) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}
