/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 11:12
* @description:
********************************************************************************/

package router

import "github.com/gin-gonic/gin"

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	{
		// TODO 还没开始写路由，先占个位
		userRouter.POST("")
	}
}
