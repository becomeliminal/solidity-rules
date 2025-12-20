// Package detectprefix provides import prefix detection for Solidity libraries.
//
// It auto-detects the canonical import prefix from a library's package.json,
// following the patterns established by major Solidity ecosystems:
//
//   - Scoped npm packages: @openzeppelin/contracts, @uniswap/v3-core
//   - Simple names: solmate, forge-std, solady
//
// Detection priority:
//  1. Nested package.json in the package directory (e.g., contracts/package.json)
//  2. Root package.json
//  3. Fallback to the target name
package detectprefix

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"strings"
)

// Detector reads package.json files from a zip archive to detect import prefixes.
type Detector struct {
	zipPath    string
	packageDir string
	name       string
}

// New creates a new Detector.
func New(zipPath, packageDir, name string) *Detector {
	return &Detector{
		zipPath:    zipPath,
		packageDir: packageDir,
		name:       name,
	}
}

// Detect returns the import prefix for the library.
// It checks package.json files in order of priority and falls back to the target name.
func (d *Detector) Detect() (string, error) {
	r, err := zip.OpenReader(d.zipPath)
	if err != nil {
		return "", fmt.Errorf("failed to open zip: %w", err)
	}
	defer r.Close()

	repoRoot := findRepoRoot(r)

	// Priority 1: Nested package.json in package directory
	nestedPath := path.Join(repoRoot, d.packageDir, "package.json")
	if prefix := readPackageName(r, nestedPath); prefix != "" {
		return prefix, nil
	}

	// Priority 2: Root package.json
	rootPath := path.Join(repoRoot, "package.json")
	if prefix := readPackageName(r, rootPath); prefix != "" {
		return prefix, nil
	}

	// Priority 3: Fallback to target name
	return d.name, nil
}

// findRepoRoot finds the repository root directory in the zip.
// GitHub zip archives typically have structure: owner-repo-revision/...
func findRepoRoot(r *zip.ReadCloser) string {
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			return strings.TrimSuffix(f.Name, "/")
		}
	}
	return ""
}

// readPackageName reads the "name" field from a package.json file in the zip.
// Returns empty string if the file doesn't exist or name is a placeholder.
func readPackageName(r *zip.ReadCloser, filePath string) string {
	for _, f := range r.File {
		if f.Name == filePath {
			return extractName(f)
		}
	}
	return ""
}

// extractName opens a zip file entry and extracts the package name.
// It applies known migrations to handle renamed/moved packages.
func extractName(f *zip.File) string {
	rc, err := f.Open()
	if err != nil {
		return ""
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return ""
	}

	var pkg struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return ""
	}

	if isPlaceholder(pkg.Name) {
		return ""
	}

	// Apply known migrations for renamed/moved packages
	return applyMigrations(pkg.Name)
}

// isPlaceholder returns true if the package name is a monorepo placeholder.
// These names indicate the package.json is a workspace root, not a publishable package.
func isPlaceholder(name string) bool {
	if name == "" {
		return true
	}
	lower := strings.ToLower(name)
	return lower == "workspace" || lower == "root" || lower == "monorepo"
}

// knownMigrations maps old/legacy package.json names to their canonical import prefixes.
// This handles cases where a library was renamed or moved to a different org, but
// older versions still have the legacy name in package.json.
var knownMigrations = map[string]string{
	// Solmate was originally published under @rari-capital (Rari Capital), but after
	// Rari merged with Fei Protocol, the canonical import became just "solmate".
	// Older tags (v6 and earlier) still have @rari-capital/solmate in package.json.
	"@rari-capital/solmate": "solmate",
}

// applyMigrations checks if a package name has a known migration and returns
// the canonical name. Returns the original name if no migration exists.
func applyMigrations(name string) string {
	if canonical, ok := knownMigrations[name]; ok {
		return canonical
	}
	return name
}

// FormatRemapping formats the import prefix as a Solidity remapping.
func FormatRemapping(prefix, outDir, name string) string {
	return fmt.Sprintf("%s/=%s/%s/", prefix, outDir, name)
}
