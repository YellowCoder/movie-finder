package config

import (
	"database/sql"
	"fmt"

	"github.com/YellowCoder/movie-finder/env"
)

type databaseConnection struct {
	*sql.DB
}

var DatabaseConnection *databaseConnection

func openDatabaseConnection() {
	db, err := sql.Open("postgres", env.Database.URL)

	if err != nil {
		fmt.Println("Error", err)
	}

	DatabaseConnection = &databaseConnection{db}
}
