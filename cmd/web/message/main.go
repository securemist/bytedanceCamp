/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 18:01
* @description:
********************************************************************************/

package main

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/util"
	"bytedanceCamp/util/initialize"
	"bytedanceCamp/web/router"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 初始化
	initialize.Init()
	// 2. 初始化routers
	routers := router.MessageRouter()
	// 3. 注册grpc服务
	registerClient := util.NewRegistryWebClient(global.ProjectConfig.Consul.Host, global.ProjectConfig.Consul.Port)
	ServiceId := fmt.Sprintf("%d", util.GenID())
	err := registerClient.Register(
		"127.0.0.1",
		global.ProjectConfig.ConsulWeb.Message.Port,
		global.ProjectConfig.ConsulWeb.Message.Name,
		global.ProjectConfig.ConsulWeb.Message.Tags,
		ServiceId,
	)
	if err != nil {
		zap.S().Errorf("注册%s服务失败: %s", global.ProjectConfig.ConsulWeb.Message.Name, err.Error())
	} else {
		zap.S().Infof("注册%s服务成功: %s:%d", global.ProjectConfig.ConsulWeb.Message.Name, "127.0.0.1", global.ProjectConfig.ConsulWeb.Message.Port)
	}
	// 4. 启动服务
	go func() {
		if err := routers.Run(fmt.Sprintf("localhost:%d", global.ProjectConfig.ConsulWeb.Message.Port)); err != nil {
			zap.S().Errorf("启动失败: %s", err.Error())
		}
	}()
	// 5. 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = registerClient.DeRegister(ServiceId)
	if err != nil {
		zap.S().Errorf("注销%s服务失败: %s", global.ProjectConfig.ConsulWeb.Message.Name, err.Error())
	}
	zap.S().Infof("注销%s服务成功: %s:%d", global.ProjectConfig.ConsulWeb.Relation.Name, "127.0.0.1", global.ProjectConfig.ConsulWeb.Message.Port)
}
