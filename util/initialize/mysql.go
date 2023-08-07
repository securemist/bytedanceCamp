/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:15
* @description:
********************************************************************************/

package initialize

import (
	"bytedanceCamp/dao/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func initMysql() {
	username := global.ProjectConfig.Mysql.UserName // 数据库用户名
	host := global.ProjectConfig.Mysql.Host         // 数据库地址
	password := global.ProjectConfig.Mysql.Password // 数据库密码
	port := global.ProjectConfig.Mysql.Port         // 数据库端口号
	DbName := global.ProjectConfig.Mysql.DbName     // 数据库名

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, DbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	var err error
	global.MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.S().Fatalf("connect database error: %s", err)
		return
	}
	sqlDB, _ := global.MysqlDB.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(50) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
}
