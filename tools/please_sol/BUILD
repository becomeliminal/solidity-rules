subinclude("///go//build_defs:go")

go_binary(
    name = "please_sol",
    srcs = ["please_sol.go"],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/solidity/go:go-cli-init",
        "//tools/please_sol/artifacts",
        "//tools/please_sol/decls",
        "//tools/please_sol/detectprefix",
        "//tools/please_sol/forgewrap",
        "//tools/please_sol/foundrytoml",
        "//tools/please_sol/results",
    ],
)
