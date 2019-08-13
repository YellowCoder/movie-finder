package movie_scraper

import (
	"fmt"

	"github.com/YellowCoder/movie-finder/model"
)

type movieSaver struct {
	movie *model.Movie
}

func CreateSaver(movie *model.Movie) *movieSaver {
	return &movieSaver{
		movie: movie,
	}
}

func (m *movieSaver) Execute() error {
	fmt.Println(m.movie.TableID)
	return nil
}
