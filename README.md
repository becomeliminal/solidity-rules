# Solidity Please Plugin

A [Please](https://please.build) plugin for building Solidity smart contracts using [Foundry](https://github.com/foundry-rs/foundry).

## Features

- **sol_library()** - Create reusable Solidity libraries with transitive dependency tracking
- **sol_contract()** - Compile contracts with ABI/bytecode extraction and optional Go bindings
- **sol_get()** - Download third-party Solidity dependencies from GitHub
- **sol_test()** - Run Foundry tests with full dependency support

## Installation

### Setup

1. Add both plugins to your `plugins/BUILD`. Forge is the
   [foundry plugin](https://github.com/becomeliminal/foundry). This repo does
   not download Foundry itself.

```python
plugin_repo(
    name = "solidity",
    owner = "becomeliminal",
    plugin = "solidity-rules",
    revision = "v0.1.0",
)

plugin_repo(
    name = "foundry",
    owner = "becomeliminal",
    revision = "7d11211b950ace6e83bb18f5df048a5e7fc31ba7",
)
```

2. Put the third-party binaries in `third_party/binary/BUILD`. `foundry()`
   fetches anvil, forge, cast, and chisel. `hashes` lists the sha256 of each
   platform's `foundry_v1.5.0_<os>_<arch>.tar.gz`; the download must match one
   of them. `solc_versions` installs the pinned solc static binaries (0.8.20
   and 0.8.23).

```python
subinclude("///foundry//build_defs:foundry")
subinclude("///solidity//build_defs:solidity")

foundry(
    name = "foundry",
    version = "1.5.0",
    hashes = [
        "sha256:5cd98f9092bcc28be087939491f786b2bf3ed55e492996a409e29519b8ab4dc8",  # linux_amd64
        "sha256:8138e1615568bfcca5999773830892d93a569370eb0ae4b7dd97db46e2af47f9",  # linux_arm64
        "sha256:493b04c731b02b85d95f3bf9aff9cdddd0acd52c277902f3313fbe55e9b154e6",  # darwin_amd64
        "sha256:17748f7b1815a66d91cce88946989b049358a874262cb848284609dd9094df06",  # darwin_arm64
    ],
    visibility = ["PUBLIC"],
)

solc_versions(versions = ["0.8.20", "0.8.23"])
```

3. Configure in `.plzconfig`. Sandbox defaults to true, so forge and solc must
   be build targets rather than files on the machine:

```ini
[Plugin "foundry"]
Target = //plugins:foundry

[Plugin "solidity"]
Target = //plugins:solidity
ForgeTool = //third_party/binary:foundry|forge
SolcPackage = //third_party/binary
```

4. (Optional) Configure Go bindings:

```ini
[Plugin "solidity"]
Target = //plugins:solidity
AbigenTool = //third_party/go:abigen
GoEthereumDep = //third_party/go:go-ethereum
```

### Local development

A host `forge` is outside the sandbox. Point `ForgeTool` at it and turn the
sandbox off. With no `SolcPackage`, forge fetches solc itself.

```ini
[Plugin "solidity"]
Target = //plugins:solidity
ForgeTool = ~/.foundry/bin/forge
Sandbox = false
```

`SolcTool` is one compiler for every rule: each rule's `solc_version` is ignored.
`SolcPackage` takes precedence over `SolcTool`. A rule whose `solc_version`
(or `DefaultSolcVersion`) has no target in that package fails to build rather
than compiling with another version. `svm()` remains for an svm binary and
requires hashes for its release tarball. `solc()` does not call it.


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
| `ForgeTool` | (none) | The forge output of the foundry plugin, `//third_party/binary:foundry|forge`. A filesystem path needs `Sandbox = false` |
| `SolcTool` | (none) | Build label for one solc binary (from `solc()`), used for every rule |
| `SolcPackage` | (none) | Package of `solc_versions()` targets; each rule uses the one its `solc_version` names. Overrides `SolcTool` |
| `SvmTool` | (none) | Build label for an `svm()` binary. `solc()` does not use it |
| `DefaultSolcVersion` | `0.8.20` | Default Solidity version when not specified per-rule |
| `AbigenTool` | (none) | Build label for abigen (required for Go bindings) |
| `GoEthereumDep` | (none) | Build label for go-ethereum (required for Go bindings) |
| `DefaultLanguages` | `go` | Default output languages for sol_contract |
| `Optimize` | `true` | Enable Solidity optimizer |
| `OptimizerRuns` | `100` | Number of optimizer runs |
| `Sandbox` | `true` | Run forge inside the Please sandbox |

## Rule Reference

Forge is not a rule in this plugin. `foundry()` from the foundry plugin
provides `//third_party/binary:foundry|forge`.

### svm

Downloads a pre-built svm binary. Put the call in `third_party/binary/BUILD`.
`hashes` is required. `solc()` does not use svm.

```python
svm(
    name = "svm",
    version = "0.5.22",
    hashes = ["sha256:<sha256 of the release tarball>"],
    visibility = [],
)
```

### solc

Installs one pinned solc static binary for the host platform. Call it from
`third_party/binary/BUILD`. `solc_versions()` creates `solc_0_8_20` and
`solc_0_8_23`. Any other version fails until it is added to the pin table.

```python
solc(
    name = "solc_0_8_20",
    version = "0.8.20",
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
    solc_version = "0.8.20", # Solidity version (a solc_versions() target)
    solc_flags = "",         # Additional solc flags
    evm_version = "",        # Empty keeps the compiler default (Shanghai on 0.8.20, which emits PUSH0)
    foundry_toml = "",       # Used for evm_version and optimizer when the rule arguments are empty
    optimizer = None,        # True or False. Empty uses foundry.toml, then the plugin default
    optimizer_runs = None,
    contract_names = [],     # Contracts, interfaces, or libraries to bind
    skip = [],               # Contracts to skip
    languages = ["go"],      # Output languages
    test_only = False,
    visibility = [],
)
```

A failed compile fails the target. `contract_names` may name a library that has an external function; comments and strings are not treated as declarations.

### sol_get

Downloads Solidity libraries from GitHub. Import remappings are automatically generated.

```python
sol_get(
    name = "openzeppelin-contracts",
    repo = "OpenZeppelin/openzeppelin-contracts",
    revision = "v5.0.2",     # Git revision (prefer commit SHA)
    hashes = ["sha256:..."], # Required. Pin the commit archive; tag zips are not stable
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
    evm_version = "",
    foundry_toml = "",
    optimizer = None,
    optimizer_runs = None,
    timeout = 0,
    labels = [],
    visibility = [],
)
```

`sol_test` runs one `forge test`. A failing test fails the target. Please sees each test's revert reason and its real duration.

## How It Works

1. **Compilation**: `forge build` runs the `ForgeTool` binary. With `SolcPackage` or `SolcTool` set, `--use` is that solc and forge does not download one. The EVM version is the compiler default unless `evm_version` or `foundry.toml` sets one. Optimizer settings on the rule win over `foundry.toml`, which wins over the plugin default. A toml that only sets `evm_version` does not turn the optimizer off.

2. **ABI/Bytecode Extraction**: After a successful compile, ABIs and bytecode are extracted from Forge's JSON artifacts. A failed compile fails the target.

3. **Go Bindings**: If configured, uses `abigen` to generate Go bindings from the compiled ABIs and bytecode.

4. **Dependency Tracking**: The plugin uses Please's `requires` and `provides` mechanism to track transitive Solidity dependencies.

## License

MIT
