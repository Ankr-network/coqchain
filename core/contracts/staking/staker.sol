// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Staker {
	uint constant max_signers = 21;
	uint public threshold;
	address[] current_signers;
	mapping (address => bool) signers;
	mapping(address => uint256) public balances;
	bool locked = false;
	
	constructor ( address[] memory sigs, uint hold){
		for (uint idx = 0;idx < sigs.length; idx++){
			current_signers.push(sigs[idx]);
			signers[sigs[idx]] = true;
		}
		threshold = hold;
	}
	
	modifier needSigner() {
        require(
            signers[msg.sender],"you must be signer!"
        );
        _;
    }
	
	function list() view public returns (address[] memory) {
	   return current_signers;	
	}
	
	function stake() payable public {
		balances[msg.sender] += msg.value;
	}
	
	function balanceOf() view public returns(uint){
	   return balances[msg.sender];	
	}
	
	function _transfer(address payable addr, uint amt) internal {
	    require(balances[address] > amt, " balance not enougth");
		require(!locked,"Reentrant call detected!");
		locked = true;
		addr.transfer(amt);	
		locked = false;
	}
	
	function commit(address[] memory sigs) public needSigner{
		
	}

}