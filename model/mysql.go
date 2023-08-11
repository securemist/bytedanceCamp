/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 14:44
* @description:
********************************************************************************/

package model

import "gorm.io/gorm"

// 用户信息
type User struct {
	Uuid          int64  `json:"uuid" gorm:"type:bigint;not null;unique;comment:用户id"`
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

// 视频信息
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

// 用户对视频的评论（可以有多条评论）
type Comment struct {
	// 根据user_id和video_id建立联合索引
	UserId  int64  `json:"user_id" gorm:"type:bigint;not null;index:idx_user_video;comment:评论者id"`
	VideoId int64  `json:"video_id" gorm:"type:bigint;not null;index:idx_user_video;comment:视频id"`
	Content string `json:"content" gorm:"type:varchar(200);not null;comment:评论内容"`
	User    User   `gorm:"foreignKey:UserId;references:Uuid"`
	Video   Video  `gorm:"foreignKey:VideoId;references:Uuid"`
	gorm.Model
}

// 用户对视频点赞
type Favorite struct {
	// 根据user_id和video_id建立联合索引
	UserId     int64 `json:"user_id" gorm:"type:bigint;not null;index:idx_user_video;comment:评论者id"`
	VideoId    int64 `json:"video_id" gorm:"type:bigint;not null;index:idx_user_video;comment:视频id"`
	IsFavorite bool  `json:"is_favorite" gorm:"type:bool;not null;comment:用户是否点赞视频"`
	User       User  `gorm:"foreignKey:UserId;references:Uuid"`
	Video      Video `gorm:"foreignKey:VideoId;references:Uuid"`
	gorm.Model
}

// 用户关注
type Relation struct {
	UserId     int64 `json:"user_id" gorm:"type:bigint;not null;index:idx_user_to_user;comment:关注者id"`
	ToUserId   int64 `json:"to_user_id" gorm:"type:bigint;not null;index:idx_user_to_user;comment:被关注者id"`
	IsRelation bool  `gorm:"type:bool;not null;comment:用户是否关注"`
	User       User  `gorm:"foreignKey:UserId;references:Uuid"`
	ToUser     User  `gorm:"foreignKey:ToUserId;references:Uuid"`
	gorm.Model
}

// 好友关系
type Friend struct {
	UserId   int64 `json:"user_id" gorm:"type:bigint;not null;index:idx_user_friend;comment:用户id"`
	FriendId int64 `json:"friend_id" gorm:"type:bigint;not null;index:idx_user_friend;comment:好友id"`
	User     User  `gorm:"foreignKey:UserId;references:Uuid"`
	Friend   User  `gorm:"foreignKey:FriendId;references:Uuid"`
	gorm.Model
}

// 消息
type Message struct {
	FromUserId int64  `json:"user_id" gorm:"type:bigint;not null;comment:发送用户id"`
	ToUserId   int64  `json:"to_user_id" gorm:"type:bigint;not null;comment:接收用户id"`
	Content    string `json:"content" gorm:"type:varchar(200);not null;comment:消息内容"`
	FromUser   User   `gorm:"foreignKey:FromUserId;references:Uuid"`
	ToUser     User   `gorm:"foreignKey:ToUserId;references:Uuid"`
	gorm.Model
}
