package main

import "github.com/YellowCoder/movie-finder/imdb_page_scraper"

func main() {
	url := "https://www.imdb.com/chart/top"

	scraper := imdb_page_scraper.CreateApplication(url)
	scraper.Execute()
}
