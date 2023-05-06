package utils

import (
	"archive/tar"
	"context"
	"github.com/docker/docker/client"
	"io"
	"jinx/types"
	"log"
	"os"
	"path"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func CopyFromContainer(globalRuntime jinxtypes.JinxGlobalRuntime, topLevelDir string, pathToCopy string) error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return err
	}

	tarReader, _, err := cli.CopyFromContainer(ctx, globalRuntime.ContainerName, pathToCopy)

	// CopyFromContainer says it's our responsibility to close this file handle.
	defer func() error {
		err := tarReader.Close()

		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		return err
	}

	tr := tar.NewReader(tarReader)

	for {
		hdr, err := tr.Next()

		if err == io.EOF {
			break // End of archive
		}

		if err != io.EOF && err != nil {
			return err
		}

		outputPath := path.Join(topLevelDir, hdr.Name)
		fileInfo := hdr.FileInfo()

		switch hdr.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(outputPath, fileInfo.Mode())
		case tar.TypeReg:
			// Typically one would use `defer outputFile.Close()` here. However, the deferred file handle closures
			// will occur after the loop has exited, so open file handles will accrue through the lifetime of this
			// loop, a possible cause of file handle exhaustion.
			outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())

			if err != nil {
				outputFile.Close()
				return err
			}

			if _, err := io.Copy(outputFile, tr); err != nil {
				outputFile.Close()
				return err
			}

			outputFile.Close()
		}
	}

	return nil
}
