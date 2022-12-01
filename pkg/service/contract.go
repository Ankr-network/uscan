package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/xiaobaiskill/solc-go"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func WriteValidateContractMetadata(metadata *types.ValidateContractMetadata) error {
	return store.WriteValidateContractMetadata(metadata)
}

func ReadValidateContractMetadata() (*types.ValidateContractMetadata, error) {
	return store.GetValidateContractMetadata()
}

func ValidateContract(req *types.ValidateContractReq) (map[string]string, error) {
	if req.ContractAddress == "" {
		response.ErrVerityContract.Msg = "contract address cannot be empty"
		return nil, response.ErrVerityContract
	}

	_, err := store.GetContract(common.HexToAddress(req.ContractAddress))
	if err != nil {
		if err == kv.NotFound {
			return nil, response.ErrRecordNotFind
		}
		return nil, err
	}
	hash, err := store.GetValidateContract(common.HexToAddress(req.ContractAddress))
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	if hash != nil && hash.CodeHash != "" {
		response.ErrVerityContract.Msg = "the contract has been verified"
		return nil, response.ErrVerityContract
	}

	if req.CompilerType == "" {
		response.ErrVerityContract.Msg = "compiler type cannot be empty"
		return nil, response.ErrVerityContract
	}

	if req.CompilerVersion == "" {
		response.ErrVerityContract.Msg = "compiler version cannot be empty"
		return nil, response.ErrVerityContract
	}

	if req.LicenseType == 0 {
		response.ErrVerityContract.Msg = "license type cannot be empty"
		return nil, response.ErrVerityContract
	}

	if req.SourceCode == "" {
		response.ErrVerityContract.Msg = "contract source code cannot be empty"
		return nil, response.ErrVerityContract
	}

	body := &types.ContractVerityTmp{
		Address:          req.ContractAddress,
		ContractName:     req.ContractName,
		CompilerType:     req.CompilerType,
		CompilerVersion:  req.CompilerVersion,
		CompilerFileName: req.CompilerFileName,
		LicenseType:      req.LicenseType,
		SourceCode:       req.SourceCode,
		Optimization:     req.Optimization,
		Runs:             req.Runs,
		EVMVersion:       req.EVMVersion,
	}

	switch req.CompilerType {
	case types.SolidityStandardJsonInput:
		contractFileName := strings.Split(req.ContractName, ":")
		if len(contractFileName) != 2 || req.ContractName == "" {
			response.ErrVerityContract.Msg = "when license type is solidity-standard-json-input, contractName example: contracts/BlindBox.sol:BlindBox"
			return nil, response.ErrVerityContract
		}
	case types.SoliditySingleFile:
		if req.ContractName == "" {
			response.ErrVerityContract.Msg = "when license type is solidity-single-file, contractName example: BlindBox"
			return nil, response.ErrVerityContract
		}
	default:
		response.ErrVerityContract.Msg = "contractName error"
		return nil, response.ErrVerityContract
	}

	if err := store.WriteValidateContractStatus(common.HexToAddress(req.ContractAddress), big.NewInt(0)); err != nil {
		return nil, err
	}

	go StartContractVerity(body)
	return map[string]string{"id": body.Address}, nil
}

func validateContract(param *types.ContractVerityTmp) error {
	input := &solc.Input{}
	switch param.CompilerType {
	case types.SoliditySingleFile:
		enabled := false
		if param.Optimization == 1 {
			enabled = true
		}
		optimizer := solc.Optimizer{
			Enabled: enabled,
			Runs:    int(param.Runs),
		}
		settings := solc.Settings{
			Optimizer:  optimizer,
			EVMVersion: "",
			OutputSelection: map[string]map[string][]string{
				"*": map[string][]string{
					"*": []string{
						"abi",
						"evm.deployedBytecode",
						"evm.methodIdentifiers",
						"metadata",
					},
				},
			},
		}
		input = &solc.Input{
			Language: "Solidity",
			Sources: map[string]solc.SourceIn{
				"contract.sol": {
					Content: param.SourceCode,
				},
			},
			Settings: settings,
		}
		param.ContractName = "contract.sol:" + param.ContractName
	case types.SolidityStandardJsonInput:
		if err := json.Unmarshal([]byte(param.SourceCode), &input); err != nil {
			return err
		}
	}

	filePath := getSolcFile(param.CompilerFileName)
	logrus.Infof("getSolcFile:%s", filePath)
	newSolc := solc.NewSolc(filePath)
	out, err := newSolc.Compile(input)
	if err != nil {
		return errors.New(fmt.Sprintf("validateContract err: %+v, filePath: %+v", err, filePath))
	}
	if out != nil && out.Errors != nil {
		return errors.New(fmt.Sprintf("validateContract out.Errors: %+v, filePath: %+v", out.Errors, filePath))
	}
	abi := make([]json.RawMessage, 0)
	object := ""

	metadata := make(map[string]string)
	contractFileName := strings.Split(param.ContractName, ":")
	if len(contractFileName) != 2 {
		return fmt.Errorf("validateContract strings.Split contractName error. contractName: %s", param.ContractName)
	}
	contract, ok := out.Contracts[contractFileName[0]]
	if !ok {
		return fmt.Errorf("out.Contracts get %+v  error.", contractFileName[0])
	}
	v, ok := contract[contractFileName[1]]
	if !ok {
		return fmt.Errorf("contract get %+v  error.", contractFileName[1])
	}

	abi = v.ABI
	object = v.EVM.Bytecode.Object

	bytecodeObject, err := hexutil.Decode("0x" + v.EVM.Bytecode.Object)
	if err != nil {
		return err
	}
	deployedBytecodeObject, err := hexutil.Decode("0x" + v.EVM.DeployedBytecode.Object)
	if err != nil {
		return err
	}

	splitOp := deployedBytecodeObject[len(deployedBytecodeObject)-32:]
	var objectByte []byte
	res := bytes.Split(bytecodeObject, splitOp)
	if len(res) == 2 {
		objectByte = append(res[0], splitOp...)
	}

	switch param.CompilerType {
	case types.SoliditySingleFile:
		metadata[contractFileName[0]] = param.SourceCode
	case types.SolidityStandardJsonInput:
		inputTmp := &solc.Input{}
		if err := json.Unmarshal([]byte(param.SourceCode), &inputTmp); err != nil {
			return err
		}

		var inputMetadata solc.Input
		if err := json.Unmarshal([]byte(v.Metadata), &inputMetadata); err != nil {
			return err
		}
		for k, _ := range inputMetadata.Sources {
			metadata[k] = inputTmp.Sources[k].Content
		}
	}
	metadataMarshal, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	abiStr, err := json.Marshal(abi)
	if err != nil {
		return err
	}

	account, err := store.GetContract(common.HexToAddress(param.Address))
	if err != nil {
		return err
	}

	codeHash := ""
	if crypto.Keccak256Hash(objectByte).Hex() == account.ByteCodeHash.Hex() {
		codeHash = hexutil.Encode(account.ByteCode)
	}

	if codeHash != "" {
		for k, v := range v.EVM.MethodIdentifiers {
			err := store.WriteMethodName(v, k)
			if err != nil {
				return err
			}
		}
		if err := store.WriteValidateContract(common.HexToAddress(param.Address), &types.ContractVerity{
			ContractName:    param.ContractName,
			CompilerVersion: param.CompilerVersion,
			Optimization:    param.Optimization,
			Runs:            param.Runs,
			EVMVersion:      param.EVMVersion,
			LicenseType:     param.LicenseType,
			ABI:             string(abiStr),
			Metadata:        string(metadataMarshal),
			Object:          object,
			CodeHash:        codeHash,
		}); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("varity contract failed")
}

func getSolcFile(compilerFileName string) string {
	return fmt.Sprintf("%s%s", "./pkg/files/", compilerFileName)
}

var ContractVerityChain = make(chan *types.ContractVerityTmp, 100)

func StartContractVerity(body *types.ContractVerityTmp) {
	ContractVerityChain <- body
}

func StartHandleContractVerity() {
	go func() {
		for {
			select {
			case contractVerityTmp := <-ContractVerityChain:
				err := validateContract(contractVerityTmp)
				if err != nil {
					logrus.Errorf("StartHandleContractVerity validateContract error. err: %+v, contract verity id:%s", err, contractVerityTmp.Address)
					if err := store.WriteValidateContractStatus(common.HexToAddress(contractVerityTmp.Address), big.NewInt(2)); err != nil {
						logrus.Errorf("StartHandleContractVerity UpdateContractVerityTmpStatus error. err: %+v, contract verity id:%s", err, contractVerityTmp.Address)
					}
				} else {
					logrus.Errorf("StartHandleContractVerity validateContract error. err: %+v, contract verity id:%s", err, contractVerityTmp.Address)
					if err := store.WriteValidateContractStatus(common.HexToAddress(contractVerityTmp.Address), big.NewInt(1)); err != nil {
						logrus.Errorf("StartHandleContractVerity UpdateContractVerityTmpStatus error. err: %+v, contract verity id:%s", err, contractVerityTmp.Address)
					}
				}
			}
		}
	}()
}

func GetValidateContractStatus(address string) (int64, error) {
	status, err := store.GetValidateContractStatus(common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	return status.Int64(), nil
}

func GetValidateContract(address common.Address) (*types.ContractVerityInfoResp, error) {
	resp := &types.ContractVerityInfoResp{}
	contract, err := store.GetValidateContract(address)
	if err != nil && err != kv.NotFound {
		return nil, err
	}
	if contract != nil {
		metadata := make(map[string]string)
		err := json.Unmarshal([]byte(contract.Metadata), &metadata)
		if err != nil {
			return nil, err
		}
		resp.Contract = &types.ContractVerityInfo{
			ContractName:    contract.ContractName,
			CompilerVersion: contract.CompilerVersion,
			Optimization:    contract.Optimization,
			Runs:            contract.Runs,
			EVMVersion:      contract.EVMVersion,
			LicenseType:     contract.LicenseType,
			ABI:             contract.ABI,
			Metadata:        metadata,
			Object:          contract.Object,
		}
		proxyContractAddress, err := store.GetProxyContract(address)
		if err != nil && err != kv.NotFound {
			return nil, err
		}
		nullAddress := common.Address{}
		if proxyContractAddress.String() != nullAddress.String() {
			resp.ProxyContractAddress = proxyContractAddress.String()
		}
		proxyContract, err := store.GetValidateContract(proxyContractAddress)
		if err != nil && err != kv.NotFound {
			return nil, err
		}
		if proxyContract != nil {
			metadata := make(map[string]string)
			err := json.Unmarshal([]byte(contract.Metadata), &metadata)
			if err != nil {
				return nil, err
			}
			resp.ProxyContract = &types.ContractVerityInfo{
				ContractName:    contract.ContractName,
				CompilerVersion: contract.CompilerVersion,
				Optimization:    contract.Optimization,
				Runs:            contract.Runs,
				EVMVersion:      contract.EVMVersion,
				LicenseType:     contract.LicenseType,
				ABI:             contract.ABI,
				Metadata:        metadata,
				Object:          contract.Object,
			}
		}
	}

	return resp, nil
}

func ReadMetaData() (*types.ValidateContractMetadata, error) {
	jsonFile, err := os.Open("/app/pkg/files/metadata.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var metadata *types.ValidateContractMetadata
	err = json.Unmarshal(byteValue, &metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func WriteMetadata() error {
	data, err := ReadMetaData()
	if err != nil {
		return err
	}
	err = store.WriteValidateContractMetadata(data)
	if err != nil {
		return err
	}
	return nil
}
