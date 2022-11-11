package rpcclient

// Config are the configuration options for structured logger the EVM
type TracerConfig struct {
	EnableMemory     bool   `json:"enableMemory"`     // enable memory capture
	DisableStack     bool   `json:"disableStack"`     // disable stack capture
	DisableStorage   bool   `json:"disableStorage"`   // disable storage capture
	EnableReturnData bool   `json:"enableReturnData"` // enable return data capture
	Tracer           string `json:"tracer,omitempty"` //
	Timeout          string `json:"timeout,omitempty"`
}
