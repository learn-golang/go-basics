package main

import (
	"./users"
	"fmt"
)

func print_state(user users.GenericUserInfoInterface) {
	fmt.Println("Name: ", user.GetName())
	fmt.Println("Surname: ",user.GetSurname())
	fmt.Println("ID: ",user.GetID())
	fmt.Println("Is admin:", user.IsAdmin())
}


func main() {

	name, surname := "Denis", "Makogon"

	user_list := new(users.Users)
	user := new(users.User)
	admin := new(users.AdminUser)

	user = user.Create(name, surname)

	fmt.Println("------before------\n")
	print_state(user)

	user_list.Append(user)

	fmt.Println("\n------after------")
	print_state(user)

	admin = admin.Create(name, surname)
	fmt.Println("------before------\n")
	print_state(admin)

	user_list.Append(admin)
	fmt.Println("\n------after------")
	print_state(admin)

	fmt.Println("------embedded anonymous interface------")

	generic := users.GenericUser{admin}
	fmt.Println(generic.GetUserInfo())
}
