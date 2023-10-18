package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var CONF *Config = nil

// Config 配置
type Config struct {
	HttpPort int32    `toml:"http_port"`
	CertPath string   `toml:"cert_path"`
	KeyPath  string   `toml:"key_path"`
	Logger   Logger   `toml:"logger"`
	Redirect Redirect `toml:"redirect"`
}

// Logger 日志
type Logger struct {
	Level   string `toml:"level"`
	Mode    string `toml:"mode"`
	Track   bool   `toml:"track"`
	MaxSize int32  `toml:"max_size"`
}

// Redirect 重定向
type Redirect struct {
	DispatchUrl string `toml:"dispatch_url"`
	SdkUrl      string `toml:"sdk_url"`
}

func InitConfig(filePath string) {
	CONF = new(Config)
	CONF.loadConfigFile(filePath)
}

func GetConfig() *Config {
	return CONF
}

// 加载配置文件
func (c *Config) loadConfigFile(filePath string) {
	_, err := toml.DecodeFile(filePath, &c)
	if err != nil {
		info := fmt.Sprintf("config file load error: %v\n", err)
		panic(info)
	}
}
