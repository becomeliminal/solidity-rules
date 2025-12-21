# Solidity Please Plugin

A [Please](https://please.build) plugin for building Solidity smart contracts using [Foundry](https://github.com/foundry-rs/foundry).

## Features

- **sol_library()** - Create reusable Solidity libraries with transitive dependency tracking
- **sol_contract()** - Compile contracts with ABI/bytecode extraction and optional Go bindings
- **sol_get()** - Download third-party Solidity dependencies from GitHub
- **sol_test()** - Run Foundry tests with full dependency support

## Installation

### Prerequisites

- [Foundry](https://github.com/foundry-rs/foundry) must be installed (provides `forge` compiler)
- Install via: `curl -L https://foundry.paradigm.xyz | bash && foundryup`

### Setup

1. Add the plugin to your `plugins/BUILD`:

```python
plugin_repo(
    name = "solidity",
    owner = "becomeliminal",
    plugin = "solidity-rules",
    revision = "<commit-sha>",
)
```

2. Configure in `.plzconfig`:

```ini
[Plugin "solidity"]
Target = //plugins:solidity
```

3. (Optional) Configure Go bindings:

```ini
[Plugin "solidity"]
Target = //plugins:solidity
AbigenTool = //third_party/go:abigen
GoEthereumDep = //third_party/go:go-ethereum
```

4. (Optional) Use local forge via the [Foundry plugin](https://github.com/becomeliminal/foundry):

```ini
[Plugin "solidity"]
Target = //plugins:solidity
ForgeTool = //third_party/binary:foundry|forge
```

5. (Optional) Use local solc via the svm rule:

```python
# In third_party/solidity/BUILD
subinclude("///solidity//build_defs:solidity")

svm(
    name = "svm",
    version = "0.5.22",
    visibility = ["PUBLIC"],
)

solc(
    name = "solc_0.8.20",
    version = "0.8.20",
    visibility = ["PUBLIC"],
)
```

```ini
[Plugin "solidity"]
Target = //plugins:solidity
SvmTool = //third_party/solidity:svm
SolcTool = //third_party/solidity:solc_0.8.20
```

## Usage

### Basic Contract

```python
subinclude("///solidity//build_defs:solidity")

sol_contract(
    name = "mycontract",
    src = "MyContract.sol",
    solc_version = "0.8.20",
    visibility = ["PUBLIC"],
)
```

### With Go Bindings

```python
sol_contract(
    name = "mycontract",
    src = "MyContract.sol",
    solc_version = "0.8.20",
    contract_names = ["MyContract"],
    languages = ["go"],
    visibility = ["PUBLIC"],
)

# Use in Go code
go_binary(
    name = "myapp",
    srcs = ["main.go"],
    deps = [":mycontract"],  # Will resolve to Go bindings
)
```

### Third-Party Dependencies

```python
sol_get(
    name = "openzeppelin-contracts",
    repo = "OpenZeppelin/openzeppelin-contracts",
    revision = "v5.0.2",
    package = "contracts",
    install = ["."],  # Install all (needed for internal relative imports)
    visibility = ["PUBLIC"],
)

sol_contract(
    name = "mytoken",
    src = "MyToken.sol",
    solc_version = "0.8.20",
    deps = [":openzeppelin-contracts"],
)
```

Then in `MyToken.sol` (import prefix defaults to `{name}/`):

```solidity
import "openzeppelin-contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20 {
    constructor() ERC20("MyToken", "MTK") {}
}
```

### Running Tests

```python
sol_test(
    name = "mycontract_test",
    src = "MyContract.t.sol",
    solc_version = "0.8.20",
    deps = [":mycontract"],
)
```

Run with:

```bash
plz test //path/to:mycontract_test
```

### Solidity Library (for shared code)

```python
sol_library(
    name = "utils",
    src = "Utils.sol",
    visibility = ["PUBLIC"],
)

sol_contract(
    name = "mycontract",
    src = "MyContract.sol",
    solc_version = "0.8.20",
    deps = [":utils"],
)
```

## Configuration

All options can be set in `.plzconfig` under `[Plugin "solidity"]`:

| Option | Default | Description |
|--------|---------|-------------|
| `ForgeTool` | `~/.foundry/bin/forge` | Path or build label for forge (assumes foundryup install) |
| `SolcTool` | (none) | Build label for solc binary (from `solc()` rule) |
| `SvmTool` | (none) | Build label for svm binary (from `svm()` rule) |
| `DefaultSolcVersion` | `0.8.20` | Default Solidity version when not specified per-rule |
| `AbigenTool` | (none) | Build label for abigen (required for Go bindings) |
| `GoEthereumDep` | (none) | Build label for go-ethereum (required for Go bindings) |
| `DefaultLanguages` | `go` | Default output languages for sol_contract |
| `Optimize` | `true` | Enable Solidity optimizer |
| `OptimizerRuns` | `100` | Number of optimizer runs |
| `Sandbox` | `false` | Enable sandbox (requires local solc) |

## Rule Reference

### svm

Downloads pre-built svm binary for managing solc versions.

```python
svm(
    name = "svm",
    version = "0.5.22",    # svm-rs release version
    visibility = [],
)
```

### solc

Downloads a specific solc version using svm. Requires SvmTool to be configured.

```python
solc(
    name = "solc_0.8.20",
    version = "0.8.20",     # Solidity version to download
    visibility = [],
)
```

### sol_library

Creates a Solidity library that can be used as a dependency.

```python
sol_library(
    name = "lib",
    src = "Lib.sol",
    deps = [],           # Other sol_library or sol_contract rules
    test_only = False,
    visibility = [],
)
```

### sol_contract

Compiles a Solidity contract using Forge.

```python
sol_contract(
    name = "contract",
    src = "Contract.sol",
    deps = [],               # Dependencies
    solc_version = "0.8.20", # Solidity version (uses svm)
    solc_flags = "",         # Additional solc flags
    contract_names = [],     # For multi-contract files
    skip = [],               # Contracts to skip
    languages = ["go"],      # Output languages
    test_only = False,
    visibility = [],
)
```

### sol_get

Downloads Solidity libraries from GitHub. Import remappings are automatically generated.

```python
sol_get(
    name = "openzeppelin-contracts",
    repo = "OpenZeppelin/openzeppelin-contracts",
    revision = "v5.0.2",     # Git revision (prefer commit SHA)
    hashes = [],             # SHA256 hashes for verification (recommended!)
    import_prefix = "",      # Defaults to "{name}/" (e.g., "openzeppelin-contracts/")
    package = "contracts",   # Source directory in repo
    install = ["."],         # Subdirectories to install
    deps = [],               # Other sol_get dependencies
    visibility = [],
)

# Then in your contract:
# import "openzeppelin-contracts/token/ERC20/ERC20.sol";
```

### sol_test

Runs Solidity tests using Forge.

```python
sol_test(
    name = "test",
    src = "Contract.t.sol",
    deps = [],
    solc_version = "0.8.20",
    solc_flags = "",
    timeout = 0,
    labels = [],
    visibility = [],
)
```

## How It Works

1. **Compilation**: Uses Foundry's `forge build --use <version>` which leverages svm (Solidity Version Manager) to automatically download and cache the specified solc version.

2. **ABI/Bytecode Extraction**: After compilation, ABIs and bytecode are extracted from Forge's JSON output files.

3. **Go Bindings**: If configured, uses `abigen` to generate Go bindings from the compiled ABIs and bytecode.

4. **Dependency Tracking**: The plugin uses Please's `requires` and `provides` mechanism to track transitive Solidity dependencies.

## License

MIT
