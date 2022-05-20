package protoc

import (
	"regexp"
	"strings"
)

type Plugin struct {
	InstallPath string
	Arg         string

	FullName string
	ArgName  string
}

// NewPlugin
// set arg to "" to use default option
func NewPlugin(installPath string, arg string) *Plugin {
	return &Plugin{
		InstallPath: installPath,
		Arg:         arg,

		FullName: getPluginFullName(installPath),
		ArgName:  getPluginName(installPath),
	}
}

func getPluginFullName(goInstallPath string) string {
	reg, _ := regexp.Compile("protoc-gen-([^/]*)")
	return reg.FindString(goInstallPath)
}

func getPluginName(goInstallPath string) string {
	pluginFullName := getPluginFullName(goInstallPath)
	return strings.Replace(pluginFullName, "protoc-gen-", "", 1)
}
