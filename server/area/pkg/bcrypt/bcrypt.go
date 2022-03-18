package bcrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePassword(plain, encry string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encry), []byte(plain))
	if err != nil {
		return false
	} else {
		return true
	}
}
