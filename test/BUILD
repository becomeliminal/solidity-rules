subinclude("//build_defs:solidity")

# Forge standard library - tests automatic import remappings
# No contract_names → defaults to languages=[] (no Go bindings)
sol_get(
    name = "forge-std",
    repo = "foundry-rs/forge-std",
    # v1.9.4
    revision = "1eea5bae12ae557d589f9f0f0edae2faa47cb262",
    hashes = ["sha256:f709d2f1811e5048c6c207ca3d1454f27f21c98104e95e8505afc2607750b786"],
    package = "src",
    visibility = ["PUBLIC"],
)

# OpenZeppelin - tests auto-detection (should detect "@openzeppelin/contracts/")
# No contract_names → defaults to languages=[] (no Go bindings)
sol_get(
    name = "openzeppelin-contracts",
    repo = "OpenZeppelin/openzeppelin-contracts",
    # v5.0.2
    revision = "dbb6104ce834628e473d2173bbc9d47f81a9eec3",
    hashes = ["sha256:bee78702745be0611eb14b4c5d06972c2ac0f138b62a26c06fcda8618a37030b"],
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
    # v6
    revision = "a9e3ea26a2dc73bfa87f0cb189687d029028e0c5",
    hashes = ["sha256:3ac66fd23e9ab279eb2f18d0369311201e11716099f8630a57e01b65bc713731"],
    package = "src",
    install = ["tokens", "utils"],
    visibility = ["PUBLIC"],
)

# Solady - tests auto-detection (should detect "solady/")
sol_get(
    name = "solady",
    repo = "Vectorized/solady",
    # v0.0.227
    revision = "1f43cc8005cc3b3c8361dd7dbdd2cdeaf0f99e66",
    hashes = ["sha256:9ac6c5c4f023c4a2764a3e4ca23f98da1aeb15cd74b5d10efcdf75add13781ff"],
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

# Hardhat - tests strip parameter to remove test/sample-projects directories
# that contain old Solidity versions incompatible with the main console.sol
sol_get(
    name = "hardhat",
    repo = "NomicFoundation/hardhat",
    revision = "1c7f8b766052101911df43d7b748fca12bcbf996",
    hashes = ["sha256:4877fb770cd415dac2e450b28af57f85d2ceb60fc4a0b4deee302b6433d0f331"],
    package = "packages/hardhat-core",
    strip = ["test", "sample-projects"],
    contract_names = ["Console"],
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)
