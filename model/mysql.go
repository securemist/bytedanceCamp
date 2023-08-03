/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 14:44
* @description:
********************************************************************************/

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid     int64  `gorm:"type:bigint;not null;index:idx_id;unique;comment 用户id"`
	NickName string `gorm:"type:varchar(32);not null;comment 用户昵称"`
	Password string `gorm:"type:varchar(100);not null;comment 登录密码"`
	jwtToken string `gorm:"type:varchar(100);not null;comment 用户鉴权token"`
	Avatar   string `gorm:"type:varchar(200);comment 用户头像"`
}
