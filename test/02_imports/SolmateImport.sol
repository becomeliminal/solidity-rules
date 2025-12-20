// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "solmate/tokens/ERC20.sol";

contract SolmateToken is ERC20 {
    constructor() ERC20("Solmate Token", "SMT", 18) {
        _mint(msg.sender, 1000000 * 10 ** 18);
    }
}
