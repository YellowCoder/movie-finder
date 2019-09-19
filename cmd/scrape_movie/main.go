package main

import (
	"os"

	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
)

func main() {
	path := os.Args[1]

	scraper := imdb_movie_scraper.CreateApplication(path)
	scraper.Execute()
}
