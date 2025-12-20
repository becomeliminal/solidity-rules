go_library(
    name = "foundrytoml",
    srcs = ["foundrytoml.go"],
    visibility = ["//tools/please_sol/..."],
    deps = ["//third_party/go:go-toml-v2"],
)

go_test(
    name = "foundrytoml_test",
    srcs = ["foundrytoml_test.go"],
    deps = [":foundrytoml"],
)
