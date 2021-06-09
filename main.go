package main

import (
	"fmt"
	"todo-module/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.Password = "testpassword"
	fmt.Println(u)
	u.CreateUser()

	// user, _ := models.GetUser(2)

	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(3)
	user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

}
