package bindings_test

import (
	"testing"

	storage "test/06_go_bindings/simplestorage"
)

func TestABIParsing(t *testing.T) {
	abi, err := storage.InstanceMetaData.GetAbi()
	if err != nil {
		t.Fatalf("Failed to parse ABI: %v", err)
	}

	// Check that expected methods exist
	if _, ok := abi.Methods["set"]; !ok {
		t.Error("Expected 'set' method in ABI")
	}
	if _, ok := abi.Methods["get"]; !ok {
		t.Error("Expected 'get' method in ABI")
	}
	if _, ok := abi.Methods["owner"]; !ok {
		t.Error("Expected 'owner' method in ABI")
	}
	if _, ok := abi.Methods["setOwner"]; !ok {
		t.Error("Expected 'setOwner' method in ABI")
	}
}

func TestBytecodeNotEmpty(t *testing.T) {
	if len(storage.InstanceBin) == 0 {
		t.Error("Bytecode is empty")
	}
}

func TestEvents(t *testing.T) {
	abi, err := storage.InstanceMetaData.GetAbi()
	if err != nil {
		t.Fatalf("Failed to parse ABI: %v", err)
	}

	// Check that expected events exist
	if _, ok := abi.Events["ValueChanged"]; !ok {
		t.Error("Expected 'ValueChanged' event in ABI")
	}
	if _, ok := abi.Events["OwnerChanged"]; !ok {
		t.Error("Expected 'OwnerChanged' event in ABI")
	}
}
