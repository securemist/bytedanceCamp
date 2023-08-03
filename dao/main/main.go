/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 22:59
* @description: 该文件用来创建具体的表
********************************************************************************/
package main

import (
	"bytedanceCamp/dao"
	"bytedanceCamp/model"
	"go.uber.org/zap"
)

func main() {
	db, err := dao.GetDB()
	if err != nil {
		zap.S().Fatalf("connect database error: %s", err)
	}
	db.AutoMigrate(&model.User{})
}
