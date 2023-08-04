/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 13:35
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/dao"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/util"
	"bytedanceCamp/web/middlewares"
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserServer struct {
	douyin_core.UnimplementedUserServer
}

// CreateUser 注册用户
/*
	1. 先查询用户是否存在，不存在才创建
	2. 将用户的密码加密后，以密文的形式存储
	3. 生成一个64位的uuid
	4. 生成一个jwt token用以鉴权
	5. 保存到数据库中
*/
func (u *UserServer) CreateUser(ctx context.Context, req *douyin_core.UserRegisterRequest) (*douyin_core.UserLoginResponse, error) {
	user := model.User{}
	// 1. 查询用户是否存在
	result := dao.GetDB().Where(model.User{UserName: req.Username}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.UserName = req.Username
	user.Password = req.Password
	// 2. 密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	// 3. 生成uuid
	user.Uuid = util.GenID()
	// 4. 生成jwt token
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		Uuid:     user.Uuid,
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Yakult",
			ExpiresAt: &jwt.NumericDate{Time: time.Now()},
			NotBefore: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24 * 30)},
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.S().Errorf("创建用户时，生成jwt token失败: err:%s", err)
	}
	user.JwtToken = token
	// 5. 保存到数据库中
	result = dao.GetDB().Create(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	response := &douyin_core.UserLoginResponse{
		StatusCode: 0,
		UserId:     user.Uuid,
		Token:      user.JwtToken,
	}
	return response, nil
}
