package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server   Server
	Mysql    Mysql
	Redis    Redis
	RabbitMQ RabbitMQ
	//Logger Logger
}

type Mysql struct {
	Ip           string `toml:"ip"`
	Port         string `toml:"port"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Database     string `toml:"database"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

type RabbitMQ struct {
	Url   string `toml: "url"`
	Queue string `toml: "queue"`
}

type Captcha struct {
	KeyLong   int `yaml:"key-long"`
	ImgWidth  int `yaml:"img-width"`
	ImgHeight int `yaml:"img-height"`
}

type Redis struct {
	Dial      string `toml:"dial"`
	Ip        string `toml:"ip"`
	Port      string `toml:"port"`
	Password  string `toml:"password"`
	PoolSize  int    `toml:"pool_size"`
	Database  int    `toml:"database"`
	MaxIdle   int    `toml:"max-idle"`
	MaxActive int    `toml:"max-active"`
}

//type Mongodb struct {
//	Ip       string
//	Port     string
//	Database string
//}
//
//func (m *Mongodb) Link() string {
//	return fmt.Sprintf("mongodb://%s:%d", m.Ip, m.Port)
//}

type Server struct {
	Port   string `toml:"port"`   // 地址
	Status string `toml:"status"` // 状态
}

//type Logger struct {
//	FilePath   string `toml:"file_path"`
//	FileName   string `toml:"file_name"` // 日志文件路径
//	MaxSize    int    `toml:"max_size"`
//	MaxAge     int    `toml:"max_age"`
//	MaxBackups int    `toml:"max_backups"`
//}

// 默认的配置文件路由
const ConfigPath = "./config/config.toml"

func InitConfig(args ...string) Config {
	configPath := ConfigPath
	if len(args) == 1 && args[0] != "" {
		configPath = args[0]
	}
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Println("config.toml配置文件读取失败:", err)
	}
	return config
}
