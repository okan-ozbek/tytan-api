package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	return string(bytes), err
}

func Compare(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
