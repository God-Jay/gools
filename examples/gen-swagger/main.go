package main

import (
	trdparty "github.com/god-jay/gools/pkg/grpc/3rdparty"
	"github.com/god-jay/gools/pkg/grpc/swagger"
)

// cd gools/examples/gen-swagger
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
