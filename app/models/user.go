package models

import "github.com/revel/revel"

type User struct {
	Username string
	Password string
}

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Username)
	v.Required(user.Password)
}
