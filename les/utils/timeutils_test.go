// Copyright 2020 The coqchain Authors
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

package utils

import (
	"testing"
	"time"

	"github.com/Ankr-network/coqchain/common/mclock"
)

func TestUpdateTimer(t *testing.T) {
	timer := NewUpdateTimer(mclock.System{}, -1)
	if timer != nil {
		t.Fatalf("Create update timer with negative threshold")
	}
	sim := &mclock.Simulated{}
	timer = NewUpdateTimer(sim, time.Second)
	if updated := timer.Update(func(diff time.Duration) bool { return true }); updated {
		t.Fatalf("Update the clock without reaching the threshold")
	}
	sim.Run(time.Second)
	if updated := timer.Update(func(diff time.Duration) bool { return true }); !updated {
		t.Fatalf("Doesn't update the clock when reaching the threshold")
	}
	if updated := timer.UpdateAt(sim.Now()+mclock.AbsTime(time.Second), func(diff time.Duration) bool { return true }); !updated {
		t.Fatalf("Doesn't update the clock when reaching the threshold")
	}
	timer = NewUpdateTimer(sim, 0)
	if updated := timer.Update(func(diff time.Duration) bool { return true }); !updated {
		t.Fatalf("Doesn't update the clock without threshold limitaion")
	}
}
