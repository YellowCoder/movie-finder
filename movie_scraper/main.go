package movie_scraper

import "github.com/YellowCoder/movie-finder/model"

type scraper struct {
	movieDetails *movieDetails
	movieSaver   *movieSaver
}

func CreateApplication(url string) *scraper {
	movie := &model.Movie{Url: url, TableID: "movies"}

	movieDetails := CreateDetails(movie)
	movieSaver := CreateSaver(movie)

	return &scraper{
		movieDetails: movieDetails,
		movieSaver:   movieSaver,
	}
}

func (s *scraper) Execute() error {
	s.movieDetails.Execute()
	s.movieSaver.Execute()
	return nil
}
