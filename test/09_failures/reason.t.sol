// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

contract PlainTest {
    function test_ok() public pure {
        require(1 == 1, "ok");
    }

    function test_bad() public pure {
        require(1 == 2, "one is not two");
    }
}
