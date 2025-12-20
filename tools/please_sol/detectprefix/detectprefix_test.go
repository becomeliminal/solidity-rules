package detectprefix

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

// createTestZip creates a temporary zip file with the given files.
// files is a map of path -> content.
func createTestZip(t *testing.T, repoRoot string, files map[string]string) string {
	t.Helper()

	tmpDir := t.TempDir()
	zipPath := filepath.Join(tmpDir, "test.zip")

	f, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("failed to create zip file: %v", err)
	}
	defer f.Close()

	w := zip.NewWriter(f)

	// Create repo root directory entry
	if repoRoot != "" {
		_, err := w.Create(repoRoot + "/")
		if err != nil {
			t.Fatalf("failed to create dir entry: %v", err)
		}
	}

	for path, content := range files {
		fullPath := path
		if repoRoot != "" {
			fullPath = repoRoot + "/" + path
		}
		fw, err := w.Create(fullPath)
		if err != nil {
			t.Fatalf("failed to create file %s: %v", path, err)
		}
		_, err = fw.Write([]byte(content))
		if err != nil {
			t.Fatalf("failed to write file %s: %v", path, err)
		}
	}

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close zip: %v", err)
	}

	return zipPath
}

func TestDetect_NestedPackageJSON(t *testing.T) {
	// OpenZeppelin pattern: contracts/package.json has the real name
	zipPath := createTestZip(t, "openzeppelin-contracts-5.0.0", map[string]string{
		"package.json":           `{"name": "openzeppelin-solidity"}`,
		"contracts/package.json": `{"name": "@openzeppelin/contracts"}`,
	})

	d := New(zipPath, "contracts", "openzeppelin-contracts")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "@openzeppelin/contracts" {
		t.Errorf("expected @openzeppelin/contracts, got %s", prefix)
	}
}

func TestDetect_RootPackageJSON(t *testing.T) {
	// Solady pattern: root package.json has the name
	zipPath := createTestZip(t, "solady-0.0.227", map[string]string{
		"package.json": `{"name": "solady"}`,
	})

	d := New(zipPath, "src", "solady")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "solady" {
		t.Errorf("expected solady, got %s", prefix)
	}
}

func TestDetect_FallbackToName(t *testing.T) {
	// No package.json at all - fall back to target name
	zipPath := createTestZip(t, "some-lib-1.0.0", map[string]string{
		"src/Lib.sol": "// solidity code",
	})

	d := New(zipPath, "src", "some-lib")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "some-lib" {
		t.Errorf("expected some-lib, got %s", prefix)
	}
}

func TestDetect_PlaceholderNames(t *testing.T) {
	testCases := []struct {
		name        string
		packageJSON string
		expected    string
	}{
		{"workspace", `{"name": "workspace"}`, "my-lib"},
		{"root", `{"name": "root"}`, "my-lib"},
		{"monorepo", `{"name": "monorepo"}`, "my-lib"},
		{"Workspace uppercase", `{"name": "Workspace"}`, "my-lib"},
		{"ROOT uppercase", `{"name": "ROOT"}`, "my-lib"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zipPath := createTestZip(t, "my-lib-1.0.0", map[string]string{
				"package.json": tc.packageJSON,
			})

			d := New(zipPath, "src", "my-lib")
			prefix, err := d.Detect()
			if err != nil {
				t.Fatalf("Detect failed: %v", err)
			}

			if prefix != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, prefix)
			}
		})
	}
}

func TestDetect_KnownMigrations(t *testing.T) {
	// Solmate v6 has @rari-capital/solmate but should return solmate
	zipPath := createTestZip(t, "solmate-6", map[string]string{
		"package.json": `{"name": "@rari-capital/solmate"}`,
	})

	d := New(zipPath, "src", "solmate")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "solmate" {
		t.Errorf("expected solmate (migrated from @rari-capital/solmate), got %s", prefix)
	}
}

func TestDetect_MalformedJSON(t *testing.T) {
	// Malformed JSON should fall back gracefully
	zipPath := createTestZip(t, "bad-lib-1.0.0", map[string]string{
		"package.json": `{not valid json`,
	})

	d := New(zipPath, "src", "bad-lib")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	// Should fall back to target name
	if prefix != "bad-lib" {
		t.Errorf("expected bad-lib (fallback), got %s", prefix)
	}
}

func TestDetect_EmptyPackageJSON(t *testing.T) {
	// Empty name field should fall back
	zipPath := createTestZip(t, "empty-name-1.0.0", map[string]string{
		"package.json": `{"name": ""}`,
	})

	d := New(zipPath, "src", "empty-name")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "empty-name" {
		t.Errorf("expected empty-name (fallback), got %s", prefix)
	}
}

func TestDetect_MissingZip(t *testing.T) {
	d := New("/nonexistent/path/to.zip", "src", "test")
	_, err := d.Detect()
	if err == nil {
		t.Error("expected error for missing zip file")
	}
}

func TestDetect_NestedTakesPrecedence(t *testing.T) {
	// When both root and nested exist, nested wins
	zipPath := createTestZip(t, "dual-pkg-1.0.0", map[string]string{
		"package.json":     `{"name": "wrong-name"}`,
		"src/package.json": `{"name": "correct-name"}`,
	})

	d := New(zipPath, "src", "dual-pkg")
	prefix, err := d.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if prefix != "correct-name" {
		t.Errorf("expected correct-name (nested), got %s", prefix)
	}
}

func TestFormatRemapping(t *testing.T) {
	result := FormatRemapping("@openzeppelin/contracts", "third_party/solidity", "openzeppelin-contracts")
	expected := "@openzeppelin/contracts/=third_party/solidity/openzeppelin-contracts/"

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestApplyMigrations(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"@rari-capital/solmate", "solmate"},
		{"solmate", "solmate"},
		{"@openzeppelin/contracts", "@openzeppelin/contracts"},
		{"forge-std", "forge-std"},
		{"unknown-lib", "unknown-lib"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := applyMigrations(tc.input)
			if result != tc.expected {
				t.Errorf("applyMigrations(%s) = %s, expected %s", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsPlaceholder(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
	}{
		{"", true},
		{"workspace", true},
		{"Workspace", true},
		{"WORKSPACE", true},
		{"root", true},
		{"Root", true},
		{"monorepo", true},
		{"Monorepo", true},
		{"solmate", false},
		{"@openzeppelin/contracts", false},
		{"forge-std", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPlaceholder(tc.name)
			if result != tc.expected {
				t.Errorf("isPlaceholder(%q) = %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}
