/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 21:37
* @description:
********************************************************************************/

package dao

import (
	"bytedanceCamp/config"
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

func GetDB() (*gorm.DB, error) {
	username := config.GetConfig().Mysql.UserName // 数据库用户名
	host := config.GetConfig().Mysql.Host         // 数据库地址
	password := config.GetConfig().Mysql.Password // 数据库密码
	port := config.GetConfig().Mysql.Port         // 数据库端口号
	DbName := config.GetConfig().Mysql.DbName     // 数据库名

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
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.S().Fatalf("connect database error: %s", err)
		return nil, err
	}
	sqlDB, _ := _db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(50) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	return _db, err
}
