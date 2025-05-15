package example

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImageAuth() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil{
		panic(err)
	}

	authConfig := types.AuthConfig{
		Username: "juzo",
		Password: "12345",
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil{
		panic(err)
	}
	
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePull(ctx, "my.private.docker:8080", types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil{
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)

}