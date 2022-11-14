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
	"context"

	"github.com/Ankr-network/uscan/pkg/field"
	"github.com/Ankr-network/uscan/pkg/kv"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/Ankr-network/uscan/share"
	"github.com/ethereum/go-ethereum/common"
)

var (
	addressKeyPrefix = []byte("/info/")
)

/*
table: accounts

/info/<address> => account info

/<address>/tx/total => num
/<address>/tx/<index> => <txhash>

/<address>/itx/total => num
/<address>/itx/<index> => InternalTxKey{txhash,index}

/<address>/erc20/total => num
/<address>/erc20/<index> => <index>(erc20 transfer index)

/<address>/erc721/total => num
/<address>/erc721/<index> => <index>(erc721 transfer index)

/<address>/erc1155/total => num
/<address>/erc1155/<index> => <index>(erc1155 transfer index)
*/

// ----------------- account info -----------------
func ReadAccount(ctx context.Context, db kv.Getter, addr common.Address) (acc *types.Account, err error) {
	var bytesRes []byte

	bytesRes, err = db.Get(ctx, append(addressKeyPrefix, addr.Bytes()...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return nil, err
	}
	acc = &types.Account{}
	err = acc.Unmarshal(bytesRes)
	return
}

func WriteAccount(ctx context.Context, db kv.Database, addr common.Address, acc *types.Account) error {
	bytesRes, err := acc.Marshal()
	if err != nil {
		return err
	}
	return db.Put(ctx, append(addressKeyPrefix, addr.Bytes()...), bytesRes, &kv.WriteOption{Table: share.AccountsTbl})
}

// ----------------- tx ----------------

func WriteAccountTxTotal(ctx context.Context, db kv.Putter, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/tx/total")...), total.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountTxTotal(ctx context.Context, db kv.Getter, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/tx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func WriteAccountTxIndex(ctx context.Context, db kv.Putter, addr common.Address, index *field.BigInt, hash common.Hash) error {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), hash.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountTxIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (hash common.Hash, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	hash.SetBytes(bytesRes)
	return
}

func ReadAccountTxByIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (tx *types.Tx, err error) {
	var hashByte []byte
	hashByte, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/tx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	hash := common.BytesToHash(hashByte)
	return ReadTx(ctx, db, hash)
}

// ------------ internal tx -------------

func WriteAccountITxTotal(ctx context.Context, db kv.Putter, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/itx/total")...), total.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountITxTotal(ctx context.Context, db kv.Getter, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/itx/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func WriteAccountITxIndex(ctx context.Context, db kv.Putter, addr common.Address, index *field.BigInt, data *types.InternalTxKey) (err error) {
	var bytesRes []byte
	bytesRes, err = data.Marshal()
	if err != nil {
		return
	}
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), bytesRes, &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountITxIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (data *types.InternalTxKey, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	data = &types.InternalTxKey{}
	err = data.Unmarshal(bytesRes)
	return
}

func ReadAccountITxByIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (itx *types.InternalTx, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/itx/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	data := &types.InternalTxKey{}
	err = data.Unmarshal(bytesRes)
	return ReadITx(ctx, db, data.TransactionHash, data.Index)
}

//  ---------------- erc20 transfer ---------------

func WriteAccountErc20Total(ctx context.Context, db kv.Putter, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc20/total")...), total.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc20Total(ctx context.Context, db kv.Getter, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc20/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}
	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func WriteAccountErc20Index(ctx context.Context, db kv.Putter, addr common.Address, index *field.BigInt, erc20TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), erc20TransferIndex.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc20Index(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (erc20TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc20TransferIndex = &field.BigInt{}
	erc20TransferIndex.SetBytes(bytesRes)
	return
}

func ReadAccountErc20ByIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (data *types.Erc20Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc20/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc20TransferIndex := &field.BigInt{}
	erc20TransferIndex.SetBytes(bytesRes)

	return ReadErc20Transfer(ctx, db, erc20TransferIndex)
}

//  ---------------- erc721 transfer ---------------

func WriteAccountErc721Total(ctx context.Context, db kv.Putter, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc721/total")...), total.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc721Total(ctx context.Context, db kv.Getter, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc721/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func WriteAccountErc721Index(ctx context.Context, db kv.Putter, addr common.Address, index *field.BigInt, erc721TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), erc721TransferIndex.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc721Index(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (erc721TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc721TransferIndex = &field.BigInt{}
	erc721TransferIndex.SetBytes(bytesRes)
	return
}

func ReadAccountErc721ByIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (data *types.Erc721Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc721/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc721TransferIndex := &field.BigInt{}
	erc721TransferIndex.SetBytes(bytesRes)

	return ReadErc721Transfer(ctx, db, erc721TransferIndex)
}

//  ---------------- erc115 transfer ---------------

func WriteAccountErc1155Total(ctx context.Context, db kv.Putter, addr common.Address, total *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc1155/total")...), total.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc1155Total(ctx context.Context, db kv.Getter, addr common.Address) (total *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), []byte("/erc1155/total")...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	total = &field.BigInt{}
	total.SetBytes(bytesRes)
	return
}

func WriteAccountErc1155Index(ctx context.Context, db kv.Putter, addr common.Address, index *field.BigInt, erc1155TransferIndex *field.BigInt) (err error) {
	return db.Put(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), erc1155TransferIndex.Bytes(), &kv.WriteOption{Table: share.AccountsTbl})
}

func ReadAccountErc1155Index(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (erc1155TransferIndex *field.BigInt, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc1155TransferIndex = &field.BigInt{}
	erc1155TransferIndex.SetBytes(bytesRes)
	return
}

func ReadAccountErc1155ByIndex(ctx context.Context, db kv.Getter, addr common.Address, index *field.BigInt) (data *types.Erc1155Transfer, err error) {
	var bytesRes []byte
	bytesRes, err = db.Get(ctx, append(append([]byte("/"), addr.Bytes()...), append([]byte("/erc1155/"), index.Bytes()...)...), &kv.ReadOption{Table: share.AccountsTbl})
	if err != nil {
		return
	}

	erc1155TransferIndex := &field.BigInt{}
	erc1155TransferIndex.SetBytes(bytesRes)

	return ReadErc1155Transfer(ctx, db, erc1155TransferIndex)
}
