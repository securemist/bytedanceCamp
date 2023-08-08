/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 18:24
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/util/initialize"
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	req := &douyin_core.UserRegisterRequest{
		Username: "Yakult",
		Password: "123456",
	}
	u := UserServer{}
	_, err := u.CreateUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
}

func TestLoginCheck(t *testing.T) {
	req := &douyin_core.UserLoginRequest{
		Username: "Yakult",
		Password: "123456",
	}
	u := UserServer{}
	res, err := u.LoginCheck(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(res.UserId)
}

func TestGetUserInfo(t *testing.T) {
	initialize.Init()
	req := &douyin_core.UserInfoRequest{UserId: 445243736461312}
	u := UserServer{}
	res, err := u.GetUserInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(res)
}
