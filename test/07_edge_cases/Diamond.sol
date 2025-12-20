// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Diamond inheritance pattern: D inherits from B and C, both inherit from A
//
//      A
//     / \
//    B   C
//     \ /
//      D

contract A {
    function foo() public virtual returns (string memory) {
        return "A";
    }

    function bar() public virtual returns (string memory) {
        return "A.bar";
    }
}

contract B is A {
    function foo() public virtual override returns (string memory) {
        return "B";
    }
}

contract C is A {
    function foo() public virtual override returns (string memory) {
        return "C";
    }

    function bar() public virtual override returns (string memory) {
        return "C.bar";
    }
}

// D must explicitly choose which implementation to use
contract D is B, C {
    function foo() public override(B, C) returns (string memory) {
        return "D";
    }

    // For bar, C overrides A, so we need to override C
    function bar() public override(A, C) returns (string memory) {
        return "D.bar";
    }

    // Access parent implementations
    function callBFoo() public returns (string memory) {
        return B.foo();
    }

    function callCFoo() public returns (string memory) {
        return C.foo();
    }
}
