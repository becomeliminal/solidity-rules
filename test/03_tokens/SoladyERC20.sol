// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "solady/tokens/ERC20.sol";

contract SoladyERC20 is ERC20 {
    string private _name;
    string private _symbol;

    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
    }

    function name() public view override returns (string memory) {
        return _name;
    }

    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}
