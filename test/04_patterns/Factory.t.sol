// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Factory.sol";

contract FactoryTest is Test {
    Factory public factory;

    function setUp() public {
        factory = new Factory();
    }

    function test_Create() public {
        address child = factory.create(42);
        assertEq(Child(child).value(), 42);
        assertEq(Child(child).factory(), address(factory));
        assertEq(factory.childCount(), 1);
    }

    function test_CreateMultiple() public {
        factory.create(1);
        factory.create(2);
        factory.create(3);
        assertEq(factory.childCount(), 3);
    }

    function test_Create2() public {
        bytes32 salt = keccak256("test");
        address predicted = factory.predictAddress(42, salt);
        address actual = factory.create2(42, salt);
        assertEq(actual, predicted);
    }

    function test_Create2DeterministicAddress() public {
        bytes32 salt = keccak256("deterministic");
        address predicted = factory.predictAddress(100, salt);

        // Create with same salt should result in same address
        address actual = factory.create2(100, salt);
        assertEq(actual, predicted);

        // Value should be set correctly
        assertEq(Child(actual).value(), 100);
    }

    function testFuzz_Create(uint256 value) public {
        address child = factory.create(value);
        assertEq(Child(child).value(), value);
    }
}
