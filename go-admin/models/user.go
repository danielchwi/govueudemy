package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (u User) SetPassword(password []byte) []byte {
	hashPassword, _ := bcrypt.GenerateFromPassword(password, 14)
	return hashPassword
}

func (u User) ComparePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(u.Password, password)
}
