/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 18:24
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/model/proto/douyin_core"
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
