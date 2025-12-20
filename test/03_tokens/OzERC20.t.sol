// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./OzERC20.sol";

contract OzERC20Test is Test {
    OzERC20 public token;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        token = new OzERC20("OpenZeppelin Token", "OZT");
        token.mint(alice, 1000 ether);
    }

    function test_Name() public view {
        assertEq(token.name(), "OpenZeppelin Token");
    }

    function test_Symbol() public view {
        assertEq(token.symbol(), "OZT");
    }

    function test_Decimals() public view {
        assertEq(token.decimals(), 18);
    }

    function test_Balance() public view {
        assertEq(token.balanceOf(alice), 1000 ether);
    }

    function test_Transfer() public {
        vm.prank(alice);
        token.transfer(bob, 100 ether);

        assertEq(token.balanceOf(alice), 900 ether);
        assertEq(token.balanceOf(bob), 100 ether);
    }

    function test_Approve() public {
        vm.prank(alice);
        token.approve(bob, 500 ether);

        assertEq(token.allowance(alice, bob), 500 ether);
    }

    function test_TransferFrom() public {
        vm.prank(alice);
        token.approve(bob, 500 ether);

        vm.prank(bob);
        token.transferFrom(alice, bob, 300 ether);

        assertEq(token.balanceOf(alice), 700 ether);
        assertEq(token.balanceOf(bob), 300 ether);
        assertEq(token.allowance(alice, bob), 200 ether);
    }

    function testFuzz_Transfer(uint256 amount) public {
        amount = bound(amount, 0, token.balanceOf(alice));
        vm.prank(alice);
        token.transfer(bob, amount);
        assertEq(token.balanceOf(bob), amount);
    }
}
