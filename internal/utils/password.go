package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CheckPassword(inputPw []byte, currentPw []byte) error {
	return bcrypt.CompareHashAndPassword(currentPw, inputPw)
}