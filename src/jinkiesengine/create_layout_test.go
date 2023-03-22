package jinkiesengine

import (
	"github.com/stretchr/testify/assert"
	jinxtypes "jinx/types"
	"os"
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
	err := CreateLayout(testDir, globalRuntime)

	assert.Nil(t, err)

	assert.FileExists(t, "Docker/Dockerfile")
	assert.FileExists(t, "version.txt")
	assert.FileExists(t, "configFiles/jinx.yml")
	assert.FileExists(t, "configFiles/containerconfig.yml")
	assert.FileExists(t, "configFiles/hostconfig.yml")
}
