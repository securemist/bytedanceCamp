/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 11:21
* @description:
********************************************************************************/

package main

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model/proto/douyin_extra_first"
	"bytedanceCamp/service"
	"bytedanceCamp/util"
	"bytedanceCamp/util/initialize"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 初始化项目
	initialize.Init()
	// 注册grpc服务
	server := grpc.NewServer()
	douyin_extra_first.RegisterFavoriteServer(server, &service.FavoriteServer{})
	port, _ := util.GetFreePort()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", port))
	if err != nil {
		zap.S().Errorf("failed to listen: %s", err.Error())
	}
	go func() {
		err = server.Serve(lis)
		if err != nil {
			zap.S().Errorf("failed to start grpc: %s", err.Error())
		}
	}()
	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	// 注册favorite-srv服务
	registerClient := util.NewRegistryServiceClient(global.ProjectConfig.Consul.Host, global.ProjectConfig.Consul.Port)
	ServiceId := fmt.Sprintf("%d", util.GenID())
	err = registerClient.Register(
		"127.0.0.1",
		port,
		global.ProjectConfig.ConsulService.Favorite.Name,
		global.ProjectConfig.ConsulService.Favorite.Tags,
		ServiceId,
	)
	if err != nil {
		zap.S().Errorf("注册%s服务失败: %s", global.ProjectConfig.ConsulService.Favorite.Name, err.Error())
	} else {
		zap.S().Infof("注册%s服务成功: %s:%d", global.ProjectConfig.ConsulService.Favorite.Name, "127.0.0.1", port)
	}
	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = registerClient.DeRegister(ServiceId)
	if err != nil {
		zap.S().Errorf("注销%s服务失败: %s", global.ProjectConfig.ConsulService.Favorite.Name, err.Error())
	} else {
		zap.S().Infof("注销%s服务成功: %s:%d", global.ProjectConfig.ConsulService.Favorite.Name, "127.0.0.1", port)
	}
}
