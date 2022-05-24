package swagger

import (
	"github.com/god-jay/gools/grpc/util"
	"log"
	"os"
	"os/exec"
	"path"
)

// Build will install protoc and plugins in the `deps` directory, and then build to generate the swagger.json files
// of the proto files in the protoDir directory.
func Build(protoDir string, swaggerDir string, proto3rdparty string) error {
	depsPath := util.SetPath()

	err := util.InstallProtoc(depsPath)
	if err != nil {
		return err
	}

	err = util.InstallProtocGen(depsPath)
	if err != nil {
		return err
	}

	return genSwagger(protoDir, swaggerDir, proto3rdparty)
}

func genSwagger(protoDir string, swaggerDir string, proto3rdparty string) error {
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
			return err
		}
	}

	log.Println("Generating swagger finished.")

	return nil
}
