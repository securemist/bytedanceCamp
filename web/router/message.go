/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 18:01
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/api"
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	relationRouter := router.Group("message").Use(middlewares.JWTAuth())
	{
		relationRouter.POST("send", api.MessageSend)
	}
	return router
}
