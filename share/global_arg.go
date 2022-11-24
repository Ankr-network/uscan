package share

import "time"

const (
	MaxChanSize  = 4096
	MaxCacheTime = 6 * time.Hour
	NeverExpired = 0
	HttpTimeout  = time.Minute
	Retry        = 3
	PartNum      = 100
	WriteTimeout = time.Minute
	ReadTimeout  = time.Minute
)
