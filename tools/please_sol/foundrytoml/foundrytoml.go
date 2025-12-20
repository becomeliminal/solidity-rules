// Package foundrytoml parses foundry.toml configuration files.
package foundrytoml

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config represents the parsed foundry.toml configuration.
type Config struct {
	Profile   map[string]Profile `toml:"profile" json:"profile,omitempty"`
	Remapping []string           `toml:"remappings" json:"remappings,omitempty"`
}

// Profile represents a foundry profile (e.g., default, ci, production).
type Profile struct {
	// Compiler settings
	SolcVersion    string `toml:"solc_version" json:"solc_version,omitempty"`
	EvmVersion     string `toml:"evm_version" json:"evm_version,omitempty"`
	ViaIR          bool   `toml:"via_ir" json:"via_ir,omitempty"`
	Optimizer      bool   `toml:"optimizer" json:"optimizer,omitempty"`
	OptimizerRuns  int    `toml:"optimizer_runs" json:"optimizer_runs,omitempty"`

	// Path settings
	Src       string   `toml:"src" json:"src,omitempty"`
	Out       string   `toml:"out" json:"out,omitempty"`
	Libs      []string `toml:"libs" json:"libs,omitempty"`
	Test      string   `toml:"test" json:"test,omitempty"`
	Cache     bool     `toml:"cache" json:"cache,omitempty"`
	CachePath string   `toml:"cache_path" json:"cache_path,omitempty"`

	// Remappings
	Remappings []string `toml:"remappings" json:"remappings,omitempty"`

	// Testing
	FuzzRuns       int  `toml:"fuzz_runs" json:"fuzz_runs,omitempty"`
	InvariantRuns  int  `toml:"invariant_runs" json:"invariant_runs,omitempty"`
	Verbosity      int  `toml:"verbosity" json:"verbosity,omitempty"`
	NoMatchTest    string `toml:"no_match_test" json:"no_match_test,omitempty"`
	MatchTest      string `toml:"match_test" json:"match_test,omitempty"`
	MatchContract  string `toml:"match_contract" json:"match_contract,omitempty"`
}

// ParseFile parses a foundry.toml file.
func ParseFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read foundry.toml: %w", err)
	}
	return Parse(data)
}

// Parse parses foundry.toml content.
func Parse(data []byte) (*Config, error) {
	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse foundry.toml: %w", err)
	}
	return &config, nil
}

// GetProfile returns the specified profile, or the default profile if not found.
func (c *Config) GetProfile(name string) *Profile {
	if name == "" {
		name = "default"
	}
	if profile, ok := c.Profile[name]; ok {
		return &profile
	}
	if profile, ok := c.Profile["default"]; ok {
		return &profile
	}
	return nil
}

// ToJSON converts the config to JSON for use in build rules.
func (c *Config) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// GetSolcVersion returns the solc version from the specified profile.
func (c *Config) GetSolcVersion(profile string) string {
	p := c.GetProfile(profile)
	if p == nil {
		return ""
	}
	return p.SolcVersion
}

// GetRemappings returns remappings from the specified profile, merged with top-level remappings.
func (c *Config) GetRemappings(profile string) []string {
	var remappings []string

	// Top-level remappings first
	remappings = append(remappings, c.Remapping...)

	// Profile-specific remappings
	p := c.GetProfile(profile)
	if p != nil {
		remappings = append(remappings, p.Remappings...)
	}

	return remappings
}

// GetOptimizerSettings returns optimizer settings from the specified profile.
func (c *Config) GetOptimizerSettings(profile string) (enabled bool, runs int) {
	p := c.GetProfile(profile)
	if p == nil {
		return false, 0
	}
	return p.Optimizer, p.OptimizerRuns
}

// GetEvmVersion returns the EVM version from the specified profile.
func (c *Config) GetEvmVersion(profile string) string {
	p := c.GetProfile(profile)
	if p == nil {
		return ""
	}
	return p.EvmVersion
}
