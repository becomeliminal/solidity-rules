// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Contract with receive and fallback functions
contract Receive {
    uint256 public receiveCalls;
    uint256 public fallbackCalls;
    bytes public lastFallbackData;

    event Received(address indexed sender, uint256 value);
    event FallbackCalled(address indexed sender, uint256 value, bytes data);

    // Called when receiving ETH with no data
    receive() external payable {
        receiveCalls++;
        emit Received(msg.sender, msg.value);
    }

    // Called when receiving ETH with data, or when function doesn't exist
    fallback() external payable {
        fallbackCalls++;
        lastFallbackData = msg.data;
        emit FallbackCalled(msg.sender, msg.value, msg.data);
    }

    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
}
