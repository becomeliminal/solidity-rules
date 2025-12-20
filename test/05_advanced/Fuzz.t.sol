// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";

contract FuzzTarget {
    uint256 public value;
    mapping(address => uint256) public balances;

    function setValue(uint256 _value) public {
        value = _value;
    }

    function deposit(address to, uint256 amount) public {
        balances[to] += amount;
    }

    function safeAdd(uint256 a, uint256 b) public pure returns (uint256) {
        return a + b; // Will overflow in old Solidity, but 0.8+ reverts
    }
}

contract FuzzTest is Test {
    FuzzTarget public target;

    function setUp() public {
        target = new FuzzTarget();
    }

    // Basic fuzz test
    function testFuzz_SetValue(uint256 x) public {
        target.setValue(x);
        assertEq(target.value(), x);
    }

    // Fuzz with multiple params
    function testFuzz_Deposit(address to, uint256 amount) public {
        vm.assume(to != address(0));
        target.deposit(to, amount);
        assertEq(target.balances(to), amount);
    }

    // Fuzz with bounds
    function testFuzz_BoundedValue(uint256 x) public {
        x = bound(x, 1, 1000);
        target.setValue(x);
        assertGe(target.value(), 1);
        assertLe(target.value(), 1000);
    }

    // Fuzz testing overflow protection
    function testFuzz_SafeAdd(uint256 a, uint256 b) public view {
        // Skip if would overflow
        unchecked {
            if (a + b < a) return;
        }
        uint256 result = target.safeAdd(a, b);
        assertEq(result, a + b);
    }

    // Fuzz with assume to filter inputs
    function testFuzz_NonZeroDeposit(address to, uint256 amount) public {
        vm.assume(to != address(0));
        vm.assume(amount > 0);
        vm.assume(amount < type(uint128).max);

        target.deposit(to, amount);
        assertGt(target.balances(to), 0);
    }
}
