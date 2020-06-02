package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)  // 获取容器拉取产生的输出到stdout
	log.Println("容器拉取成功")

	// 创建容器
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	log.Println("容器镜像创建成功")

	// 执行容器
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	log.Println("容器执行")

	// 等待容器运行结束
	wait, err := cli.ContainerWait(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	log.Println("容器运行结束 id: ",wait)

	// 获取容器中产生的stdout
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	log.Println("重定向输出")
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}