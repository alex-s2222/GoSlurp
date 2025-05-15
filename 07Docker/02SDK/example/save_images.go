package example

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func SaveImage(imageId string, filePath string){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil{
		panic(err)
	}
	images := make([]string, 1)
	images[0] = imageId
	
	reader, err := cli.ImageSave(ctx, images)
	if err != err{
		panic(err)
	}
	defer reader.Close()

	file, err := os.Create(filePath)
	writterBytes, err := io.Copy(file, reader)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Bytes written %v", writterBytes)
}