// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract OzToken is ERC20 {
    constructor() ERC20("OZ Token", "OZT") {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}
