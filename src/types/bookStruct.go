package types

import "github.com/jinzhu/gorm"

type BookStruct struct {
	gorm.Model
	Title string
	Disc string
	Value int
	Author string
	IsPublished bool
}

