package utils

import "golang.org/x/crypto/bcrypt"

const bcryptFactor = 10

func HasHCompare(hashed, notHashed, salt string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(notHashed+salt)); err != nil {
		return false
	}
	return true
}

func HashWithSalt(value, salt string) string {
	epw, err := bcrypt.GenerateFromPassword([]byte(value+salt), bcryptFactor)

	if err != nil {
		return ""
	}

	return string(epw)
}
