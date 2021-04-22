package data

import (
	"fmt"
	"gorm.io/gorm"
)

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

func (u UserData) IsLoginExists(login string) (bool, error) {
	var user User
	result := u.db.Where("login", login).Find(&user)
	if result.Error != nil {
		return false, fmt.Errorf("can't get user from database, error: %w", result.Error)
	}
	if user.Login == "" {
		return false, nil
	}
	return true, nil
}
