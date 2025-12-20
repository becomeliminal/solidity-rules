// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Storage.sol";

contract StorageTest is Test {
    Storage public store;
    address public alice = address(0x1);

    function setUp() public {
        store = new Storage();
    }

    function test_Owner() public view {
        assertEq(store.owner(), address(this));
    }

    function test_SetNumber() public {
        store.setNumber(42);
        assertEq(store.number(), 42);
    }

    function test_SetFlag() public {
        store.setFlag(true);
        assertTrue(store.flag());
    }

    function test_SetData() public {
        bytes32 testData = keccak256("hello");
        store.setData(testData);
        assertEq(store.data(), testData);
    }

    function test_SetName() public {
        store.setName("Test Name");
        assertEq(store.name(), "Test Name");
    }

    function test_SetRawData() public {
        bytes memory testData = hex"deadbeef";
        store.setRawData(testData);
        assertEq(store.rawData(), testData);
    }

    function test_PushNumber() public {
        store.pushNumber(1);
        store.pushNumber(2);
        store.pushNumber(3);

        assertEq(store.getNumbersLength(), 3);
        assertEq(store.numbers(0), 1);
        assertEq(store.numbers(1), 2);
        assertEq(store.numbers(2), 3);
    }

    function test_SetBalance() public {
        store.setBalance(alice, 100);
        assertEq(store.balances(alice), 100);
    }

    function testFuzz_SetNumber(uint256 x) public {
        store.setNumber(x);
        assertEq(store.number(), x);
    }

    function testFuzz_SetBalance(address addr, uint256 amount) public {
        store.setBalance(addr, amount);
        assertEq(store.balances(addr), amount);
    }
}
