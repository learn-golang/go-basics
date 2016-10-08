package users


type User struct {
	name string
	surname string
	id int64
}

type Users struct {
	users []*User
	index int64
}
