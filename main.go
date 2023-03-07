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

	fmt.Printf("redis database: %v\n", config.Redis.Database)
	fmt.Printf("redis port: %v\n", config.Redis.Port)
	fmt.Printf("mysql database: %v\n", config.Mysql.Database)
	fmt.Printf("mysql port: %v\n", config.Mysql.Port)
}
