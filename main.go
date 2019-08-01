package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/YellowCoder/movie-finder/config"
	"github.com/YellowCoder/movie-finder/model"
)

func main() {
	// This is only and example
	category := &model.Category{}

	sqlStatement := `SELECT name FROM categories WHERE id=$1;`

	row := config.DatabaseConnection.QueryRow(sqlStatement, 1)
	row.Scan(&category.Name)

	fmt.Println(category)
}
