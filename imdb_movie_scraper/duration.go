package imdb_movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type DurationScraper struct {
	query string
}

func CreateDurationScraper() *DurationScraper {
	return &DurationScraper{query: ".subtext time"}
}

func (c *DurationScraper) FindValue(doc *goquery.Document, movie *model.Movie) error {
	movie.Duration = doc.Find(c.query).Text()
	return nil
}
