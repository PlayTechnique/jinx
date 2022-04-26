package utils

import (
	"archive/tar"
	"context"
	"github.com/docker/docker/client"
	"io"
	"jinx/types"
	"os"
	"path"
)

func CopyFromContainer(globalRuntime jinxtypes.JinxData, topLevelDir string, pathToCopy string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	tarReader, _, err := cli.CopyFromContainer(ctx, globalRuntime.ContainerName, pathToCopy)

	// CopyFromContainer says it's our responsibility to close this file handle.
	defer func() {
		err := tarReader.Close()

		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	tr := tar.NewReader(tarReader)

	for {
		hdr, err := tr.Next()

		if err == io.EOF {
			break // End of archive
		}

		if err != io.EOF && err != nil {
			panic(err)
		}

		outputPath := path.Join(topLevelDir, hdr.Name)
		fileInfo := hdr.FileInfo()

		switch hdr.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(outputPath, fileInfo.Mode())
		case tar.TypeReg:
			outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())

			if err != nil {
				outputFile.Close()
				panic(err)
			}

			if _, err := io.Copy(outputFile, tr); err != nil {
				outputFile.Close()
				panic(err)
			}

			outputFile.Close()
		}
	}
}
