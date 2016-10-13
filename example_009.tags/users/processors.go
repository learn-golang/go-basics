package users


// USER API //////////////////////////////////////////
func (user *User) Create(name, surname string) (*User, error) {
	_new_user := &User{name: name, surname:surname}
	err := _new_user.Validate()
	if err != nil {
		return nil, err
	} else {
		return _new_user, nil
	}
}

// Admin USER API //////////////////////////////////////////

func (admin *AdminUser) Create(name, surname string) (*AdminUser, error) {
	_new_user := &User{name: name, surname:surname}
	err := _new_user.Validate()
	if err != nil {
		return nil, err
	} else {
		return &AdminUser{User:_new_user, admin:true}, nil
	}
}
