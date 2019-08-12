package imdb_movie_scraper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/model"
)

type ReleaseDate struct {
	query string
}

var validReleaseDate = regexp.MustCompile(`releaseinfo`)

func CreateReleaseDate() *ReleaseDate {
	return &ReleaseDate{query: ".title_wrapper .subtext a"}
}

func (c *ReleaseDate) FindValue(doc *goquery.Document, movie *model.Movie) error {

	doc.Find(c.query).Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")

		if validReleaseDate.MatchString(href) {
			// name := item.Text()
			// fmt.Println(name)
		}
	})

	return nil
}
