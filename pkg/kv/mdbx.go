package kv

import (
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/share"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

type MdbxDB struct {
	env    *mdbx.Env
	path   string
	tables map[string]mdbx.DBI
}

var schemas = []string{
	share.AccountsTbl,
	share.BodiesTbl,
	share.ContractsTbl,
	share.Erc1155Tbl,
	share.Erc20Tbl,
	share.Erc721Tbl,
	share.TxLookupTbl,
}

func NewMdbx(path string) *MdbxDB {
	env, err := mdbx.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	env.SetOption(mdbx.OptMaxDB, 1024)
	env.SetGeometry(-1, -1, 1<<37, 1<<30, -1, 1<<16)
	if err = env.Open(path, mdbx.Create, 0766); err != nil {
		log.Fatal(err)
	}

	d := &MdbxDB{
		path:   path,
		tables: make(map[string]mdbx.DBI),
	}
	d.env = env

	// init all tables
	env.Update(func(txn *mdbx.Txn) error {
		for _, name := range schemas {
			dbi, err := txn.CreateDBI(name)
			if err != nil {
				log.Fatal(err)
			}
			d.tables[name] = dbi
		}
		return nil
	})

	return nil
}
