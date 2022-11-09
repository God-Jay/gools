package main

import (
	trdparty "github.com/god-jay/gools/grpc/3rdparty"
	"github.com/god-jay/gools/grpc/protoc"
)

func main() {
	proto3rdparty := "../../3rdparty"
	trdparty.Publish(proto3rdparty)

	protoDir := "../../pb/proto"
	buildPbDir := "../../pb"

	err := protoc.Build(protoDir, buildPbDir, proto3rdparty)
	if err != nil {
		panic(err)
	}

	err = protoc.Build(protoDir, buildPbDir, proto3rdparty,
		protoc.NewPlugin("github.com/god-jay/gools/cmd/protoc-gen-gin-http@latest", ""))
	if err != nil {
		panic(err)
	}

	err = protoc.Build(protoDir, buildPbDir, proto3rdparty,
		protoc.NewPlugin("google.golang.org/grpc/cmd/protoc-gen-go-grpc", ""))
	if err != nil {
		panic(err)
	}
}
