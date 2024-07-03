// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/SimulationToken.sol";

contract MyTokenTest is Test {
    SimulationToken token;

    function setUp() public {
        token = new SimulationToken(1000000 * 10 ** 18);
    }

    function testInitialSupply() public {
        assertEq(token.totalSupply(), 1000000 * 10 ** 18);
    }
}
