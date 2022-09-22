// Copyright (c) 2022 coqchain team
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

package container

import (
	"math/big"
	"testing"
)

func TestPut(t *testing.T) {
	var (
		size = 4
		idx  int64
	)
	hi := New(size)
	for idx = 1; idx < int64(size); idx++ {
		rs := hi.Put(big.NewInt(idx))
		if rs != nil {
			t.Error("rs should be nil")
		}
	}
	rs := hi.Put(big.NewInt(4))
	if rs != nil {
		t.Error("rs should be nil, but got: ", rs)
	}

	for idx = 5; idx < 10; idx++ {
		rs := hi.Put(big.NewInt(idx))
		if rs.Int64() != (idx - int64(size)) {
			if rs != nil {
				t.Errorf("rs should be %d, but got: %d \n", idx-int64(size), rs)
			}
		}
		t.Logf("rs input %d want %d, got: %d \n", idx, idx-int64(size), rs)
	}
	t.Logf("rest: %v \n", hi.List())
}
