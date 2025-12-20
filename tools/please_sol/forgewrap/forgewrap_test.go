package forgewrap

import (
	"strings"
	"testing"
)

func TestEnhanceError_ImportResolution(t *testing.T) {
	tests := []struct {
		name       string
		stderr     string
		remappings []string
		wantHints  []string
	}{
		{
			name:   "openzeppelin import",
			stderr: `Error: Unable to resolve import "@openzeppelin/contracts/token/ERC20/ERC20.sol"`,
			wantHints: []string{
				"HINT: Import resolution failed",
				"//third_party/solidity:openzeppelin-contracts",
			},
		},
		{
			name:   "forge-std import",
			stderr: `Error: Unable to resolve import "forge-std/Test.sol"`,
			wantHints: []string{
				"HINT: Import resolution failed",
				"//third_party/solidity:forge-std",
			},
		},
		{
			name:   "solmate import",
			stderr: `Error: Unable to resolve import "@rari-capital/solmate/src/tokens/ERC20.sol"`,
			wantHints: []string{
				"HINT: Import resolution failed",
				"//third_party/solidity:solmate",
			},
		},
		{
			name:       "with existing remappings",
			stderr:     `Error: Unable to resolve import "mylib/Foo.sol"`,
			remappings: []string{"mylib/=third_party/mylib/"},
			wantHints: []string{
				"HINT: Import resolution failed",
				"remapping exists: mylib/=third_party/mylib/",
			},
		},
		{
			name:   "multiple imports",
			stderr: `Error: Unable to resolve import "@openzeppelin/contracts/token/ERC20/ERC20.sol". Error: Unable to resolve import "forge-std/Test.sol"`,
			wantHints: []string{
				"//third_party/solidity:openzeppelin-contracts",
				"//third_party/solidity:forge-std",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.remappings)
			enhanced := w.enhanceError(tt.stderr)

			for _, hint := range tt.wantHints {
				if !strings.Contains(enhanced, hint) {
					t.Errorf("expected hint %q in output:\n%s", hint, enhanced)
				}
			}
		})
	}
}

func TestEnhanceError_CompilerVersion(t *testing.T) {
	w := New(nil)
	stderr := "Error: Source file requires different compiler version"
	enhanced := w.enhanceError(stderr)

	if !strings.Contains(enhanced, "HINT: Compiler version mismatch") {
		t.Errorf("expected compiler version hint in output:\n%s", enhanced)
	}
	if !strings.Contains(enhanced, "solc_version") {
		t.Errorf("expected solc_version suggestion in output:\n%s", enhanced)
	}
}

func TestEnhanceError_StackTooDeep(t *testing.T) {
	w := New(nil)
	stderr := "CompilerError: Stack too deep when compiling inline assembly"
	enhanced := w.enhanceError(stderr)

	if !strings.Contains(enhanced, "HINT: Stack too deep") {
		t.Errorf("expected stack too deep hint in output:\n%s", enhanced)
	}
	if !strings.Contains(enhanced, "smaller functions") {
		t.Errorf("expected suggestion about smaller functions in output:\n%s", enhanced)
	}
}

func TestEnhanceError_NoMatch(t *testing.T) {
	w := New(nil)
	stderr := "Some other error message"
	enhanced := w.enhanceError(stderr)

	// Should just return the original message without any hints
	if enhanced != stderr {
		t.Errorf("expected original stderr, got:\n%s", enhanced)
	}
}

func TestSuggestDeps(t *testing.T) {
	tests := []struct {
		name       string
		importPath string
		remappings []string
		wantEmpty  bool
		wantDep    string
	}{
		{
			name:       "openzeppelin full path",
			importPath: "@openzeppelin/contracts/token/ERC20/ERC20.sol",
			wantDep:    "openzeppelin-contracts",
		},
		{
			name:       "openzeppelin short",
			importPath: "@openzeppelin/access/Ownable.sol",
			wantDep:    "openzeppelin-contracts",
		},
		{
			name:       "forge-std",
			importPath: "forge-std/Test.sol",
			wantDep:    "forge-std",
		},
		{
			name:       "solmate direct",
			importPath: "solmate/tokens/ERC20.sol",
			wantDep:    "solmate",
		},
		{
			name:       "solmate rari-capital",
			importPath: "@rari-capital/solmate/src/tokens/ERC20.sol",
			wantDep:    "solmate",
		},
		{
			name:       "solady",
			importPath: "solady/tokens/ERC20.sol",
			wantDep:    "solady",
		},
		{
			name:       "unknown import",
			importPath: "unknown/lib/Foo.sol",
			wantEmpty:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.remappings)
			suggestions := w.suggestDeps(tt.importPath)

			if tt.wantEmpty {
				if len(suggestions) > 0 {
					t.Errorf("expected no suggestions, got: %v", suggestions)
				}
				return
			}

			found := false
			for _, s := range suggestions {
				if strings.Contains(s, tt.wantDep) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected suggestion containing %q, got: %v", tt.wantDep, suggestions)
			}
		})
	}
}
