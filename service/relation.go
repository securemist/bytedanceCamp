/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 10:15
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/model/proto/douyin_extra_second"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RelationService struct {
	douyin_extra_second.UnimplementedRelationServer
}

func (r *RelationService) RelationAction(ctx context.Context, req *douyin_extra_second.RelationActionRequest) (
	*douyin_extra_second.RelationActionResponse, error) {
	relation := model.Relation{}
	result := global.MysqlDB.Where(&model.Relation{UserId: req.UserId, ToUserId: req.ToUserId}).First(&relation)
	if result.RowsAffected == 0 {
		relation.UserId = req.UserId
		relation.ToUserId = req.ToUserId
		if req.ActionType == 1 {
			relation.IsRelation = true
		} else {
			relation.IsRelation = false
		}
		result = global.MysqlDB.Create(&relation)
		if result.RowsAffected == 0 {
			if req.ActionType == 1 {
				zap.S().Errorf("关注用户失败: %s", result.Error)
				return nil, status.Errorf(codes.Internal, "关注用户失败: %s", result.Error)
			} else {
				zap.S().Errorf("取消关注失败: %s", result.Error)
				return nil, status.Errorf(codes.Internal, "取消关注失败: %s", result.Error)
			}
		}
	} else {
		var tmp bool
		if req.ActionType == 1 {
			tmp = true
		} else {
			tmp = false
		}
		result = global.MysqlDB.Model(&model.Relation{}).Where("user_id = ? and to_user_id = ?", req.UserId, req.ToUserId).
			Update("is_relation", tmp)
		if result.RowsAffected == 0 {
			if req.ActionType == 1 {
				zap.S().Errorf("关注用户失败: %s", result.Error)
				return nil, status.Errorf(codes.Internal, "关注用户失败: %s", result.Error)
			} else {
				zap.S().Errorf("取消关注失败: %s", result.Error)
				return nil, status.Errorf(codes.Internal, "取消关注失败: %s", result.Error)
			}
		}
	}
	return &douyin_extra_second.RelationActionResponse{StatusCode: 0}, nil
}

func (r *RelationService) RelationFollowList(ctx context.Context, req *douyin_extra_second.RelationFollowListRequest) (
	*douyin_extra_second.RelationFollowListResponse, error) {
	var relations []model.Relation
	result := global.MysqlDB.Where("user_id = ?", req.UserId).Find(&relations)
	if result.Error != nil {
		zap.S().Errorf("查询关注列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "查询关注列表失败: %s", result.Error)
	}
	var followList []int64
	for _, relation := range relations {
		if relation.IsRelation {
			followList = append(followList, relation.ToUserId)
		}
	}
	return &douyin_extra_second.RelationFollowListResponse{FollowList: followList}, nil
}

func (r *RelationService) RelationFansList(ctx context.Context, req *douyin_extra_second.RelationFansListRequest) (
	*douyin_extra_second.RelationFansListResponse, error) {
	var relations []model.Relation
	result := global.MysqlDB.Where("to_user_id = ?", req.UserId).Find(&relations)
	if result.Error != nil {
		zap.S().Errorf("查询粉丝列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "查询粉丝列表失败: %s", result.Error)
	}
	var fansList []int64
	for _, relation := range relations {
		if relation.IsRelation {
			fansList = append(fansList, relation.UserId)
		}
	}
	return &douyin_extra_second.RelationFansListResponse{FanList: fansList}, nil
}
func (r *RelationService) RelationFriendList(ctx context.Context, req *douyin_extra_second.RelationFriendListRequest) (
	*douyin_extra_second.RelationFriendListResponse, error) {
	var friends []model.Friend
	result := global.MysqlDB.Where("user_id = ?", req.UserId).Find(&friends)
	if result.Error != nil {
		zap.S().Errorf("查询好友列表失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "查询好友列表失败: %s", result.Error)
	}
	var friendList []int64
	for _, friend := range friends {
		friendList = append(friendList, friend.FriendId)
	}
	return &douyin_extra_second.RelationFriendListResponse{FriendList: friendList}, nil
}

func (r *RelationService) RelationAddFriend(ctx context.Context, req *douyin_extra_second.RelationAddFriendRequest) (
	*douyin_extra_second.RelationAddFriendResponse, error) {
	friend := model.Friend{}
	result := global.MysqlDB.Where(&model.Friend{UserId: req.UserId, FriendId: req.ToUserId}).First(&friend)
	if result.RowsAffected == 0 {
		friend.UserId = req.UserId
		friend.FriendId = req.ToUserId
		result = global.MysqlDB.Create(&friend)
		if result.RowsAffected == 0 {
			zap.S().Errorf("添加好友失败: %s", result.Error)
			return nil, status.Errorf(codes.Internal, "添加好友失败: %s", result.Error)
		}
	}
	return &douyin_extra_second.RelationAddFriendResponse{StatusCode: 0}, nil
}
