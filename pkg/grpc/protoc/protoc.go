package protoc

import (
	"fmt"
	"github.com/god-jay/gools/pkg/grpc/util"
	"log"
	"os"
	"os/exec"
	"path"
)

func Build(protoDir string, pbDir string, proto3rdparty string) error {
	depsPath := util.SetPath()

	err := util.InstallProtoc(depsPath)

	err = util.InstallProtocGen(depsPath)

	genProtoc(protoDir, pbDir, proto3rdparty)

	return err
}

func genProtoc(protoDir string, pbDir string, proto3rdparty string) {
	log.Println("Generating protoc files ......")

	protoFiles, _ := util.FindProtoFiles(protoDir)
	for _, protoFile := range protoFiles {
		log.Println("Generating protoc for", protoFile)

		cmd := exec.Command(
			"protoc",
			"-I", path.Dir(protoFile), "-I", proto3rdparty,
			fmt.Sprintf("--go_out=paths=source_relative:%s", pbDir),
			protoFile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Generating protoc files finished.")
}
