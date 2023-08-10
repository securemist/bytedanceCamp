/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:26
* @description:
********************************************************************************/

package initialize

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/model/proto/douyin_extra_first"
	"bytedanceCamp/model/proto/douyin_extra_second"
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
		zap.S().Errorf("[InitSrvConn] 连接 [%s失败]: %s", global.ProjectConfig.ConsulService.User.Name, err.Error())
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
		zap.S().Errorf("[InitSrvConn] 连接 [%s失败]: %s", global.ProjectConfig.ConsulService.Feed.Name, err.Error())
		return
	}
	global.FeedSrvClient = douyin_core.NewFeedClient(feedConn)
	// 连接favorite-srv
	favoriteConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ProjectConfig.Consul.Host,
			global.ProjectConfig.Consul.Port,
			global.ProjectConfig.ConsulService.Favorite.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[InitSrvConn] 连接 [%s失败]: %s", global.ProjectConfig.ConsulService.Favorite.Name, err.Error())
		return
	}
	global.FavoriteSrvClient = douyin_extra_first.NewFavoriteClient(favoriteConn)
	// 连接comment-srv
	commentConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ProjectConfig.Consul.Host,
			global.ProjectConfig.Consul.Port,
			global.ProjectConfig.ConsulService.Comment.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[InitSrvConn] 连接 [%s失败]: %s", global.ProjectConfig.ConsulService.Comment.Name, err.Error())
		return
	}
	global.CommentSrvClient = douyin_extra_first.NewCommentClient(commentConn)
	// 连接relation-srv
	relationConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ProjectConfig.Consul.Host,
			global.ProjectConfig.Consul.Port,
			global.ProjectConfig.ConsulService.Relation.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[InitSrvConn] 连接 [%s失败]: %s", global.ProjectConfig.ConsulService.Relation.Name, err.Error())
		return
	}
	global.RelationSrvClient = douyin_extra_second.NewRelationClient(relationConn)
}
