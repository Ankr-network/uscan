/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package kv

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/share"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

var _ Database = (*MdbxDB)(nil)

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

func (d *MdbxDB) Begin(context.Context) context.Context {
	return nil
}
func (d *MdbxDB) Commit(context.Context)   {}
func (d *MdbxDB) RollBack(context.Context) {}

func (d *MdbxDB) Put(ctx context.Context, key []byte, val []byte, opts *WriteOption) error {
	return d.env.Update(func(txn *mdbx.Txn) error {
		return txn.Put(d.tables[opts.Table], key, val, mdbx.Upsert)
	})
}

func (d *MdbxDB) Get(ctx context.Context, key []byte, opts *ReadOption) (rs []byte, err error) {
	d.env.View(func(txn *mdbx.Txn) error {
		rs, err = txn.Get(d.tables[opts.Table], key)
		if mdbx.IsNotFound(err) {
			err = NotFound
		}
		return nil
	})

	return
}

func (d *MdbxDB) Has(key []byte, opts *ReadOption) (rs bool, err error) {
	d.env.View(func(txn *mdbx.Txn) error {
		var res []byte
		res, err = txn.Get(d.tables[opts.Table], key)
		if mdbx.IsNotFound(err) {
			return nil
		}
		if err != nil {
			if len(res) != 0 {
				rs = true
			}
		}
		return nil
	})

	return
}

func (d *MdbxDB) Close() error {
	for _, v := range d.tables {
		d.env.CloseDBI(v)
	}
	d.env.Close()
	return nil
}
