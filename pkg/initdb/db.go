package initdb

import (
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/kv/mdbx"
	"github.com/Ankr-network/uscan/share"
	"github.com/spf13/viper"
)

var DB kv.Database

func NewDB() {
	DB = mdbx.NewMdbx(viper.GetString(share.MdbxPath))
}
