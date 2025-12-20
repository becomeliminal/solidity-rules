subinclude("//build_defs:solidity")

# Helper library for local import tests
sol_library(
    name = "helper",
    src = "Helper.sol",
    visibility = ["PUBLIC"],
)

# Local import test
sol_contract(
    name = "local_import",
    src = "LocalImport.sol",
    solc_version = "0.8.20",
    deps = [":helper"],
    languages = [],
    visibility = ["PUBLIC"],
)

# forge-std import test
sol_test(
    name = "forge_std_test",
    src = "ForgeStdTest.t.sol",
    solc_version = "0.8.20",
    deps = ["//test:forge-std"],
)

# OpenZeppelin import test (auto-detects @openzeppelin/contracts/)
sol_contract(
    name = "oz_import",
    src = "OzImport.sol",
    solc_version = "0.8.20",
    deps = ["//test:openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

# Solmate import test (auto-detects solmate/)
sol_contract(
    name = "solmate_import",
    src = "SolmateImport.sol",
    solc_version = "0.8.20",
    deps = ["//test:solmate"],
    languages = [],
    visibility = ["PUBLIC"],
)

# Solady import test (auto-detects solady/)
sol_contract(
    name = "solady_import",
    src = "SoladyImport.sol",
    solc_version = "0.8.20",
    deps = ["//test:solady"],
    languages = [],
    visibility = ["PUBLIC"],
)

# Multiple library imports
sol_contract(
    name = "multi_lib",
    src = "MultiLib.sol",
    solc_version = "0.8.20",
    deps = [
        "//test:openzeppelin-contracts",
        "//test:solmate",
    ],
    languages = [],
    visibility = ["PUBLIC"],
)
