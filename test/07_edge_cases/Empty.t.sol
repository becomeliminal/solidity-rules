// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Empty.sol";

contract EmptyTest is Test {
    Empty public empty;

    function setUp() public {
        empty = new Empty();
    }

    function test_DeploySucceeds() public view {
        // Empty contract should deploy successfully
        assertTrue(address(empty) != address(0));
    }

    function test_CodeExists() public view {
        // Empty contract still has code (just the constructor)
        assertTrue(address(empty).code.length > 0);
    }
}
