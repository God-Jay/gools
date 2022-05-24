package protoc

import (
	"fmt"
	"github.com/god-jay/gools/grpc/util"
	"log"
	"os"
	"os/exec"
	"path"
)

// Build will install protoc and plugins in the `deps` directory, and then build to generate the pb.go files
// of the proto files in the protoDir directory.
func Build(protoDir string, pbDir string, proto3rdparty string, plugins ...*Plugin) error {
	depsPath := util.SetPath()

	err := util.InstallProtoc(depsPath)
	if err != nil {
		return err
	}

	err = util.InstallProtocGen(depsPath)
	if err != nil {
		return err
	}

	for _, plugin := range plugins {
		err = util.InstallPlugin(depsPath, plugin.InstallPath, plugin.FullName)
		if err != nil {
			return err
		}
	}

	return genProtoc(protoDir, pbDir, proto3rdparty, plugins...)
}

func genProtoc(protoDir string, pbDir string, proto3rdparty string, plugins ...*Plugin) error {
	log.Println("Generating protoc files ......")

	protoFiles, _ := util.FindProtoFiles(protoDir)
	for _, protoFile := range protoFiles {
		log.Println("Generating protoc for", protoFile)

		protocArgs := []string{
			"-I", path.Dir(protoFile), "-I", proto3rdparty,
			fmt.Sprintf("--go_out=paths=source_relative:%s", pbDir),
		}
		for _, plugin := range plugins {
			var arg string
			if plugin.Arg == "" {
				arg = fmt.Sprintf("--%s_out=paths=source_relative:%s", plugin.ArgName, pbDir)
			} else {
				arg = fmt.Sprintf("--%s_out=%s", plugin.ArgName, plugin.Arg)
			}
			protocArgs = append(protocArgs, arg)
		}

		protocArgs = append(protocArgs, protoFile)

		cmd := exec.Command(
			"protoc",
			protocArgs...,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	log.Println("Generating protoc files finished.")

	return nil
}
