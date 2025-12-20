// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Storage {
    // Value types
    uint256 public number;
    address public owner;
    bool public flag;
    bytes32 public data;

    // Reference types
    string public name;
    bytes public rawData;
    uint256[] public numbers;

    // Mapping
    mapping(address => uint256) public balances;

    constructor() {
        owner = msg.sender;
    }

    function setNumber(uint256 _number) public {
        number = _number;
    }

    function setFlag(bool _flag) public {
        flag = _flag;
    }

    function setData(bytes32 _data) public {
        data = _data;
    }

    function setName(string memory _name) public {
        name = _name;
    }

    function setRawData(bytes memory _rawData) public {
        rawData = _rawData;
    }

    function pushNumber(uint256 _number) public {
        numbers.push(_number);
    }

    function getNumbersLength() public view returns (uint256) {
        return numbers.length;
    }

    function setBalance(address _addr, uint256 _balance) public {
        balances[_addr] = _balance;
    }
}
