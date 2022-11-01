load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test")
go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "monorepo/src/go/services/video_capture/main",
    deps = [
        "//services/video_capture/router:router"
    ],
)
go_test(
    name = "mainTest",
    srcs = ["main_test.go"],
)