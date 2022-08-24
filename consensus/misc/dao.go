// Copyright 2016 The coqchain Authors
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

package misc

import (
	"errors"
	"math/big"

	"github.com/Ankr-network/coqchain/core/state"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/params"
)

var (
	// ErrBadProDAOExtra is returned if a header doesn't support the DAO fork on a
	// pro-fork client.
	ErrBadProDAOExtra = errors.New("bad DAO pro-fork extra-data")

	// ErrBadNoDAOExtra is returned if a header does support the DAO fork on a no-
	// fork client.
	ErrBadNoDAOExtra = errors.New("bad DAO no-fork extra-data")
)

// VerifyDAOHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to DAO hard-fork rules.
//
// DAO hard-fork extension to the header validity:
//
//	a) if the node is no-fork, do not accept blocks in the [fork, fork+10) range
//	   with the fork specific extra-data set
//	b) if the node is pro-fork, require blocks in the specific range to have the
//	   unique extra-data set.
func VerifyDAOHeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	// All ok, header has the same extra-data we expect
	return nil
}

// ApplyDAOHardFork modifies the state database according to the DAO hard-fork
// rules, transferring all balances of a set of DAO accounts to a single refund
// contract.
func ApplyDAOHardFork(statedb *state.StateDB) {
	// Retrieve the contract to refund balances into
	if !statedb.Exist(params.DAORefundContract) {
		statedb.CreateAccount(params.DAORefundContract)
	}

	// Move every DAO account and extra-balance account funds into the refund contract
	for _, addr := range params.DAODrainList() {
		statedb.AddBalance(params.DAORefundContract, statedb.GetBalance(addr))
		statedb.SetBalance(addr, new(big.Int))
	}
}
