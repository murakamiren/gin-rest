package types

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string
	Disc string
	Value int
	Author string
	IsPublished bool
}

