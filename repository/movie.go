package repository

import (
	"github.com/YellowCoder/movie-finder/scrape_model"
	"github.com/jinzhu/gorm"
)

type Movie struct {
	ID   uint `gorm:"primary_key"`
	Name string
	Url  string
	Rate float64
}

type movieRepository struct {
	db *gorm.DB
}

func (m *movieRepository) Create(scraped_movie *scrape_model.Movie) {
	movie := &Movie{}
	movie.Name = scraped_movie.Title
	movie.Url = scraped_movie.Url
	movie.Rate = scraped_movie.Rate

	m.db.Create(movie)
}
