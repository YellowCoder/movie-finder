package main

import "github.com/YellowCoder/movie-finder/imdb_page_scraper"

var pagePaths = [2]string{"/chart/top", "/chart/moviemeter"}

func main() {
	for _, path := range pagePaths {
		scraper := imdb_page_scraper.CreateApplication(path)
		scraper.Execute()
	}
}
