load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["feedback.go"],
    importpath = "github.com/golang/dep/internal/feedback",
    visibility = ["//vendor/github.com/golang/dep:__subpackages__"],
    deps = ["//vendor/github.com/golang/dep/gps:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
