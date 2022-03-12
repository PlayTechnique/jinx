package jinkiesengine

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func NewJinkies(info ContainerInfo) Jinkies {
	return Jinkies{info, context.Background()}
}

type Jinkies struct {
	ContainerInfo
	Ctx context.Context
}

type ContainerInfo struct {
	AutoRemove    bool
	ImageName     string
	ContainerName string
	ContainerPort nat.Port
	HostIp        string
	HostPort      string
	PullImages    bool
}

func (jinkies *Jinkies) RunRunRun(hostConfig container.HostConfig) container.ContainerCreateCreatedBody {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := jinkies.ImageName

	if jinkies.PullImages {
		out, err := cli.ImagePull(jinkies.Ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}

		defer out.Close()
		io.Copy(os.Stdout, out) // write to stdout
	}

	resp, err := cli.ContainerCreate(jinkies.Ctx, &container.Config{
		Image: imageName,
	}, &hostConfig,
		nil, nil, jinkies.ContainerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(jinkies.Ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)

	return resp
}

func (jinkies *Jinkies) StopGirl() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	if stopErr := cli.ContainerStop(jinkies.Ctx, jinkies.ContainerName, nil); err != nil {
		panic(stopErr)
	}
}
