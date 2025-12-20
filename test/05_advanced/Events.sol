// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Events {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Log(string message);
    event LogWithData(string indexed message, uint256 data);

    function emitTransfer(address from, address to, uint256 value) public {
        emit Transfer(from, to, value);
    }

    function emitApproval(address owner, address spender, uint256 value) public {
        emit Approval(owner, spender, value);
    }

    function emitLog(string memory message) public {
        emit Log(message);
    }

    function emitLogWithData(string memory message, uint256 data) public {
        emit LogWithData(message, data);
    }
}
