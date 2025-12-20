// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract OwnableContract is Ownable {
    uint256 public value;

    constructor() Ownable(msg.sender) {}

    function setValue(uint256 _value) public onlyOwner {
        value = _value;
    }

    function publicFunction() public pure returns (string memory) {
        return "anyone can call";
    }
}
