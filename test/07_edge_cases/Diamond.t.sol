// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Diamond.sol";

contract DiamondTest is Test {
    A public a;
    B public b;
    C public c;
    D public d;

    function setUp() public {
        a = new A();
        b = new B();
        c = new C();
        d = new D();
    }

    function test_BaseContractA() public {
        assertEq(a.foo(), "A");
        assertEq(a.bar(), "A.bar");
    }

    function test_ContractBOverridesFoo() public {
        assertEq(b.foo(), "B");
        assertEq(b.bar(), "A.bar"); // bar not overridden in B
    }

    function test_ContractCOverridesBoth() public {
        assertEq(c.foo(), "C");
        assertEq(c.bar(), "C.bar");
    }

    function test_DiamondContractD() public {
        // D overrides both, resolving diamond
        assertEq(d.foo(), "D");
        assertEq(d.bar(), "D.bar");
    }

    function test_DiamondCallParentB() public {
        // D can explicitly call B's implementation
        assertEq(d.callBFoo(), "B");
    }

    function test_DiamondCallParentC() public {
        // D can explicitly call C's implementation
        assertEq(d.callCFoo(), "C");
    }

    function test_InheritanceHierarchy() public {
        // Verify the inheritance chain works correctly
        // D should be usable as A, B, or C
        A aFromD = A(address(d));
        assertEq(aFromD.foo(), "D"); // Still calls D's implementation

        B bFromD = B(address(d));
        assertEq(bFromD.foo(), "D");

        C cFromD = C(address(d));
        assertEq(cFromD.foo(), "D");
    }
}
