package contract

import (
	"math/big"

	"github.com/Ankr-network/uscan/pkg/contract/eip"
	"github.com/ethereum/go-ethereum/core/types"
)

type Contractor interface {
	GetEIP721Meta(contract string, tokenID string) (uri string, err error)
	GetEIP1155Meta(contract string, tokenID string) (uri string, err error)
	IsEIP721(contract string) bool
	IsEIP1155(contract string) bool
	Erc20Transfer(contract string, log types.Log) (*eip.Erc20Transfer, error)
	Erc721Transfer(contract string, log types.Log) (*eip.Ieip721Transfer, error)
	Erc1155TransferSingle(contract string, log types.Log) (*eip.Ieip1155TransferSingle, error)
	Erc1155TransferBatch(contract string, log types.Log) (*eip.Ieip1155TransferBatch, error)
	GetContractName(contract string) (string, error)
	GetContractSymbol(contract string) (string, error)
	GetContractDecimals(contract string) (*big.Int, error)
	GetContractTotalSupply(contract string) (*big.Int, error)
	//CheckLog(log *types.Log) (*model.EventTransferData, error)
}
