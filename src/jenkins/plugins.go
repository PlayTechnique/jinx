package jenkins

import (
	"archive/tar"
	"context"
	"github.com/docker/docker/client"
	"io"
	jinxtypes "jinx/types"
	"log"
	"os"
	"path"
)

func Plugins(globalRuntime jinxtypes.JinxData) {
	// ToDo: pathToCopy should be populated by a call inside the container userland to `exec`.
	pathToCopy := "/var/jenkins_home/plugins"
	// ToDo: topLevelDir should be passed in as a cli argument, with a reasonable default
	topLevelDir := "tmp"

	f, err := os.OpenFile("./outfile.tar", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	tarReader, _, err := cli.CopyFromContainer(ctx, globalRuntime.ContainerName, pathToCopy)

	if err != nil {
		panic(err)
	}

	defer func() {
		err := tarReader.Close()

		if err != nil {
			panic(err)
		}
	}()

	tr := tar.NewReader(tarReader)

	for {
		hdr, err := tr.Next()

		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}

		outputPath := path.Join(topLevelDir, hdr.Name)
		fileInfo := hdr.FileInfo()

		switch hdr.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(outputPath, fileInfo.Mode())
		case tar.TypeReg:
			outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())
			defer outputFile.Close()
			if err != nil {
				panic(err)
			}

			if _, err := io.Copy(outputFile, tr); err != nil {
				panic(err)
			}
		}
	}
}
