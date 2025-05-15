package example

import (
	"context"
	"fmt"
	"os"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)


func RunContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil{
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io.library/alpine", types.ImagePullOptions{})
	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd: []string{"echo", "hello word"},
		Tty: false,
	}, nil, nil, nil, "FreshContainer")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil{
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <- errCh:
		if err != nil{
			panic(err)
		}
	case status := <- statusCh:
		fmt.Println("Status", status)
	}

}