// SPDX-License-Identifier: MIT
pragma solidity =0.8.20;

// Skipped by its rule: it gets no binding.
contract Unwanted {
    function nothing() external pure returns (uint256) {
        return 0;
    }
}

contract Vault {
    uint256 public total;

    function deposit(uint256 amount) external {
        total += amount;
    }
}
