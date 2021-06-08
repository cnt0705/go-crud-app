package main

import (
	"fmt"
	"log"
	"todo-module/config"
)

func main() {
	fmt.Println(config.Config.Port)

	log.Println("test")
}
