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

func (d *MdbxDB) SGet(ctx context.Context, key []byte, offset, limit uint64, opts *kv.ReadOption) (rs [][]byte, err error) {
	d.env.View(func(txn *mdbx.Txn) error {
		c, err := txn.OpenCursor(d.tables[opts.Table])
		if err != nil {
			return err
		}
		defer c.Close()
		var k, v []byte
		k, v, err = c.Get(key, nil, mdbx.Set)
		if err != nil {
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
			}
			return err
		}
		count, err := c.Count()
		if err != nil {
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
			}
			return err
		}
		k, v, err = c.Get(k, v, mdbx.LastDup)
		if err != nil {
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
			}
			return err
		}
		rs = make([][]byte, 0)
		begin := offset + 1
		end := offset + limit
		var i uint64
		i = 1
		for k, v, err := c.Get(k, v, mdbx.GetCurrent); k != nil && err == nil; k, v, err = c.Get(nil, nil, mdbx.PrevDup) {
			if i >= begin {
				rs = append(rs, v)
			}
			if i == end || i == count {
				break
			}
			i++
		}
		return err
	})

	return
}

func (d *MdbxDB) SCount(ctx context.Context, key []byte, opts *kv.ReadOption) (count uint64, err error) {
	d.env.View(func(txn *mdbx.Txn) error {
		c, err := txn.OpenCursor(d.tables[opts.Table])
		if err != nil {
			return err
		}
		defer c.Close()
		_, _, err = c.Get(key, nil, mdbx.Set)
		if err != nil {
			if mdbx.IsNotFound(err) {
				err = kv.NotFound
			}
			return err
		}
		count, err = c.Count()
		return err
	})
	return
}
