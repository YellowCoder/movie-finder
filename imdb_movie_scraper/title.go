package imdb_movie_scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type TitleScraper struct {
	query string
	value string
}

func CreateTitleScraper() *TitleScraper {
	return &TitleScraper{query: ".title_wrapper h1"}
}

func (t *TitleScraper) FindValue(doc *goquery.Document) error {
	t.value = doc.Find(t.query).Text()
	return nil
}

func (t *TitleScraper) GetValue() map[string]interface{} {
	return map[string]interface{}{
		"title": t.value,
	}
}

func (t *TitleScraper) PersistValue(movieURL string) error {
	fmt.Println(movieURL)
	return nil
}
