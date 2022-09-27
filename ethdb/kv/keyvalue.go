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

package kv

import (
	"encoding/binary"

	"github.com/sunvim/utils/cachem"
)

type KeyValue struct {
	Key []byte
	Val []byte
}

func (v *KeyValue) Marshal() []byte {
	kl := len(v.Key)
	kls := cachem.Malloc(2)
	defer cachem.Free(kls)
	binary.BigEndian.PutUint16(kls, uint16(kl))
	return append(append(kls, v.Key...), v.Val...)
}

func (v *KeyValue) Unmarshal(data []byte) error {
	kl := binary.BigEndian.Uint16(data[:2])
	v.Key = make([]byte, kl, kl)
	copy(v.Key, data[2:kl])
	v.Val = make([]byte, len(data)-2-int(kl))
	copy(v.Val, data[2+kl:])
	return nil
}
