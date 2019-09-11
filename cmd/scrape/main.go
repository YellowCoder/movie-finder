package main

import (
	"github.com/YellowCoder/movie-finder/imdb_scraper"
)

func main() {
	// url := "https://www.imdb.com/title/tt5886046/"
	url := "https://www.imdb.com/title/tt1369706/"

	scraper := imdb_scraper.CreateApplication(url)
	scraper.Execute()
}
