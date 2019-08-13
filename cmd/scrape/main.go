package main

import (
	"github.com/YellowCoder/movie-finder/movie_scraper"
)

func main() {
	url := "https://www.imdb.com/title/tt5886046/"

	scraper := movie_scraper.CreateApplication(url)
	scraper.Execute()
}
