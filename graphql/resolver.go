package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"github.com/colachg/pallas"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProject(ctx context.Context, input pallas.CreateProject) (*pallas.Project, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateProject(ctx context.Context, input pallas.UpdateProject) (*pallas.Project, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Projects(ctx context.Context) ([]*pallas.Project, error) {
	panic("not implemented")
}
