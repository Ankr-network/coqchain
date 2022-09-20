// Copyright (c) 2022 mobus sunsc0220@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rss

import (
	"encoding/json"
	"math/big"

	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
)

type RpcBlock struct {
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  *big.Int      `json:"totalDifficulty"`
	Transactions     []interface{} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
	BaseFee          string        `json:"baseFeePerGas,omitempty"`
}

func (b *RpcBlock) ToBlock() *types.Block {
	oh := &types.Header{}
	oh.ParentHash = common.HexToHash(b.ParentHash)
	oh.UncleHash = common.HexToHash(b.Sha3Uncles)
	oh.Coinbase = common.HexToAddress(b.Miner)
	oh.Root = common.HexToHash(b.StateRoot)
	oh.TxHash = common.HexToHash(b.TransactionsRoot)
	oh.ReceiptHash = common.HexToHash(b.ReceiptsRoot)
	oh.Bloom = types.BytesToBloom(common.Hex2Bytes(b.LogsBloom))
	oh.Difficulty, _ = new(big.Int).SetString(b.Difficulty[2:], 16)
	oh.Number, _ = new(big.Int).SetString(b.Number[2:], 16)
	tnum, _ := new(big.Int).SetString(b.GasLimit[2:], 16)
	oh.GasLimit = tnum.Uint64()
	tnum, _ = new(big.Int).SetString(b.GasUsed[2:], 16)
	oh.GasUsed = tnum.Uint64()
	tnum, _ = new(big.Int).SetString(b.Timestamp[2:], 16)
	oh.Time = tnum.Uint64()
	oh.Extra = common.Hex2Bytes(b.ExtraData)
	oh.MixDigest = common.HexToHash(b.MixHash)
	tnum, _ = new(big.Int).SetString(b.Nonce[2:], 16)
	oh.Nonce = types.EncodeNonce(tnum.Uint64())
	if b.BaseFee != "" {
		oh.BaseFee, _ = new(big.Int).SetString(b.BaseFee[2:], 16)
	}
	ob := &types.Block{}
	ob.SetHeader(oh)
	return ob
}

type RpcTransaction struct {
	tx *types.Transaction
	TxExtraInfo
}

type TxExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *RpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.TxExtraInfo)
}
