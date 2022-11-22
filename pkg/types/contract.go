package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Contract struct {
	Owner                 common.Address `rlp:"-"`
	ByteCode              []byte
	ByteCodeHash          common.Hash
	ConstructorArguements []byte
	DeployedCode          []byte
}

func (b *Contract) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *Contract) Unmarshal(bin []byte) error {
	return rlp.DecodeBytes(bin, &b)
}

type CompilerVersion struct {
	Name     string `json:"name"`
	FileName string `json:"fileName"`
}

type LicenseType struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

type ValidateContractMetadata struct {
	CompilerVersions []*CompilerVersion `json:"compilerVersions"`
	LicenseTypes     []*LicenseType     `json:"licenseTypes"`
	CompilerTypes    []string           `json:"compilerTypes"`
}

func (b *ValidateContractMetadata) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *ValidateContractMetadata) Unmarshal(bin []byte) (err error) {
	return rlp.DecodeBytes(bin, &b)
}

type ContractVerityTmp struct {
	//ID      string
	Address          string
	ContractName     string
	CompilerType     string
	CompilerVersion  string
	CompilerFileName string
	LicenseType      uint64
	SourceCode       string
	Optimization     uint64
	Runs             uint64
	EVMVersion       string // 当前默认default
	Status           int    // 0:handling  1 success 2 fail
}

func (b *ContractVerityTmp) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *ContractVerityTmp) Unmarshal(bin []byte) (err error) {
	return rlp.DecodeBytes(bin, &b)
}

type ContractVerity struct {
	ContractName    string `json:"contractName"`
	CompilerVersion string `json:"compilerVersion"`
	Optimization    uint64 `json:"optimization"`
	Runs            uint64 `json:"runs"`
	EVMVersion      string `gorm:"column:evm_version"`
	LicenseType     uint64 `json:"licenseType"`
	ABI             string `json:"abi"`
	Metadata        string `json:"metadata"`
	CodeHash        string `json:"codeHash" gorm:"uniqueIndex:code_hash,priority:1;type:varchar(255)"`
	Object          string `json:"object"`
}

func (b *ContractVerity) Marshal() ([]byte, error) {
	return rlp.EncodeToBytes(b)
}

func (b *ContractVerity) Unmarshal(bin []byte) (err error) {
	return rlp.DecodeBytes(bin, &b)
}
