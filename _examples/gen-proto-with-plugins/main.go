package main

import (
	"github.com/god-jay/gools/grpc/3rdparty"
	"github.com/god-jay/gools/grpc/protoc"
)

// cd gools/_examples/gen-proto-with-plugins
// go run main.go
func main() {
	proto3rdparty := "../../3rdparty"
	trdparty.Publish(proto3rdparty)

	protoDir := "./pb/proto"
	buildPbDir := "./pb"

	err := protoc.Build(protoDir, buildPbDir, proto3rdparty,
		protoc.NewPlugin("github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2", ""),
	)
	if err != nil {
		panic(err)
	}
}
