// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

error InsufficientBalance(uint256 requested, uint256 available);
error Unauthorized(address caller);
error InvalidInput();

contract Errors {
    mapping(address => uint256) public balances;
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    function deposit() public payable {
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 amount) public {
        if (balances[msg.sender] < amount) {
            revert InsufficientBalance(amount, balances[msg.sender]);
        }
        balances[msg.sender] -= amount;
        (bool success,) = payable(msg.sender).call{value: amount}("");
        require(success, "transfer failed");
    }

    function onlyOwner() public view {
        if (msg.sender != owner) {
            revert Unauthorized(msg.sender);
        }
    }

    function requireNonZero(uint256 value) public pure {
        if (value == 0) {
            revert InvalidInput();
        }
    }
}
