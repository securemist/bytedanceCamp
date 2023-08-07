package main

import (
	"bytedanceCamp/config"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/service"
	"bytedanceCamp/util"
	"bytedanceCamp/util/log"
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
	// 1. 初始化日志
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	// 注册grpc服务
	server := grpc.NewServer()
	douyin_core.RegisterUserServer(server, &service.UserServer{})
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
	// 注册user-srv服务
	registerClient := util.NewRegistryClient(config.GetConfig().Consul.Host, config.GetConfig().Consul.Port)
	ServiceId := fmt.Sprintf("%d", util.GenID())
	err = registerClient.Register(
		"127.0.0.1",
		port,
		config.GetConfig().ConsulService.User.Name,
		config.GetConfig().ConsulService.User.Tags,
		ServiceId,
	)
	if err != nil {
		zap.S().Errorf("注册user-srv服务失败: %s", err.Error())
	} else {
		zap.S().Infof("user-srv服务注册成功: %s:%d", "127.0.0.1", port)
	}
	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = registerClient.DeRegister(ServiceId)
	if err != nil {
		zap.S().Errorf("user-srv服务注销失败: %s", err.Error())
	} else {
		zap.S().Infof("user-srv服务注销成功: %s:%d", "127.0.0.1", port)
	}
}