/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 14:21
* @description:
********************************************************************************/

package service

import (
	"bytedanceCamp/model/proto/douyin_extra_first"
	"bytedanceCamp/util/initialize"
	"context"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	initialize.Init()
	req := &douyin_extra_first.FavoriteActionRequest{
		VideoId:    1854278998167552,
		UserId:     1533847262990336,
		ActionType: 2,
	}
	f := FavoriteServer{}
	res, err := f.FavoriteAction(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestFavoriteList(t *testing.T) {
	initialize.Init()
	req := &douyin_extra_first.FavoriteListRequest{UserId: 1533847262990336}
	f := FavoriteServer{}
	res, err := f.FavoriteList(context.Background(), req)
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
