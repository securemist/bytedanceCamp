/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 22:59
* @description: 该文件用来创建具体的表
********************************************************************************/
package main

import (
	"bytedanceCamp/dao"
	"bytedanceCamp/model"
)

func main() {
	db := dao.GetDB()
	db.AutoMigrate(&model.Video{})
}
