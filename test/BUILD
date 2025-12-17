subinclude("//build_defs:solidity")

# Test svm rule (downloads svm binary)
svm(
    name = "svm",
    version = "0.5.22",
    visibility = ["PUBLIC"],
)

# Test solc rule (uses svm to download solc)
solc(
    name = "solc_0.8.20",
    version = "0.8.20",
    svm_tool = ":svm",
    visibility = ["PUBLIC"],
)

# Test sol_contract (uses default version download)
sol_contract(
    name = "counter",
    src = "Counter.sol",
    solc_version = "0.8.20",
    contract_names = ["Counter"],
    languages = [],  # No Go bindings in basic test
    visibility = ["PUBLIC"],
)

# Test sol_contract with local solc
sol_contract(
    name = "counter_local",
    src = "Counter.sol",
    solc_tool = ":solc_0.8.20",
    contract_names = ["Counter"],
    languages = [],
    visibility = ["PUBLIC"],
)

# Test sol_test (uses default version download)
sol_test(
    name = "counter_test",
    src = "Counter.t.sol",
    solc_version = "0.8.20",
    deps = [":counter"],
)

# Test sol_test with local solc
sol_test(
    name = "counter_test_local",
    src = "Counter.t.sol",
    solc_tool = ":solc_0.8.20",
    deps = [":counter_local"],
)
