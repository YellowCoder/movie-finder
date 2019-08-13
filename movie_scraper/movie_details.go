package movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
	"github.com/YellowCoder/movie-finder/model"
)

type movieDetails struct {
	url      string
	doc      *goquery.Document
	movie    *model.Movie
	elements []MovieElement
}

func CreateDetails(movie *model.Movie) *movieDetails {
	doc, _ := goquery.NewDocument(movie.Url)

	return &movieDetails{
		doc:   doc,
		movie: movie,
		elements: []MovieElement{
			imdb_movie_scraper.CreateTitleScraper(),
			imdb_movie_scraper.CreateRateScraper(),
			imdb_movie_scraper.CreateCategoryScraper(),
			imdb_movie_scraper.CreateDurationScraper(),
		},
	}
}

func (m *movieDetails) Execute() (err error) {
	m.findElements()
	return nil
}

func (m *movieDetails) findElements() error {
	for _, element := range m.elements {
		element.FindValue(m.doc, m.movie)
	}
	return nil
}
