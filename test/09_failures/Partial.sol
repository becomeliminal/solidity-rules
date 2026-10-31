// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

// Parses, then fails type checking, so forge has already created out/.
contract Partial {
    function f() public pure returns (uint256) {
        return missing;
    }
}
