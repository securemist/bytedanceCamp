/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:26
* @description:
********************************************************************************/

package initialize

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model/proto/douyin_core"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 这行代码很重要
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initSrvConn() {
	// 连接user-srv
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ProjectConfig.Consul.Host,
			global.ProjectConfig.Consul.Port,
			global.ProjectConfig.ConsulService.User.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Error("[InitSrvConn] 连接 [use-srv失败]", err.Error())
		return
	}
	global.UserSrvClient = douyin_core.NewUserClient(userConn)
	// 连接feed-srv
	feedConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ProjectConfig.Consul.Host,
			global.ProjectConfig.Consul.Port,
			global.ProjectConfig.ConsulService.Feed.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Error("[InitSrvConn] 连接 [feed-srv失败]", err.Error())
		return
	}
	global.FeedSrvClient = douyin_core.NewFeedClient(feedConn)
}
