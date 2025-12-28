package user

import (
	"time"
)

type UserDTO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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
