package repository

var Movie *movieRepository

func init() {
	Movie = CreateMovieRepository()
}
