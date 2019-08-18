package scrape_model

type Movie struct {
	ID          int
	TableID     string
	Categories  []string
	Duration    string
	Rate        float64
	ReleaseDate string
	Title       string
	Url         string
}
