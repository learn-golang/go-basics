package main

import (
	"./users"
	"fmt"
	"strings"
)


func main() {

	name, surname := "Denis", "Makogon"

	user := new(users.User)

	fmt.Println("Attempt #1")
	user, err := user.Create(name, surname)
	fmt.Println("User:", user)
	fmt.Println("Error:", err)
	fmt.Println("\n\n")

	fmt.Println("Attempt #2")
	user, err = user.Create(strings.ToUpper(name), surname)
	fmt.Println("User:", user)
	fmt.Println("Error", err)
	fmt.Println("\n\n")

	fmt.Println("Attempt #3")
	user, err = user.Create(strings.ToUpper(name), strings.ToUpper(surname))
	fmt.Println("User:", user)
	fmt.Println("Error", err)
	fmt.Println("\n\n")

}
