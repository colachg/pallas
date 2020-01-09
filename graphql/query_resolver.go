package graphql

import (
	"context"
	"github.com/colachg/pallas/models"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (q *queryResolver) Projects(ctx context.Context) ([]*models.Project, error) {
	return q.ProjectRepo.GetProjects()
}
