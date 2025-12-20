subinclude("//build_defs:solidity")

# OpenZeppelin ERC20
sol_contract(
    name = "oz_erc20",
    src = "OzERC20.sol",
    solc_version = "0.8.20",
    deps = ["//test:openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "oz_erc20_test",
    src = "OzERC20.t.sol",
    solc_version = "0.8.20",
    deps = [":oz_erc20", "//test:forge-std", "//test:openzeppelin-contracts"],
)

# OpenZeppelin ERC721
sol_contract(
    name = "oz_erc721",
    src = "OzERC721.sol",
    solc_version = "0.8.20",
    deps = ["//test:openzeppelin-contracts"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "oz_erc721_test",
    src = "OzERC721.t.sol",
    solc_version = "0.8.20",
    deps = [":oz_erc721", "//test:forge-std", "//test:openzeppelin-contracts"],
)

# Solmate ERC20
sol_contract(
    name = "solmate_erc20",
    src = "SolmateERC20.sol",
    solc_version = "0.8.20",
    deps = ["//test:solmate"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "solmate_erc20_test",
    src = "SolmateERC20.t.sol",
    solc_version = "0.8.20",
    deps = [":solmate_erc20", "//test:forge-std", "//test:solmate"],
)

# Solady ERC20
sol_contract(
    name = "solady_erc20",
    src = "SoladyERC20.sol",
    solc_version = "0.8.20",
    deps = ["//test:solady"],
    languages = [],
    visibility = ["PUBLIC"],
)

sol_test(
    name = "solady_erc20_test",
    src = "SoladyERC20.t.sol",
    solc_version = "0.8.20",
    deps = [":solady_erc20", "//test:forge-std", "//test:solady"],
)
