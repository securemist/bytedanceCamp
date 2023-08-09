/********************************************************************************
* @author: Yakult
* @date: 2023/8/8 12:13
* @description:
********************************************************************************/

package api

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/web/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// 获取视频流
func GetFeed(ctx *gin.Context) {
	reqTime := ctx.DefaultQuery("last_time", "")
	var lastTime int64
	if reqTime == "" {
		lastTime = time.Now().Unix()
	} else {
		lastTime, _ = strconv.ParseInt(reqTime, 10, 64)
	}
	res, err := global.FeedSrvClient.GetFeed(context.Background(), &douyin_core.FeedRequest{
		LatestTime: &lastTime,
	})
	if err != nil {
		zap.S().Errorf("获取视频流出错: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res.VideoList,
	})
}

// 获取所有投稿
func GetPublishList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	res, err := global.FeedSrvClient.PublishList(context.Background(), &douyin_core.PublishListRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("获取投稿列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res.VideoList,
	})
}

// 投稿视频
func PublishVideo(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	var video model.VideoData
	if err := ctx.ShouldBind(&video); err != nil {
		zap.S().Errorf("投稿视频出错: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	_, err := global.FeedSrvClient.PublishVideo(context.Background(), &douyin_core.PublishVideoRequest{
		Data:   []byte(video.Data),
		Title:  video.Title,
		UserId: userId,
	})
	if err != nil {
		zap.S().Errorf("投稿视频出错: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "视频投稿成功",
	})
}
