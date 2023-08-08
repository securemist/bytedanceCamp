/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 14:44
* @description:
********************************************************************************/

package model

import "gorm.io/gorm"

type User struct {
	Uuid          int64  `json:"uuid,omitempty" gorm:"type:bigint;not null;unique;comment:用户id"`
	UserName      string `json:"username" binding:"required,max=32" gorm:"type:varchar(32);column:username;not null;unique;comment:用户名"`
	Password      string `json:"password" binding:"required,max=32" gorm:"type:varchar(100);not null;comment:登录密码"`
	JwtToken      string `json:"jwt_token" gorm:"type:varchar(250);not null;comment:用户鉴权token"`
	Avatar        string `json:"avatar" gorm:"type:varchar(200);comment:用户头像"`
	Followers     int64  `json:"followers" gorm:"type:int;default:0;comment:粉丝数"`
	Followings    int64  `json:"followings" gorm:"type:int;default:0;comment:关注数"`
	Signature     string `json:"signature" gorm:"type:varchar(200);comment:个人简介"`
	TotalFavorite int64  `json:"total_favorite" gorm:"type:int;default:0;comment:获赞数量"`
	WorkCount     int64  `json:"work_count" gorm:"type:int;default:0;comment:作品数量"`
	FavoriteCount int64  `json:"favorite_count" gorm:"type:int;default:0;comment:喜欢数量"`
	gorm.Model
}

type Video struct {
	Uuid          int64  `json:"uuid" gorm:"type:bigint;not null;unique;comment:视频id"`
	AuthorId      int64  `json:"author_id" gorm:"type:bigint;not null;comment:视频作者id"`
	FavoriteCount int64  `json:"favorite_count" gorm:"type:int;default:0;comment:视频的点赞总数"`
	CommentCount  int64  `json:"comment_count" gorm:"type:int;default:0;comment:视频的评论总数"`
	Title         string `json:"title" gorm:"type:varchar(200);not null;comment:视频标题"`
	PlayUrl       string `json:"play_url" gorm:"type:varchar(200);not null;comment:视频播放地址"`
	CoverUrl      string `json:"cover_url" gorm:"type:varchar(200);not null;comment:视频封面地址"`
	gorm.Model
}
