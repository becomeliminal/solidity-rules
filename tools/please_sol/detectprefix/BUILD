subinclude("///go//build_defs:go")

go_library(
    name = "detectprefix",
    srcs = ["detectprefix.go"],
    visibility = ["//tools/please_sol/..."],
)

go_test(
    name = "detectprefix_test",
    srcs = ["detectprefix_test.go"],
    deps = [":detectprefix"],
)
