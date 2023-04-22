package jinxengine

import (
	_ "embed"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	jinxtypes "jinx/types"
	"log"
	"os"
)

//go:embed embed_files/version.txt
var version []byte

//go:embed embed_files/Dockerfile
var dockerfile []byte

// Initialise first verifies if the directory exists. If it does, it returns os.ErrExist.
func Initialise(containerName string, topLevelDir string) (jinxtypes.JinxGlobalRuntime, []string, error) {

	if _, err := os.Stat(topLevelDir); errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir(topLevelDir, 0700)
		if err != nil {
			return jinxtypes.JinxGlobalRuntime{}, nil, err
		}
	} else {
		log.Fatal(topLevelDir + " already exists. Cowardly refusing to proceed...")
		return jinxtypes.JinxGlobalRuntime{}, nil, fmt.Errorf("Directory already exists: %s. Cowardly refusing to proceed.", topLevelDir)
	}

	globalRuntime, createdFiles, err := createFiles(topLevelDir, containerName)

	return globalRuntime, createdFiles, err
}

func createFiles(topLevelDir string, containerName string) (jinxtypes.JinxGlobalRuntime, []string, error) {
	var createdFiles []string
	var globalRuntime jinxtypes.JinxGlobalRuntime

	filename, err := writeDockerFile(topLevelDir+"/Docker", "Dockerfile")

	if err != nil {
		log.Fatal(err)
		return globalRuntime, createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeVersionFile(topLevelDir + "/version.txt")

	if err != nil {
		log.Fatal(err)
		return globalRuntime, createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	globalRuntime, filename, err = writeJinxConfig(topLevelDir+"/configFiles", "jinx.yml", containerName)

	if err != nil {
		log.Fatal(err)
		return globalRuntime, createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeContainerConfig(topLevelDir+"/configFiles", "containerconfig.yml")

	if err != nil {
		log.Fatal(err)
		return globalRuntime, createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	filename, err = writeHostConfig(topLevelDir+"/configFiles", "hostconfig.yml")

	if err != nil {
		log.Fatal(err)
		return globalRuntime, createdFiles, err
	}

	createdFiles = append(createdFiles, filename)

	return globalRuntime, createdFiles, err
}

func writeDockerFile(dir string, filename string) (string, error) {

	err := os.MkdirAll(dir, 0700)

	if err != nil {
		log.Fatal(err)
		return dir, err
	}

	err = os.WriteFile(dir+"/"+filename, []byte(dockerfile), 0700)

	if err != nil {
		log.Fatal(err)
		// return's on the next line
	}

	return dir + "/" + filename, err
}

func writeVersionFile(filename string) (string, error) {

	err := os.WriteFile(filename, []byte(version), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return filename, err
}

func writeJinxConfig(dir string, filename string, containerName string) (jinxtypes.JinxGlobalRuntime, string, error) {
	globalRuntime := jinxtypes.JinxGlobalRuntime{ContainerName: containerName, PullImages: false}

	// Convert the struct to YAML
	data, err := yaml.Marshal(&globalRuntime)

	if err != nil {
		fmt.Println(err)
		return globalRuntime, "", err
	}

	err = os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
		return globalRuntime, dir, err
	}

	jinxConfigPath := dir + "/" + filename

	err = os.WriteFile(jinxConfigPath, data, 0644)

	return globalRuntime, jinxConfigPath, err
}

func writeContainerConfig(dir string, filename string) (string, error) {
	config := `---
autoremove: true
containerport: "8080/tcp"
Env:
  #The JINKIES_* environment variables are documented in https://github.com/gwynforthewyn/jinkies/blob/main/Docker/build.env
  #- JINKIES_SEED_DESCRIPTION=
  #- JINKIES_SEED_JENKINSFILE=
hostip: "0.0.0.0"
hostport: "8090/tcp"
image: "jamandbees/jinkies:local"
`
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + filename, err

}

func writeHostConfig(dir string, filename string) (string, error) {
	config := `---
AutoRemove: true
PublishAllPorts: true
ExposedPorts: "8091/tcp"
PortBindings:
  8080/tcp:
    HostIp: "0.0.0.0"
    HostPort: "8091/tcp"
`
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+"/"+filename, []byte(config), 0700)

	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + filename, err
}
