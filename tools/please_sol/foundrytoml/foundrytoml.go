// Package foundrytoml parses foundry.toml configuration files.
package foundrytoml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	SolcVersion   string `toml:"solc_version" json:"solc_version,omitempty"`
	EvmVersion    string `toml:"evm_version" json:"evm_version,omitempty"`
	ViaIR         bool   `toml:"via_ir" json:"via_ir,omitempty"`
	Optimizer     bool   `toml:"optimizer" json:"optimizer,omitempty"`
	OptimizerRuns int    `toml:"optimizer_runs" json:"optimizer_runs,omitempty"`

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
	FuzzRuns      int    `toml:"fuzz_runs" json:"fuzz_runs,omitempty"`
	InvariantRuns int    `toml:"invariant_runs" json:"invariant_runs,omitempty"`
	Verbosity     int    `toml:"verbosity" json:"verbosity,omitempty"`
	NoMatchTest   string `toml:"no_match_test" json:"no_match_test,omitempty"`
	MatchTest     string `toml:"match_test" json:"match_test,omitempty"`
	MatchContract string `toml:"match_contract" json:"match_contract,omitempty"`
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

// FlagInput is how a build rule and an optional foundry.toml combine into
// forge flags. An empty explicit value means "not set on the rule". A key
// absent from the toml does not override the defaults: optimizer = false is
// only applied when the file actually contains that key, so a toml that only
// sets evm_version keeps the plugin's optimizer default.
type FlagInput struct {
	FoundryTOML          []byte
	Profile              string
	EVMVersion           string
	Optimize             string
	OptimizerRuns        string
	DefaultOptimize      bool
	DefaultOptimizerRuns int
}

// CompileFlags returns a shell-safe forge flag string.
func CompileFlags(in FlagInput) (string, error) {
	profile := in.Profile
	if profile == "" {
		profile = "default"
	}
	table, err := profileTable(in.FoundryTOML, profile)
	if err != nil {
		return "", err
	}

	evm := strings.TrimSpace(in.EVMVersion)
	if evm == "" {
		if v, ok := table["evm_version"]; ok {
			s, ok := v.(string)
			if !ok {
				return "", fmt.Errorf("evm_version is %T, want string", v)
			}
			evm = strings.TrimSpace(s)
		}
	}
	if evm != "" && !evmVersionRe.MatchString(evm) {
		return "", fmt.Errorf("evm_version %q has characters forge flags cannot carry safely", evm)
	}

	optimize := in.DefaultOptimize
	if in.Optimize != "" {
		switch in.Optimize {
		case "true":
			optimize = true
		case "false":
			optimize = false
		default:
			return "", fmt.Errorf("optimize must be true, false, or empty, got %q", in.Optimize)
		}
	} else if v, ok := table["optimizer"]; ok {
		b, ok := v.(bool)
		if !ok {
			return "", fmt.Errorf("optimizer is %T, want bool", v)
		}
		optimize = b
	}

	runs := in.DefaultOptimizerRuns
	if runs == 0 {
		runs = 100
	}
	if in.OptimizerRuns != "" {
		n, err := strconv.Atoi(in.OptimizerRuns)
		if err != nil || n < 0 {
			return "", fmt.Errorf("optimizer_runs %q is not a non-negative integer", in.OptimizerRuns)
		}
		runs = n
	} else if v, ok := table["optimizer_runs"]; ok {
		n, err := asInt(v)
		if err != nil {
			return "", fmt.Errorf("optimizer_runs: %w", err)
		}
		runs = n
	}

	var args []string
	if optimize {
		args = append(args, "--optimize", "--optimizer-runs", strconv.Itoa(runs))
	}
	if evm != "" {
		args = append(args, "--evm-version", evm)
	}
	return strings.Join(args, " "), nil
}

var evmVersionRe = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

func profileTable(data []byte, profile string) (map[string]interface{}, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return map[string]interface{}{}, nil
	}
	var raw struct {
		Profile map[string]map[string]interface{} `toml:"profile"`
	}
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("foundry.toml: %w", err)
	}
	if p, ok := raw.Profile[profile]; ok {
		return p, nil
	}
	if profile != "default" {
		if p, ok := raw.Profile["default"]; ok {
			return p, nil
		}
	}
	return map[string]interface{}{}, nil
}

func asInt(v interface{}) (int, error) {
	switch n := v.(type) {
	case int64:
		return int(n), nil
	case int:
		return n, nil
	case uint64:
		return int(n), nil
	case float64:
		return int(n), nil
	default:
		return 0, fmt.Errorf("%T is not an integer", v)
	}
}
