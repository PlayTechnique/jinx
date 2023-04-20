package jinxengine

import (
	"github.com/stretchr/testify/assert"
	jinxtypes "jinx/types"
	"os"
	"strings"
	"testing"
)

var globalRuntime = jinxtypes.JinxGlobalRuntime{
	ContainerName: "containerofdestiny",
	PullImages:    false,
}

func TestCreateLayout(t *testing.T) {

	testDir, _ := os.MkdirTemp("", "")
	os.Chdir(testDir)
	defer os.RemoveAll(testDir)

	jinxDir := "jeffrey"

	_, _, err := Initialise("jinx", jinxDir)

	os.Chdir(jinxDir)

	assert.Nil(t, err)

	assert.FileExists(t, "Docker/Dockerfile")
	assert.FileExists(t, "version.txt")
	assert.FileExists(t, "configFiles/jinx.yml")
	assert.FileExists(t, "configFiles/containerconfig.yml")
	assert.FileExists(t, "configFiles/hostconfig.yml")

}

func TestVerifyStringEntry(t *testing.T) {

	testDir, _ := os.MkdirTemp("", "")
	os.Chdir(testDir)
	defer os.RemoveAll(testDir)

	jinxDir := "jeffrey"
	containerName := "flibble"
	_, _, err := Initialise(containerName, jinxDir)

	os.Chdir(jinxDir)

	assert.Nil(t, err)

	jinxConfig := "configFiles/jinx.yml"
	assert.FileExists(t, jinxConfig)

	content, err := os.ReadFile(jinxConfig)

	assert.True(t, strings.Contains(string(content), containerName), jinxConfig+" should contain "+containerName)

}
