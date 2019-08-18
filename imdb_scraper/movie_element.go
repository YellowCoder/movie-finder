package imdb_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type MovieElement interface {
	FindValue(doc *goquery.Document, movie *scrape_model.Movie) error
}
