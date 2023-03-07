package config

import (
	"fmt"
	// "log"

	"github.com/spf13/viper"
)

// 定义 Redis 配置文件结构体
type RedisCfg struct {
	Port     int    `yaml:"port" json:"port"`
	Host     string `yaml:"host" json:"host"`
	Database string `yaml:"database" json:"database"`
	Password string `yaml:"password" json:"password"`
}

// 定义 MySQL 配置文件结构体
type MysqlCfg struct {
	Port     int    `yaml:"port" json:"port"`
	Host     string `yaml:"host" json:"host"`
	Database string `yaml:"database" json:"database"`
	Password string `yaml:"password" json:"password"`
}

// Config 结构体用于存储所有配置数据
type Config struct {
	Redis RedisCfg
	Mysql MysqlCfg
}

// 定义全局配置文件变量
var config Config

func InitConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/config.yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println(err.Error())
		// log.Fatalf("failed to read config file: %v", err)
		fmt.Printf("failed to read config file: %v", err)
		return nil, err
	}

	// 解析配置文件
	if err := viper.Unmarshal(&config); err != nil {
		// log.Fatalf("failed to unmarshal config file: %v", err)
		fmt.Printf("failed to unmarshall config: %v", err)
		return nil, err
	}

	fmt.Printf("all config contents: %v\n",config)

	// 将所有配置文件返回
	return &config, nil
}
