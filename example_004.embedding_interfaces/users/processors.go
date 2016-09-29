package users

import (
	"errors"
)


// USER API //////////////////////////////////////////
func (user *User) Create(name, surname string) *User {
	return &User{name: name, surname:surname}
}


func (user *User) GetName() string {
	return user.name
}


func (user *User) GetSurname() string {
	return user.surname
}


func (user *User) GetID() int64 {
	return user.id
}

func (user *User) SetID(new_id int64) {
	user.id = new_id
}

func (user *User) IsAdmin() bool {
	return false
}


// Admin USER API //////////////////////////////////////////

func (admin *AdminUser) Create(name, surname string) *AdminUser {
	return &AdminUser{User:User{name:name, surname:surname}, admin:true}
}

func (admin *AdminUser) IsAdmin() bool {
	return admin.admin
}

//////////////////////////////////////////////////////

// USERS API /////////////////////////////////////////

func _add(users *Users, v UserInterface)  {
	v.SetID(users.index)
	users.users = append(users.users, v)
}

func (users *Users) Append(user UserInterface) error {
	users.index += 1
	switch v := user.(type) {
		case *User, *AdminUser:
			_add(users, v)
			return nil
		default:
			return errors.New("Unknown type")
	}
}

func (users *Users) GetIndex() int64 {
	return users.index
}

func (users *Users) GetUsers() []UserInterface {
	return users.users
}


// Anon interface API

func (generic *GenericUser) GetUserInfo() (string, string, int64, bool) {
	return generic.GetName(), generic.GetSurname(), generic.GetID(), generic.IsAdmin()
}
