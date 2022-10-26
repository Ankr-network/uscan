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
package rawdb

import (
	"math/big"
	"sync"

	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/utils"
	"github.com/Ankr-network/uscan/share"
	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/ethereum/go-ethereum/common"
)

var (
	bmpool = sync.Pool{
		New: func() interface{} {
			return roaring64.New()
		},
	}
)

func ReadBody(db kv.Getter, number *big.Int) ([]common.Hash, error) {
	bm := bmpool.Get().(*roaring64.Bitmap)
	defer bmpool.Put(bm)
	bm.Clear()

	bs, err := db.Get(utils.EncodeVarint(number.Uint64()), &kv.ReadOption{Table: share.BodiesTbl})
	if err != nil {
		return nil, err
	}

	err = bm.UnmarshalBinary(bs)
	if err != nil {
		return nil, err
	}

	bna := bm.ToArray()
	rs := make([]common.Hash, len(bna))

	for i, n := range bna {
		rs[i] = common.BigToHash(big.NewInt(0).SetUint64(n))
	}

	return rs, nil
}

func WriteBody(db kv.Putter, number *big.Int, hs []common.Hash) error {

	bm := bmpool.Get().(*roaring64.Bitmap)
	defer bmpool.Put(bm)
	bm.Clear()

	for _, h := range hs {
		bm.Add(h.Big().Uint64())
	}
	hbs, _ := bm.ToBytes()
	return db.Put(utils.EncodeVarint(number.Uint64()), hbs, &kv.WriteOption{Table: share.BodiesTbl})
}
