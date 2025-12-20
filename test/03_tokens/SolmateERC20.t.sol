// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./SolmateERC20.sol";

contract SolmateERC20Test is Test {
    SolmateERC20 public token;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        token = new SolmateERC20("Solmate Token", "SMT");
        token.mint(alice, 1000 ether);
    }

    function test_Name() public view {
        assertEq(token.name(), "Solmate Token");
    }

    function test_Symbol() public view {
        assertEq(token.symbol(), "SMT");
    }

    function test_Decimals() public view {
        assertEq(token.decimals(), 18);
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
