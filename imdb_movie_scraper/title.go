package imdb_movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
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

func (t *TitleScraper) PersistValue(movie *model.Movie) error {
	movie.Title = t.value
	return nil
}
