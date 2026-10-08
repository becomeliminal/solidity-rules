subinclude("//build_defs:solidity")

# Each rule compiles with the solc its solc_version names (SolcPackage). The
# pragmas are exact, so a rule built with any other version fails to compile.
sol_contract(
    name = "pinned_0820",
    src = "Pinned0820.sol",
    solc_version = "0.8.20",
    languages = [],
    visibility = ["PUBLIC"],
)

sol_contract(
    name = "pinned_0823",
    src = "Pinned0823.sol",
    solc_version = "0.8.23",
    languages = [],
    visibility = ["PUBLIC"],
)

# The bytecode itself says which solc compiled it, independent of the pragmas.
gentest(
    name = "solc_versions_test",
    data = [
        "check_versions.sh",
        ":_pinned_0820#forge_build",
        ":_pinned_0823#forge_build",
    ],
    no_test_output = True,
    test_cmd = "$(location check_versions.sh)",
)
