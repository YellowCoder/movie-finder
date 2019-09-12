package imdb_movie_scraper

import (
	"fmt"

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

func (m *movieSaver) Execute() {
	_, err := repository.MovieRepository.FindOrCreate(m.movie)

	if err != nil {
		fmt.Println(err)
	}
}

func (m *movieSaver) findMovie() error {
	return nil
}
