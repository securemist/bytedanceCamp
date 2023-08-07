/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:10
* @description:
********************************************************************************/

package initialize

import "bytedanceCamp/dao/global"

func Init() {
	// 1. 初始化配置
	initConfig()
	// 2. 初始化日志
	initLogger(global.ProjectConfig.Log.Path, global.ProjectConfig.Log.Level)
	// 3. 初始化数据库
	initMysql()
	// 4. 初始化与service端的连接
	initSrvConn()
}
