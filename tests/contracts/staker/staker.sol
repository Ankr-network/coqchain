// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Staker {
    mapping(address => uint256) balances;
    mapping(address => bool) admin;
    address payable owner;

    event Consume(address sender, uint amt);
    event Withdraw(address to, uint amount);
    event Transfer(address from, address to, uint amt);

    constructor() payable {
        owner = payable(msg.sender);
    }

    function setadmin(address _admin) public {
        require(msg.sender == owner, "only owner can invoke");
        admin[_admin] = true;
    }

    function consume() public payable {
        require(msg.value != 0, "wrong value");
        balances[msg.sender] += msg.value;
        emit Consume(msg.sender, msg.value);
    }

    function transfer(address to, uint amount) public payable {
        require(
            balances[msg.sender] >= amount,
            "transfer amount should be less or equal your balance"
        );
        balances[msg.sender] -= amount;
        balances[to] += amount;

        emit Transfer(msg.sender, to, amount);
    }

    function withdraw() public {
        require(msg.sender == owner, "only owner");

        uint amount = address(this).balance;

        (bool success, ) = owner.call{value: amount}("");
        require(success, "Failed to send");

        balances[msg.sender] = 0;
        emit Withdraw(msg.sender, amount);
    }

    function balanceOf(address addr) public view returns (uint) {
        return balances[addr];
    }
}
