package repository

import "github.com/jinzhu/gorm"

type Category struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

type categoryRepository struct {
	db *gorm.DB
}
