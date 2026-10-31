package bindings_test

import (
	"testing"

	mathlib "test/06_go_bindings/mathlib"
)

func TestLibraryBinding(t *testing.T) {
	abi, err := mathlib.InstanceMetaData.GetAbi()
	if err != nil {
		t.Fatalf("Failed to parse library ABI: %v", err)
	}
	if _, ok := abi.Methods["add"]; !ok {
		t.Error("Expected 'add' method on MathLib")
	}
	if len(mathlib.InstanceBin) == 0 {
		t.Error("Library bytecode is empty")
	}
}
