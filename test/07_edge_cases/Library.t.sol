// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Library.sol";

contract LibraryTest is Test {
    MathUser public mathUser;

    function setUp() public {
        mathUser = new MathUser();
    }

    function test_LibraryAddition() public view {
        (uint256 sum, , ) = mathUser.calculate(10, 5);
        assertEq(sum, 15);
    }

    function test_LibrarySubtraction() public view {
        (, uint256 diff, ) = mathUser.calculate(10, 5);
        assertEq(diff, 5);
    }

    function test_LibraryMultiplication() public view {
        (, , uint256 product) = mathUser.calculate(10, 5);
        assertEq(product, 50);
    }

    function test_LibraryWithReversedInputs() public view {
        // Test when b > a for subtraction
        (, uint256 diff, ) = mathUser.calculate(5, 10);
        assertEq(diff, 5); // abs(5 - 10) = 5
    }

    function test_LibraryWithZero() public view {
        (uint256 sum, uint256 diff, uint256 product) = mathUser.calculate(0, 10);
        assertEq(sum, 10);
        assertEq(diff, 10);
        assertEq(product, 0);
    }

    function test_FuzzCalculate(uint256 a, uint256 b) public view {
        // Prevent overflow in addition and multiplication
        vm.assume(a < type(uint128).max);
        vm.assume(b < type(uint128).max);

        (uint256 sum, uint256 diff, uint256 product) = mathUser.calculate(a, b);

        assertEq(sum, a + b);
        assertEq(diff, a > b ? a - b : b - a);
        assertEq(product, a * b);
    }
}
