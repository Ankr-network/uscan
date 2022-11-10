package types

type HomeBlock struct {
	Number            string `json:"number"`
	Timestamp         uint64 `json:"timestamp"`
	Miner             string `json:"miner"`
	GasUsed           uint64 `json:"gasUsed"`
	TransactionsTotal int64  `json:"transactionsTotal"`
}

type HomeTx struct {
	Hash        string  `json:"hash"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	GasPrice    *string `json:"gasPrice"`
	Gas         *string `json:"gas"`
	CreatedTime uint64  `json:"createTime"`
}
