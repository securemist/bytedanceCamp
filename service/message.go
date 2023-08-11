/********************************************************************************
* @author: Yakult
* @date: 2023/8/10 17:59
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

type MessageService struct {
	douyin_extra_second.UnimplementedMessageServer
}

// 发送信息
func (m *MessageService) MessageSend(ctx context.Context, req *douyin_extra_second.MessageSendRequest) (
	*douyin_extra_second.MessageSendResponse, error) {
	message := model.Message{
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	}
	result := global.MysqlDB.Create(&message)
	if result.Error != nil {
		zap.S().Errorf("消息发送失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "消息发送失败: %s", result.Error)
	}
	return &douyin_extra_second.MessageSendResponse{
		StatusCode: 0,
	}, nil
}

// 查询所有聊天记录
func (m *MessageService) MessageChat(ctx context.Context, req *douyin_extra_second.MessageChatRequest) (
	*douyin_extra_second.MessageChatResponse, error) {
	var messages []model.Message
	result := global.MysqlDB.Where("from_user_id = ? and to_user_id = ?", req.UserId, req.ToUserId).Find(&messages)
	if result.Error != nil {
		zap.S().Errorf("消息记录查询失败: %s", result.Error)
		return nil, status.Errorf(codes.Internal, "消息记录查询失败: %s", result.Error)
	}
	var messageList []*douyin_extra_second.MessageInfo
	for _, v := range messages {
		time := v.CreatedAt.Format("2006-01-02 15:04:05")
		messageList = append(messageList, &douyin_extra_second.MessageInfo{
			ToUserId:   v.ToUserId,
			FromUserId: v.FromUserId,
			Content:    v.Content,
			CreateTime: &time,
		})
	}
	return &douyin_extra_second.MessageChatResponse{
		StatusCode:  0,
		MessageList: messageList,
	}, nil
}
