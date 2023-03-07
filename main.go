package main

import (
	"fmt"
	"init_parse_config/config"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to get init config file: %v", err)
		return
	}

	fmt.Printf("redis database: %v", config.Redis.Database)
	fmt.Printf("redis port: %v", config.Redis.Port)
	fmt.Printf("mysql database: %v", config.Mysql.Database)
	fmt.Printf("mysql port: %v", config.Mysql.Port)
}
