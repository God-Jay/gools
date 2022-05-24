# Gools 项目

[English](README.md) | 简体中文

Gools (go tools) 是一组开发工具包，让您更便利地进行应用开发。

## gools

gools 提供了几个开发工具，帮助您快速构建应用：

- [publisher](publisher) - 使用 `//go:embed` 生成文件的发布工具
- [grpc](grpc) - 包含了几个 grpc 相关的工具
    - [3rdparty](grpc/3rdparty) - 3rdparty 提供了一个 `Publish` 接口，可以将包中的 3rdparty 文件发布到自己的应用中
    - [protoc](grpc/protoc) - 用来将 proto 文件编译成 pb.go 文件，同时支持安装使用 protoc 插件
    - [swagger](grpc/swagger) - 用来将 proto 文件生成 swagger 文档
- [kafka](kafka) - 对 kafka consumer group 的封装，您只需要关注自己的业务逻辑

## 安装

`go get -u github.com/god-jay/gools`

## 示例

- [publisher](_examples/publisher)
- grpc
    - [protoc](_examples/gen-proto)
    - [protoc-with-plugins](_examples/gen-proto-with-plugins)
    - [swagger](_examples/gen-swagger)
- [kafka](_examples/kafka)
