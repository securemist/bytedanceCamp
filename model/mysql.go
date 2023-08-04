/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 14:44
* @description:
********************************************************************************/

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid     int64  `json:"-" gorm:"type:bigint;not null;unique;comment 用户id"`
	UserName string `json:"username" binding:"required,max=32" gorm:"type:varchar(32);column:username;not null;unique;comment 用户名"`
	Password string `json:"password" binding:"required,max=32" gorm:"type:varchar(100);not null;comment 登录密码"`
	JwtToken string `json:"-" gorm:"type:varchar(250);not null;comment 用户鉴权token"`
	Avatar   string `json:"-" gorm:"type:varchar(200);comment 用户头像"`
}
