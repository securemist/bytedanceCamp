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
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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
func (u *UserServer) CreateUser(ctx context.Context, req *douyin_core.UserRegisterRequest) (*douyin_core.UserRegisterResponse, error) {
	user := model.User{}
	// 1. 查询用户是否存在
	result := dao.GetDB().Where(model.User{UserName: req.Username}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.UserName = req.Username
	user.Password = req.Password
	// 2. 密码加密
	user.Password = util.EncodePassword(req.Password)
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
	response := &douyin_core.UserRegisterResponse{
		StatusCode: 0,
		UserId:     user.Uuid,
		Token:      user.JwtToken,
	}
	return response, nil
}

// LoginCheck 用户登录
func (u *UserServer) LoginCheck(ctx context.Context, req *douyin_core.UserLoginRequest) (*douyin_core.UserLoginResponse, error) {
	user := model.User{}
	result := dao.GetDB().Where(model.User{UserName: req.Username}).First(&user)
	if result.RowsAffected == 0 {
		zap.S().Errorf("%s用户不存在", req.Username)
		return nil, status.Errorf(codes.Internal, "%s用户不存在", req.Username)
	}
	// 验证密码是否正确
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(user.Password, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	if check == false {
		zap.S().Errorf("密码错误,请重试")
		return nil, status.Errorf(codes.InvalidArgument, "密码错误,请重试")
	}
	response := douyin_core.UserLoginResponse{
		StatusCode: 0,
		UserId:     user.Uuid,
		Token:      user.JwtToken,
	}
	return &response, nil
}

// GetUserInfo 用户信息
func (u *UserServer) GetUserInfo(ctx context.Context, req *douyin_core.UserInfoRequest) (*douyin_core.UserInfoResponse, error) {
	user := model.User{}
	result := dao.GetDB().Where(model.User{Uuid: req.UserId}).First(&user)
	if result.RowsAffected == 0 {
		zap.S().Errorf("%s用户不存在", req.UserId)
		return nil, status.Errorf(codes.Internal, "%d用户不存在", req.UserId)
	}
	response := douyin_core.UserInfoResponse{
		StatusCode: 0,
		UserInfo: &douyin_core.UserInfo{
			UserId:        user.Uuid,
			Name:          user.UserName,
			FollowCount:   user.Followings,
			FollowerCount: user.Followers,
			Avatar:        user.Avatar,
			Signature:     user.Signature,
			TotalFavorite: user.TotalFavorite,
			WorkCount:     user.WorkCount,
			FavoriteCount: user.FavoriteCount,
		},
	}
	return &response, nil
}
