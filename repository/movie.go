package repository

import (
	"github.com/YellowCoder/movie-finder/database_model"
)

type movieRepository struct{}

type movieRepositoryResult struct {
	Movie *database_model.Movie
	Error error
}

func CreateMovieRepository() *movieRepository {
	return &movieRepository{}
}

func (m *movieRepository) RealAll() (result *movieRepositoryResult) {
	// rows, _ := config.DatabaseConnection.Query("Select * from movies")

	// for rows.Next() {
	// 	movie := &model.Movie{}
	// 	fmt.Println(rows.Scan(&movie.ID, &movie.Title))
	// }
	return nil
}
