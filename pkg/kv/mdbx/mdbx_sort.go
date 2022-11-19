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

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/torquem-ch/mdbx-go/mdbx"
)

var _ kv.Sorter = (*MdbxDB)(nil)

func (d *MdbxDB) SPut(ctx context.Context, key []byte, val []byte, opts *kv.WriteOption) error {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		return out.Put(d.tables[opts.Table], key, val, mdbx.Upsert)
	} else {
		return d.env.Update(func(txn *mdbx.Txn) error {
			return txn.Put(d.tables[opts.Table], key, val, mdbx.Upsert)
		})
	}
}

func (d *MdbxDB) SDel(ctx context.Context, key, val []byte, opts *kv.WriteOption) (err error) {
	out, ok := ctx.Value(txKey{}).(*mdbx.Txn)
	if ok {
		err = out.Del(d.tables[opts.Table], key, val)
		if mdbx.IsNotFound(err) {
			err = nil
		}
	} else {
		return d.env.Update(func(txn *mdbx.Txn) error {
			err = txn.Del(d.tables[opts.Table], key, val)
			if mdbx.IsNotFound(err) {
				err = nil
			}
			return err
		})
	}
	return
}

func (d *MdbxDB) SGet(ctx context.Context, key []byte, page, pageSize uint64, opts *kv.ReadOption) (rs [][]byte, err error) {
	// d.env.View(func(txn *mdbx.Txn) error {
	// 	c, err := txn.OpenCursor(d.tables[opts.Table])
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer c.Close()
	// 	_, _, err = c.Get(key, nil, mdbx.SetRange)
	// 	count, err = c.Count()
	// 	return err
	// })
	return nil, nil
}

func (d *MdbxDB) SCount(ctx context.Context, key []byte, opts *kv.ReadOption) (count uint64, err error) {
	d.env.View(func(txn *mdbx.Txn) error {
		c, err := txn.OpenCursor(d.tables[opts.Table])
		if err != nil {
			return err
		}
		defer c.Close()
		_, _, err = c.Get(key, nil, mdbx.SetRange)
		count, err = c.Count()
		return err
	})
	return
}
