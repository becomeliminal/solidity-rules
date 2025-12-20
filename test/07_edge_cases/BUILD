subinclude("//build_defs:solidity")

# Empty contract
sol_contract(
    name = "empty",
    src = "Empty.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

# Interface only
sol_contract(
    name = "interface",
    src = "Interface.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

# Abstract contract
sol_contract(
    name = "abstract",
    src = "Abstract.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

# Solidity library
sol_contract(
    name = "library",
    src = "Library.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

# Diamond inheritance
sol_contract(
    name = "diamond",
    src = "Diamond.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

# Receive and fallback
sol_contract(
    name = "receive",
    src = "Receive.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "receive_test",
    src = "Receive.t.sol",
    solc_version = "0.8.20",
    deps = [":receive", "//test:forge-std"],
)

# Edge case tests
sol_test(
    name = "empty_test",
    src = "Empty.t.sol",
    solc_version = "0.8.20",
    deps = [":empty", "//test:forge-std"],
)

sol_test(
    name = "interface_test",
    src = "Interface.t.sol",
    solc_version = "0.8.20",
    deps = [":interface", "//test:forge-std"],
)

sol_test(
    name = "abstract_test",
    src = "Abstract.t.sol",
    solc_version = "0.8.20",
    deps = [":abstract", "//test:forge-std"],
)

sol_test(
    name = "library_test",
    src = "Library.t.sol",
    solc_version = "0.8.20",
    deps = [":library", "//test:forge-std"],
)

sol_test(
    name = "diamond_test",
    src = "Diamond.t.sol",
    solc_version = "0.8.20",
    deps = [":diamond", "//test:forge-std"],
)
