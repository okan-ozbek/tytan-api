package auth

import (
	"fmt"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func NewAuth() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func Generate(id int, username string, email string) (string, error) {
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"id":       id,
		"username": username,
		"email":    email,
	})

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
