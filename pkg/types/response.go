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

type ListBlockResp struct {
	Number            string `json:"number"`
	Timestamp         uint64 `json:"timestamp"`
	TransactionsTotal uint64 `json:"transactionsTotal"`
	Miner             string `json:"miner"`
	GasLimit          string `json:"gasLimit"` // int64 改成string
	GasUsed           string `json:"gasUsed"`  // int64 改成string
}

type TxResp struct {
	Hash                 string               `json:"hash"`
	Method               string               `json:"method"`
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
	Input                string               `json:"input"`
	Nonce                string               `json:"nonce"` // change string
	TransactionIndex     *string              `json:"transactionIndex"`
	Type                 *string              `json:"type"`
	ChainID              *string              `json:"chainID"`
	V                    string               `json:"v"`
	R                    string               `json:"r"`
	S                    string               `json:"s"`
	TotalLogs            int                  `json:"totalLogs"`
	TokensTransferred    []*TokensTransferred `json:"tokensTransferred"`
	BaseFeePerGas        *string              `json:"baseFeePerGas"`
	GasLimit             string               `json:"gasLimit"` // change string
	MethodName           string               `json:"methodName"`
	Logs                 []*RtLogResp         `json:"logs"`
	RtResp                                    // 新增
}

type RtLogResp struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type RtResp struct {
	ContractAddress       *string `json:"contractAddress"`
	ContractAddressName   string  `json:"contractAddressName"`
	ContractAddressSymbol string  `json:"contractAddressSymbol"`
	CumulativeGasUsed     *string `json:"cumulativeGasUsed"`
	EffectiveGasPrice     *string `json:"effectiveGasPrice"`
	GasUsed               string  `json:"gasUsed"`
	LogsBloom             *string `json:"logsBloom"`
	Root                  string  `json:"root"`
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

type ListTransactionResp struct {
	Hash        string  `json:"hash"` // transaction Hash
	Method      *string `json:"method"`
	BlockHash   *string `json:"blockHash"`
	BlockNumber *string `json:"blockNumber"`
	From        string  `json:"from"`
	FromName    string  `json:"fromName"`
	FromSymbol  string  `json:"fromSymbol"`
	FromCode    string  `json:"fromCode"`
	To          string  `json:"to"`
	ToName      string  `json:"toName"`
	ToSymbol    string  `json:"toSymbol"`
	ToCode      string  `json:"toCode"`
	Gas         *string `json:"gas"`
	GasPrice    *string `json:"gasPrice"`
	Value       *string `json:"value"`
	CreatedTime uint64  `json:"createTime"`
}

type TransactionBaseResp struct {
	Hash     string `json:"hash"` // transaction Hash
	Nonce    string `json:"nonce"`
	GasUsed  string `json:"gasUsed"`
	GasLimit string `json:"gasLimit"`
	Status   uint64 `json:"status"`
}

type AccountResp struct {
	Owner            string  `json:"owner"`
	Balance          string  `json:"balance"`
	BlockNumber      *string `json:"blockNumber"`
	Creator          *string `json:"creator"`
	TxHash           *string `json:"txHash"`
	Code             *string `json:"code"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	TokenTotalSupply *string `json:"tokenTotalSupply"`
	NftTotalSupply   *string `json:"nftTotalSupply"`
	Decimals         uint64  `json:"decimals"`
	//CreatedTime      uint64  `json:"createdTime"`
}

type InternalTxResp struct {
	TransactionHash string  `json:"transactionHash"`
	BlockNumber     string  `json:"blockNumber"`
	Status          bool    `json:"status"`
	CallType        string  `json:"callType"`
	Depth           string  `json:"depth"`
	From            string  `json:"from"`
	To              *string `json:"to"`
	Amount          string  `json:"amount"`
	GasLimit        string  `json:"gasLimit"`
	CreatedTime     uint64  `json:"createdTime"`
}

type Erc20TxResp struct {
	TransactionHash string `json:"transactionHash"`
	BlockHash       string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	Contract        string `json:"contract"`
	ContractName    string `json:"contractName"`
	ContractSymbol  string `json:"contractSymbol"`
	Method          string `json:"method"`
	From            string `json:"from"`
	FromName        string `json:"fromName"`
	FromSymbol      string `json:"fromSymbol"`
	FromCode        string `json:"fromCode"`
	To              string `json:"to"`
	ToName          string `json:"toName"`
	ToSymbol        string `json:"toSymbol"`
	ToCode          string `json:"toCode"`
	Value           string `json:"value"`
	CreatedTime     uint64 `json:"createdTime"`
}

type Erc721TxResp struct {
	TransactionHash string `json:"transactionHash"`
	BlockHash       string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	Contract        string `json:"contract"`
	ContractName    string `json:"contractName"`
	ContractSymbol  string `json:"contractSymbol"`
	Method          string `json:"method"`
	From            string `json:"from"`
	FromName        string `json:"fromName"`
	FromSymbol      string `json:"fromSymbol"`
	FromCode        string `json:"fromCode"`
	To              string `json:"to"`
	ToName          string `json:"toName"`
	ToSymbol        string `json:"toSymbol"`
	ToCode          string `json:"toCode"`
	TokenID         uint64 `json:"tokenID"`
	CreatedTime     uint64 `json:"createdTime"`
}

type Erc1155TxResp struct {
	TransactionHash string `json:"transactionHash"`
	BlockHash       string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	Contract        string `json:"contract"`
	ContractName    string `json:"contractName"`
	ContractSymbol  string `json:"contractSymbol"`
	Method          string `json:"method"`
	From            string `json:"from"`
	FromName        string `json:"fromName"`
	FromSymbol      string `json:"fromSymbol"`
	FromCode        string `json:"fromCode"`
	To              string `json:"to"`
	ToName          string `json:"toName"`
	ToSymbol        string `json:"toSymbol"`
	ToCode          string `json:"toCode"`
	TokenID         uint64 `json:"tokenID"`
	Value           string `json:"value"`
	CreatedTime     uint64 `json:"createdTime"`
}

type TraceTxResp struct {
	Res    string `json:"res"`
	LogNum string `json:"logNum"`
}

type TraceTx2Resp struct {
	Res string `json:"res"`
}
