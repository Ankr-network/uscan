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

// import (
// 	"github.com/Ankr-network/uscan/pkg/field"
// 	"github.com/Ankr-network/uscan/pkg/kv"
// 	"github.com/Ankr-network/uscan/pkg/types"
// 	"github.com/Ankr-network/uscan/share"
// 	"github.com/ethereum/go-ethereum/common"
// )

// // demo
// func ReadAccountBalance(db kv.Getter, addr common.Address) (uint64, error) {
// 	bs, err := db.Get(addr.Bytes(), &kv.ReadOption{Table: share.AccountsTbl})
// 	if err != nil {
// 		return 0, err
// 	}
// 	a := &types.Account{}
// 	a.Unmarshal(bs)
// 	return a.Balance.ToUint64(), nil
// }

// // demo
// func WriteAccountBalance(db kv.Database, addr common.Address, balance uint64) error {
// 	bs, err := db.Get(addr.Bytes(), &kv.ReadOption{Table: share.AccountsTbl})
// 	if err != nil {
// 		return err
// 	}
// 	a := &types.Account{}
// 	a.Unmarshal(bs)
// 	a.Balance = field.NewInt(int64(balance))
// 	abs, _ := a.Marshal()
// 	return db.Put(addr.Bytes(), abs, &kv.WriteOption{Table: share.AccountsTbl})
// }
