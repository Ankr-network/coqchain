// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Staker {

	mapping(address => uint256) public balances;
	mapping(address => bool) public signers;
    uint epoch;
	uint stakeAmt = 1000;
    constructor() payable {}
	
	function vote(kickout address,confirm bool) {
		
	}

    // Function to deposit Ether into this contract.
    // Call this function along with some Ether.
    // The balance of this contract will be automatically updated.
    function deposit() public payable {
		require(msg.value != 0,"wrong value");
		balances[msg.sender] += msg.value;
	}

    function withdraw(address payable _to, uint _amount) public{

		// _amount should be less or equal _to balance
		require(balances[_to] >= _amount, "withdraw amount should be less or equal your balance");

        (bool success, ) = _to.call{value: _amount}("");
        require(success, "Failed to send");

    }
	
	function balanceOf(address addr) view public returns (uint) {
	     return balances[addr];	
	}
}
