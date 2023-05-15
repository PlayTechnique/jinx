package jinxengine

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"io/fs"
	jinxtypes "jinx/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed embed_files/*
var jinxsupportembed embed.FS

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// Initialise first verifies if the directory exists. If it does, it returns os.ErrExist.
func Initialise(containerName string, topLevelDir string) (jinxtypes.JinxGlobalRuntime, error) {

	if _, err := os.Stat(topLevelDir); errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir(topLevelDir, 0755)
		if err != nil {
			return jinxtypes.JinxGlobalRuntime{}, err
		}

	} else {
		log.Print(topLevelDir + " already exists. Cowardly refusing to proceed...")
		return jinxtypes.JinxGlobalRuntime{}, fmt.Errorf("Directory already exists: %s. Cowardly refusing to proceed.", topLevelDir)
	}

	topLevelDir, _ = filepath.Abs(topLevelDir)
	// Return to this directory at the end of creating files.
	defer os.Chdir(topLevelDir)

	globalRuntime, err := createFiles(topLevelDir, containerName)

	return globalRuntime, err
}

// Parameters:
// 1. topLevelDir: absolute path to the top level directory that will contain this jinx project.
// 2. containerName: name of the container to populate in the various config files.
func createFiles(topLevelDir string, containerName string) (jinxtypes.JinxGlobalRuntime, error) {
	globalRuntime := jinxtypes.JinxGlobalRuntime{ContainerName: containerName}

	_, err := os.Stat(topLevelDir)

	if os.IsNotExist(err) {
		return globalRuntime, err
	}

	err = writeEmbedFs(jinxsupportembed, topLevelDir, globalRuntime)

	if err != nil {
		return globalRuntime, err
	}

	return globalRuntime, err
}

func writeEmbedFs(fsys embed.FS, topLevelDir string, globalRuntime jinxtypes.JinxGlobalRuntime) error {

	os.Chdir(topLevelDir)

	dirWalker := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// The embed.fs has embed_files as the top level directory for all paths.
		// Skip embed_files by itself, and strip it from everything else.
		if path == "." || path == "embed_files" {
			return nil
		}

		destPath := strings.TrimPrefix(path, "embed_files/")

		if d.IsDir() {
			err = os.MkdirAll(destPath, 0700)
			if err != nil {
				return err
			}
		} else {
			switch path {
			case "embed_files/configFiles/jinx.yml":
				templ, err := template.ParseFS(fsys, path)
				if err != nil {
					return err
				}

				outputFile, err := os.Create(destPath)
				if err != nil {
					return err
				}

				defer outputFile.Close()

				err = templ.Execute(outputFile, globalRuntime)

				if err != nil {
					return err
				}

			default:
				content, err := fs.ReadFile(fsys, path)

				if err != nil {
					return err
				}
				err = os.WriteFile(destPath, content, 0755)

				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	err := fs.WalkDir(fsys, "embed_files", dirWalker)

	if err != nil {
		return err
	}

	return nil
}
