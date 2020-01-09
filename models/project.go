package models

type Project struct {
	//gorm.Model
	ID          string
	Name        string
	Version     string
	Description string
}
