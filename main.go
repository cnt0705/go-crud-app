package main

import (
	"fmt"
	"todo-module/config"
)

func main() {
	fmt.Println(config.Config.Port)
}
