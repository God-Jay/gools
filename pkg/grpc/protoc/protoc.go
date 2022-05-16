package protoc

import (
	"fmt"
	"github.com/god-jay/gools/pkg/grpc/util"
	"log"
	"os"
	"os/exec"
	"path"
)

func Build(protoDir string, pbDir string, proto3rdparty string, plugins ...string) error {
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
		err = util.InstallPlugin(depsPath, plugin)
		if err != nil {
			return err
		}
	}

	var pluginNames []string
	for _, plugin := range plugins {
		pluginName := util.GetPluginName(plugin)
		pluginNames = append(pluginNames, pluginName)
	}

	return genProtoc(protoDir, pbDir, proto3rdparty, pluginNames...)
}

func genProtoc(protoDir string, pbDir string, proto3rdparty string, pluginNames ...string) error {
	log.Println("Generating protoc files ......")

	protoFiles, _ := util.FindProtoFiles(protoDir)
	for _, protoFile := range protoFiles {
		log.Println("Generating protoc for", protoFile)

		protocArgs := []string{
			"-I", path.Dir(protoFile), "-I", proto3rdparty,
			fmt.Sprintf("--go_out=paths=source_relative:%s", pbDir),
		}
		for _, pluginName := range pluginNames {
			protocArgs = append(protocArgs, fmt.Sprintf("--%s_out=paths=source_relative:%s", pluginName, pbDir))
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
