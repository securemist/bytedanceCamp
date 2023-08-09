/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 18:54
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

func CommentAction(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId := claims.(*middlewares.CustomClaims).Uuid
	videoId, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	actType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	if actType == 1 {
		commentText := ctx.Query("comment_text")
		res, err := global.CommentSrvClient.CommentAction(context.Background(), &douyin_extra_first.CommentActionRequest{
			VideoId:     videoId,
			UserId:      userId,
			ActionType:  1,
			CommentText: &commentText,
		})
		if err != nil {
			zap.S().Errorf("评论视频失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res.Comment,
		})
	} else {
		commentId, _ := strconv.ParseInt(ctx.Query("comment_id"), 10, 64)
		_, err := global.CommentSrvClient.CommentAction(context.Background(), &douyin_extra_first.CommentActionRequest{
			VideoId:    videoId,
			UserId:     userId,
			ActionType: 2,
			CommentId:  &commentId,
		})
		if err != nil {
			zap.S().Errorf("删除评论失败: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "评论删除成功",
		})
	}
}

func CommentList(ctx *gin.Context) {
	videoId, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	res, err := global.CommentSrvClient.CommentList(context.Background(), &douyin_extra_first.CommentListRequest{VideoId: videoId})
	if err != nil {
		zap.S().Errorf("获取评论列表失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"评论": res.CommentList,
	})
}
