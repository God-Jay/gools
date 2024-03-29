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

// NewPlugin requires the installation path of the plugin and the build argument.
// Set arg to "" to use default build option.
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
	pluginFullName := reg.FindString(goInstallPath)

	// your-plugin@latest etc.
	return strings.Split(pluginFullName, "@")[0]
}

func getPluginName(goInstallPath string) string {
	pluginFullName := getPluginFullName(goInstallPath)
	return strings.Replace(pluginFullName, "protoc-gen-", "", 1)
}
