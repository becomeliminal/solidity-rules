// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

contract Push0 {
    function f() public pure returns (uint256) {
        uint256 x;
        assembly {
            x := 0
        }
        return x;
    }
}
