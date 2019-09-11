package repository

import (
	"github.com/YellowCoder/movie-finder/config"
	"github.com/YellowCoder/movie-finder/database_model"
	"github.com/YellowCoder/movie-finder/scrape_model"
	"github.com/jinzhu/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

type movieRepositoryResult struct {
	Movie *database_model.Movie
	Error error
}

func CreateMovieRepository() *movieRepository {
	return &movieRepository{config.DatabaseConnection}
}

func (m *movieRepository) Create(movie *scrape_model.Movie) {
	m.db.Create(&database_model.Movie{Name: movie.Title})
}
