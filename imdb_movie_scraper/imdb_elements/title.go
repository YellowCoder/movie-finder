package imdb_elements

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type TitleScraper struct {
	query string
}

func CreateTitleScraper() *TitleScraper {
	return &TitleScraper{query: ".title_wrapper h1"}
}

func (t *TitleScraper) FindValue(doc *goquery.Document, movie *scrape_model.Movie) error {
	movie.Title = doc.Find(t.query).Text()
	return nil
}
