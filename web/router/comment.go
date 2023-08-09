/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 18:28
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/api"
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	commentRouter := router.Group("comment")
	{
		commentRouter.POST("action", middlewares.JWTAuth(), api.CommentAction)
		commentRouter.GET("list", middlewares.JWTAuth(), api.CommentList)
	}
	return router
}
