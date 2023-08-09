/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 18:26
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

type CommentService struct {
	douyin_extra_first.UnimplementedCommentServer
}

func (c *CommentService) CommentAction(ctx context.Context, req *douyin_extra_first.CommentActionRequest) (
	*douyin_extra_first.CommentActionResponse, error) {
	// 发布评论
	if req.ActionType == 1 {
		comment := model.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: *req.CommentText,
		}
		result := global.MysqlDB.Create(&comment)
		if result.RowsAffected == 0 {
			zap.S().Errorf("评论视频失败: %s", result.Error)
			return nil, status.Errorf(codes.Internal, "评论视频失败: %s", result.Error)
		}
		return &douyin_extra_first.CommentActionResponse{
			StatusCode: 0,
			Comment: &douyin_extra_first.CommentInfo{
				Id:         int64(comment.ID),
				UserId:     comment.UserId,
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		}, nil
	} else {
		// 删除评论
		result := global.MysqlDB.Delete(&model.Comment{}, req.CommentId)
		if result.RowsAffected == 0 {
			zap.S().Errorf("删除评论失败: %s", result.Error)
			return nil, status.Errorf(codes.Internal, "删除评论失败: %s", result.Error)
		}
		return &douyin_extra_first.CommentActionResponse{
			StatusCode: 0,
		}, nil
	}
}

func (c *CommentService) CommentList(ctx context.Context, req *douyin_extra_first.CommentListRequest) (
	*douyin_extra_first.CommentListResponse, error) {
	var comment []model.Comment
	result := global.MysqlDB.Where(&model.Comment{VideoId: req.VideoId}).Find(&comment)
	if result.Error != nil {
		zap.S().Errorf("获取视频评论列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "获取视频评论列表失败: %s", result.Error)
	}
	var commentList []*douyin_extra_first.CommentInfo
	for _, c := range comment {
		commentList = append(commentList, &douyin_extra_first.CommentInfo{
			Id:         int64(c.ID),
			UserId:     c.UserId,
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &douyin_extra_first.CommentListResponse{
		StatusCode:  0,
		CommentList: commentList,
	}, nil
}
