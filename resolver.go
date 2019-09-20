package graphql

import (
	"context"

	"github.com/YellowCoder/movie-finder/graphql"
	"github.com/YellowCoder/movie-finder/repository"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMovie(ctx context.Context, input graphql.NewMovie) (*graphql.Movie, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
// 	panic("not implemented")
// }

func (r *queryResolver) Movies(ctx context.Context) ([]repository.Movie, error) {

}
