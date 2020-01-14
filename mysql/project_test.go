package mysql

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/colachg/pallas/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ProjectRepo
	project    models.Project
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = ProjectRepo{DB: s.DB}
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestProjectRepo_CreateProject() {
	var (
		t1        = time.Now()
		name      = "A"
		version   = "v20200114"
		createdAt = t1
		updatedAt = t1
	)

	project := models.Project{
		Name:      name,
		Version:   version,
		CreatedAt: t1,
		UpdatedAt: t1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `projects`").
		WithArgs(name, version, createdAt, updatedAt).
		WillReturnResult(
			sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	_, err := s.repository.CreateProject(&project)
	require.NoError(s.T(), err)
}
