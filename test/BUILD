subinclude("//build_defs:solidity")

# Test svm rule (downloads svm binary)
svm(
    name = "svm",
    version = "0.5.22",
    visibility = ["PUBLIC"],
)

# Test solc rule (requires SvmTool config - see .plzconfig)
solc(
    name = "solc_0.8.20",
    version = "0.8.20",
    visibility = ["PUBLIC"],
)

# Forge standard library - tests automatic import remappings
# No contract_names → defaults to languages=[] (no Go bindings)
sol_get(
    name = "forge-std",
    repo = "foundry-rs/forge-std",
    revision = "v1.9.4",
    package = "src",
    visibility = ["PUBLIC"],
)

# OpenZeppelin - tests auto-detection (should detect "@openzeppelin/contracts/")
# No contract_names → defaults to languages=[] (no Go bindings)
sol_get(
    name = "openzeppelin-contracts",
    repo = "OpenZeppelin/openzeppelin-contracts",
    revision = "v5.0.2",
    package = "contracts",
    install = ["."],
    visibility = ["PUBLIC"],
)

# Test sol_contract (uses default version download or SolcTool if configured)
sol_contract(
    name = "counter",
    src = "Counter.sol",
    solc_version = "0.8.20",
    contract_names = ["Counter"],
    languages = [],  # No Go bindings in basic test
    visibility = ["PUBLIC"],
)

# Test contract using OpenZeppelin with default import_prefix
sol_contract(
    name = "mytoken",
    src = "MyToken.sol",
    solc_version = "0.8.20",
    deps = [":openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

# Test sol_test
sol_test(
    name = "counter_test",
    src = "Counter.t.sol",
    solc_version = "0.8.20",
    deps = [":counter", ":forge-std"],
)

# Solmate - v6 has @rari-capital/solmate in package.json, but please_sol
# auto-detects this and maps it to "solmate/" via knownMigrations
sol_get(
    name = "solmate",
    repo = "transmissions11/solmate",
    revision = "v6",
    package = "src",
    install = ["tokens", "utils"],
    visibility = ["PUBLIC"],
)

# Solady - tests auto-detection (should detect "solady/")
sol_get(
    name = "solady",
    repo = "Vectorized/solady",
    revision = "v0.0.227",
    package = "src",
    install = ["auth", "tokens", "utils"],
    visibility = ["PUBLIC"],
)

# Test sol_contract with go bindings (simple contract, no deps)
sol_contract(
    name = "counter_go",
    src = "Counter.sol",
    solc_version = "0.8.20",
    contract_names = ["Counter"],
    languages = ["go"],
    visibility = ["PUBLIC"],
)
