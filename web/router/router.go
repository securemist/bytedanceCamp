/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 11:12
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	apiGroup := router.Group("")
	InitUserRouter(apiGroup)
	InitFeedRouter(apiGroup)
	return router
}
