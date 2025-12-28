package user

import (
	"time"
	hashUtil "tytan-api/util/hash"
)

type UserDTO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserForm struct {
	Username string `json:"username" validate:"required,alpha_space"`
	Password string `json:"password" validate:"required,max=100"`
}

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserModel(form *UserForm) *User {
	password, err := hashUtil.Hash(form.Password)
	if err != nil {
		return nil
	}

	return &User{
		Username:  form.Username,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
