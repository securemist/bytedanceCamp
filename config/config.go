/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 15:35
* @description:
********************************************************************************/

package config

type Config struct {
	Log           LogConfig           `mapstructure:"log"`
	Mysql         MysqlConfig         `mapstructure:"mysql"`
	Jwt           JwtConfig           `mapstructure:"jwt"`
	Consul        ConsulConfig        `mapstructure:"consul"`
	ConsulService ConsulServiceConfig `mapstructure:"consul-service"`
	ConsulWeb     ConsulWebConfig     `mapstructure:"consul-web"`
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

// JwtConfig jwt token相关
type JwtConfig struct {
	Secret string `mapstructure:"secret"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type ConsulServiceConfig struct {
	User     ConsulServiceCommon `mapstructure:"user"`
	Feed     ConsulServiceCommon `mapstructure:"feed"`
	Favorite ConsulServiceCommon `mapstructure:"favorite"`
	Comment  ConsulServiceCommon `mapstructure:"comment"`
}
type ConsulWebConfig struct {
	User     ConsulWebCommon `mapstructure:"user"`
	Feed     ConsulWebCommon `mapstructure:"feed"`
	Favorite ConsulWebCommon `mapstructure:"favorite"`
	Comment  ConsulWebCommon `mapstructure:"comment"`
}
type ConsulServiceCommon struct {
	Name string   `mapstructure:"name"`
	Tags []string `mapstructure:"tags"`
}

type ConsulWebCommon struct {
	Name string   `mapstructure:"name"`
	Tags []string `mapstructure:"tags"`
	Port int      `mapstructure:"port"`
}
