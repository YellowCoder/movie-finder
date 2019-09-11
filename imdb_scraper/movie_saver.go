package imdb_scraper

import (
	"github.com/YellowCoder/movie-finder/repository"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type movieSaver struct {
	movie *scrape_model.Movie
}

func CreateSaver(movie *scrape_model.Movie) *movieSaver {
	return &movieSaver{
		movie: movie,
	}
}

func (m *movieSaver) Execute() error {
	repository.MovieRepository.Create(m.movie)
	return nil
}

func (m *movieSaver) findMovie() error {
	return nil
}
