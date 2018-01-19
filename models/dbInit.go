package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // db dialect
)

// DB : database connection
var DB gorm.DB

// InitDb : initialize the database connection
func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "prayr.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	db.SingularTable(true)
	db.AutoMigrate(
		&Church{},
		&User{},
	)

	DB = *db
	return db
}
