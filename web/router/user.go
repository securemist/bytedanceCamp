/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 11:12
* @description:
********************************************************************************/

package router

import (
	"bytedanceCamp/web/api"
	"bytedanceCamp/web/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRouter() *gin.Engine {
	router := gin.Default()
	// 健康检查
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	// 配置跨域
	router.Use(middlewares.Cors())
	userRouter := router.Group("user")
	{
		userRouter.GET("", middlewares.JWTAuth(), api.GetUserInfo)
		userRouter.POST("register", api.CreateUser)
		userRouter.POST("login", api.LoginCheck)
	}
	return router
}
