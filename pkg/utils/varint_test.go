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
