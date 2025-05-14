package example

import (
	"context"
	"os"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImage() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil{
		panic(err)
	}

	out, err := cli.ImagePull(ctx, "alpine", types.ImagePullOptions{})
	if err != nil{
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)

}