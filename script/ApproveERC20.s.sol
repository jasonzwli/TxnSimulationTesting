// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ApproveERC20 is Script {
    function run() public view {
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        address spender = vm.envAddress("SPENDER");
        uint256 amount = vm.envUint("AMOUNT");

        IERC20 token = IERC20(tokenAddress);
        bytes memory data = abi.encodeWithSelector(token.approve.selector, spender, amount);

        console.logBytes(data);
    }
}
