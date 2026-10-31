// Package artifacts reads Forge's out/ directory and writes one ABI and one
// bytecode file per contract, flat, for abigen.
package artifacts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type artifact struct {
	ABI      json.RawMessage `json:"abi"`
	Bytecode bytecode        `json:"bytecode"`
}

type bytecode struct {
	Object string `json:"object"`
}

// Extract walks forge's output directory and writes <name>.abi and <name>.bin
// into dest. build-info and any JSON without bytecode are skipped. Two
// contracts with the same name is an error.
func Extract(outDir, dest string) error {
	if err := os.MkdirAll(dest, 0o755); err != nil {
		return err
	}
	var wrote []string
	err := filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == "build-info" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var art artifact
		if err := json.Unmarshal(buf, &art); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if len(art.ABI) == 0 || string(art.ABI) == "null" {
			return nil
		}
		if art.Bytecode.Object == "" && !jsonHasBytecode(buf) {
			return nil
		}
		name := strings.TrimSuffix(info.Name(), ".json")
		abiPath := filepath.Join(dest, name+".abi")
		binPath := filepath.Join(dest, name+".bin")
		if _, err := os.Stat(abiPath); err == nil {
			return fmt.Errorf("two artifacts named %s (latest at %s)", name, path)
		}
		abi, err := json.Marshal(json.RawMessage(art.ABI))
		if err != nil {
			return fmt.Errorf("%s: abi: %w", path, err)
		}
		if err := os.WriteFile(abiPath, abi, 0o644); err != nil {
			return err
		}
		hex := strings.TrimPrefix(strings.TrimSpace(art.Bytecode.Object), "0x")
		if err := os.WriteFile(binPath, []byte(hex), 0o644); err != nil {
			return err
		}
		wrote = append(wrote, name)
		return nil
	})
	if err != nil {
		return err
	}
	if len(wrote) == 0 {
		return fmt.Errorf("no contract artifacts in %s", outDir)
	}
	return nil
}

// jsonHasBytecode reports whether the object has a bytecode key at all, so an
// interface whose object is "0x" still produces a .bin file.
func jsonHasBytecode(buf []byte) bool {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(buf, &raw); err != nil {
		return false
	}
	_, ok := raw["bytecode"]
	return ok
}
