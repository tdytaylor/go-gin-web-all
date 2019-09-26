package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"default:'tom'"`
	Password string
	Email    string
	Phone    string
}

func (u User) TableName() string {
	return "user"
}

func SelectUserByName(name string) (user *User) {
	db.Where("name = ?", name).First(user)
	return
}
