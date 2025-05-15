package example

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	"fmt"
)


func ListAllImages() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	filtersList := filters.NewArgs()
	// filtersList.Add("label", "Myserver")

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All: true,
		Filters: filtersList,
	})
	if err != nil{
		panic(err)
	}
	
	for _, image := range images{
		fmt.Printf("Image from repo %s with ID %s\n", image.RepoTags[0], image.ID)
	} 

}
