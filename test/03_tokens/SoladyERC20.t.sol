// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./SoladyERC20.sol";

contract SoladyERC20Test is Test {
    SoladyERC20 public token;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        token = new SoladyERC20("Solady Token", "SDY");
        token.mint(alice, 1000 ether);
    }

    function test_Name() public view {
        assertEq(token.name(), "Solady Token");
    }

    function test_Symbol() public view {
        assertEq(token.symbol(), "SDY");
    }

    function test_Transfer() public {
        vm.prank(alice);
        token.transfer(bob, 100 ether);

        assertEq(token.balanceOf(alice), 900 ether);
        assertEq(token.balanceOf(bob), 100 ether);
    }

    function test_TransferFrom() public {
        vm.prank(alice);
        token.approve(bob, 500 ether);

        vm.prank(bob);
        token.transferFrom(alice, bob, 300 ether);

        assertEq(token.balanceOf(bob), 300 ether);
    }
}
