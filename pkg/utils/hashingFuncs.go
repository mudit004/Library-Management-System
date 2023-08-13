package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd string) string {
	var saltRounds int = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), saltRounds)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func Check(inputPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
