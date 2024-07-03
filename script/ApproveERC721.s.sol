// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract ApproveERC721 is Script {
    function run() public view {
        address nftAddress = vm.envAddress("NFT_ADDRESS");
        address approved = vm.envAddress("APPROVED");
        uint256 tokenId = vm.envUint("TOKEN_ID");

        IERC721 nft = IERC721(nftAddress);
        bytes memory data = abi.encodeWithSelector(nft.approve.selector, approved, tokenId);

        console.logBytes(data);
    }
}
