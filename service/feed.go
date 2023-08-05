/********************************************************************************
* @author: Yakult
* @date: 2023/8/5 15:16
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/dao"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/util"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type FeedServer struct {
	douyin_core.UnimplementedFeedServer
}

// GetFeed 视频流接口
func (f *FeedServer) GetFeed(ctx context.Context, req *douyin_core.FeedRequest) (*douyin_core.FeedResponse, error) {
	if req.LatestTime == nil {
		nowTime := time.Now().Unix()
		req.LatestTime = &nowTime
	}
	var videos []model.Video
	result := dao.GetDB().Order("created_at desc").Limit(30).Find(&videos)
	if result.Error != nil {
		zap.S().Errorf("获取视频流失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "获取视频流失败: %s", result.Error)
	}
	response := &douyin_core.FeedResponse{
		StatusCode: 0,
	}
	var videoList []*douyin_core.VideoInfo
	var nextTime int64
	for _, v := range videos {
		videoList = append(videoList, &douyin_core.VideoInfo{
			VideoId:       v.Uuid,
			AuthorId:      v.AuthorId,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		})
		if nextTime == 0 {
			nextTime = v.CreatedAt.Unix()
		} else {
			if nextTime > v.CreatedAt.Unix() {
				nextTime = v.CreatedAt.Unix()
			}
		}
	}
	response.VideoList = videoList
	response.NextTime = &nextTime
	return response, nil
}

// PublishVideo 投稿视频
func (f *FeedServer) PublishVideo(ctx context.Context, req *douyin_core.PublishVideoRequest) (*douyin_core.PublishVideoResponse, error) {
	//TODO 要叫视频数据存储到一个服务器上，得到一个url地址，这里随便给一个
	video := model.Video{
		Uuid:     util.GenID(),
		AuthorId: req.UserId,
		Title:    req.Title,
		PlayUrl:  "http://play_url_test.com",
		CoverUrl: "http://cover_url_test.com",
	}
	result := dao.GetDB().Create(&video)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "%s视频投稿失败: %s", video.Title, result.Error.Error())
	}
	response := douyin_core.PublishVideoResponse{
		StatusCode: 0,
	}
	return &response, nil
}

// PublishList 投稿列表
func (f *FeedServer) PublishList(ctx context.Context, req *douyin_core.PublishListRequest) (*douyin_core.PublishListResponse, error) {
	var videos []model.Video
	result := dao.GetDB().Where(model.Video{AuthorId: req.UserId}).Find(&videos)
	if result.Error != nil {
		zap.S().Errorf("获取投稿列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "获取投稿列表失败: %s", result.Error)
	}
	response := &douyin_core.PublishListResponse{
		StatusCode: 0,
	}
	var videoList []*douyin_core.VideoInfo
	var nextTime int64
	for _, v := range videos {
		videoList = append(videoList, &douyin_core.VideoInfo{
			VideoId:       v.Uuid,
			AuthorId:      v.AuthorId,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		})
		if nextTime == 0 {
			nextTime = v.CreatedAt.Unix()
		} else {
			if nextTime > v.CreatedAt.Unix() {
				nextTime = v.CreatedAt.Unix()
			}
		}
	}
	response.VideoList = videoList
	return response, nil
}
