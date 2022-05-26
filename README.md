# The Gools Project

English | [简体中文](README.zh-cn.md)

Gools (go tools) is a set of go tools for building your applications more convenient.

## gools

The gools project provides several tools for you to build your application:

- [publisher](publisher) - a tool using `//go:embed` to publish package files to your project.
- [grpc](grpc) - grpc provides several grpc related tools.
    - [3rdparty](grpc/3rdparty) - 3rdparty provides a `Publish` api to publish the third_party proto files to your
      project.
    - [protoc](grpc/protoc) - protoc can be used to generate the proto files as well as installing and using third part
      plugins.
    - [swagger](grpc/swagger) - swagger can be used to generate the swagger files.
- [kafka](kafka) - a kafka consumer group package, you only need to focus on your own consume logic.

## installation

`go get -u github.com/god-jay/gools`

## examples

- [publisher](_examples/publisher)
- grpc
    - [protoc](_examples/gen-proto)
    - [protoc-with-plugins](_examples/gen-proto-with-plugins)
    - [swagger](_examples/gen-swagger)
- [kafka](_examples/kafka)
