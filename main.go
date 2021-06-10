package main

import (
	"fmt"
	"todo-module/app/models"
)

func main() {
	fmt.Println(models.Db)

	t, _ := models.GetTodo(1)
	t.Content = "updated todo"
	t.UpdateTodo()
}
