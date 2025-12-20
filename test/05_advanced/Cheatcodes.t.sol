// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";

contract Target {
    address public lastCaller;
    uint256 public lastTimestamp;
    uint256 public lastBlockNumber;

    function recordCaller() public {
        lastCaller = msg.sender;
        lastTimestamp = block.timestamp;
        lastBlockNumber = block.number;
    }

    function onlyAlice() public view returns (bool) {
        return msg.sender == address(0x1);
    }
}

contract Reverter {
    function doRevert() public pure {
        revert("always reverts");
    }

    function revertWithMessage(string memory message) public pure {
        revert(message);
    }
}

contract CheatcodesTest is Test {
    Target public target;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        target = new Target();
    }

    // vm.prank - spoof msg.sender for next call
    function test_Prank() public {
        vm.prank(alice);
        target.recordCaller();
        assertEq(target.lastCaller(), alice);
    }

    // vm.startPrank / vm.stopPrank - spoof for multiple calls
    function test_StartStopPrank() public {
        vm.startPrank(alice);
        target.recordCaller();
        assertEq(target.lastCaller(), alice);
        target.recordCaller();
        assertEq(target.lastCaller(), alice);
        vm.stopPrank();

        target.recordCaller();
        assertEq(target.lastCaller(), address(this));
    }

    // vm.deal - set ETH balance
    function test_Deal() public {
        assertEq(alice.balance, 0);
        vm.deal(alice, 100 ether);
        assertEq(alice.balance, 100 ether);
    }

    // vm.warp - set block.timestamp
    function test_Warp() public {
        uint256 newTime = 1000000;
        vm.warp(newTime);
        target.recordCaller();
        assertEq(target.lastTimestamp(), newTime);
    }

    // vm.roll - set block.number
    function test_Roll() public {
        uint256 newBlock = 12345;
        vm.roll(newBlock);
        target.recordCaller();
        assertEq(target.lastBlockNumber(), newBlock);
    }

    // vm.expectRevert - expect next external call to revert
    function test_ExpectRevert() public {
        Reverter reverter = new Reverter();
        vm.expectRevert();
        reverter.doRevert();
    }

    // vm.expectRevert with message
    function test_ExpectRevertWithMessage() public {
        Reverter reverter = new Reverter();
        vm.expectRevert("specific error");
        reverter.revertWithMessage("specific error");
    }

    // Warp forward in time
    function test_WarpForward() public {
        uint256 startTime = block.timestamp;
        vm.warp(startTime + 1 days);
        assertEq(block.timestamp, startTime + 1 days);
    }

    // vm.label - label an address for better traces
    function test_Label() public {
        vm.label(alice, "Alice");
        vm.label(bob, "Bob");
        // Labels show in traces
        assertEq(alice, address(0x1));
    }

    // Combined cheatcodes
    function test_CombinedCheatcodes() public {
        vm.deal(alice, 10 ether);
        vm.warp(1000);
        vm.roll(100);
        vm.prank(alice);

        target.recordCaller();

        assertEq(target.lastCaller(), alice);
        assertEq(target.lastTimestamp(), 1000);
        assertEq(target.lastBlockNumber(), 100);
        assertEq(alice.balance, 10 ether);
    }
}
