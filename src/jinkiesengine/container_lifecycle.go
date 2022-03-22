package jinkiesengine

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunRunRun(containerName string, pullImages bool, containerConfig container.Config, hostConfig container.HostConfig) container.ContainerCreateCreatedBody {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := containerConfig.Image

	if pullImages {
		out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}

		defer out.Close()
		io.Copy(os.Stdout, out) // write to stdout
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)

	return resp
}

func StopGirl(containerName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	if stopErr := cli.ContainerStop(ctx, containerName, nil); err != nil {
		panic(stopErr)
	}
}
