package artifacts

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtract_WritesABIAndBin(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out", "C.sol")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}
	info := filepath.Join(root, "out", "build-info")
	if err := os.MkdirAll(info, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(info, "abc.json"), []byte(`{"id":"nope"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	body := `{"abi":[{"type":"function","name":"f"}],"bytecode":{"object":"0x60805f"}}`
	if err := os.WriteFile(filepath.Join(out, "C.json"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	empty := `{"abi":[],"bytecode":{"object":"0x"}}`
	if err := os.WriteFile(filepath.Join(out, "I.json"), []byte(empty), 0o644); err != nil {
		t.Fatal(err)
	}

	dest := filepath.Join(root, "dest")
	if err := Extract(filepath.Join(root, "out"), dest); err != nil {
		t.Fatal(err)
	}
	bin, err := os.ReadFile(filepath.Join(dest, "C.bin"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bin) != "60805f" {
		t.Errorf("bin %q", bin)
	}
	abi, err := os.ReadFile(filepath.Join(dest, "C.abi"))
	if err != nil {
		t.Fatal(err)
	}
	if string(abi) == "" || string(abi) == "null" {
		t.Errorf("abi %s", abi)
	}
	if _, err := os.Stat(filepath.Join(dest, "I.bin")); err != nil {
		t.Errorf("interface bin missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dest, "abc.abi")); err == nil {
		t.Error("build-info was extracted")
	}
}

func TestExtract_DuplicateName(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"a", "b"} {
		d := filepath.Join(root, "out", dir)
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
		body := `{"abi":[],"bytecode":{"object":"0x60"}}`
		if err := os.WriteFile(filepath.Join(d, "C.json"), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	err := Extract(filepath.Join(root, "out"), filepath.Join(root, "dest"))
	if err == nil {
		t.Fatal("expected duplicate name to fail")
	}
}

func TestExtract_NoArtifacts(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}
	err := Extract(out, filepath.Join(root, "dest"))
	if err == nil {
		t.Fatal("expected empty out to fail")
	}
}
