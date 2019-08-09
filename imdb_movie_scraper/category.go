package imdb_movie_scraper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type CategoryScraper struct {
	query string
}

var validCategory = regexp.MustCompile(`genres=`)

func CreateCategoryScraper() *CategoryScraper {
	return &CategoryScraper{query: ".title_wrapper .subtext a"}
}

func (c *CategoryScraper) FindValue(doc *goquery.Document, movie *model.Movie) error {
	categories := []string{}

	doc.Find(c.query).Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")

		if validCategory.MatchString(href) {
			name := item.Text()
			categories = append(categories, name)
		}
	})

	movie.Categories = categories

	return nil
}
