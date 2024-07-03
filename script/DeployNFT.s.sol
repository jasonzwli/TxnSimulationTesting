// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/SimulationNFT.sol";

contract DeployNFT is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        SimulationNFT nft = new SimulationNFT();
        nft.createNFT();

        vm.stopBroadcast();
    }
}
