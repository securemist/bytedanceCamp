/********************************************************************************
* @author: Yakult
* @date: 2023/8/9 21:04
* @description:
********************************************************************************/

package model

type VideoData struct {
	Data  string `json:"data"`
	Title string `json:"title"`
}

type MessageData struct {
	ToUserId int64  `json:"to_user_id,omitempty" form:"to_user_id"`
	Content  string `json:"content,omitempty" form:"content"`
}
