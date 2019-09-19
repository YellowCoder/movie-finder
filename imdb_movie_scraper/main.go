package imdb_movie_scraper

import (
	"fmt"
	"regexp"

	"github.com/YellowCoder/movie-finder/scrape_model"
)

type scraper struct {
	movieDetails *movieDetails
	movieSaver   *movieSaver
}

const HOST = "https://imdb.com"
const URL_REGEX = "([^\\?]+)(\\?.*)?"

func CreateApplication(path string) *scraper {
	movie := &scrape_model.Movie{Url: stripQueryStringFromUrl(path)}

	fmt.Println("Scraping:", movie.Url)

	movieDetails := CreateDetails(movie)
	movieSaver := CreateSaver(movie)

	return &scraper{
		movieDetails: movieDetails,
		movieSaver:   movieSaver,
	}
}

func (s *scraper) Execute() {
	s.movieDetails.Execute()
	s.movieSaver.Execute()
}

func stripQueryStringFromUrl(url string) string {
	fullUrl := fmt.Sprintf("%s%s", HOST, url)

	r := regexp.MustCompile(URL_REGEX)
	return r.FindStringSubmatch(fullUrl)[1]
}
