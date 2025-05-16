package example

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)


func RunContainerBackground() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil{
		panic(err)
	}
	
	imageName := "docker.io/library/alpine"

	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd: []string{"echo", "hello word"},
		Tty: false,
	}, nil, nil, nil, "FreshContainer")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil{
		panic(err)
	}

	fmt.Println(resp.ID)
}