// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "solmate/tokens/ERC20.sol";

contract MultiLibToken is ERC20, Ownable {
    constructor() ERC20("Multi Lib Token", "MLT", 18) Ownable(msg.sender) {
        _mint(msg.sender, 1000000 * 10 ** 18);
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
}
