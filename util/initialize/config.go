/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 15:12
* @description:
********************************************************************************/

package initialize

import (
	"bytedanceCamp/dao/global"
	"fmt"
	"github.com/spf13/viper"
)

func initConfig() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath("/Users/yakult/Documents/code/GoLang/bytedanceCamp/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&global.ProjectConfig)
	if err != nil {
		panic(fmt.Errorf("unmarshal config file error: %s", err))
	}
}
