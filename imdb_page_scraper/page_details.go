package imdb_page_scraper

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/imdb_movie_scraper"
)

type pageDetails struct {
	url   string
	query string
}

const HOST = "https://imdb.com"

func CreateDetails(pageUrl string, query string) *pageDetails {
	return &pageDetails{url: pageUrl, query: query}
}

func (p *pageDetails) Execute() {
	doc, _ := goquery.NewDocument(p.url)

	doc.Find(p.query).Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		href = strings.TrimSpace(href)

		fullMovieUrl := fmt.Sprintf("%s%s", HOST, href)

		scraper := imdb_movie_scraper.CreateApplication(fullMovieUrl)
		scraper.Execute()
		fmt.Println("Scraping:", fullMovieUrl)
	})
}
