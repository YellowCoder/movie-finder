package imdb_elements

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/scrape_model"
)

type ReleaseDate struct {
	query string
}

var validReleaseDate = regexp.MustCompile(`releaseinfo`)

func CreateReleaseDate() *ReleaseDate {
	return &ReleaseDate{query: ".title_wrapper .subtext a"}
}

func (c *ReleaseDate) FindValue(doc *goquery.Document, movie *scrape_model.Movie) error {

	doc.Find(c.query).Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")

		if validReleaseDate.MatchString(href) {
			movie.ReleaseDate = item.Text()
		}
	})

	return nil
}
