// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Events.sol";

contract EventsTest is Test {
    Events public events;
    address public alice = address(0x1);
    address public bob = address(0x2);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Log(string message);
    event LogWithData(string indexed message, uint256 data);

    function setUp() public {
        events = new Events();
    }

    function test_EmitTransfer() public {
        vm.expectEmit(true, true, false, true);
        emit Transfer(alice, bob, 100);
        events.emitTransfer(alice, bob, 100);
    }

    function test_EmitApproval() public {
        vm.expectEmit(true, true, false, true);
        emit Approval(alice, bob, 500);
        events.emitApproval(alice, bob, 500);
    }

    function test_EmitLog() public {
        vm.expectEmit(false, false, false, true);
        emit Log("hello");
        events.emitLog("hello");
    }

    function test_EmitLogWithData() public {
        vm.expectEmit(true, false, false, true);
        emit LogWithData("test", 42);
        events.emitLogWithData("test", 42);
    }
}
