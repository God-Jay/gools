package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func SetPath() string {
	pwd, _ := os.Getwd()
	depsPath := pwd + "/deps"
	if strings.Contains(os.Getenv("PATH"), depsPath) {
		return depsPath
	}
	os.Setenv("GOBIN", depsPath)
	os.Setenv("PATH", depsPath+"/:"+os.Getenv("PATH"))

	return depsPath
}

func InstallProtoc(depsPath string) error {
	_, err := os.Stat(depsPath + "/protoc")
	if err == nil {
		return nil
	}

	log.Println("Installing protoc ......")
	wgetCmd := exec.Command(
		"wget", fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/%s/%s", protocVer, protocZip),
	)
	wgetCmd.Stdout = os.Stdout
	wgetCmd.Stderr = os.Stderr
	err = wgetCmd.Run()

	unzipCmd := exec.Command(
		"unzip", "-j", protocZip, "bin/protoc", "-d", depsPath,
	)
	unzipCmd.Stderr = os.Stderr
	err = unzipCmd.Run()

	exec.Command("rm", protocZip).Run()

	return err
}

func InstallProtocGen(depsPath string) error {
	_, err := os.Stat(depsPath + "/protoc-gen-install-success")
	if err == nil {
		return nil
	}

	log.Println("Installing protoc-gen-go ......")
	command := exec.Command(
		"go", "install",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
	)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()

	if err == nil {
		err = os.WriteFile(depsPath+"/protoc-gen-install-success", nil, os.ModePerm)
	}

	return err
}

func FindProtoFiles(pbDir string) ([]string, error) {
	log.Println("Finding proto files ......")
	cmd := exec.Command("find", pbDir, "-name", "*.proto")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Run()

	var protoFiles []string
	for _, line := range strings.Split(out.String(), "\n") {
		if line != "" {
			protoFiles = append(protoFiles, line)
		}
	}

	return protoFiles, nil
}
