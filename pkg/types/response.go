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
	ExtraData         []byte   `json:"extraData"`
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
	Number            string  `json:"number"`
	Timestamp         uint64  `json:"timestamp"`
	TransactionsTotal uint64  `json:"transactionsTotal"`
	Miner             string  `json:"miner"`
	GasLimit          string  `json:"gasLimit"` // int64 改成string
	GasUsed           string  `json:"gasUsed"`  // int64 改成string
	BaseFeePerGas     *string `json:"baseFeePerGas"`
}

type TxResp struct {
	Hash                 string               `json:"hash"`
	Method               string               `json:"method"`
	BlockHash            string               `json:"blockHash"`
	BlockNumber          string               `json:"blockNumber"`
	From                 string               `json:"from"`
	FromName             string               `json:"fromName"`
	FromSymbol           string               `json:"fromSymbol"`
	FromContract         bool                 `json:"fromContract"`
	To                   string               `json:"to"`
	ToName               string               `json:"toName"`
	ToSymbol             string               `json:"toSymbol"`
	ToContract           bool                 `json:"toContract"`
	Gas                  string               `json:"gas"` // change string
	GasPrice             string               `json:"gasPrice"`
	Value                string               `json:"value"`
	CreatedTime          uint64               `json:"createTime"`
	MaxFeePerGas         *string              `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *string              `json:"maxPriorityFeePerGas"`
	Input                string               `json:"input"`
	Nonce                string               `json:"nonce"` // change string
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
	Address  string   `json:"address"`
	Topics   []string `json:"topics"`
	Data     string   `json:"data"`
	LogIndex uint64   `json:"logIndex"`
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
	From            string `json:"from"`
	FromHex         string `json:"fromHex"`
	To              string `json:"to"`
	ToHex           string `json:"toHex"`
	Address         string `json:"address"`
	AddressName     string `json:"addressName"`
	AddressSymbol   string `json:"addressSymbol"`
	AddressDecimals uint64 `json:"addressDecimals"`
	AddressValue    string `json:"addressValue"`
}

type ListTransactionResp struct {
	Hash        string  `json:"hash"` // transaction Hash
	Method      string  `json:"method"`
	BlockHash   string  `json:"blockHash"`
	BlockNumber string  `json:"blockNumber"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	ToName      string  `json:"toName"`
	ToSymbol    string  `json:"toSymbol"`
	ToContract  bool    `json:"toContract"`
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
	Erc20   bool `json:"erc20"`
	Erc721  bool `json:"erc721"`
	Erc1155 bool `json:"erc1155"`
}

type InternalTxResp struct {
	TransactionHash string `json:"transactionHash"`
	BlockNumber     string `json:"blockNumber"`
	Status          bool   `json:"status"`
	CallType        string `json:"callType"`
	Depth           string `json:"depth"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	GasLimit        string `json:"gasLimit"`
	CreatedTime     uint64 `json:"createdTime"`
}

type Erc20TxResp struct {
	TransactionHash  string `json:"transactionHash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Contract         string `json:"contract"`
	ContractName     string `json:"contractName"`
	ContractSymbol   string `json:"contractSymbol"`
	ContractDecimals uint64 `json:"contractDecimals"`
	Method           string `json:"method"`
	From             string `json:"from"`
	FromName         string `json:"fromName"`
	FromSymbol       string `json:"fromSymbol"`
	FromContract     bool   `json:"fromContract"`
	To               string `json:"to"`
	ToName           string `json:"toName"`
	ToSymbol         string `json:"toSymbol"`
	ToContract       bool   `json:"toContract"`
	Value            string `json:"value"`
	CreatedTime      uint64 `json:"createdTime"`
}

type Erc721TxResp struct {
	TransactionHash  string `json:"transactionHash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Contract         string `json:"contract"`
	ContractName     string `json:"contractName"`
	ContractSymbol   string `json:"contractSymbol"`
	ContractDecimals uint64 `json:"contractDecimals"`
	Method           string `json:"method"`
	From             string `json:"from"`
	FromName         string `json:"fromName"`
	FromSymbol       string `json:"fromSymbol"`
	FromContract     bool   `json:"fromContract"`
	To               string `json:"to"`
	ToName           string `json:"toName"`
	ToSymbol         string `json:"toSymbol"`
	ToContract       bool   `json:"toContract"`
	TokenID          string `json:"tokenID"`
	CreatedTime      uint64 `json:"createdTime"`
}

type Erc1155TxResp struct {
	TransactionHash  string `json:"transactionHash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Contract         string `json:"contract"`
	ContractName     string `json:"contractName"`
	ContractSymbol   string `json:"contractSymbol"`
	ContractDecimals uint64 `json:"contractDecimals"`
	Method           string `json:"method"`
	From             string `json:"from"`
	FromName         string `json:"fromName"`
	FromSymbol       string `json:"fromSymbol"`
	FromContract     bool   `json:"fromContract"`
	To               string `json:"to"`
	ToName           string `json:"toName"`
	ToSymbol         string `json:"toSymbol"`
	ToContract       bool   `json:"toContract"`
	TokenID          string `json:"tokenID"`
	Value            string `json:"value"`
	CreatedTime      uint64 `json:"createdTime"`
}

type TraceTxResp struct {
	Res    string `json:"res"`
	LogNum string `json:"logNum"`
}

type TraceTx2Resp struct {
	Res string `json:"res"`
}
type HolderResp struct {
	Address  string
	Quantity string
}

type InventoryResp struct {
	Address string
	TokenID string
}

type ContractVerityInfo struct {
	ContractName    string            `json:"contractName"`
	CompilerVersion string            `json:"compilerVersion"`
	Optimization    uint64            `json:"optimization"`
	Runs            uint64            `json:"runs"`
	EVMVersion      string            `json:"evmVersion"`
	LicenseType     uint64            `json:"licenseType"`
	ABI             string            `json:"abi"`
	Metadata        map[string]string `json:"metadata"`
	Object          string            `json:"object"`
}

type ContractVerityInfoResp struct {
	Contract             *ContractVerityInfo `json:"contract"`
	ProxyContractAddress string              `json:"proxyContractAddress"`
	ProxyContract        *ContractVerityInfo `json:"proxyContract"`
}
