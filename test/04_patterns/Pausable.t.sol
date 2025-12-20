// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Pausable.sol";

contract PausableTest is Test {
    PausableContract public pausable;

    function setUp() public {
        pausable = new PausableContract();
    }

    function test_NotPausedByDefault() public view {
        assertFalse(pausable.paused());
    }

    function test_SetValueWhenNotPaused() public {
        pausable.setValue(42);
        assertEq(pausable.value(), 42);
    }

    function test_Pause() public {
        pausable.pause();
        assertTrue(pausable.paused());
    }

    function test_Unpause() public {
        pausable.pause();
        pausable.unpause();
        assertFalse(pausable.paused());
    }

    function test_RevertWhen_SetValueWhilePaused() public {
        pausable.pause();
        vm.expectRevert();
        pausable.setValue(42);
    }

    function test_SetValueAfterUnpause() public {
        pausable.pause();
        pausable.unpause();
        pausable.setValue(42);
        assertEq(pausable.value(), 42);
    }
}
