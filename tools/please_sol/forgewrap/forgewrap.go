// Package forgewrap wraps forge commands to provide enhanced error messages.
package forgewrap

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Result contains the outcome of running a forge command.
type Result struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Enhanced string // Enhanced error message if applicable
}

// Wrapper wraps forge commands to provide enhanced error messages.
type Wrapper struct {
	remappings []string
}

// New creates a new Wrapper with the given available remappings.
func New(remappings []string) *Wrapper {
	return &Wrapper{remappings: remappings}
}

// Run executes a forge command and enhances any error output.
func (w *Wrapper) Run(forgePath string, args []string) (*Result, error) {
	cmd := exec.Command(forgePath, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		result.ExitCode = exitErr.ExitCode()
		result.Enhanced = w.enhanceError(stderr.String())
	} else if err != nil {
		return nil, fmt.Errorf("failed to run forge: %w", err)
	}

	return result, nil
}

// enhanceError parses forge error output and adds helpful hints.
func (w *Wrapper) enhanceError(stderr string) string {
	var enhanced strings.Builder
	enhanced.WriteString(stderr)

	// Check for unresolved import errors
	importPattern := regexp.MustCompile(`Unable to resolve import "([^"]+)"`)
	matches := importPattern.FindAllStringSubmatch(stderr, -1)

	if len(matches) > 0 {
		enhanced.WriteString("\n\n" + strings.Repeat("=", 60) + "\n")
		enhanced.WriteString("HINT: Import resolution failed\n")
		enhanced.WriteString(strings.Repeat("=", 60) + "\n\n")

		for _, match := range matches {
			importPath := match[1]
			enhanced.WriteString(fmt.Sprintf("  Import: %s\n", importPath))

			// Suggest potential deps based on import path
			suggestions := w.suggestDeps(importPath)
			if len(suggestions) > 0 {
				enhanced.WriteString("  Suggested deps:\n")
				for _, suggestion := range suggestions {
					enhanced.WriteString(fmt.Sprintf("    - %s\n", suggestion))
				}
			}
			enhanced.WriteString("\n")
		}

		if len(w.remappings) > 0 {
			enhanced.WriteString("Available remappings:\n")
			for _, r := range w.remappings {
				enhanced.WriteString(fmt.Sprintf("  %s\n", r))
			}
		}
	}

	// Check for compiler version errors
	if strings.Contains(stderr, "Source file requires different compiler version") {
		enhanced.WriteString("\n\nHINT: Compiler version mismatch.\n")
		enhanced.WriteString("Try specifying solc_version in your sol_contract rule.\n")
	}

	// Check for stack too deep
	if strings.Contains(stderr, "Stack too deep") {
		enhanced.WriteString("\n\nHINT: Stack too deep error.\n")
		enhanced.WriteString("Consider:\n")
		enhanced.WriteString("  - Breaking up the function into smaller functions\n")
		enhanced.WriteString("  - Using structs to group variables\n")
		enhanced.WriteString("  - Enabling optimizer with higher runs\n")
	}

	return enhanced.String()
}

// suggestDeps suggests deps based on an import path.
func (w *Wrapper) suggestDeps(importPath string) []string {
	var suggestions []string

	// Map common prefixes to deps
	prefixToDep := map[string]string{
		"@openzeppelin/contracts": "openzeppelin-contracts",
		"@openzeppelin":           "openzeppelin-contracts",
		"openzeppelin-contracts":  "openzeppelin-contracts",
		"forge-std":               "forge-std",
		"solmate":                 "solmate",
		"@solmate":                "solmate",
		"@rari-capital/solmate":   "solmate",
		"solady":                  "solady",
	}

	for prefix, dep := range prefixToDep {
		if strings.HasPrefix(importPath, prefix) {
			suggestions = append(suggestions, fmt.Sprintf("//third_party/solidity:%s", dep))
			break
		}
	}

	// Also check current remappings for partial matches
	for _, remap := range w.remappings {
		parts := strings.SplitN(remap, "=", 2)
		if len(parts) == 2 {
			prefix := strings.TrimSuffix(parts[0], "/")
			if strings.HasPrefix(importPath, prefix) {
				// Remapping exists but might not be in deps
				suggestions = append(suggestions, fmt.Sprintf("(remapping exists: %s)", remap))
			}
		}
	}

	return suggestions
}
