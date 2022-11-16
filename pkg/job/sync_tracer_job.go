package job

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SyncTracerJob struct {
	Completed bool
	Status    bool // tx exec result
	block     uint64
	tx        common.Hash
	client    rpcclient.RpcClient
	CallFrame *types.CallFrame
	Error     string

	InternalTxs []*types.InternalTx
	// address => map
	ContractOrMemberData map[common.Address]*types.Account
	ContractInfoMap      map[common.Address]*types.Contract
}

func NewSyncTracerJob(block uint64,
	tx common.Hash,
	client rpcclient.RpcClient,
) *SyncTracerJob {
	return &SyncTracerJob{
		block:                block,
		tx:                   tx,
		client:               client,
		InternalTxs:          make([]*types.InternalTx, 0, 1),
		ContractOrMemberData: make(map[common.Address]*types.Account),
		ContractInfoMap:      make(map[common.Address]*types.Contract),
	}
}

func (e *SyncTracerJob) Execute() {
	var err error
	for {
		e.CallFrame, err = e.client.GetTracerCall(context.Background(), e.tx)
		if err != nil {
			log.Errorf("get transaction(%s) data failed: %v", e.tx.Hex(), err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	if e.CallFrame.Error == "" {
		e.Status = true
	} else {
		if bytes.HasPrefix(e.CallFrame.Output, []byte("0x08c379a0")) {
			res, _ := decodeParameter("string", hexutil.MustDecode(fmt.Sprintf("0x%s", e.CallFrame.Output[10:])))
			e.Error = res[0].(string)
		} else if len(e.CallFrame.Output) != 0 {
			e.Error = e.CallFrame.Output.String()
		}
	}
	e.handleCall("0", e.CallFrame)

	if !e.Status {
		for _, v := range e.InternalTxs {
			v.Status = e.Status
		}
	}

	e.Completed = true
}

func (e *SyncTracerJob) handleCall(prefix string, data *types.CallFrame) {
	if data == nil {
		return
	}

	if data.Type == "CREATE2" || data.Type == "CREATE" {
		in, arg := getByteCodeAndArg(data.To, data.Input, data.Output)
		e.ContractOrMemberData[data.To] = &types.Account{
			Owner:       data.To,
			BlockNumber: field.NewInt(int64(e.block)),
			Creator:     &data.From,
			TxHash:      &e.tx,
		}
		e.ContractInfoMap[data.To] = &types.Contract{
			ByteCode:              in,
			ConstructorArguements: arg,
			DeployedCode:          data.Output,
		}
	}

	if data.Value != nil && data.Value.String() != "0x0" {
		var status bool
		if data.Error == "" {
			status = true
			if _, ok := e.ContractOrMemberData[data.From]; !ok {
				e.ContractOrMemberData[data.From] = &types.Account{
					Owner:       data.From,
					BlockNumber: field.NewInt(int64(e.block)),
				}
			}

			if _, ok := e.ContractOrMemberData[data.To]; !ok {
				e.ContractOrMemberData[data.To] = &types.Account{
					Owner:       data.To,
					BlockNumber: field.NewInt(int64(e.block)),
				}
			}
		}
		if prefix != "0" {
			e.InternalTxs = append(e.InternalTxs, &types.InternalTx{
				TransactionHash: e.tx,
				BlockNumber:     field.NewInt(int64(e.block)),
				Status:          status,
				CallType:        data.Type,
				Depth:           prefix,
				From:            data.From,
				To:              &data.To,
				Amount:          data.Value,
				GasLimit:        data.Gas,
			})
		}
	}

	e.handleCalls(prefix+"_1", data.Calls)
}

func (e *SyncTracerJob) handleCalls(prefix string, data []*types.CallFrame) {
	for _, v := range data {
		e.handleCall(prefix, v)
	}
}

func decodeParameter(typ string, data []byte) ([]interface{}, error) {
	t, _ := abi.NewType(typ, "", nil)
	args := abi.Arguments{{
		Type: t,
	},
	}
	return args.UnpackValues(data)
}

func getByteCodeAndArg(to common.Address, in []byte, out []byte) (bytecode []byte, arg []byte) {
	splitOp := out[len(out)-32:]

	res := bytes.Split(in, splitOp)
	if len(res) == 2 {
		return append(res[0], splitOp...), res[1]
	}
	log.Errorf("Bytecode does not match: %s", to.Hex())
	return in, []byte{}
}
