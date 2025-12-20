// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Receive.sol";

contract ReceiveTest is Test {
    Receive public target;

    function setUp() public {
        target = new Receive();
        vm.deal(address(this), 10 ether);
    }

    function test_ReceiveETH() public {
        (bool success,) = address(target).call{value: 1 ether}("");
        assertTrue(success);
        assertEq(target.getBalance(), 1 ether);
        assertEq(target.receiveCalls(), 1);
    }

    function test_FallbackWithData() public {
        bytes memory data = abi.encodeWithSignature("nonExistentFunction()");
        (bool success,) = address(target).call{value: 0.5 ether}(data);
        assertTrue(success);
        assertEq(target.fallbackCalls(), 1);
        assertEq(target.lastFallbackData(), data);
    }

    function test_MultipleReceives() public {
        (bool success1,) = address(target).call{value: 1 ether}("");
        (bool success2,) = address(target).call{value: 2 ether}("");
        (bool success3,) = address(target).call{value: 0.5 ether}("");

        assertTrue(success1 && success2 && success3);
        assertEq(target.receiveCalls(), 3);
        assertEq(target.getBalance(), 3.5 ether);
    }
}
