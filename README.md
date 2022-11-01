# Monorepo
Consist of all the backend services

Why are we using Bazel's build system and not "make"?
- depedencies are tracked much better
- build every application separately
- build dependencies / tools in other languages than Go
[Source](https://www.reddit.com/r/golang/comments/bxpso2/whats_the_deal_with_using_bazel_for/)

# Setup

1) Install Bazel - https://docs.bazel.build/versions/master/install.html
2) Install Go Modules
```shell script
go mod download
```

# Build
All applications inside apps folder:
1) my Memories App
2) video capture app

## my memories App

Build:
```shell script from go dir
bazel build services/my_memories/main
```
Run:
```
bazel run services/my_memories/main
```
Bin will be here:

` bazel-bin/services/my_memories/darwin_amd64_stripped/main`

Test:
```shell script
bazel test services/my_memories/...:all
```
## Video Capture App

Build:
```shell script from go dir
bazel build services/video_capture/main
```
Run:
```
bazel run services/video_capture/main
```
Bin will be here:

`bazel-bin/services/video_capture/darwin_amd64_stripped/main`

Test:
```shell script
bazel test services/video_capture/...:all
```

# All application

Test:
```shell script
bazel test //services/...:all
```