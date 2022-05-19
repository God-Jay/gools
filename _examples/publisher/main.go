package main

import (
	"embed"
	"github.com/god-jay/gools/publisher"
)

//go:embed src
var src embed.FS

// cd gools/_examples/publisher
// go run main.go
func main() {
	// publisher will create dst/src/...
	err := publisher.CopyTo("dst", src)
	if err != nil {
		panic(err)
	}
}
