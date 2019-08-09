package movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type MovieElement interface {
	FindValue(doc *goquery.Document, movie *model.Movie) error
}
