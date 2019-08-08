package movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
	"github.com/YellowCoder/movie-finder/model"
)

type ScrapedMovie struct {
	url      string
	doc      *goquery.Document
	movie    *model.Movie
	elements []MovieElement
}

func CreateMovieScraper(movieURL string) ScrapedMovie {
	doc, _ := goquery.NewDocument(movieURL)

	return ScrapedMovie{
		url:   movieURL,
		doc:   doc,
		movie: &model.Movie{Url: movieURL},
		elements: []MovieElement{
			imdb_movie_scraper.CreateTitleScraper(),
			imdb_movie_scraper.CreateRateScraper(),
		},
	}
}

func (m *ScrapedMovie) Execute() (err error) {
	m.findElements()
	m.updateMovie()

	return nil
}

func (m *ScrapedMovie) findElements() error {
	for _, element := range m.elements {
		element.FindValue(m.doc)
	}
	return nil
}

func (m *ScrapedMovie) updateMovie() error {
	for _, element := range m.elements {
		element.PersistValue(m.movie)
	}
	return nil
}
