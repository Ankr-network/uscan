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
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeVarint(t *testing.T) {
	var tbls = []struct {
		input uint64
		want  []byte
	}{
		{
			want:  []byte{1},
			input: 1,
		},
		{
			want:  []byte{111},
			input: 111,
		},
		{
			want:  []byte{169, 18},
			input: 2345,
		},
		{
			want:  []byte{135, 173, 75},
			input: 1234567,
		},
		{
			want:  []byte{130, 239, 133, 27},
			input: 56719234,
		},
	}

	for _, v := range tbls {
		bs := EncodeVarint(v.input)
		assert.Equal(t, v.want, bs, "equal")
	}

}

func TestDecodeVarint(t *testing.T) {

	var tbls = []struct {
		input []byte
		want  uint64
	}{
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{111},
			want:  111,
		},
		{
			input: []byte{169, 18},
			want:  2345,
		},
		{
			input: []byte{135, 173, 75},
			want:  1234567,
		},
		{
			input: []byte{130, 239, 133, 27},
			want:  56719234,
		},
	}

	for _, v := range tbls {
		n, _ := DecodeVarint(v.input)
		assert.Equal(t, v.want, n, "equal")
	}

}
