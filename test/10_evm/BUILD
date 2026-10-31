subinclude("//build_defs:solidity")

sol_contract(
    name = "push0_default",
    src = "Push0.sol",
    solc_version = "0.8.20",
    languages = [],
)

sol_contract(
    name = "push0_paris",
    src = "Push0.sol",
    solc_version = "0.8.20",
    evm_version = "paris",
    languages = [],
)

sol_contract(
    name = "push0_toml",
    src = "Push0.sol",
    solc_version = "0.8.20",
    foundry_toml = "paris.toml",
    languages = [],
)

gentest(
    name = "push0_test",
    data = [
        "check_push0.sh",
        ":_push0_default#forge_build",
        ":_push0_paris#forge_build",
        ":_push0_toml#forge_build",
    ],
    no_test_output = True,
    test_cmd = "$(location check_push0.sh)",
)
