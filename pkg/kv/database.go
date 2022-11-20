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
)

type ReadOption struct {
	Table      string
	Latest     bool
	NextIfNone bool
}

type WriteOption struct {
	Table string
}

type Writer interface {
	Put(ctx context.Context, key, val []byte, opts *WriteOption) error
	Del(ctx context.Context, key []byte, opts *WriteOption) error
}

type Reader interface {
	Has(ctx context.Context, key []byte, opts *ReadOption) (bool, error)
	Get(ctx context.Context, key []byte, opts *ReadOption) ([]byte, error)
}
type Closer interface {
	Close() error
}
type Transactioner interface {
	BeginTx(context.Context) (context.Context, error)
	Commit(context.Context)
	RollBack(context.Context)
}

type Sorter interface {
	SPut(ctx context.Context, key, val []byte, opts *WriteOption) error
	SDel(ctx context.Context, key, val []byte, opts *WriteOption) error
	SCount(ctx context.Context, key []byte, opts *ReadOption) (uint64, error)
	SGet(ctx context.Context, key []byte, offset, limit uint64, opts *ReadOption) ([][]byte, error)
}

type Database interface {
	Transactioner
	Writer
	Reader
	Sorter
	Closer
}
