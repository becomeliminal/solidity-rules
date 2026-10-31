package decls

import (
	"strings"
	"testing"
)

func TestExtract_SkipsCommentsAndStrings(t *testing.T) {
	src := []byte(`
// this contract is not a declaration
/* contract CommentTrick {} */
/*
    contract AlsoInAComment {}
*/
contract Real {
    string note = "contract Fake";
}
abstract /* gap */ contract Abs {
    function a() public virtual;
}
interface IFace {
    function b() external;
}
library Lib {
    function c(uint256 x) external pure returns (uint256) { return x; }
}
`)
	got, err := Extract(src)
	if err != nil {
		t.Fatal(err)
	}
	want := []Decl{
		{"contract", "Real"},
		{"abstract contract", "Abs"},
		{"interface", "IFace"},
		{"library", "Lib"},
	}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("decl %d: got %#v, want %#v", i, got[i], want[i])
		}
	}
	text := Format(got)
	if strings.Contains(text, "Fake") || strings.Contains(text, "AlsoInAComment") || strings.Contains(text, "CommentTrick") {
		t.Errorf("comment or string was treated as a declaration:\n%s", text)
	}
}

func TestExtract_EdgeShapes(t *testing.T) {
	// Shapes from test/07_edge_cases: empty contract, interface, abstract, library.
	src := []byte(`
pragma solidity ^0.8.20;

contract Empty {}

interface IERC {
    function balanceOf(address who) external view returns (uint256);
}

abstract contract Base {
    function name() public view virtual returns (string memory);
}

library MathLib {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }
}
`)
	got, err := Extract(src)
	if err != nil {
		t.Fatal(err)
	}
	names := make([]string, len(got))
	for i, d := range got {
		names[i] = d.Kind + " " + d.Name
	}
	want := "contract Empty, interface IERC, abstract contract Base, library MathLib"
	if strings.Join(names, ", ") != want {
		t.Fatalf("got %s", strings.Join(names, ", "))
	}
}

func TestFormat_TabSeparated(t *testing.T) {
	text := Format([]Decl{{Kind: "abstract contract", Name: "Abs"}})
	if text != "abstract contract\tAbs\n" {
		t.Fatalf("got %q", text)
	}
}
