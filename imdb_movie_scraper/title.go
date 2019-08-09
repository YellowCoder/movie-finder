package imdb_movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type TitleScraper struct {
	query string
}

func CreateTitleScraper() *TitleScraper {
	return &TitleScraper{query: ".title_wrapper h1"}
}

func (t *TitleScraper) FindValue(doc *goquery.Document, movie *model.Movie) error {
	movie.Title = doc.Find(t.query).Text()
	return nil
}
