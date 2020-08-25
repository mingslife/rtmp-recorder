package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) string {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return ""
	} else {
		return string(hp)
	}
}
