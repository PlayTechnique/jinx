package jinkiesengine

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/viper"
	"io"
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

func hydrateFromConfig[T any](configPath string, config *T) {

	viper.AddConfigPath("./")
	viper.SetConfigType("yml")
	viper.SetConfigName(configPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	viper.Unmarshal(&config)
}

func RunRunRun(containerName string, pullImages bool, containerConfigPath string, hostConfigPath string) container.ContainerCreateCreatedBody {

	containerConfig := container.Config{}
	hostConfig := container.HostConfig{}

	if containerConfigPath == "" {
		containerConfig = defaultContainerConfig
	} else {
		hydrateFromConfig(containerConfigPath, &containerConfig)
	}

	if hostConfigPath == "" {
		hostConfig = defaultHostConfig
	} else {
		hydrateFromConfig(hostConfigPath, &hostConfig)
	}

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
