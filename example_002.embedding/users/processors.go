package users

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

// Admin USER API //////////////////////////////////////////

func (admin *AdminUser) Create(name, surname string) *AdminUser {
	return &AdminUser{User:User{name:name, surname:surname}, admin:true}
}

func (admin *AdminUser) IsAdmin() bool {
	return admin.admin
}

//////////////////////////////////////////////////////

// USERS API /////////////////////////////////////////

func (users *Users) Append(user *User) {
	users.index += 1
	user.id = users.index
	users.users = append(users.users, user)
}


func (users *Users) GetIndex() int64 {
	return users.index
}

func (users *Users) GetUsers() []*User {
	return users.users
}
