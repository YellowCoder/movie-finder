package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/YellowCoder/movie-finder/env"
)

var DatabaseConnection *gorm.DB

func configureDatabase() {
	var err error
	DatabaseConnection, err = gorm.Open("postgres", env.Database.URL)

	if err != nil {
		fmt.Println("Error", err)
	}
}
