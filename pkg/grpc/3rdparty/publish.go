package trdparty

import (
	"embed"
	"github.com/god-jay/gools/pkg/publisher"
)

//go:embed google
var google embed.FS

//go:embed protoc-gen-openapiv2
var protocGenOpenapiv2 embed.FS

//go:embed validate
var validate embed.FS

func Publish(dstDir string) error {
	err := publisher.CopyTo(dstDir, google)
	if err != nil {
		return err
	}
	err = publisher.CopyTo(dstDir, protocGenOpenapiv2)
	if err != nil {
		return err
	}
	return publisher.CopyTo(dstDir, validate)
}
