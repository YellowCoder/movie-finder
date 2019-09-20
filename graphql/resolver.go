package graphql

import (
	"context"
	"fmt"

	"github.com/YellowCoder/movie-finder/repository"
)

type Resolver struct {
	movies []*Movie
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMovie(ctx context.Context, input NewMovie) (*Movie, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Movies(ctx context.Context) ([]*Movie, error) {
	for _, repMovie := range repository.MovieRepository.All() {
		movie := &Movie{}
		movie.ID = repMovie.ID
		movie.Title = repMovie.Title
		fmt.Println(repMovie)

		r.movies = append(r.movies, movie)
	}

	return r.movies, nil
}
