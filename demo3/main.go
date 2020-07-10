package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func main() {
	// init
	ctx := context.Background()
	cli, err := client.NewEnvClient() // 连接本地docker
	if err != nil {
		log.Fatalln(err)
	}

	// 获取镜像列表
	list, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _,v := range list {
		fmt.Printf("ID: %s NAME: %s TAG:%s SIZE: %d\n",v.ID,v.RepoDigests,v.RepoTags,v.Size)
	}


	// pull image
	// pull image with authentication
	authConfig := types.AuthConfig{
		Username: "xxxx",
		Password: "xxxx",
	}
	authJsonByte, err := json.Marshal(authConfig)
	if err != nil {
		log.Fatalln(err)
	}
	authStr := base64.URLEncoding.EncodeToString(authJsonByte)
	authStr = authStr
	/**
	cli.ImagePull(ctx,"alpine",types.ImagePullOptions{RegistryAuth: authStr}) // 私有仓库  0.0.0.0:8080/xxx:xxx
	 */


	// Commit container
	/*
		// create container
	createResp,err := cli.ContainerCreate(ctx,&container.Config{
		Image: "apline",
	},nil,nil,"") // hostConfig,networkingConfig,containerName
	if err != nil {
		log.Fatalln(err)
	}
		// commit container
	commit, err := cli.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{
		Reference: "helloworld",
		Author: authStr,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(commit.ID)
	*/

	// 获取网络
	networkList, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _,v := range networkList{
		fmt.Printf("ID: %s, Name: %s DRIVER: %s SCOPE: %s \n",v.ID,v.Name,v.Driver,v.Scope)
	}

	// 轻量级的跨网络通信可以采用  overlay (非加密效率损失48%加密效率损失99%)   在同一台机器还是可以承受的
	// 在 生产环境 使用跨机 网络 推荐使用 Calico overlay 插件 效率87%
	//cli.NetworkCreate(ctx,"cpx",types.NetworkCreate{Driver: "Calico overlay"})
}
