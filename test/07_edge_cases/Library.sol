// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Solidity library - can't have state, all functions are internal/external
library MathLib {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "underflow");
        return a - b;
    }

    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        return a * b;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "division by zero");
        return a / b;
    }

    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a >= b ? a : b;
    }

    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a <= b ? a : b;
    }
}

// Contract using the library
contract MathUser {
    using MathLib for uint256;

    function calculate(uint256 a, uint256 b) public pure returns (uint256 sum, uint256 diff, uint256 product) {
        sum = a.add(b);
        diff = a > b ? a.sub(b) : b.sub(a);
        product = a.mul(b);
    }
}
