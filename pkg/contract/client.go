package contract

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Ankr-network/uscan/pkg/contract/eip"
	log "github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	eip721Inf          = "0x80ac58cd" // ieip721
	eip721MetadataInf  = "0x5b5e139f"
	eip1155Inf         = "0xd9b67a26" // ieip1155
	eip1155MetadataInf = "0x0e89341c"
)

var (
	TransferBatchEventTopic  = common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb")
	TransferSingleEventTopic = common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62")
	TransferEventTopic       = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

	ErrinvalidTopic   = errors.New("invalid topic")
	ErrNotNftContract = errors.New("non NFT contract")
)

type Client struct {
	client         rpcclient.RpcClient
	eip1155Cache   *utils.Cache
	eip721Cache    *utils.Cache
	eip1155Abi     *abi.ABI
	eip1155MetaAbi *abi.ABI
	eip721Abi      *abi.ABI
	eip721MetaAbi  *abi.ABI
}

func NewClient(client rpcclient.RpcClient) *Client {
	eip1155Abi, err := abi.JSON(strings.NewReader(eip.Ieip1155ABI))
	if err != nil {
		log.Fatalf("get abi from eip1155: %v", err)
	}
	eip1155MetaAbi, err := abi.JSON(strings.NewReader(eip.Ieip1155metaABI))
	if err != nil {

		log.Fatalf("get abi from eip1155meta: %v", err)
	}

	eip721Abi, err := abi.JSON(strings.NewReader(eip.Ieip721ABI))
	if err != nil {
		log.Fatalf("get abi from eip721: %v", err)
	}
	eip721MetaAbi, err := abi.JSON(strings.NewReader(eip.Ieip721metaABI))
	if err != nil {

		log.Fatalf("get abi from eip721: %v", err)
	}

	return &Client{
		client:         client,
		eip1155Cache:   utils.NewCache(),
		eip721Cache:    utils.NewCache(),
		eip1155Abi:     &eip1155Abi,
		eip1155MetaAbi: &eip1155MetaAbi,
		eip721Abi:      &eip721Abi,
		eip721MetaAbi:  &eip721MetaAbi,
	}
}

func (e *Client) IsEIP721(contract string) bool {
	contract = strings.ToLower(contract)
	if value, ok := e.eip721Cache.Get(contract); ok {
		return value.(bool)
	}
	ctr, err := eip.NewIeip165Caller(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {

		log.Errorf("NewIeip165Caller failed: %v; contract:%s", err, contract)
		return false
	}
	var data [4]byte
	copy(data[:], hexutil.MustDecode(eip721Inf)[0:4])
	eip721Res, eip721Err := ctr.SupportsInterface(nil, data)
	if eip721Err != nil {
		e.eip721Cache.Add(contract, false)
		log.Errorf("Check whether the contract complies with the EIP721 standard: %v; contract:%s", eip721Err, contract)
		return false
	}

	copy(data[:], hexutil.MustDecode(eip721MetadataInf)[0:4])
	eip721MetadataRes, eip721MetadataErr := ctr.SupportsInterface(nil, data)
	if eip721MetadataErr != nil {
		e.eip721Cache.Add(contract, false)
		log.Errorf("Check whether the contract complies with the EIP721Metadata standard: %v; contract:%s", eip721MetadataErr, contract)
		return false
	}
	res := eip721Res && eip721MetadataRes
	e.eip721Cache.Add(contract, res)

	return res
}

func (e *Client) GetEIP721Meta(contract string, tokenID string) (string, error) {
	ctr, err := eip.NewIeip721metaCaller(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		log.Errorf("NewIeip721metaCaller failed: %v; contract:%s", err, contract)
		return "", err
	}
	id, _ := hexutil.DecodeBig(tokenID)
	return ctr.TokenURI(nil, id)
}

func (e *Client) IsEIP1155(contract string) bool {
	contract = strings.ToLower(contract)
	if value, ok := e.eip1155Cache.Get(contract); ok {
		return value.(bool)
	}

	ctr, err := eip.NewIeip165Caller(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		log.Errorf("NewIeip165Caller failed: %v; contract:%s", err, contract)
		return false
	}
	var data [4]byte
	copy(data[:], hexutil.MustDecode(eip1155Inf)[0:4])
	eip1155Res, eip1155Err := ctr.SupportsInterface(nil, data)
	if eip1155Err != nil {
		e.eip1155Cache.Add(contract, false)
		log.Errorf("Check whether the contract complies with the EIP1155 standard: %v; contract:%s", eip1155Err, contract)
		return false
	}

	copy(data[:], hexutil.MustDecode(eip1155MetadataInf)[0:4])
	eip1155MetadataEes, eip1155MetadataErr := ctr.SupportsInterface(nil, data)

	if eip1155MetadataErr != nil {
		e.eip1155Cache.Add(contract, false)
		log.Errorf("Check whether the contract complies with the EIP1155Metadata standard: %v; contract:%s", eip1155MetadataErr, contract)
		return false
	}

	res := eip1155Res && eip1155MetadataEes
	e.eip1155Cache.Add(contract, res)
	return res
}

func (e *Client) GetEIP1155Meta(contract string, tokenID string) (string, error) {
	ctr, err := eip.NewIeip1155metaCaller(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		log.Errorf("NewIeip1155metaCaller failed: %v; contract:%s", err, contract)
		return "", err
	}
	id, _ := hexutil.DecodeBig(tokenID)
	return ctr.Uri(nil, id)
}

func (e *Client) Erc20Transfer(contract string, log types.Log) (*eip.Erc20Transfer, error) {
	filter, err := eip.NewErc20Filterer(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return nil, err
	}
	res, err := filter.ParseTransfer(log)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Client) Erc721Transfer(contract string, log types.Log) (*eip.Ieip721Transfer, error) {
	filter, err := eip.NewIeip721Filterer(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return nil, err
	}
	res, err := filter.ParseTransfer(log)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Client) Erc1155TransferSingle(contract string, log types.Log) (*eip.Ieip1155TransferSingle, error) {
	filter, err := eip.NewIeip1155Filterer(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return nil, err
	}
	res, err := filter.ParseTransferSingle(log)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Client) Erc1155TransferBatch(contract string, log types.Log) (*eip.Ieip1155TransferBatch, error) {
	filter, err := eip.NewIeip1155Filterer(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return nil, err
	}
	res, err := filter.ParseTransferBatch(log)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Client) GetContractName(contract string) (string, error) {
	meta, err := eip.NewMeta(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return "", err
	}
	name, err := meta.Name(nil)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (e *Client) GetContractSymbol(contract string) (string, error) {
	meta, err := eip.NewMeta(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return "", err
	}
	symbol, err := meta.Symbol(nil)
	if err != nil {
		return "", err
	}
	return symbol, nil
}

func (e *Client) GetContractDecimals(contract string) (*big.Int, error) {
	meta, err := eip.NewMeta(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return nil, err
	}
	decimals, err := meta.Decimals(nil)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(decimals)), nil
}

func (e *Client) GetContractTotalSupply(contract string) (*big.Int, error) {
	meta, err := eip.NewMeta(common.HexToAddress(contract), e.client.GetClient())
	if err != nil {
		return big.NewInt(0), err
	}
	totalSupply, err := meta.TotalSupply(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return totalSupply, nil
}

func (e *Client) GetNumWith721ByContactOwnerTokenID(owner, contract common.Address, tokenID *big.Int, block *big.Int) (*big.Int, error) {
	ctr, err := eip.NewIeip721Caller(contract, e.client.GetClient())
	if err != nil {
		log.Errorf("NewIeip721Caller failed: %v; contract:%s; owner: %s; tokenId: %s", err, contract.Hex(), owner.Hex(), tokenID.String())
		return nil, err
	}
	from, err := ctr.OwnerOf(&bind.CallOpts{
		BlockNumber: block,
	}, tokenID)
	if err != nil {
		log.Errorf("eip721 OwnerOf: %v; contract:%s; owner: %s; tokenId: %s", err, contract.Hex(), owner.Hex(), tokenID.String())
		return nil, err
	}
	if from == owner {
		return big.NewInt(1), nil
	}
	return big.NewInt(0), nil
}

func (e *Client) GetNumWith1155ByContactOwnerTokenID(owner, contract common.Address, tokenID *big.Int, block *big.Int) (*big.Int, error) {
	ctr, err := eip.NewIeip1155(contract, e.client.GetClient())
	if err != nil {
		log.Errorf("NewIeip1155Caller OwnerOf: %v; contract:%s; owner: %s; tokenId: %s", err, contract.Hex(), owner.Hex(), tokenID.String())
		return nil, err
	}
	amount, err := ctr.BalanceOf(&bind.CallOpts{
		BlockNumber: block,
	}, owner, tokenID)
	if err != nil {
		log.Errorf("eip1155 BalanceOf: %v; contract:%s; owner: %s; tokenId: %s", err, contract.Hex(), owner.Hex(), tokenID.String())
		return nil, err
	}
	return amount, nil
}
