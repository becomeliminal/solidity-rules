subinclude("//build_defs:solidity")

# Ownable pattern
sol_contract(
    name = "ownable",
    src = "Ownable.sol",
    solc_version = "0.8.20",
    deps = ["//test:openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "ownable_test",
    src = "Ownable.t.sol",
    solc_version = "0.8.20",
    deps = [":ownable", "//test:forge-std", "//test:openzeppelin-contracts"],
)

# Pausable pattern
sol_contract(
    name = "pausable",
    src = "Pausable.sol",
    solc_version = "0.8.20",
    deps = ["//test:openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "pausable_test",
    src = "Pausable.t.sol",
    solc_version = "0.8.20",
    deps = [":pausable", "//test:forge-std", "//test:openzeppelin-contracts"],
)

# Factory pattern
sol_contract(
    name = "factory",
    src = "Factory.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "factory_test",
    src = "Factory.t.sol",
    solc_version = "0.8.20",
    deps = [":factory", "//test:forge-std"],
)
