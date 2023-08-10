/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 10:17
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RelationRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	feedRouter := router.Group("relation")
	{
		feedRouter.GET("")
	}
	return router
}
