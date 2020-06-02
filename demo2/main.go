package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := client.NewEnvClient() // 连接本机docker  default: "unix:///var/run/docker.sock"
	if err != nil {
		log.Fatalln("client err: ",err)
	}

	// 获取目标Img
	out, err := client.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		log.Fatalln("image pull err: ",err)
	}
	io.Copy(os.Stdout,out)

	// 创建容器
	container, err := client.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd: []string{"sh"},
	}, nil, nil, "thisIsTest") // hostConfig,networkingConfig,containerName
	if err != nil {
		log.Fatalln("container create error: ",err)
	}

	// 启动容器
	if err := client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalln("Run container err: ",err)
	}
	fmt.Println(container.ID)
}