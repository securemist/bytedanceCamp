/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 15:35
* @description:
********************************************************************************/

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Log   LogConfig   `mapstructure:"log"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

// MysqlConfig mysql相关
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

// LogConfig log相关
type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

var c Config

func init() {
	// 设置文件名
	viper.SetConfigName("mysql")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("unmarshal config file error: %s", err))
	}
}

func GetConfig() Config {
	return c
}
