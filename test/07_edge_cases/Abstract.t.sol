// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Abstract.sol";

contract AbstractTest is Test {
    ConcreteImpl public impl;

    function setUp() public {
        impl = new ConcreteImpl();
    }

    function test_ConcreteImplDeploys() public view {
        assertTrue(address(impl) != address(0));
    }

    function test_InheritedSetValue() public {
        // Test inherited implemented function from abstract base
        impl.setValue(42);
        assertEq(impl.value(), 42);
    }

    function test_OverriddenProcess() public {
        // Test that abstract function is correctly implemented
        impl.setValue(10);
        assertEq(impl.process(), 20); // value * 2
    }

    function test_OverriddenCompute() public view {
        // Test that abstract function with params is correctly implemented
        assertEq(impl.compute(5, 7), 12); // a + b
    }

    function test_FuzzProcess(uint256 value) public {
        // Fuzz test the process function
        vm.assume(value < type(uint256).max / 2); // Prevent overflow
        impl.setValue(value);
        assertEq(impl.process(), value * 2);
    }

    function test_FuzzCompute(uint256 a, uint256 b) public view {
        // Fuzz test compute - need to prevent overflow
        vm.assume(a < type(uint256).max - b);
        assertEq(impl.compute(a, b), a + b);
    }
}
