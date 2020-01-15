package graphql

import (
	"context"
	"github.com/colachg/pallas/models"
	"strconv"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (m *mutationResolver) CreateProject(ctx context.Context, input models.CreateProject) (*models.Project, error) {
	project := &models.Project{
		Name:    input.Name,
		Version: input.Version,
	}
	return m.ProjectRepo.CreateProject(project)
}

func (m *mutationResolver) UpdateProject(ctx context.Context, id string, input models.UpdateProject) (*models.Project, error) {
	ID, _ := strconv.Atoi(id)
	project := &models.Project{
		ID:      ID,
		Name:    input.Name,
		Version: input.Version,
	}
	return m.ProjectRepo.UpdateProject(project)
}

func (m *mutationResolver) DeleteProject(ctx context.Context, id string) (bool, error) {
	ID, _ := strconv.Atoi(id)
	return m.ProjectRepo.DeleteProject(ID)
}
