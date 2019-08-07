package main

import (
	"github.com/YellowCoder/movie-finder/movie_scraper"
)

func main() {
	url := "https://www.imdb.com/title/tt0111161/?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=e31d89dd-322d-4646-8962-327b42fe94b1&pf_rd_r=DJVPQSBSMW27TEWNS65C&pf_rd_s=center-1&pf_rd_t=15506&pf_rd_i=top&ref_=chttp_tt_1"

	scraper := movie_scraper.CreateMovieScraper(url)
	scraper.Execute()
}
