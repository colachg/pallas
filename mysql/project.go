package mysql

import (
	"github.com/colachg/pallas/models"
	"github.com/jinzhu/gorm"
	"log"
)

type ProjectRepo struct {
	DB *gorm.DB
}

func (p *ProjectRepo) GetProjects() ([]*models.Project, error) {
	return nil, nil
}

func (p *ProjectRepo) GetByID(id string) (*models.Project, error) {
	return nil, nil
}

//Todo: generate ID automatically
func (p *ProjectRepo) CreateProject(project *models.Project) (*models.Project, error) {
	err := p.DB.Create(project).Error

	if err != nil {
		log.Println("create project error:", err)
		return nil, err
	}
	return project, nil
}

func (p *ProjectRepo) UpdateProject(project *models.Project) (*models.Project, error) {
	return nil, nil
}

func (p *ProjectRepo) Delete(project *models.Project) error {
	return nil
}
