package imdb_page_scraper

import (
	"fmt"
)

type scraper struct {
	pageDetails *pageDetails
}

const HOST = "https://imdb.com"

func CreateApplication(pagePath string) *scraper {
	pageUrl := stripQueryStringFromUrl(pagePath)
	pageDetails := CreateDetails(pageUrl, ".lister-list .titleColumn a")

	return &scraper{
		pageDetails: pageDetails,
	}
}

func (s *scraper) Execute() {
	s.pageDetails.Execute()
}

func stripQueryStringFromUrl(url string) string {
	return fmt.Sprintf("%s%s", HOST, url)
}
