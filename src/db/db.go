package db

import (
	"gin-rest/src/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	db, err := gorm.Open(sqlite.Open("src/db/data/test.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open db")
	}
	db.AutoMigrate(&types.Book{})

	DB = db
}