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

type ReadOption struct {
	Table      string
	Latest     bool
	NextIfNone bool
}

type WriteOption struct {
	Table string
}

type Putter interface {
	Put(key, val []byte, opts *WriteOption) error
}

type Getter interface {
	Get(key []byte, opts *ReadOption) ([]byte, error)
}

type Closer interface {
	Close() error
}

type Database interface {
	Putter
	Getter
	Closer
}
