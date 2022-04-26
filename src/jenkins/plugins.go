package jenkins

import (
	"fmt"
	"io/fs"
	"jinx/src/utils"
	jinxtypes "jinx/types"
	"os"
	"path/filepath"
)

type pluginMetaInf struct {
	ManifestVersion       string `yaml:"Manifest-Version,omitempty"`
	HudsonVersion         string `yaml:"Hudson-Version,omitempty"`
	PluginDependencies    string `yaml:"Plugin-Dependencies,omitempty"`
	ImplementationTitle   string `yaml:"Implementation-Title,omitempty"`
	LongName              string `yaml:"Long-Name, omitempty"`
	ImplementationVersion string `yaml:"Implementation-Version,omitempty"`
	GroupId               string `yaml:"Group-Id,omitempty"`
	//MinimumJavaVersion string `yaml:"Minimum-Java-Version,ignore"`
	//PluginLicenseName string `yaml:"Plugin-License-Name,ignore"`
	//Specification-Title: Multi-configuration (matrix) project type.
	//Plugin-ScmUrl: https://github.com/jenkinsci/matrix-project-plugin
	PluginVersion string `yaml:"Plugin-Version,omitempty"`
	//Jenkins-Version: 2.289.1
	//Url: https://github.com/jenkinsci/matrix-project-plugin
	//Short-Name: matrix-project
	//Plugin-License-Url: https://opensource.org/licenses/MIT
	//Plugin-Developers:
	//Extension-Name: matrix-project
	//Build-Jdk-Spec: 1.8
	//Created-By: Maven Archiver 3.5.2
}

func Plugins(globalRuntime jinxtypes.JinxData, topLevelDir string) {
	// ToDo: pathToCopy should be populated by a call inside the container userland to `exec`.
	//manifests := []pluginMetaInf

	pathToCopy := "/var/jenkins_home/plugins"
	removePlugins := false

	if topLevelDir == "" {
		topLevelDir = os.TempDir()
		removePlugins = true
	}

	// We'll let the umask take care of security here.
	os.MkdirAll(topLevelDir, 0777)

	if removePlugins {
		defer os.RemoveAll(topLevelDir)
	}

	utils.CopyFromContainer(globalRuntime, topLevelDir, pathToCopy)

	err := filepath.WalkDir(topLevelDir, func(path string, info fs.DirEntry, err error) error {
		if info.Name() == "MANIFEST.MF" {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
