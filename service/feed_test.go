/********************************************************************************
* @author: Yakult
* @date: 2023/8/5 16:18
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/model/proto/douyin_core"
	"bytedanceCamp/util/initialize"
	"context"
	"testing"
	"time"
)

func TestPublishVideo(t *testing.T) {
	initialize.Init()
	req := &douyin_core.PublishVideoRequest{
		Data:   []byte("hello Gopher!"),
		Title:  "投稿测试1",
		UserId: 1533847262990336,
	}
	f := FeedServer{}
	res, err := f.PublishVideo(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestGetFeed(t *testing.T) {
	initialize.Init()
	tie := time.Now().Unix()
	req := &douyin_core.FeedRequest{
		LatestTime: &tie,
	}
	f := FeedServer{}
	res, err := f.GetFeed(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestPublishList(t *testing.T) {
	initialize.Init()
	req := &douyin_core.PublishListRequest{
		UserId: 1533847262990336,
	}
	f := FeedServer{}
	res, err := f.PublishList(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
