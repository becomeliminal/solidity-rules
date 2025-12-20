// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract PausableContract is Pausable, Ownable {
    uint256 public value;

    constructor() Ownable(msg.sender) {}

    function setValue(uint256 _value) public whenNotPaused {
        value = _value;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }
}
