package main

import (
	"fmt"
	"todo-module/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testpassword"
	// fmt.Println(u)
	// u.CreateUser()

	u, _ := models.GetUser(1)
	// fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}
