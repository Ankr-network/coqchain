// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

library Array {
    struct ArrayMap {
        bytes32[] values;
        mapping(bytes32 => bool) map;
    }

    function push(ArrayMap storage _arr, bytes32 _value) internal {
        if (!_arr.map[_value]) {
            _arr.map[_value] = true;
            _arr.values.push(_value);
        }
    }

    function contains(ArrayMap storage _arr, bytes32 _value)
        internal
        view
        returns (bool)
    {
        return _arr.map[_value];
    }

    function remove(ArrayMap storage _arr, bytes32 _value) internal {
        if (_arr.map[_value]) {
            delete _arr.map[_value];
            for (uint256 i = 0; i <= _arr.values.length; i++) {
                if (_arr.values[i] == _value) {
                    _arr.values[i] = _arr.values[_arr.values.length - 1];
                    _arr.values.pop();
                    break;
                }
            }
        }
    }

    function list(ArrayMap storage _arr)
        internal
        view
        returns (bytes32[] memory)
    {
        return _arr.values;
    }

    function length(ArrayMap storage _arr) internal view returns (uint256) {
        return _arr.values.length;
    }
}

contract Staker {
    using Array for Array.ArrayMap;
    uint256 public epoch;
    uint256 public threshold;
    uint256 public fineRatio; // 10 = 10%

    Array.ArrayMap signers;
    mapping(address => uint256) public balances;

    enum VoteType {
        UNKNOW,
        JOIN,
        EXIT
    }

    enum VoteRes {
        UNKNOW,
        AGREE,
        AGAINST
    }

    struct Proposal {
        VoteType voteType;
        mapping(address => VoteRes) voteMaps;
        address[] votes;
    }

    // epoch => (votee => Proposal)
    mapping(uint256 => mapping(address => Proposal)) public epochProposals;

    // epoch => [votee]
    mapping(uint256 => address[]) public epochProposalVotees;

    mapping(uint256 => bool) epochVoted;

    function stake() external payable {
        require(msg.value > 0, "invalid amount");
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 _amount) external {
        require(
            !signers.contains(bytes32(bytes20(msg.sender))),
            "staking, unable to withdraw"
        );
        require(balances[msg.sender] >= _amount, "insufficient amount");
        balances[msg.sender] -= _amount;
        payable(msg.sender).transfer(_amount);
    }

    function signerList() external view returns (address[] memory) {
        bytes32[] memory sigsBytes_ = signers.list();
        address[] memory sig_ = new address[](sigsBytes_.length);
        for (uint256 i = 0; i < sigsBytes_.length; i++) {
            sig_[i] = address(bytes20(sigsBytes_[i]));
        }
        return sig_;
    }

    function signerContains(address _signer) public view returns (bool) {
        return signers.contains(bytes32(bytes20(_signer)));
    }

    function getCycle(uint256 _blockNumber) public view returns (uint256) {
        return _blockNumber / epoch;
    }

    function epochVotedByBlockNumber(uint256 _blockNumber)
        external
        view
        returns (bool)
    {
        return epochVoted[getCycle(_blockNumber)];
    }

    function checkVoteStatus(
        uint256 _blockNumber,
        address _votee,
        address _voter
    ) external view returns (VoteRes) {
        return epochProposals[getCycle(_blockNumber)][_votee].voteMaps[_voter];
    }
}
