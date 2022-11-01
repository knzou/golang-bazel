load("@io_bazel_rules_go//go:def.bzl", "go_library")
go_library(
    name = "router",
    srcs = ["router.go"],
    importpath = "monorepo/src/go/services/video_capture/router",
    visibility = ["//visibility:public"],
    deps = [
         "//services/lib/handlers/health:health",
          "@com_github_gin_gonic_gin//:go_default_library"
    ]
)