package example

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func StopContainer(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil{
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil{
		panic(err)
	}
	
	for _, container := range containers{
		fmt.Print("Stopping container", container.ID[:10], "...")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil{
			panic(err)
		}
		fmt.Println("Success")
	}

}