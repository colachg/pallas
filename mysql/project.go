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
		log.Fatal("Get projects error:", err)
		return nil, err
	}
	return projects, nil
}

func (p *ProjectRepo) GetByID(id string) (*models.Project, error) {

	project := &models.Project{}
	err := p.DB.LogMode(true).Where("id = ?", id).First(project).Error
	if err != nil {
		log.Fatal("Get project by ID error:", err)
		return nil, err
	}
	return project, nil
}

func (p *ProjectRepo) CreateProject(project *models.Project) (*models.Project, error) {
	err := p.DB.LogMode(true).Create(project).Error
	if err != nil {
		log.Fatal("Create project error:", err)
		return nil, err
	}
	return project, nil
}

func (p *ProjectRepo) UpdateProject(project *models.Project) (*models.Project, error) {
	err := p.DB.LogMode(true).Model(project).Update(project).Error
	if err != nil {
		log.Fatal("Update project error:", err)
		return nil, err
	}
	return project, nil
}

func (p *ProjectRepo) DeleteProject(id int) (bool, error) {
	project := &models.Project{}

	err := p.DB.Where("id = ?", id).Delete(project).Error
	if err != nil {
		log.Fatal("Delete project error:", err)
		return false, err
	}
	return true, nil
}
