// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./Helper.sol";

contract LocalImport {
    using Helper for uint256;

    function calculate(uint256 a, uint256 b) public pure returns (uint256 sum, uint256 product) {
        sum = a.add(b);
        product = a.mul(b);
    }
}
