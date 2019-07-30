package main

import (
	"fmt"
	"os"

	"github.com/YellowCoder/movie-finder/env"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	switch os.Args[1] {
	case "migrate":
		migrateDatabase(env.Database.URL)
	default:
		fmt.Println("Invalid action")
	}
}

func migrateDatabase(databaseURL string) error {
	migrations, err := migrate.New("file://./migrations", databaseURL)

	if err != nil {
		return fmt.Errorf("Could not create migrations: %v", err)
	}

	err = migrations.Up()
	if err != nil {
		return fmt.Errorf("Could not migrate the database: %v", err)
	}

	return nil
}
