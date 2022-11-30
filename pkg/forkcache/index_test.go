package forkcache

import (
	"github.com/Ankr-network/uscan/pkg/field"
	"testing"
)

func TestIndex(t *testing.T) {
	var indexMap = make(map[string]*field.BigInt, 0)
	if indexMap["x"] == nil {
		indexMap["x"] = field.NewInt(1)
	}
	indexMap["x"].Add(field.NewInt(1))
	t.Log(indexMap["x"].Bytes())
}
