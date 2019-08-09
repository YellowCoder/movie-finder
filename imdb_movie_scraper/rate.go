package imdb_movie_scraper

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type RateScraper struct {
	query string
}

func CreateRateScraper() *RateScraper {
	return &RateScraper{query: ".ratingValue strong span"}
}

func (r *RateScraper) FindValue(doc *goquery.Document, movie *model.Movie) error {
	rate, _ := strconv.ParseFloat(doc.Find(r.query).Text(), 32)
	movie.Rate = rate

	return nil
}
