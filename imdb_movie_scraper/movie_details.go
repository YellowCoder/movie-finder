package imdb_movie_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper/imdb_elements"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type movieDetails struct {
	url      string
	doc      *goquery.Document
	movie    *scrape_model.Movie
	elements []MovieElement
}

func CreateDetails(movie *scrape_model.Movie) *movieDetails {
	doc, _ := goquery.NewDocument(movie.Url)

	return &movieDetails{
		doc:   doc,
		movie: movie,
		elements: []MovieElement{
			imdb_elements.CreateTitleScraper(),
			imdb_elements.CreateRateScraper(),
			imdb_elements.CreateCategoryScraper(),
			imdb_elements.CreateDurationScraper(),
			imdb_elements.CreateReleaseDate(),
		},
	}
}

func (m *movieDetails) Execute() (err error) {
	for _, element := range m.elements {
		element.FindValue(m.doc, m.movie)
	}
	return nil
}
