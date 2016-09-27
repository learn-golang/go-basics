package main

import (
	"./users"
	"fmt"
)

func main() {

	name, surname := "Denis", "Makogon"

	user_list := new(users.Users)
	user := new(users.User)

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
}
