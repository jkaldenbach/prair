package models

import (
	"github.com/jinzhu/gorm"
)

// User : user model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users : list of users
type Users []User
