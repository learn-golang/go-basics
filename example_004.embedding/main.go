package main

import (
	"./users"
	"fmt"
)

func main() {

	name, surname := "Denis", "Makogon"

	user_list := new(users.Users)
	user := new(users.User)
	admin := new(users.AdminUser)

	user = user.Create(name, surname)

	fmt.Println("------before------\n")
	fmt.Println("Name: ", user.GetName())
	fmt.Println("Surname: ",user.GetSurname())
	fmt.Println("ID: ",user.GetID())

	user_list.Append(user)

	fmt.Println("\n------after------")
	fmt.Println("Name: ", user.GetName())
	fmt.Println("Surname: ", user.GetSurname())
	fmt.Println("ID: ", user.GetID())

	admin = admin.Create(name, surname)
	fmt.Println("Admin name: ", admin.GetName())
	fmt.Println("Admin surname: ", admin.GetSurname())
	fmt.Println("Admin ID: ", admin.GetID())
	fmt.Println("Is admin?: ", admin.IsAdmin())
}
