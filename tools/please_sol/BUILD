subinclude("///go//build_defs:go")

go_binary(
    name = "please_sol",
    srcs = ["please_sol.go"],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:go-cli-init",
        "//tools/please_sol/detectprefix",
        "//tools/please_sol/forgewrap",
        "//tools/please_sol/foundrytoml",
    ],
)
