package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Category struct {
	Movies []*Movie `gorm:"many2many:categories_movies;"`
	ID     uint     `gorm:"primary_key"`
	Name   string
}

type categoryRepository struct {
	db *gorm.DB
}

func (r *categoryRepository) FirstOrCreate(name string) (*Category, error) {
	category := &Category{}
	requestedCategory := &Category{Name: name}

	query := r.db.Where(requestedCategory)
	if err := query.FirstOrCreate(&category).Error; err != nil {
		return nil, fmt.Errorf("Category FirstOrCreate error: %v", err)
	}

	return category, nil
}
