package user

import (
	"time"
)

type UserForm struct {
	Username string `json:"username" validate:"required,alpha_space"`
	Password string `json:"password" validate:"required,max=100"`
}

func ToUserModel(form *UserForm) *User {
	return &User{
		Username:  form.Username,
		Password:  form.Password,
		CreatedAt: time.Now(),
	}
}
