/********************************************************************************
* @author: Yakult
* @date: 2023/8/11 17:04
* @description:
********************************************************************************/

package api

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_extra_second"
	"bytedanceCamp/web/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func MessageSend(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	message := model.MessageData{}
	if err := ctx.ShouldBind(&message); err != nil {
		zap.S().Errorf("获取消息内容和对方id失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	_, err := global.MessageSrvClient.MessageSend(context.Background(), &douyin_extra_second.MessageSendRequest{
		UserId:   userId,
		ToUserId: message.ToUserId,
		Content:  message.Content,
	})
	if err != nil {
		zap.S().Errorf("消息发送失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "消息发送成功",
	})
}

func MessageChat(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	message := model.MessageData{}
	if err := ctx.ShouldBind(&message); err != nil {
		zap.S().Errorf("获取对方id失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 我发送给对方的消息
	res1, err := global.MessageSrvClient.MessageChat(context.Background(), &douyin_extra_second.MessageChatRequest{
		UserId:   userId,
		ToUserId: message.ToUserId,
	})
	if err != nil {
		zap.S().Errorf("获取我发送给对方的消息列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 对方发送给我的消息
	res2, err := global.MessageSrvClient.MessageChat(context.Background(), &douyin_extra_second.MessageChatRequest{
		UserId:   message.ToUserId,
		ToUserId: userId,
	})
	if err != nil {
		zap.S().Errorf("获取对方发送给我的消息列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"我发送的消息": res1.MessageList,
		"我接收的消息": res2.MessageList,
	})
}
