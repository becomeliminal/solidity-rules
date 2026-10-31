subinclude("//build_defs:solidity")

# Simple storage contract with Go bindings
sol_contract(
    name = "simple_storage",
    src = "SimpleStorage.sol",
    solc_version = "0.8.20",
    contract_names = ["SimpleStorage"],
    languages = ["go"],
    visibility = ["PUBLIC"],
)

# Solidity test for the contract
sol_test(
    name = "simple_storage_test",
    src = "SimpleStorage.t.sol",
    solc_version = "0.8.20",
    deps = [":simple_storage", "//test:forge-std"],
)

# External library. extract-decls has to see `library`, or this has no Go file.
sol_contract(
    name = "mathlib",
    src = "MathLib.sol",
    solc_version = "0.8.20",
    contract_names = ["MathLib"],
    languages = ["go"],
    visibility = ["PUBLIC"],
)

# Go test using generated bindings
go_test(
    name = "bindings_test",
    srcs = ["bindings_test.go"],
    deps = [
        ":simple_storage",
        "//third_party/solidity/go:go-ethereum",
    ],
)

go_test(
    name = "mathlib_binding_test",
    srcs = ["mathlib_test.go"],
    deps = [
        ":mathlib",
        "//third_party/solidity/go:go-ethereum",
    ],
)
