// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "forge-std/console.sol";

contract Caller {
    function getSender() public view returns (address) {
        return msg.sender;
    }
}

contract ForgeStdTest is Test {
    function test_Console() public pure {
        console.log("Hello from forge-std!");
    }

    function test_Vm() public {
        vm.warp(1000);
        assertEq(block.timestamp, 1000);
    }

    function test_Deal() public {
        address alice = address(0x1);
        vm.deal(alice, 100 ether);
        assertEq(alice.balance, 100 ether);
    }

    function test_Prank() public {
        address alice = address(0x1);
        Caller caller = new Caller();
        vm.prank(alice);
        assertEq(caller.getSender(), alice);
    }
}
