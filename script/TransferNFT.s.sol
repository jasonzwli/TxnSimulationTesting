// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract TransferNFT is Script {
    function run() public view {
        address sender = vm.envAddress("SENDER");
        address recipient = vm.envAddress("RECIPIENT");
        uint256 tokenId = vm.envUint("TOKEN_ID");

        bytes4 selector = bytes4(keccak256("safeTransferFrom(address,address,uint256)"));
        bytes memory data = abi.encodeWithSelector(
            selector,
            sender,
            recipient,
            tokenId
        );

        console.logBytes(data);
    }
}
