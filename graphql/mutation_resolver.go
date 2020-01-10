package graphql

import (
	"context"
	"fmt"
	"github.com/colachg/pallas/models"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (m *mutationResolver) CreateProject(ctx context.Context, input models.CreateProject) (*models.Project, error) {
	project := &models.Project{
		Name:        input.Name,
		Version:     input.Version,
		Description: input.Description,
	}
	fmt.Println("1.---------->", project)
	return m.ProjectRepo.CreateProject(project)
}

func (m *mutationResolver) UpdateProject(ctx context.Context, id string, input models.UpdateProject) (*models.Project, error) {
	panic("implement me")
}

func (m *mutationResolver) DeleteProject(ctx context.Context, id string) (bool, error) {
	panic("implement me")
}
