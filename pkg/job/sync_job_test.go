package job

import (
	"testing"
)

func TestBlockJob(t *testing.T) {
	GlobalInit(2)
	sj := NewSyncJob(1600937, testRpc)
	sj.Execute()

	t.Log("blockData", sj.BlockData)
	t.Log("TransactionDatas", sj.TransactionDatas)
	t.Log("ReceiptDatas", sj.ReceiptDatas)

	for _, v := range sj.ReceiptDatas {
		t.Log(v.TxHash.Hex(), " err: ", v.ReturnErr)
		t.Log(v.TxHash.Hex(), " existInternalTx: ", v.ExistInternalTx)
	}

	t.Log("ContractOrMemberData", sj.ContractOrMemberData)

	for i, v := range sj.InternalTxs {
		t.Log("InternalTxs(", i, "): ", v)
	}
}
