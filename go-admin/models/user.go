package models

type User struct {
	Id        uint
	FirstName string
	Lastname  string
	Email     string `gorm:"unique"`
	Password  []byte
}
