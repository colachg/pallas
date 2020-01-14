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
	var projects []*models.Project

	err := p.DB.LogMode(true).Find(&projects).Error
	if err != nil {
		log.Fatal("Get project error:", err)
		return nil, err
	}
	return projects, nil
}

func (p *ProjectRepo) GetByID(id string) (*models.Project, error) {
	return nil, nil
}

func (p *ProjectRepo) CreateProject(project *models.Project) (*models.Project, error) {
	err := p.DB.LogMode(true).Create(project).Error
	if err != nil {
		log.Fatal("create project error:", err)
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
