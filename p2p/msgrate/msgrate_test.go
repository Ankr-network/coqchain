// Copyright 2021 The coqchain Authors
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

package msgrate

import "testing"

func TestCapacityOverflow(t *testing.T) {
	tracker := NewTracker(nil, 1)
	tracker.Update(1, 1, 100000)
	cap := tracker.Capacity(1, 10000000)
	if int32(cap) < 0 {
		t.Fatalf("Negative: %v", int32(cap))
	}
}
