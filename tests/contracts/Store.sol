// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

contract Store {
  event ItemSet(string key, string value);

  string public version;
  mapping (string => string) public items;

  constructor (string memory _version) {
    version = _version;
  }

  function setItem(string memory key, string memory value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
  
  function getItem(string memory key) external view returns(string memory) {
	return items[key];
  }
}