package foundrytoml

import (
	"encoding/json"
	"testing"
)

func TestParse_BasicConfig(t *testing.T) {
	input := `
[profile.default]
src = "src"
out = "out"
libs = ["lib"]
solc_version = "0.8.20"
optimizer = true
optimizer_runs = 200
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	profile := config.GetProfile("default")
	if profile == nil {
		t.Fatal("expected default profile")
	}

	if profile.SolcVersion != "0.8.20" {
		t.Errorf("expected solc_version '0.8.20', got %q", profile.SolcVersion)
	}
	if !profile.Optimizer {
		t.Error("expected optimizer to be true")
	}
	if profile.OptimizerRuns != 200 {
		t.Errorf("expected optimizer_runs 200, got %d", profile.OptimizerRuns)
	}
	if profile.Src != "src" {
		t.Errorf("expected src 'src', got %q", profile.Src)
	}
}

func TestParse_WithRemappings(t *testing.T) {
	input := `
remappings = [
  "@openzeppelin/=lib/openzeppelin-contracts/",
  "forge-std/=lib/forge-std/src/"
]

[profile.default]
remappings = [
  "solmate/=lib/solmate/src/"
]
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	remappings := config.GetRemappings("default")
	if len(remappings) != 3 {
		t.Errorf("expected 3 remappings, got %d: %v", len(remappings), remappings)
	}

	// Check top-level remappings come first
	if remappings[0] != "@openzeppelin/=lib/openzeppelin-contracts/" {
		t.Errorf("expected openzeppelin remapping first, got %q", remappings[0])
	}
}

func TestParse_MultipleProfiles(t *testing.T) {
	input := `
[profile.default]
solc_version = "0.8.20"
optimizer = false

[profile.ci]
solc_version = "0.8.21"
optimizer = true
optimizer_runs = 1000

[profile.production]
solc_version = "0.8.20"
optimizer = true
optimizer_runs = 1000000
via_ir = true
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Test default profile
	defProfile := config.GetProfile("default")
	if defProfile.SolcVersion != "0.8.20" {
		t.Errorf("default: expected solc 0.8.20, got %q", defProfile.SolcVersion)
	}
	if defProfile.Optimizer {
		t.Error("default: expected optimizer to be false")
	}

	// Test ci profile
	ciProfile := config.GetProfile("ci")
	if ciProfile.SolcVersion != "0.8.21" {
		t.Errorf("ci: expected solc 0.8.21, got %q", ciProfile.SolcVersion)
	}
	if !ciProfile.Optimizer {
		t.Error("ci: expected optimizer to be true")
	}
	if ciProfile.OptimizerRuns != 1000 {
		t.Errorf("ci: expected 1000 runs, got %d", ciProfile.OptimizerRuns)
	}

	// Test production profile
	prodProfile := config.GetProfile("production")
	if !prodProfile.ViaIR {
		t.Error("production: expected via_ir to be true")
	}
	if prodProfile.OptimizerRuns != 1000000 {
		t.Errorf("production: expected 1000000 runs, got %d", prodProfile.OptimizerRuns)
	}
}

func TestGetProfile_FallbackToDefault(t *testing.T) {
	input := `
[profile.default]
solc_version = "0.8.20"
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Request non-existent profile should fall back to default
	profile := config.GetProfile("nonexistent")
	if profile == nil {
		t.Fatal("expected to fall back to default profile")
	}
	if profile.SolcVersion != "0.8.20" {
		t.Errorf("expected solc 0.8.20 from default, got %q", profile.SolcVersion)
	}

	// Empty string should return default
	profile = config.GetProfile("")
	if profile == nil {
		t.Fatal("expected default profile for empty name")
	}
}

func TestGetProfile_NoProfiles(t *testing.T) {
	input := `
# Empty foundry.toml with just remappings
remappings = ["foo/=lib/foo/"]
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	profile := config.GetProfile("default")
	if profile != nil {
		t.Error("expected nil profile when no profiles defined")
	}
}

func TestGetOptimizerSettings(t *testing.T) {
	input := `
[profile.default]
optimizer = true
optimizer_runs = 500
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	enabled, runs := config.GetOptimizerSettings("default")
	if !enabled {
		t.Error("expected optimizer to be enabled")
	}
	if runs != 500 {
		t.Errorf("expected 500 runs, got %d", runs)
	}
}

func TestGetEvmVersion(t *testing.T) {
	input := `
[profile.default]
evm_version = "paris"
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	evmVersion := config.GetEvmVersion("default")
	if evmVersion != "paris" {
		t.Errorf("expected evm_version 'paris', got %q", evmVersion)
	}
}

func TestToJSON(t *testing.T) {
	input := `
[profile.default]
solc_version = "0.8.20"
optimizer = true
optimizer_runs = 200
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	jsonBytes, err := config.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	// Parse it back to verify valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
		t.Fatalf("JSON output is invalid: %v", err)
	}
}

func TestParse_RealWorldExample(t *testing.T) {
	// Based on actual foundry.toml from common projects
	input := `
[profile.default]
src = "src"
out = "out"
libs = ["lib"]
test = "test"
cache = true
cache_path = "cache"

# Compiler settings
solc_version = "0.8.23"
evm_version = "paris"
optimizer = true
optimizer_runs = 200
via_ir = false

# Testing
fuzz_runs = 256
invariant_runs = 256
verbosity = 2

# Remappings
remappings = [
    "@openzeppelin/contracts/=lib/openzeppelin-contracts/contracts/",
    "@openzeppelin/contracts-upgradeable/=lib/openzeppelin-contracts-upgradeable/contracts/",
    "forge-std/=lib/forge-std/src/",
    "solmate/=lib/solmate/src/"
]

[profile.ci]
fuzz_runs = 10000
invariant_runs = 10000

[profile.production]
optimizer_runs = 1000000
via_ir = true
`
	config, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Verify default profile
	def := config.GetProfile("default")
	if def.SolcVersion != "0.8.23" {
		t.Errorf("expected solc 0.8.23, got %q", def.SolcVersion)
	}
	if def.EvmVersion != "paris" {
		t.Errorf("expected evm paris, got %q", def.EvmVersion)
	}
	if len(def.Remappings) != 4 {
		t.Errorf("expected 4 remappings, got %d", len(def.Remappings))
	}
	if def.FuzzRuns != 256 {
		t.Errorf("expected 256 fuzz runs, got %d", def.FuzzRuns)
	}

	// Verify CI profile overrides
	ci := config.GetProfile("ci")
	if ci.FuzzRuns != 10000 {
		t.Errorf("ci: expected 10000 fuzz runs, got %d", ci.FuzzRuns)
	}

	// Verify production profile
	prod := config.GetProfile("production")
	if !prod.ViaIR {
		t.Error("production: expected via_ir to be true")
	}
}
