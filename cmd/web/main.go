package main

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/YellowCoder/movie-finder/repository"
)

type data struct {
	Movies []*repository.Movie
}

func main() {
	http.HandleFunc("/movies", moviesHandler)
	http.ListenAndServe(":80", nil)
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	message, _ := ioutil.ReadFile("views/movies.gohtml")
	t, _ := template.New("UsersPage").Parse(string(message))

	data := &data{Movies: repository.MovieRepository.All()}
	t.Execute(w, data)
}
