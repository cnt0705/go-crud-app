package main

import (
	"fmt"
	"todo-module/app/controllers"
	"todo-module/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
