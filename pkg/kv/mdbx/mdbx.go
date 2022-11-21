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
package mdbx

import (
	"context"
	"runtime"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/share"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

var _ kv.Database = (*MdbxDB)(nil)

type txKey struct{}

type MdbxDB struct {
	env    *mdbx.Env
	path   string
	tables map[string]mdbx.DBI
}

var schemas = []string{
	share.AccountsTbl,
	share.HomeTbl,
	share.TxTbl,
	share.BlockTbl,
	share.TraceLogTbl,
	share.TransferTbl,
	share.HolderTbl,
	share.ValidateContractTbl,
}

var schemasSort = []string{
	share.HolderSortTabl,
	share.InventorySortTabl,
}

var DB *MdbxDB

func NewDB(path string) {
	DB = NewMdbx(path)
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

		for _, name := range schemasSort {
			dbi, err := txn.OpenDBI(name, mdbx.Create|mdbx.DupSort, nil, nil)
			if err != nil {
				log.Fatal(err)
			}
			d.tables[name] = dbi
		}
		return nil
	})

	return d
}

func (d *MdbxDB) BeginTx(ctx context.Context) (context.Context, error) {
	runtime.LockOSThread()
	tnx, err := d.env.BeginTxn(nil, 0)
	if err != nil {
		runtime.UnlockOSThread()
		return nil, err
	}

	return context.WithValue(ctx, txKey{}, tnx), nil
}

func (d *MdbxDB) Commit(ctx context.Context) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		out.Commit()
	}
	runtime.UnlockOSThread()
}

func (d *MdbxDB) RollBack(ctx context.Context) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		out.Abort()
	}
	runtime.UnlockOSThread()
}

func (d *MdbxDB) Put(ctx context.Context, key []byte, val []byte, opts *kv.WriteOption) error {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		return out.Put(d.tables[opts.Table], key, val, mdbx.Upsert)
	} else {
		return d.env.Update(func(txn *mdbx.Txn) error {
			return txn.Put(d.tables[opts.Table], key, val, mdbx.Upsert)
		})
	}
}

func (d *MdbxDB) Del(ctx context.Context, key []byte, opts *kv.WriteOption) (err error) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		err = out.Del(d.tables[opts.Table], key, nil)
		if err != nil && mdbx.IsNotFound(err) {
			err = nil
		}
	} else {
		return d.env.Update(func(txn *mdbx.Txn) error {
			err = txn.Del(d.tables[opts.Table], key, nil)
			if err != nil && mdbx.IsNotFound(err) {
				err = nil
			}
			return err
		})
	}
	return
}

func (d *MdbxDB) Get(ctx context.Context, key []byte, opts *kv.ReadOption) (rs []byte, err error) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		rs, err = out.Get(d.tables[opts.Table], key)
		if mdbx.IsNotFound(err) {
			err = kv.NotFound
		}
	} else {
		d.env.View(func(txn *mdbx.Txn) error {
			rs, err = txn.Get(d.tables[opts.Table], key)
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
			}
			return nil
		})
	}

	return
}

func (d *MdbxDB) Has(ctx context.Context, key []byte, opts *kv.ReadOption) (rs bool, err error) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		var res []byte
		res, err = out.Get(d.tables[opts.Table], key)
		if mdbx.IsNotFound(err) {
			err = kv.NotFound
			return
		}
		if err == nil {
			if len(res) != 0 {
				rs = true
			}
		}
	} else {
		d.env.View(func(txn *mdbx.Txn) error {
			var res []byte
			res, err = txn.Get(d.tables[opts.Table], key)
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
				return nil
			}
			if err == nil {
				if len(res) != 0 {
					rs = true
				}
			}
			return nil
		})
	}

	return
}

func (d *MdbxDB) Close() error {
	for _, v := range d.tables {
		d.env.CloseDBI(v)
	}
	d.env.Close()
	return nil
}
