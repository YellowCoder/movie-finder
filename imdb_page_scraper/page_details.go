package imdb_page_scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
)

type pageDetails struct {
	url   string
	query string
}

const MAX_GO_ROUTINES = 20

func CreateDetails(pageUrl string, query string) *pageDetails {
	return &pageDetails{url: pageUrl, query: query}
}

func (p *pageDetails) Execute() {
	doc, _ := goquery.NewDocument(p.url)

	done := make(chan bool, MAX_GO_ROUTINES)

	doc.Find(p.query).Each(func(index int, item *goquery.Selection) {
		moviePath, _ := item.Attr("href")
		moviePath = strings.TrimSpace(moviePath)

		done <- true
		go func() {
			scraper := imdb_movie_scraper.CreateApplication(moviePath)
			scraper.Execute()
			<-done
		}()

	})
}
