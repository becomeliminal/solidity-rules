// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Interface.sol";

// Mock implementation of IToken for testing
contract MockToken is IToken {
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;

    constructor(uint256 initialSupply) {
        _totalSupply = initialSupply;
        _balances[msg.sender] = initialSupply;
    }

    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    function transfer(address to, uint256 amount) external override returns (bool) {
        _balances[msg.sender] -= amount;
        _balances[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function allowance(address owner, address spender) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 amount) external override returns (bool) {
        _allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) external override returns (bool) {
        _allowances[from][msg.sender] -= amount;
        _balances[from] -= amount;
        _balances[to] += amount;
        emit Transfer(from, to, amount);
        return true;
    }
}

contract InterfaceTest is Test {
    MockToken public token;
    address public alice = address(0x1);
    address public bob = address(0x2);

    function setUp() public {
        token = new MockToken(1000);
    }

    function test_InterfaceImplementation() public {
        // Test that MockToken correctly implements IToken interface
        assertEq(token.totalSupply(), 1000);
        assertEq(token.balanceOf(address(this)), 1000);
    }

    // Define events locally for expectEmit
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    function test_TransferEmitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit Transfer(address(this), alice, 100);
        token.transfer(alice, 100);
    }

    function test_ApprovalEmitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit Approval(address(this), alice, 500);
        token.approve(alice, 500);
    }

    function test_InterfaceTypeCheck() public view {
        // Verify the contract is castable to IToken
        IToken itoken = IToken(address(token));
        assertEq(itoken.totalSupply(), 1000);
    }
}
