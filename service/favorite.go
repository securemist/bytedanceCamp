/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 10:57
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_extra_first"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FavoriteServer struct {
	douyin_extra_first.UnimplementedFavoriteServer
}

// 视频点赞
func (f *FavoriteServer) FavoriteAction(ctx context.Context, req *douyin_extra_first.FavoriteActionRequest) (
	*douyin_extra_first.FavoriteActionResponse, error) {
	favorite := model.Favorite{
		VideoId: req.VideoId,
		UserId:  req.UserId,
	}
	if req.ActionType == 1 {
		favorite.IsFavorite = true
	} else {
		favorite.IsFavorite = false
	}
	// 先查询是否已经保存过这条记录
	var fav model.Favorite
	result := global.MysqlDB.Where(&model.Favorite{
		VideoId: req.VideoId,
		UserId:  req.UserId,
	}).Find(&fav)
	// 已经保存过，则直接更新
	if result.RowsAffected == 1 {
		global.MysqlDB.Model(&model.Favorite{}).Where("video_id = ? and user_id = ?", req.VideoId, req.UserId).
			Update("is_favorite", favorite.IsFavorite)
	} else {
		result = global.MysqlDB.Create(&favorite)
		if result.RowsAffected == 0 {
			zap.S().Errorf("点赞视频失败: %s", result.Error)
			return nil, status.Errorf(codes.Internal, "点赞视频失败: %s", result.Error)
		}
	}
	response := douyin_extra_first.FavoriteActionResponse{StatusCode: 0}
	return &response, nil
}

// 点赞视频列表
func (f *FavoriteServer) FavoriteList(ctx context.Context, req *douyin_extra_first.FavoriteListRequest) (
	*douyin_extra_first.FavoriteListResponse, error) {
	var favorites []model.Favorite
	result := global.MysqlDB.Where(model.Favorite{UserId: req.UserId}).Find(&favorites)
	if result.Error != nil {
		zap.S().Errorf("获取点赞列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "获取点赞列表失败: %s", result.Error)
	}
	var videoList []int64
	for _, v := range favorites {
		videoList = append(videoList, v.VideoId)
	}
	response := douyin_extra_first.FavoriteListResponse{
		VideoList:  videoList,
		StatusCode: 0,
	}
	return &response, nil
}
