// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Staker {
    uint256 public epoch;
    uint256 public threshold;
    uint256 public fineRatio; // 10 = 10%

    address[] signers;
    mapping(address => uint256) public balances;

    struct VoteInfo {
        address vote;
        bool authorize;
    }

    mapping(address => VoteInfo[]) public proposalVotes;

    function stake() external payable {
        require(msg.value > 0, "invalid amount");
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 _amount) external {
        require(!signerContains(msg.sender), "staking, unable to withdraw");
        require(balances[msg.sender] >= _amount, "insufficient amount");
        balances[msg.sender] -= _amount;
        payable(msg.sender).transfer(_amount);
    }

    function signerList() external view returns (address[] memory) {
        return signers;
    }

    function signerContains(address _signer) public view returns (bool) {
        for (uint256 i = 0; i < signers.length; i++) {
            if (signers[i] == _signer) {
                return true;
            }
        }
        return false;
    }
}
