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
	movie, err := repository.MovieRepository.FindOrCreate(m.movie)

	if err == nil {
		m.saveCategories(movie)
	} else {
		fmt.Println(err)
	}
}

func (m *movieSaver) saveCategories(movie *repository.Movie) {
	repository.MovieRepository.ClearCategories(movie)

	for _, categoryName := range m.movie.Categories {
		category, _ := repository.CategoryRepository.FirstOrCreate(categoryName)
		repository.MovieRepository.AppendCategory(movie, category)
	}
}
