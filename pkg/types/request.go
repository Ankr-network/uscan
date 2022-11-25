package types

import "strconv"

type Pager struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

func (f *Pager) Complete() {
	if f.Offset <= 0 {
		f.Offset = 0
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Limit > 100 {
		f.Limit = 100
	}
}

type SearchFilter struct {
	Type    int    `query:"type"` // 1 all filter 2 address
	Keyword string `query:"keyword"`
}

type ValidateContractTmpReq struct {
	ContractAddress  []string `json:"contractAddress"`
	ContractName     []string `json:"contractName"`
	CompilerType     []string `json:"compilerType"` // solidity-single-file / solidity-standard-json-input
	CompilerVersion  []string `json:"compilerVersion"`
	CompilerFileName []string `json:"compilerFileName"`
	LicenseType      []string `json:"licenseType"` // int
	SourceCode       []string `json:"sourceCode"`
	Optimization     []string `json:"optimization"` // bool
	Runs             []string `json:"runs"`         // int
	EVMVersion       []string `json:"evmVersion"`   // 默认：default
}

func (r *ValidateContractTmpReq) ToValidateContractReq() *ValidateContractReq {
	resp := &ValidateContractReq{}
	if r.ContractAddress != nil && len(r.ContractAddress) > 0 {
		resp.ContractAddress = r.ContractAddress[0]
	}
	if r.ContractName != nil && len(r.ContractName) > 0 {
		resp.ContractName = r.ContractName[0]
	}

	if r.CompilerType != nil && len(r.CompilerType) > 0 {
		resp.CompilerType = r.CompilerType[0]
	}
	if r.CompilerVersion != nil && len(r.CompilerVersion) > 0 {
		resp.CompilerVersion = r.CompilerVersion[0]
	}
	if r.CompilerFileName != nil && len(r.CompilerFileName) > 0 {
		resp.CompilerFileName = r.CompilerFileName[0]
	}
	if r.LicenseType != nil && len(r.LicenseType) > 0 {
		licenseType, _ := strconv.ParseUint(r.LicenseType[0], 10, 64)
		resp.LicenseType = licenseType
	}
	if r.SourceCode != nil && len(r.SourceCode) > 0 {
		resp.SourceCode = r.SourceCode[0]
	}
	if r.Optimization != nil && len(r.Optimization) > 0 {
		optimization, _ := strconv.ParseUint(r.Optimization[0], 10, 64)
		resp.Optimization = optimization
	}
	if r.Runs != nil && len(r.Runs) > 0 {
		runs, _ := strconv.ParseUint(r.Runs[0], 10, 64)
		resp.Runs = runs
	}
	if r.EVMVersion != nil && len(r.EVMVersion) > 0 {
		resp.EVMVersion = r.EVMVersion[0]
	}
	return resp
}

type ValidateContractReq struct {
	ContractAddress  string `json:"contractAddress"`
	ContractName     string `json:"contractName"`
	CompilerType     string `json:"compilerType"` // solidity-single-file / solidity-standard-json-input
	CompilerVersion  string `json:"compilerVersion"`
	CompilerFileName string `json:"compilerFileName"`
	LicenseType      uint64 `json:"licenseType"` // int
	SourceCode       string `json:"sourceCode"`
	Optimization     uint64 `json:"optimization"` // bool
	Runs             uint64 `json:"runs"`         // int
	EVMVersion       string `json:"evmVersion"`   // 默认：default
}

const (
	SoliditySingleFile        = "solidity-single-file"
	SolidityStandardJsonInput = "solidity-standard-json-input"
)
