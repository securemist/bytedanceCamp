/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 11:21
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/api"
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	favoriteRouter := router.Group("favorite")
	{
		favoriteRouter.POST("action", middlewares.JWTAuth(), api.FavoriteAction)
		favoriteRouter.GET("list", middlewares.JWTAuth(), api.FavoriteList)
	}
	return router
}
