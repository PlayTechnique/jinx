package jinkiesengine

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateLayout(t *testing.T) {
	testDir, _ := os.MkdirTemp("", "")
	os.Chdir(testDir)
	defer os.RemoveAll(testDir)
	err := CreateLayout(testDir)

	assert.Nil(t, err)

	assert.FileExists(t, "Docker/Dockerfile")
	assert.FileExists(t, "version.txt")
	assert.FileExists(t, "configFiles/jinx.yml")
	assert.FileExists(t, "configFiles/containerconfig.yml")
	assert.FileExists(t, "configFiles/hostconfig.yml")
}
