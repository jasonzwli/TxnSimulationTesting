// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract SimulationNFT is ERC721 {
    uint256 public tokenCounter;

    constructor() ERC721("SimulationNFT", "SNFT") {
        tokenCounter = 0;
    }

    function createNFT() public returns (uint256) {
        uint256 newTokenId = tokenCounter;
        _safeMint(msg.sender, newTokenId);
        tokenCounter += 1;
        return newTokenId;
    }
}
