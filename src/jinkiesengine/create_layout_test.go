package jinkiesengine

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io/fs"
	jinxtypes "jinx/types"
	"os"
	"path/filepath"
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
	err := CreateLayout(testDir, globalRuntime)

	assert.Nil(t, err)

	assert.FileExists(t, "Docker/Dockerfile")
	assert.FileExists(t, "version.txt")
	assert.FileExists(t, "configFiles/jinx.yml")
	assert.FileExists(t, "configFiles/containerconfig.yml")
	assert.FileExists(t, "configFiles/hostconfig.yml")
}

func TestCreateLayoutHasContentInAllVariables(t *testing.T) {
	testDir, _ := os.MkdirTemp("", "")
	os.Chdir(testDir)
	defer os.RemoveAll(testDir)
	err := CreateLayout(testDir, globalRuntime)

	assert.Nil(t, err)

	filepath.WalkDir(testDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		} else {
			f, err := os.Open(path)

			if err != nil {
				return err
			}

			defer func() error {
				err = f.Close()

				if err != nil {
					panic(err)
				}

				return nil
			}()

			// Splits on newlines by default.
			scanner := bufio.NewScanner(f)

			// https://golang.org/pkg/bufio/#Scanner.Scan
			for scanner.Scan() {
				line := scanner.Text()

				if strings.Contains(line, "{{") {
					assert.FailNowf(t, "File "+path+" contains a template in line <"+line+">", "All template variables are populated")
				}
			}
		}
		return nil
	})
}
