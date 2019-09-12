package imdb_elements

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type DurationScraper struct {
	query string
}

func CreateDurationScraper() *DurationScraper {
	return &DurationScraper{query: ".subtext time"}
}

func (c *DurationScraper) FindValue(doc *goquery.Document, movie *scrape_model.Movie) error {
	movie.Duration = doc.Find(c.query).Text()
	return nil
}
