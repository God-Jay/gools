package main

import (
	"github.com/god-jay/gools/grpc/3rdparty"
	"github.com/god-jay/gools/grpc/protoc"
)

// cd gools/_examples/gen-proto
// go run main.go
func main() {
	proto3rdparty := "../../3rdparty"
	trdparty.Publish(proto3rdparty)

	protoDir := "./pb/proto"
	buildPbDir := "./pb"

	err := protoc.Build(protoDir, buildPbDir, proto3rdparty)
	if err != nil {
		panic(err)
	}
}
