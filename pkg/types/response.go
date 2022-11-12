package types

type HomeBlock struct {
	Number            string `json:"number"`
	Timestamp         uint64 `json:"timestamp"`
	Miner             string `json:"miner"`
	GasUsed           string `json:"gasUsed"`
	TransactionsTotal uint64 `json:"transactionsTotal"`
}

type HomeTx struct {
	Hash        string  `json:"hash"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	GasPrice    *string `json:"gasPrice"`
	Gas         *string `json:"gas"`
	CreatedTime uint64  `json:"createTime"`
}

type BlockResp struct {
	BaseFeePerGas     *string  `json:"baseFeePerGas"`
	Difficulty        string   `json:"difficulty"`
	ExtraData         string   `json:"extraData"`
	GasLimit          string   `json:"gasLimit"` // int64 改成string
	GasUsed           string   `json:"gasUsed"`  // int64 改成string
	Hash              string   `json:"hash"`
	LogsBloom         string   `json:"logsBloom"`
	Miner             string   `json:"miner"`
	MixHash           string   `json:"mixHash"`
	Nonce             string   `json:"nonce"`
	Number            string   `json:"number"`
	ParentHash        string   `json:"parentHash"`
	ReceiptsRoot      string   `json:"receiptsRoot"`
	Sha3Uncles        string   `json:"sha3Uncles"`
	Size              string   `json:"size"`
	StateRoot         string   `json:"stateRoot"`
	Timestamp         uint64   `json:"timestamp"`
	TotalDifficulty   uint64   `json:"totalDifficulty"`
	Transactions      []string `json:"transactions"`
	TransactionsTotal uint64   `json:"transactionsTotal"`
	TransactionsRoot  string   `json:"transactionsRoot"`
	//CreatedTime       uint64   `json:"createdTime"`
}

type TransactionResp struct {
	Hash                 string               `json:"hash"`
	Method               *string              `json:"method"`
	BlockHash            *string              `json:"blockHash"`
	BlockNumber          *string              `json:"blockNumber"`
	From                 string               `json:"from"`
	FromName             string               `json:"fromName"`
	FromSymbol           string               `json:"fromSymbol"`
	FromCode             string               `json:"fromCode"`
	To                   string               `json:"to"`
	ToName               string               `json:"toName"`
	ToSymbol             string               `json:"toSymbol"`
	ToCode               string               `json:"toCode"`
	Gas                  string               `json:"gas"` // change string
	GasPrice             string               `json:"gasPrice"`
	Value                string               `json:"value"`
	CreatedTime          uint64               `json:"createTime"`
	MaxFeePerGas         *string              `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *string              `json:"maxPriorityFeePerGas"`
	Input                *string              `json:"input"`
	Nonce                string               `json:"nonce"` // change string
	TransactionIndex     *string              `json:"transactionIndex"`
	Type                 *string              `json:"type"`
	ChainID              *string              `json:"chainID"`
	V                    string               `json:"v"`
	R                    string               `json:"r"`
	S                    string               `json:"s"`
	TotalLogs            int64                `json:"totalLogs"`
	TokensTransferred    []*TokensTransferred `json:"tokensTransferred"`
	BaseFeePerGas        *string              `json:"baseFeePerGas"`
	GasLimit             string               `json:"gasLimit"` // change string
	MethodName           string               `json:"methodName"`
	ReceiptResp
}

type ReceiptResp struct {
	ContractAddress       *string `json:"contractAddress"`
	ContractAddressName   string  `json:"contractAddressName"`
	ContractAddressSymbol string  `json:"contractAddressSymbol"`
	CumulativeGasUsed     *uint64 `json:"cumulativeGasUsed"`
	EffectiveGasPrice     *string `json:"effectiveGasPrice"`
	GasUsed               uint64  `json:"gasUsed"`
	LogsBloom             *string `json:"logsBloom"`
	Root                  *string `json:"root"`
	Status                uint64  `json:"status"`
	ErrorReturn           string  `json:"errorReturn"`
}

type TokensTransferred struct {
	From          string `json:"from"`
	FromHex       string `json:"fromHex"`
	To            string `json:"to"`
	ToHex         string `json:"toHex"`
	Address       string `json:"address"`
	AddressName   string `json:"addressName"`
	AddressSymbol string `json:"addressSymbol"`
}
