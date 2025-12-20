// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Errors.sol";

contract ErrorsTest is Test {
    Errors public errors;
    address public alice = address(0x1);

    // Allow this contract to receive ETH
    receive() external payable {}

    function setUp() public {
        errors = new Errors();
    }

    function test_Deposit() public {
        vm.deal(alice, 1 ether);
        vm.prank(alice);
        errors.deposit{value: 0.5 ether}();
        assertEq(errors.balances(alice), 0.5 ether);
    }

    function test_Withdraw() public {
        vm.deal(alice, 1 ether);
        vm.prank(alice);
        errors.deposit{value: 1 ether}();

        vm.prank(alice);
        errors.withdraw(0.5 ether);

        assertEq(errors.balances(alice), 0.5 ether);
    }

    function test_RevertWith_InsufficientBalance() public {
        vm.expectRevert(
            abi.encodeWithSelector(InsufficientBalance.selector, 100, 0)
        );
        errors.withdraw(100);
    }

    function test_RevertWith_Unauthorized() public {
        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSelector(Unauthorized.selector, alice)
        );
        errors.onlyOwner();
    }

    function test_RevertWith_InvalidInput() public {
        vm.expectRevert(InvalidInput.selector);
        errors.requireNonZero(0);
    }

    function test_NoRevert_ValidInput() public {
        // This should not revert
        errors.requireNonZero(1);
    }
}
