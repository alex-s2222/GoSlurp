package example

import (
	"context"
	"io"
	"os"
	
	"github.com/docker/docker/client"
)

func GetStats(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil{
		panic(err)
	}
	

	stats, err := cli.ContainerStats(ctx, "92n3452n32", false)
	if err != nil{
		panic(err)
	}
	defer stats.Body.Close()
	io.Copy(os.Stdout, stats.Body)
}