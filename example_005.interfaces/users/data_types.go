package users


type User struct {
	name string
	surname string
	id int64
}

type Users struct {
	users []UserInterface
	index int64
}


type AdminUser struct {
	User
	admin bool
}
