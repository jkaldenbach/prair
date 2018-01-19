package models

import (
	"github.com/jinzhu/gorm"
)

// Church : church model
type Church struct {
	gorm.Model
	Name string `json:"name"`
}

// Churches : list of churches
type Churches []Church
