package types

import "strconv"

type Pager struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

func (p *Pager) Complete() {
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 10
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
		licenseType, _ := strconv.Atoi(r.LicenseType[0])
		resp.LicenseType = licenseType
	}
	if r.SourceCode != nil && len(r.SourceCode) > 0 {
		resp.SourceCode = r.SourceCode[0]
	}
	if r.Optimization != nil && len(r.Optimization) > 0 {
		optimization, _ := strconv.Atoi(r.Optimization[0])
		resp.Optimization = optimization
	}
	if r.Runs != nil && len(r.Runs) > 0 {
		runs, _ := strconv.Atoi(r.Runs[0])
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
	LicenseType      int    `json:"licenseType"` // int
	SourceCode       string `json:"sourceCode"`
	Optimization     int    `json:"optimization"` // bool
	Runs             int    `json:"runs"`         // int
	EVMVersion       string `json:"evmVersion"`   // 默认：default
}

const (
	SoliditySingleFile        = "solidity-single-file"
	SolidityStandardJsonInput = "solidity-standard-json-input"
)
