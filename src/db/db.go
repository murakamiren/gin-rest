package db

import (
	"gin-rest/src/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbInit() {
	Db, err := gorm.Open(sqlite.Open("src/db/data/test.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open db")
	}
	Db.AutoMigrate(&types.Book{})
}