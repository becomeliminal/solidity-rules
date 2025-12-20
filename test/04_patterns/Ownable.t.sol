// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Ownable.sol";

contract OwnableTest is Test {
    OwnableContract public ownable;
    address public alice = address(0x1);

    function setUp() public {
        ownable = new OwnableContract();
    }

    function test_Owner() public view {
        assertEq(ownable.owner(), address(this));
    }

    function test_SetValue() public {
        ownable.setValue(42);
        assertEq(ownable.value(), 42);
    }

    function test_RevertWhen_NonOwnerSetsValue() public {
        vm.prank(alice);
        vm.expectRevert();
        ownable.setValue(42);
    }

    function test_TransferOwnership() public {
        ownable.transferOwnership(alice);
        assertEq(ownable.owner(), alice);
    }

    function test_RenounceOwnership() public {
        ownable.renounceOwnership();
        assertEq(ownable.owner(), address(0));
    }

    function test_PublicFunction() public view {
        assertEq(ownable.publicFunction(), "anyone can call");
    }
}
