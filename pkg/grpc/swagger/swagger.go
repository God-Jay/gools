package swagger

import (
	"github.com/god-jay/gools/pkg/grpc/util"
	"log"
	"os"
	"os/exec"
	"path"
)

func Build(protoDir string, swaggerDir string, proto3rdparty string) error {
	depsPath := util.SetPath()

	err := util.InstallProtoc(depsPath)

	err = util.InstallProtocGen(depsPath)

	genSwagger(protoDir, swaggerDir, proto3rdparty)

	return err
}

func genSwagger(protoDir string, swaggerDir string, proto3rdparty string) {
	log.Println("Generating swagger ......")

	protoFiles, _ := util.FindProtoFiles(protoDir)
	for _, protoFile := range protoFiles {
		log.Println("Generating swagger for", protoFile)

		cmd := exec.Command(
			"protoc",
			"-I", path.Dir(protoFile), "-I", proto3rdparty,
			"--openapiv2_out", swaggerDir,
			"--openapiv2_opt", "json_names_for_fields=false",
			protoFile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Generating swagger finished.")
}
