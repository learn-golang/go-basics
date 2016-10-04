package main

import (
	"./users"
	"fmt"
	"reflect"
)


func process(generic users.GenericUserInfoInterface, match_type reflect.Type) {
	type_processor := users.TypeProcessor{Generic: generic}
	datarefs := type_processor.GetFieldsByType(match_type)
	fmt.Println("Fields that are matching to search criteria", datarefs)
	all_refs := type_processor.GetAllFieldsDef()
	fmt.Println("All data fileds of a type", all_refs)
}


func main() {

	name, surname := "Denis", "Makogon"

	user_list := new(users.Users)
	user := new(users.User)
	admin := new(users.AdminUser)

	user = user.Create(name, surname)
	user_list.Append(user)

	admin = admin.Create(name, surname)
	user_list.Append(admin)

	process(user, reflect.TypeOf(""))

	process(admin, reflect.TypeOf(false))
}
