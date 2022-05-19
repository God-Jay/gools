package main

import (
	"github.com/god-jay/gools/grpc/3rdparty"
	"github.com/god-jay/gools/grpc/swagger"
)

// cd gools/_examples/gen-swagger
// go run main.go
func main() {
	proto3rdparty := "../../3rdparty"
	trdparty.Publish(proto3rdparty)

	protoDir := "./pb/proto"
	buildSwaggerDir := "./pb/swagger"

	err := swagger.Build(protoDir, buildSwaggerDir, proto3rdparty)
	if err != nil {
		panic(err)
	}
}
