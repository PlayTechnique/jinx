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

type ContainerInfo struct {
	autoRemove    bool
	imageName     string
	containerName string
	containerPort nat.Port
	hostIp        string
	hostPort      string
}

var jenkins = ContainerInfo{autoRemove: true, imageName: "jamandbees/jinkies", containerName: "jinkies",
	containerPort: "8080/tcp", hostIp: "0.0.0.0", hostPort: "8090"}

func RunRunRun() container.ContainerCreateCreatedBody {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := jenkins.imageName
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out) // write to stdout

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, &container.HostConfig{
		AutoRemove:   jenkins.autoRemove,
		PortBindings: nat.PortMap{jenkins.containerPort: {{HostIP: jenkins.hostIp, HostPort: jenkins.hostPort}}},
	}, nil, nil, jenkins.containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)

	return resp
}

func StopGirl() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	if stopErr := cli.ContainerStop(ctx, jenkins.containerName, nil); err != nil {
		panic(stopErr)
	}
}
