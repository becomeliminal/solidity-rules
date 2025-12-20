// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "solady/auth/Ownable.sol";

contract SoladyOwnable is Ownable {
    uint256 public value;

    constructor() {
        _initializeOwner(msg.sender);
    }

    function setValue(uint256 _value) public onlyOwner {
        value = _value;
    }
}
