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

package common

const (
	TxTypeLength     = 1
	CrossChainLength = TxTypeLength + HashLength
)
const (
	COQ_TX byte = iota + 1
	ETH_TX
	BSC_TX
)

// ParseData tx type: 1 byte  tx id: 32 bytes
func ParseData(data []byte) (byte, *Hash, []byte) {
	h := BytesToHash(data[TxTypeLength : HashLength+TxTypeLength])
	return data[0], &h, data[CrossChainLength:]
}
