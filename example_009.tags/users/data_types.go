package users


type User struct {
	name string `validate:"length:[0,255];case:upper_only"`
	surname string `validate:"length:[0,255];case:upper_only"`
	id int64 `validate:"value:zero_only"`
}


type AdminUser struct {
	*User
	admin bool
}
