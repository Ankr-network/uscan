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
