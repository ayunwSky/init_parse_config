package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 是全局的配置变量
var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.SetConfigFile("./config.yaml")

	if err := Config.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		return
	}
}
