package main

import "github.com/YellowCoder/movie-finder/imdb_movie_scraper"

func main() {
	// url := "https://www.imdb.com/title/tt5886046/"
	url := "https://www.imdb.com/title/tt1369706/"

	scraper := imdb_movie_scraper.CreateApplication(url)
	scraper.Execute()
}
