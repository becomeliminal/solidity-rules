// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Abstract contract - some implemented, some not
abstract contract AbstractBase {
    uint256 public value;

    // Implemented function
    function setValue(uint256 _value) public {
        value = _value;
    }

    // Abstract function - must be implemented by child
    function process() public virtual returns (uint256);

    // Abstract function with parameters
    function compute(uint256 a, uint256 b) public virtual returns (uint256);
}

// Concrete implementation
contract ConcreteImpl is AbstractBase {
    function process() public view override returns (uint256) {
        return value * 2;
    }

    function compute(uint256 a, uint256 b) public pure override returns (uint256) {
        return a + b;
    }
}
