package movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
)

type Movie struct {
	url      string
	doc      *goquery.Document
	elements []MovieElement
}

func CreateMovieScraper(movieURL string) Movie {
	doc, _ := goquery.NewDocument(movieURL)

	return Movie{
		url: movieURL,
		doc: doc,
		elements: []MovieElement{
			imdb_movie_scraper.CreateTitleScraper(),
			imdb_movie_scraper.CreateRateScraper(),
		},
	}
}

func (m *Movie) Execute() (err error) {
	m.findElements()
	m.persistElements()

	return nil
}

func (m *Movie) findElements() error {
	for _, element := range m.elements {
		element.FindValue(m.doc)
	}
	return nil
}

func (m *Movie) persistElements() error {
	for _, element := range m.elements {
		element.PersistValue(m.url)
	}
	return nil
}
