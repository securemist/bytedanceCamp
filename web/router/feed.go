/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 11:20
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/api"
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FeedRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	feedRouter := router.Group("feed")
	{
		feedRouter.GET("", api.GetFeed)
		feedRouter.GET("publish/list", middlewares.JWTAuth(), api.GetPublishList)
		feedRouter.POST("publish/action", middlewares.JWTAuth(), api.PublishVideo)
	}
	return router
}
