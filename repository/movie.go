package repository

import (
	"fmt"

	"github.com/YellowCoder/movie-finder/scrape_model"
	"github.com/jinzhu/gorm"
)

type Movie struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Url         string
	Rate        float64
	ReleaseDate string
	Duration    string
}

type movieRepository struct {
	db *gorm.DB
}

func (r *movieRepository) Create(scraped_movie *scrape_model.Movie) (*Movie, error) {
	movie := &Movie{}
	movie.Title = scraped_movie.Title
	movie.Url = scraped_movie.Url
	movie.Rate = scraped_movie.Rate
	movie.Duration = scraped_movie.Duration
	movie.ReleaseDate = scraped_movie.ReleaseDate

	if err := r.db.Create(movie).Error; err != nil {
		return nil, fmt.Errorf("Movie Create error: %v", err)
	}

	return movie, nil
}

func (r *movieRepository) FindByUrl(url string) *Movie {
	movie := &Movie{}

	if err := r.db.Where(&Movie{Url: url}).First(&movie).Error; err != nil {
		fmt.Errorf("Movie FindByUrl error: %v", err)
		return nil
	}

	return movie
}

func (r *movieRepository) FindOrCreate(scraped_movie *scrape_model.Movie) (*Movie, error) {
	requestedMovie := r.FindByUrl(scraped_movie.Url)

	if requestedMovie == nil {
		movie, err := r.Create(scraped_movie)

		if err != nil {
			return nil, err
		}

		return movie, nil
	}

	return requestedMovie, nil
}

func (r *movieRepository) FirstOrCreate(scraped_movie *scrape_model.Movie) (*Movie, error) {
	movie := &Movie{}
	requestedMovie := &Movie{Url: scraped_movie.Url}

	query := r.db.Where(requestedMovie)
	if err := query.FirstOrCreate(&movie).Error; err != nil {
		return nil, fmt.Errorf("Movie FirstOrCreate error: %v", err)
	}

	return movie, nil
}
