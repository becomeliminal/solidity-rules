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

	"tools/please_sol/detectprefix"
	"tools/please_sol/forgewrap"
	"tools/please_sol/foundrytoml"
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
		ForgePath     string   `short:"f" long:"forge" required:"true" description:"Path to the forge binary"`
		RemappingFile string   `short:"r" long:"remapping-file" description:"Path to file containing remappings (one per line)"`
		Args          []string `positional-args:"true" description:"Arguments to pass to forge"`
	} `command:"forge-wrap" description:"Run forge with enhanced error messages"`

	ParseFoundry struct {
		File    string `short:"f" long:"file" required:"true" description:"Path to foundry.toml file"`
		Profile string `short:"p" long:"profile" description:"Profile to extract (default: default)"`
		Output  string `short:"o" long:"output" description:"Output format: json, solc-version, remappings, optimizer (default: json)"`
	} `command:"parse-foundry" description:"Parse foundry.toml and extract configuration"`
}{
	Usage: `
please_sol is used by the solidity build rules to perform complex parsing operations.

Supported commands:
  detect-prefix  Auto-detect import prefixes from package.json files
  forge-wrap     Run forge with enhanced error messages
  parse-foundry  Parse foundry.toml configuration files
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
		result, err := wrapper.Run(fw.ForgePath, fw.Args)
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
}

func main() {
	command := flags.ParseFlagsOrDie("please_sol", &opts, nil)
	os.Exit(subCommands[command]())
}
