package imdb_movie_scraper

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type RateScraper struct {
	query string
	value float64
}

func CreateRateScraper() *RateScraper {
	return &RateScraper{query: ".ratingValue strong span"}
}

func (n *RateScraper) FindValue(doc *goquery.Document) error {
	rate, _ := strconv.ParseFloat(doc.Find(n.query).Text(), 32)
	n.value = rate

	return nil
}

func (n *RateScraper) PersistValue(movie *model.Movie) error {
	movie.Rate = n.value
	return nil
}
