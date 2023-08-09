/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 11:20
* @description:
********************************************************************************/

package api

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model/proto/douyin_extra_first"
	"bytedanceCamp/web/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func FavoriteAction(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	videoId, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	actType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	_, err := global.FavoriteSrvClient.FavoriteAction(context.Background(),
		&douyin_extra_first.FavoriteActionRequest{
			VideoId:    videoId,
			UserId:     userId,
			ActionType: int32(actType),
		})
	if actType == 1 {
		if err != nil {
			zap.S().Errorf("点赞失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "点赞成功",
		})
	} else {
		if err != nil {
			zap.S().Errorf("取消点赞失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "取消点赞成功",
		})
	}
}

func FavoriteList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	res, err := global.FavoriteSrvClient.FavoriteList(context.Background(),
		&douyin_extra_first.FavoriteListRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("获取喜欢列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res.VideoList,
	})
}
