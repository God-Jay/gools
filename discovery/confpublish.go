package discovery

import (
	"embed"
	"github.com/god-jay/gools/publisher"
)

//go:embed etc/config_server.yaml
var serverConf embed.FS

//go:embed etc/config_client.yaml
var clientConf embed.FS

//go:embed etc/example_*
var exampleConf embed.FS

func PublishServerConf(dstDir string) error {
	return publisher.CopyTo(dstDir, serverConf)
}

func PublishClientConf(dstDir string) error {
	return publisher.CopyTo(dstDir, clientConf)
}

// PublishExampleConf publish the config examples in the ./etc directory
func PublishExampleConf(dstDir string) error {
	return publisher.CopyTo(dstDir, exampleConf)
}
