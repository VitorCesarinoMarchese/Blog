package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func ConnectDb() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Maybe your db isn't working, just maybe")
	}
	return db
}