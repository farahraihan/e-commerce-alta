package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtilityInterface interface {
	GeneratePassword(string) ([]byte, error)
	CheckPassword([]byte, []byte) error
}

type passwordUtility struct{}

func NewPasswordUtility() PasswordUtilityInterface {
	return &passwordUtility{}
}

func (pu *passwordUtility) GeneratePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pu *passwordUtility) CheckPassword(inputPw []byte, currentPw []byte) error {
	return bcrypt.CompareHashAndPassword(currentPw, inputPw)
}
