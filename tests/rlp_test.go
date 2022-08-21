// Copyright 2015 The coqchain Authors
// This file is part of the coqchain library.
//
// The coqchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The coqchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the coqchain library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"testing"
)

func TestRLP(t *testing.T) {
	t.Parallel()
	tm := new(testMatcher)
	tm.walk(t, rlpTestDir, func(t *testing.T, name string, test *RLPTest) {
		if err := tm.checkFailure(t, test.Run()); err != nil {
			t.Error(err)
		}
	})
}
