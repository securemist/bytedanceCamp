/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 11:10
* @description:
********************************************************************************/

package api

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model/proto/douyin_extra_second"
	"bytedanceCamp/web/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func RelationAction(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	if actType == 1 {
		_, err := global.RelationSrvClient.RelationAction(context.Background(), &douyin_extra_second.RelationActionRequest{
			UserId:     userId,
			ToUserId:   toUserId,
			ActionType: 1,
		})
		if err != nil {
			zap.S().Errorf("关注失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "关注成功",
		})
	} else {
		_, err := global.RelationSrvClient.RelationAction(context.Background(), &douyin_extra_second.RelationActionRequest{
			UserId:     userId,
			ToUserId:   toUserId,
			ActionType: 2,
		})
		if err != nil {
			zap.S().Errorf("取消关注失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "取消关注成功",
		})
	}
}

func RelationFollowList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	res, err := global.RelationSrvClient.RelationFollowList(context.Background(), &douyin_extra_second.RelationFollowListRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("获取关注列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"关注列表": res.FollowList,
	})
}
func RelationFanList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	res, err := global.RelationSrvClient.RelationFansList(context.Background(), &douyin_extra_second.RelationFansListRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("获取粉丝列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"粉丝列表": res.FanList,
	})
}
func RelationFriendList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	res, err := global.RelationSrvClient.RelationFriendList(context.Background(), &douyin_extra_second.RelationFriendListRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("获取好友列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"好友列表": res.FriendList,
	})
}

func RelationAddFriend(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	_, err := global.RelationSrvClient.RelationAddFriend(context.Background(), &douyin_extra_second.RelationAddFriendRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	if err != nil {
		zap.S().Errorf("添加好友失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "添加好友成功",
	})
}
