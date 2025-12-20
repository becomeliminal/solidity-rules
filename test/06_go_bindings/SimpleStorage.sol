// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract SimpleStorage {
    uint256 private _value;
    address public owner;

    event ValueChanged(uint256 indexed oldValue, uint256 indexed newValue);
    event OwnerChanged(address indexed previousOwner, address indexed newOwner);

    constructor() {
        owner = msg.sender;
    }

    function set(uint256 value) external {
        uint256 oldValue = _value;
        _value = value;
        emit ValueChanged(oldValue, value);
    }

    function get() external view returns (uint256) {
        return _value;
    }

    function setOwner(address newOwner) external {
        require(msg.sender == owner, "not owner");
        address oldOwner = owner;
        owner = newOwner;
        emit OwnerChanged(oldOwner, newOwner);
    }
}
