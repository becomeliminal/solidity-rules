// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./OzERC721.sol";

contract OzERC721Test is Test {
    OzERC721 public nft;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        nft = new OzERC721("OpenZeppelin NFT", "OZNFT");
    }

    function test_Name() public view {
        assertEq(nft.name(), "OpenZeppelin NFT");
    }

    function test_Symbol() public view {
        assertEq(nft.symbol(), "OZNFT");
    }

    function test_Mint() public {
        uint256 tokenId = nft.mint(alice);
        assertEq(nft.ownerOf(tokenId), alice);
        assertEq(nft.balanceOf(alice), 1);
    }

    function test_Transfer() public {
        uint256 tokenId = nft.mint(alice);

        vm.prank(alice);
        nft.transferFrom(alice, bob, tokenId);

        assertEq(nft.ownerOf(tokenId), bob);
        assertEq(nft.balanceOf(alice), 0);
        assertEq(nft.balanceOf(bob), 1);
    }

    function test_Approve() public {
        uint256 tokenId = nft.mint(alice);

        vm.prank(alice);
        nft.approve(bob, tokenId);

        assertEq(nft.getApproved(tokenId), bob);
    }

    function test_ApproveForAll() public {
        vm.prank(alice);
        nft.setApprovalForAll(bob, true);

        assertTrue(nft.isApprovedForAll(alice, bob));
    }
}
