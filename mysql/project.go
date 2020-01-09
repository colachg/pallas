package mysql

import (
	"github.com/colachg/pallas/models"
	"github.com/jinzhu/gorm"
)

type ProjectRepo struct {
	DB *gorm.DB
}

func (p *ProjectRepo) GetProjects() ([]*models.Project, error) {

	projects := p.DB.Find(&models.Project{})
	return nil, nil
}

func (p *ProjectRepo) GetByID(id string) (*models.Project, error) {
	return nil, nil
}

func (p *ProjectRepo) CreateProject(project *models.Project) (*models.Project, error) {
	p.DB.NewRecord(project)
	p.DB.CreateTable(project)
	return project, nil
}

func (p *ProjectRepo) UpdateProject(project *models.Project) (*models.Project, error) {
	return nil, nil
}

func (p *ProjectRepo) Delete(project *models.Project) error {
	return nil
}
