// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Child {
    address public factory;
    uint256 public value;

    constructor(uint256 _value) {
        factory = msg.sender;
        value = _value;
    }
}

contract Factory {
    address[] public children;

    event ChildCreated(address indexed child, uint256 value);

    function create(uint256 _value) public returns (address) {
        Child child = new Child(_value);
        children.push(address(child));
        emit ChildCreated(address(child), _value);
        return address(child);
    }

    function create2(uint256 _value, bytes32 salt) public returns (address) {
        Child child = new Child{salt: salt}(_value);
        children.push(address(child));
        emit ChildCreated(address(child), _value);
        return address(child);
    }

    function predictAddress(uint256 _value, bytes32 salt) public view returns (address) {
        bytes memory bytecode = abi.encodePacked(
            type(Child).creationCode,
            abi.encode(_value)
        );
        bytes32 hash = keccak256(
            abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode))
        );
        return address(uint160(uint256(hash)));
    }

    function childCount() public view returns (uint256) {
        return children.length;
    }
}
