package imdb_elements

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type RateScraper struct {
	query string
}

func CreateRateScraper() *RateScraper {
	return &RateScraper{query: ".ratingValue strong span"}
}

func (r *RateScraper) FindValue(doc *goquery.Document, movie *scrape_model.Movie) error {
	rate, _ := strconv.ParseFloat(doc.Find(r.query).Text(), 32)
	movie.Rate = rate

	return nil
}
