package data

import "gorm.io/gorm"

type User struct {
	ID       uint64
	Login    string
	Password string
	Email    string
}

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db: db}
}
