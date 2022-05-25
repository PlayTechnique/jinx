package jinkiesengine

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"io"
	"jinx/src/utils"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	defaultContainerConfig = container.Config{
		ExposedPorts: nat.PortSet{"8090/tcp": {}},
		Image:        "jamandbees/jinkies",
	}

	defaultHostConfig = container.HostConfig{
		AutoRemove:   true,
		PortBindings: nat.PortMap{"8080/tcp": {{HostIP: "0.0.0.0", HostPort: "8090/tcp"}}},
	}
)

func RunRunRun(containerName string, pullImages bool, containerConfigPath string, hostConfigPath string) (*container.ContainerCreateCreatedBody, error) {

	containerConfig := container.Config{}
	hostConfig := container.HostConfig{}

	if containerConfigPath == "" {
		containerConfig = defaultContainerConfig
	} else {
		utils.HydrateFromConfig(containerConfigPath, &containerConfig)
	}

	if hostConfigPath == "" {
		hostConfig = defaultHostConfig
	} else {
		utils.HydrateFromConfig(hostConfigPath, &hostConfig)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	imageName := containerConfig.Image

	if pullImages {
		out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			return nil, err
		}

		defer out.Close()
		io.Copy(os.Stdout, out) // write to stdout
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, containerName)
	if err != nil {
		return nil, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	fmt.Println(resp.ID)

	return &resp, err
}

func StopGirl(containerName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	if stopErr := cli.ContainerStop(ctx, containerName, nil); err != nil {
		return stopErr
	}

	return nil
}
