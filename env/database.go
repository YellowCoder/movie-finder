package env

import (
	"fmt"

	"github.com/codingconcepts/env"
)

var Database = databaseConfig{}

type databaseConfig struct {
	URL            string `env:"DATABASE_URL"`
	MaxConnections int    `env:"DATABASE_MAX_CONNECTIONS" required:"true"`
}

func setDatabase() {
	if err := env.Set(&Database); err != nil {
		fmt.Println("Database configuration not found:", err)
	}

	if len(Database.URL) == 0 {
		fmt.Println("Database URL not found.")
	}
}
