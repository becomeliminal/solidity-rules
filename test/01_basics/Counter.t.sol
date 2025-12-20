// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Counter.sol";

contract CounterTest is Test {
    Counter public counter;

    function setUp() public {
        counter = new Counter();
    }

    function test_InitialCount() public view {
        assertEq(counter.count(), 0);
    }

    function test_Increment() public {
        counter.increment();
        assertEq(counter.count(), 1);
    }

    function test_Decrement() public {
        counter.increment();
        counter.decrement();
        assertEq(counter.count(), 0);
    }

    function test_SetCount() public {
        counter.setCount(42);
        assertEq(counter.count(), 42);
    }

    function test_RevertWhen_DecrementBelowZero() public {
        vm.expectRevert();
        counter.decrement();
    }

    function testFuzz_SetCount(uint256 x) public {
        counter.setCount(x);
        assertEq(counter.count(), x);
    }
}
