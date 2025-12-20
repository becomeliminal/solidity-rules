subinclude("//build_defs:solidity")

# Basic counter contract - state, functions, events
sol_contract(
    name = "counter",
    src = "Counter.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "counter_test",
    src = "Counter.t.sol",
    solc_version = "0.8.20",
    deps = [":counter", "//test:forge-std"],
)

# Storage contract - all storage types
sol_contract(
    name = "storage",
    src = "Storage.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "storage_test",
    src = "Storage.t.sol",
    solc_version = "0.8.20",
    deps = [":storage", "//test:forge-std"],
)
