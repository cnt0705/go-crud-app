package main

import (
	"fmt"
	"todo-module/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testpassword"
	fmt.Println(u)
	u.CreateUser()

	user, _ := models.GetUser(2)

	user.CreateTodo("First Todo")
}
