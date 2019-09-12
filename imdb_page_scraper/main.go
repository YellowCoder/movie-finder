package imdb_page_scraper

type scraper struct {
	pageDetails *pageDetails
}

func CreateApplication(pageUrl string) *scraper {
	pageDetails := CreateDetails(pageUrl, ".lister-list td a")

	return &scraper{
		pageDetails: pageDetails,
	}
}

func (s *scraper) Execute() {
	s.pageDetails.Execute()
}
