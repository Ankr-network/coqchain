/*
 * Copyright (C) 2022  mobus <sv0220@163.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestLegacyTx(t *testing.T) {

	var ltx LegacyTx
	ltx.Type = EthTx
	ltx.TxID = common.HexToHash("0xe45177fc06771e30764e55942b02c45033528a705a35095855a7ef8d8aad84d8")
	ltx.Nonce = 20
	ltx.Gas = 21000
	ltx.GasPrice = big.NewInt(789000)
	ltx.To = nil
	ltx.Value = big.NewInt(20)

	buf := bytes.NewBuffer([]byte{})

	rlp.Encode(buf, ltx)

	t.Logf("ltx encode: %x \n", buf.String())

	var tx LegacyTx

	rlp.Decode(buf, &tx)

	t.Logf("tx decode: %+v \n", tx)

}
