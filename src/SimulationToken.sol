// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract SimulationToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("SimulationToken", "STK") {
        _mint(msg.sender, initialSupply);
    }

    function mint(uint256 amount) public {
        _mint(msg.sender, amount);
    }
}