package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(p string) string {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func CheckBcrypt(h, p []byte) bool {

	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}
