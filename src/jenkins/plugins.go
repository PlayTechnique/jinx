package jenkins

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"jinx/src/utils"
	jinxtypes "jinx/types"
	"os"
	"path/filepath"
	"regexp"
	"syscall"
	"text/template"
)

type PluginMetadata struct {
	ModelVersion string `xml:"modelVersion"`
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
}

type pluginsData struct {
	Collection []PluginMetadata
}

func (self *pluginsData) gatherPlugins(path string, info fs.DirEntry, err error) error {
	var plugin PluginMetadata

	if info.Name() == "pom.xml" {
		content, err := os.ReadFile(path)

		if err != nil {
			panic(err)
		}

		xml.Unmarshal(content, &plugin)
		self.Collection = append(self.Collection, plugin)
	}

	return nil
}

func Plugins(globalRuntime jinxtypes.JinxData, topLevelDir string, outputFormat string) {
	validOutputs, err := regexp.Compile("plugins\\.txt|build\\.gradle")

	if err != nil {
		panic(err)
	}

	if !validOutputs.MatchString(outputFormat) {
		fmt.Println(fmt.Errorf("Valid output formats are plugins.txt and build.gradle. You supplied %s", outputFormat))
		syscall.Exit(2)
	}

	// ToDo: pathToCopy should be populated by a call inside the container userland to `exec`.
	pathToCopy := "/var/jenkins_home/plugins"
	removePlugins := false
	pluginsCollection := pluginsData{}

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

	err = filepath.WalkDir(topLevelDir, pluginsCollection.gatherPlugins)

	if err != nil {
		panic(err)
	}

	var formatter = template.New(outputFormat)

	switch outputFormat {
	case "plugins.txt":
		template.Must(formatter.Parse("{{.ArtifactId}} {{.Version}}\n"))
	case "build.gradle":
		template.Must(formatter.Parse("implementation group: '{{.GroupId}}', " +
			"name: '{{.ArtifactId}}',  version: '{{.Version}}'\n"))
	}

	for _, plugin := range pluginsCollection.Collection {
		err = formatter.Execute(os.Stdout, plugin)
		if err != nil {
			panic(err)
		}

	}
}
