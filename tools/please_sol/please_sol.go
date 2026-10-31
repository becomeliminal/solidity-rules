// please_sol is a tool for Solidity build operations.
// Used by the solidity Please plugin for complex parsing operations.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterebden/go-cli-init/v5/flags"

	"tools/please_sol/artifacts"
	"tools/please_sol/decls"
	"tools/please_sol/detectprefix"
	"tools/please_sol/forgewrap"
	"tools/please_sol/foundrytoml"
	"tools/please_sol/results"
)

var opts = struct {
	Usage string

	DetectPrefix struct {
		ZipPath string `short:"z" long:"zip" required:"true" description:"Path to the zip file containing the Solidity library"`
		Package string `short:"p" long:"package" required:"true" description:"The package directory within the repository"`
		Name    string `short:"n" long:"name" required:"true" description:"The name of the sol_get target"`
		OutDir  string `short:"o" long:"out_dir" required:"true" description:"The output directory (PKG path)"`
	} `command:"detect-prefix" description:"Detect import prefix from a Solidity library's package.json"`

	ForgeWrap struct {
		ForgePath     string `short:"f" long:"forge" required:"true" description:"Path to the forge binary"`
		RemappingFile string `short:"r" long:"remapping-file" description:"Path to file containing remappings (one per line)"`
		// positional-args belongs on a struct. A slice field with that tag is
		// ignored, and ParseFlagsOrDie then rejects forge's own flags.
		Args struct {
			Rest []string `description:"Arguments passed to forge"`
		} `positional-args:"yes" required:"yes"`
	} `command:"forge-wrap" description:"Run forge with enhanced error messages"`

	ParseFoundry struct {
		File    string `short:"f" long:"file" required:"true" description:"Path to foundry.toml file"`
		Profile string `short:"p" long:"profile" description:"Profile to extract (default: default)"`
		Output  string `short:"o" long:"output" description:"Output format: json, solc-version, remappings, optimizer (default: json)"`
	} `command:"parse-foundry" description:"Parse foundry.toml and extract configuration"`

	ExtractDecls struct {
		Paths struct {
			Files []string `description:"Solidity files or directories"`
		} `positional-args:"yes" required:"yes"`
	} `command:"extract-decls" description:"List contract, interface, and library declarations"`

	ExtractArtifacts struct {
		Out  string `long:"out" required:"true" description:"Forge out directory"`
		Dest string `long:"dest" required:"true" description:"Directory for .abi and .bin files"`
	} `command:"extract-artifacts" description:"Flatten Forge artifacts to ABI and bytecode files"`

	ForgeResults struct {
		JSON string `long:"json" required:"true" description:"Path to forge test --json output"`
		Out  string `long:"out" required:"true" description:"Go test results file to write"`
	} `command:"forge-results" description:"Convert forge test JSON to Go test results"`

	ResolveFlags struct {
		FoundryTOML          string `long:"foundry-toml" description:"Optional foundry.toml"`
		Profile              string `long:"profile" description:"Profile (default: default)"`
		EVMVersion           string `long:"evm-version" description:"Rule evm_version, wins over the toml"`
		Optimize             string `long:"optimize" description:"true, false, or empty"`
		OptimizerRuns        string `long:"optimizer-runs" description:"Rule optimizer_runs, or empty"`
		DefaultOptimize      string `long:"default-optimize" description:"Plugin optimizer default: true or false"`
		DefaultOptimizerRuns int    `long:"default-optimizer-runs" description:"Plugin optimizer_runs default"`
		DefaultEVMVersion    string `long:"default-evm-version" description:"Plugin evm_version default, or empty for solc's own"`
	} `command:"resolve-flags" description:"Print forge compiler flags"`
}{
	Usage: `
please_sol is used by the solidity build rules to perform complex parsing operations.

Supported commands:
  detect-prefix      Auto-detect import prefixes from package.json files
  forge-wrap         Run forge with enhanced error messages
  parse-foundry      Parse foundry.toml configuration files
  extract-decls      List contract, interface, and library declarations
  extract-artifacts  Flatten Forge artifacts to ABI and bytecode files
  forge-results      Convert forge test JSON to Go test results
  resolve-flags      Print forge compiler flags
`,
}

var subCommands = map[string]func() int{
	"detect-prefix": func() int {
		dp := opts.DetectPrefix
		detector := detectprefix.New(dp.ZipPath, dp.Package, dp.Name)
		prefix, err := detector.Detect()
		if err != nil {
			log.Fatalf("failed to detect import prefix: %v", err)
		}
		fmt.Println(detectprefix.FormatRemapping(prefix, dp.OutDir, dp.Name))
		return 0
	},
	"forge-wrap": func() int {
		fw := opts.ForgeWrap

		// Load remappings from file if provided
		var remappings []string
		if fw.RemappingFile != "" {
			f, err := os.Open(fw.RemappingFile)
			if err == nil {
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					line := strings.TrimSpace(scanner.Text())
					if line != "" {
						remappings = append(remappings, line)
					}
				}
				f.Close()
			}
		}

		wrapper := forgewrap.New(remappings)
		result, err := wrapper.Run(fw.ForgePath, fw.Args.Rest)
		if err != nil {
			log.Fatalf("failed to run forge: %v", err)
		}

		// Output stdout
		if result.Stdout != "" {
			fmt.Print(result.Stdout)
		}

		// Output stderr (enhanced if applicable)
		if result.Enhanced != "" {
			fmt.Fprint(os.Stderr, result.Enhanced)
		} else if result.Stderr != "" {
			fmt.Fprint(os.Stderr, result.Stderr)
		}

		return result.ExitCode
	},
	"parse-foundry": func() int {
		pf := opts.ParseFoundry

		config, err := foundrytoml.ParseFile(pf.File)
		if err != nil {
			log.Fatalf("failed to parse foundry.toml: %v", err)
		}

		profile := pf.Profile
		if profile == "" {
			profile = "default"
		}

		output := pf.Output
		if output == "" {
			output = "json"
		}

		switch output {
		case "json":
			jsonBytes, err := config.ToJSON()
			if err != nil {
				log.Fatalf("failed to convert to JSON: %v", err)
			}
			fmt.Println(string(jsonBytes))

		case "solc-version":
			version := config.GetSolcVersion(profile)
			if version != "" {
				fmt.Println(version)
			}

		case "remappings":
			remappings := config.GetRemappings(profile)
			for _, r := range remappings {
				fmt.Println(r)
			}

		case "optimizer":
			enabled, runs := config.GetOptimizerSettings(profile)
			if enabled {
				fmt.Printf("true %d\n", runs)
			} else {
				fmt.Println("false 0")
			}

		case "evm-version":
			evmVersion := config.GetEvmVersion(profile)
			if evmVersion != "" {
				fmt.Println(evmVersion)
			}

		default:
			log.Fatalf("unknown output format: %s", output)
		}

		return 0
	},
	"extract-decls": func() int {
		found, err := decls.ExtractPaths(opts.ExtractDecls.Paths.Files)
		if err != nil {
			log.Fatalf("extract-decls: %v", err)
		}
		fmt.Print(decls.Format(found))
		return 0
	},
	"extract-artifacts": func() int {
		a := opts.ExtractArtifacts
		if err := artifacts.Extract(a.Out, a.Dest); err != nil {
			log.Fatalf("extract-artifacts: %v", err)
		}
		return 0
	},
	"forge-results": func() int {
		f := opts.ForgeResults
		in, err := os.Open(f.JSON)
		if err != nil {
			log.Fatalf("forge-results: %v", err)
		}
		defer in.Close()
		out, err := os.Create(f.Out)
		if err != nil {
			log.Fatalf("forge-results: %v", err)
		}
		defer out.Close()
		if err := results.Convert(in, out); err != nil {
			log.Fatalf("forge-results: %v", err)
		}
		return 0
	},
	"resolve-flags": func() int {
		rf := opts.ResolveFlags
		var data []byte
		if rf.FoundryTOML != "" {
			var err error
			data, err = os.ReadFile(rf.FoundryTOML)
			if err != nil {
				log.Fatalf("resolve-flags: %v", err)
			}
		}
		defaultOptimize, err := parseBoolFlag("default-optimize", rf.DefaultOptimize, true)
		if err != nil {
			log.Fatalf("resolve-flags: %v", err)
		}
		flags, err := foundrytoml.CompileFlags(foundrytoml.FlagInput{
			FoundryTOML:          data,
			Profile:              rf.Profile,
			EVMVersion:           rf.EVMVersion,
			Optimize:             rf.Optimize,
			OptimizerRuns:        rf.OptimizerRuns,
			DefaultOptimize:      defaultOptimize,
			DefaultOptimizerRuns: rf.DefaultOptimizerRuns,
			DefaultEVMVersion:    rf.DefaultEVMVersion,
		})
		if err != nil {
			log.Fatalf("resolve-flags: %v", err)
		}
		fmt.Println(flags)
		return 0
	},
}

func parseBoolFlag(name, value string, fallback bool) (bool, error) {
	switch value {
	case "", "default":
		return fallback, nil
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("%s must be true or false, got %q", name, value)
	}
}

func main() {
	command := flags.ParseFlagsOrDie("please_sol", &opts, nil)
	os.Exit(subCommands[command]())
}
