/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 20:31
* @description:
********************************************************************************/

package api

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_core"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// 根据user_id查询用户信息
func GetUserInfo(ctx *gin.Context) {
	userId, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	res, err := global.UserSrvClient.GetUserInfo(context.Background(), &douyin_core.UserInfoRequest{UserId: userId})
	if err != nil {
		zap.S().Errorf("[GetUserInfo] 查询用户信息失败: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": fmt.Sprintf("查询%d用户信息失败", userId),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res.UserInfo,
		"followings": res.UserInfo.FollowCount,
	})
}

// 用户注册
func CreateUser(ctx *gin.Context) {
	registerForm := model.User{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		zap.S().Errorf("注册用户出错: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	user, err := global.UserSrvClient.CreateUser(context.Background(), &douyin_core.UserRegisterRequest{
		Username: registerForm.UserName,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("注册用户出错: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 用户登录
func LoginCheck(ctx *gin.Context) {
	loginForm := model.User{}
	if err := ctx.ShouldBind(&loginForm); err != nil {
		zap.S().Errorf("登录失败: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	user, err := global.UserSrvClient.LoginCheck(context.Background(), &douyin_core.UserLoginRequest{
		Username: loginForm.UserName,
		Password: loginForm.Password,
	})
	if err != nil {
		zap.S().Errorf("登录失败: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":     "登录成功",
		"user_id": user.UserId,
	})
}
