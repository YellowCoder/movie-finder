package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/YellowCoder/movie-finder/config"
)

func main() {
	url := "https://www.imdb.com/chart/top?ref_=nv_mv_250"

	doc, err := goquery.NewDocument(url)

	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	doc.Find(".lister table tbody tr").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(".titleColumn a").Attr("title")
		url, _ := s.Find(".titleColumn a").Attr("href")

		sqlStatement := `INSERT INTO movies (name, url) VALUES ($1, $2)`

		_, err = config.DatabaseConnection.Exec(sqlStatement, title, url)

		if err != nil {
			fmt.Println("Inserting error:", err)
		}
	})

}
