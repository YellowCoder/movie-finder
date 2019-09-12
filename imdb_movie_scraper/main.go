package imdb_movie_scraper

import "github.com/YellowCoder/movie-finder/scrape_model"

type scraper struct {
	movieDetails *movieDetails
	movieSaver   *movieSaver
}

func CreateApplication(url string) *scraper {
	movie := &scrape_model.Movie{Url: url}

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
