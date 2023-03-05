package main

import (
	"fmt"

	"init_parse_config/config"
)

func main() {
	fmt.Println(config.Config.GetString("endpoint"))
}
