/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 22:59
* @description: 该文件用来创建具体的表
********************************************************************************/
package main

import (
	"bytedanceCamp/dao/global"
	"bytedanceCamp/model"
	"bytedanceCamp/util/initialize"
)

func main() {
	initialize.Init()
	global.MysqlDB.AutoMigrate(&model.Message{})
}
