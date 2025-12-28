package user

type UserForm struct {
	Username string `json:"username" validate:"required,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100"`
}

func ToUserModel(form *UserForm) *User {
	return &User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}
}
