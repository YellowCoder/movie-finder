package movie_scraper

import "github.com/PuerkitoBio/goquery"

type MovieElement interface {
	FindValue(doc *goquery.Document) error
	GetValue() map[string]interface{}
	PersistValue(movieURL string) error
}
