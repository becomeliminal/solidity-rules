// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./SimpleStorage.sol";

contract SimpleStorageTest is Test {
    SimpleStorage public store;

    function setUp() public {
        store = new SimpleStorage();
    }

    function test_InitialValue() public view {
        assertEq(store.get(), 0);
    }

    function test_Set() public {
        store.set(42);
        assertEq(store.get(), 42);
    }

    function test_Owner() public view {
        assertEq(store.owner(), address(this));
    }

    function testFuzz_Set(uint256 value) public {
        store.set(value);
        assertEq(store.get(), value);
    }
}
