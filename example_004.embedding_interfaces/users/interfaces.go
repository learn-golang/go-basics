package users


type UserInterface interface {
	GetName() string
	GetSurname() string
	GetID() int64
	SetID(new_id int64)
}


type GenericUserInfoInterface interface {
	UserInterface
	IsAdmin() bool
}
