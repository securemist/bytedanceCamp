/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 11:20
* @description:
********************************************************************************/

package router

import "github.com/gin-gonic/gin"

func InitFeedRouter(router *gin.RouterGroup) {
	feedRouter := router.Group("feed")
	{
		// TODO 还没开始写路由，先占个位
		feedRouter.POST("")
	}
}
