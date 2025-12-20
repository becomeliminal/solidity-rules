subinclude("//build_defs:solidity")

# Events contract and tests
sol_library(
    name = "events_lib",
    src = "Events.sol",
    visibility = ["PUBLIC"],
)

sol_contract(
    name = "events",
    src = "Events.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "events_test",
    src = "Events.t.sol",
    solc_version = "0.8.20",
    deps = [":events_lib", "//test:forge-std"],
)

# Custom errors contract and tests
sol_library(
    name = "errors_lib",
    src = "Errors.sol",
    visibility = ["PUBLIC"],
)

sol_contract(
    name = "errors",
    src = "Errors.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "errors_test",
    src = "Errors.t.sol",
    solc_version = "0.8.20",
    deps = [":errors_lib", "//test:forge-std"],
)

# Fuzz testing
sol_test(
    name = "fuzz_test",
    src = "Fuzz.t.sol",
    solc_version = "0.8.20",
    deps = ["//test:forge-std"],
)

# Cheatcodes testing
sol_test(
    name = "cheatcodes_test",
    src = "Cheatcodes.t.sol",
    solc_version = "0.8.20",
    deps = ["//test:forge-std"],
)
