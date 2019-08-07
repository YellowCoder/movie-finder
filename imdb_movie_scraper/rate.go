package imdb_movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
)

type RateScraper struct {
	query string
	value string
}

func CreateRateScraper() *RateScraper {
	return &RateScraper{query: ".ratingValue strong span"}
}

func (n *RateScraper) FindValue(doc *goquery.Document) error {
	n.value = doc.Find(n.query).Text()
	return nil
}

func (n *RateScraper) GetValue() map[string]interface{} {
	return map[string]interface{}{
		"rate": n.value,
	}
}

func (t *RateScraper) PersistValue(movieURL string) error {
	return nil
}
