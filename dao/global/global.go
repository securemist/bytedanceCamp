/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:02
* @description:
********************************************************************************/

package global

import (
	"bytedanceCamp/config"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/model/proto/douyin_extra_first"
	"gorm.io/gorm"
)

var (
	MysqlDB           *gorm.DB               // Mysql
	ProjectConfig     config.Config          // 配置文件
	UserSrvClient     douyin_core.UserClient // service端微服务
	FeedSrvClient     douyin_core.FeedClient
	FavoriteSrvClient douyin_extra_first.FavoriteClient
	CommentSrvClient  douyin_extra_first.CommentClient
)
