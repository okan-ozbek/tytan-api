package auth

type AuthForm struct {
	Username string `json:"username" validate:"required,alpha_space"`
	Password string `json:"password" validate:"required,max=100"`
}
