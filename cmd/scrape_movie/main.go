package main

import "github.com/YellowCoder/movie-finder/imdb_movie_scraper"

func main() {
	path := "/title/tt1369706/"

	scraper := imdb_movie_scraper.CreateApplication(path)
	scraper.Execute()
}
