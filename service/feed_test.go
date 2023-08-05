/********************************************************************************
* @author: Yakult
* @date: 2023/8/5 16:18
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/model/proto/douyin_core"
	"context"
	"testing"
)

func TestPublishVideo(t *testing.T) {
	req := &douyin_core.PublishVideoRequest{
		Data:   []byte("hello Gopher!"),
		Title:  "投稿测试2",
		UserId: 445243736461312,
	}
	f := FeedServer{}
	res, err := f.PublishVideo(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestGetFeed(t *testing.T) {
	req := &douyin_core.FeedRequest{
		LatestTime: nil,
	}
	f := FeedServer{}
	res, err := f.GetFeed(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestPublishList(t *testing.T) {
	req := &douyin_core.PublishListRequest{
		UserId: 445243736461312,
	}
	f := FeedServer{}
	res, err := f.PublishList(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
